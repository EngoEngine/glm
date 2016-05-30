package geo

import (
	"fmt"
	"github.com/engoengine/glm"
	"github.com/engoengine/glm/geo/internal/qhull"
	"github.com/luxengine/math"
)

// Quickhull returns the convex hull of the given points.
func Quickhull(points []glm.Vec3) {
	// TODO(hydroflame): finish implementation

	// 1 Find Initial tetrahedron
	// 1.1 Find Initial Triangle
	extremumIndices, extremums := qhull.FindExtremums(points)

	// 0.2 calculate the epsilon
	epsilon := qhull.CalculateEpsilon(extremums)

	// 1.2 Find the 3 most extreme points
	var triangleIndices [3]int
	var maxArea float32
	for i := 0; i < len(extremums); i++ {
		for j := i + 1; j < len(extremums); j++ {
			for k := j + 1; k < len(extremums); k++ {
				l1, l2, l3 := extremums[i].Sub(&extremums[j]), extremums[i].Sub(&extremums[k]), extremums[j].Sub(&extremums[k])
				area := TriangleAreaFromLengths(l1.Len(), l2.Len(), l3.Len())
				if area > maxArea {
					maxArea = area
					triangleIndices = [3]int{i, j, k}
				}
			}
		}
	}

	// 1.3 Finish the tetrahedron

	// find the last point
	l1, l2 := extremums[triangleIndices[1]].Sub(&extremums[triangleIndices[0]]), extremums[triangleIndices[2]].Sub(&extremums[triangleIndices[0]])
	dir := l1.Cross(&l2)

	imin, imax := ExtremePointsAlongDirection(&dir, points)
	vmin, vmax := points[imin].Sub(&extremums[triangleIndices[0]]), points[imax].Sub(&extremums[triangleIndices[0]])
	p1, p2 := math.Abs(vmin.Dot(&dir)), math.Abs(vmax.Dot(&dir))

	var tetraIndex int
	if p1 > p2 {
		tetraIndex = imin
	} else {
		tetraIndex = imax
	}

	faces := qhull.BuildInitialTetrahedron(tetraIndex, extremumIndices[triangleIndices[0]], extremumIndices[triangleIndices[1]], extremumIndices[triangleIndices[2]], points)

	for _, face := range faces {
		for n := range points {
			if dist := DistToTriangle(&points[n], &points[face.Vertices[0]], &points[face.Vertices[1]], &points[face.Vertices[2]]); dist > epsilon {
				face.Conflicts = append(face.Conflicts, qhull.Conflict{Distance: dist, Index: n})
			}
		}
	}

	{ // debug
		fmt.Println("debugging initial tetrahedron")
		for _, face := range faces {
			fmt.Println("face:      ", fmt.Sprintf("%p", face), points[face.Vertices[0]], points[face.Vertices[1]], points[face.Vertices[2]])
			fmt.Println("edges:     ", face.Edges)
			fmt.Println("conflicts: ", face.Conflicts)
			fmt.Println()
		}
	}

	var cnt int
	for iface, iconflict := qhull.NextConflict(faces); iface != -1; iface, iconflict = qhull.NextConflict(faces) {
		fmt.Println("conflict", points[faces[iface].Conflicts[iconflict].Index])

		qhull.CleanVisited(faces)
		for _, face := range faces {
			fmt.Printf("%p %t\n", face, face.Visited)
		}

		horizon := qhull.FindHorizon(faces[iface], &points[faces[iface].Conflicts[iconflict].Index])
		fmt.Println("horizon: ", horizon)
		newfaces := make([]qhull.Face, len(horizon), len(horizon))

		for n, edge := range horizon {
			newfaces[n].Vertices = [3]int{faces[iface].Conflicts[iconflict].Index, edge.Tail, edge.Twin.Tail}
			newfaces[n].Edges = [3]*qhull.Edge{
				{Tail: newfaces[n].Vertices[1], Face: &newfaces[n]},
				{Tail: newfaces[n].Vertices[0], Face: &newfaces[n], Twin: edge},
				{Tail: newfaces[n].Vertices[2], Face: &newfaces[n]},
			}
			edge.Twin = newfaces[n].Edges[1]

			newfaces[n].Edges[0].Next = newfaces[n].Edges[1]
			newfaces[n].Edges[1].Next = newfaces[n].Edges[2]
			newfaces[n].Edges[2].Next = newfaces[n].Edges[0]

			newfaces[n].Edges[2].Prev = newfaces[n].Edges[1]
			newfaces[n].Edges[0].Prev = newfaces[n].Edges[2]
			newfaces[n].Edges[1].Prev = newfaces[n].Edges[0]

			newfaces[n].Faces = [3]*qhull.Face{nil, edge.Twin.Face, nil}
			l1, l2 := points[newfaces[n].Vertices[1]].Sub(&points[newfaces[n].Vertices[0]]), points[newfaces[n].Vertices[2]].Sub(&points[newfaces[n].Vertices[0]])
			newfaces[n].Normal = l2.Cross(&l1)
			newfaces[n].Point = points[newfaces[n].Vertices[0]]

			for _, c := range faces[iface].Conflicts {
				if dist := DistToTriangle(&points[c.Index], &points[newfaces[n].Vertices[0]], &points[newfaces[n].Vertices[1]], &points[newfaces[n].Vertices[2]]); dist > epsilon {
					newfaces[n].Conflicts = append(newfaces[n].Conflicts, qhull.Conflict{Distance: dist, Index: c.Index})
				}
			}

			for _, c := range faces[iface].Conflicts {
				if dist := DistToTriangle(&points[c.Index], &points[newfaces[n].Vertices[0]], &points[newfaces[n].Vertices[1]], &points[newfaces[n].Vertices[2]]); dist > epsilon {
					newfaces[n].Conflicts = append(newfaces[n].Conflicts, qhull.Conflict{Distance: dist, Index: c.Index})
				}
			}

			// check surrounding faces in case we suddently see something new
			for _, face := range faces[iface].Faces {
				for _, c := range face.Conflicts {
					if dist := DistToTriangle(&points[c.Index], &points[newfaces[n].Vertices[0]], &points[newfaces[n].Vertices[1]], &points[newfaces[n].Vertices[2]]); dist > epsilon {
						newfaces[n].Conflicts = append(newfaces[n].Conflicts, qhull.Conflict{Distance: dist, Index: c.Index})
					}
				}
			}

		}

		for n := range newfaces {
			newfaces[n].Edges[0].Twin = newfaces[(n-1+len(newfaces))%len(newfaces)].Edges[2]
			newfaces[n].Edges[2].Twin = newfaces[(n+1)%len(newfaces)].Edges[1]
			newfaces[n].Faces[0] = &newfaces[(n-1+len(newfaces))%len(newfaces)]
			newfaces[n].Faces[2] = &newfaces[(n+1)%len(newfaces)]

			faces = append(faces, &newfaces[n])
		}

		//clean old faces
		var oldfaces []*qhull.Face
		for _, e := range horizon {
			var found bool
			for _, face := range oldfaces {
				if e.Face == face {
					found = true
					break
				}
			}

			if !found {
				oldfaces = append(oldfaces, e.Face)
			}
		}

		for _, face := range oldfaces {
			for n := 0; n < len(faces); n++ {
				if face == faces[n] {
					faces = faces[:n+copy(faces[n:], faces[n+1:])]
					n--
				}
			}
		}

		{ // debug
			for _, face := range faces {
				fmt.Println("face:      ", fmt.Sprintf("%p", face), points[face.Vertices[0]], points[face.Vertices[1]], points[face.Vertices[2]])
				fmt.Println("edges:     ", face.Edges)
				fmt.Println("conflicts: ", face.Conflicts)
				fmt.Println()
			}
		}

		if cnt >= 1 {
			break
		}
		cnt++

	}

	/*for _, face := range faces {
		fmt.Printf("v %f %f %f\n", points[face.Vertices[0]][0], points[face.Vertices[0]][1], points[face.Vertices[0]][2])
		fmt.Printf("v %f %f %f\n", points[face.Vertices[1]][0], points[face.Vertices[1]][1], points[face.Vertices[1]][2])
		fmt.Printf("v %f %f %f\n", points[face.Vertices[2]][0], points[face.Vertices[2]][1], points[face.Vertices[2]][2])
	}

	for n := range faces {
		fmt.Printf("f %d %d %d\n", (n)*3+3, (n)*3+2, (n)*3+1)
	}*/

}
