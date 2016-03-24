package quickhull3

import (
	"fmt"
	"github.com/luxengine/glm"
	"github.com/luxengine/math"
)

const (
	epsilonbase = 0.0001 // arbitrary value that will require testing
)

type (
	// Edge is a quickhull utility struct for edges
	Edge struct {
		Tail             int
		Prev, Next, Twin *Edge
		Face             *Face
	}
	// Face is a quickhull utility struct for faces
	Face struct {
		Edges     [3]*Edge
		Faces     [3]*Face
		Vertices  [3]int
		Conflicts []Conflict

		// tmp variables
		Visited bool
		Normal  glm.Vec3
		Point   glm.Vec3
	}

	// Conflict is a vertex that isn't inside the convex hull yet
	Conflict struct {
		Distance float32
		Index    int
	}
)

func (f *Face) canSee(point *glm.Vec3) bool {
	ap := point.Sub(&f.Point)
	return ap.Dot(&f.Normal) > 0
}

func CleanVisited(faces []*Face) {
	for _, face := range faces {
		face.Visited = false
	}
}

// FindHorizon finds the horizon of the conflict
func FindHorizon(face *Face, point *glm.Vec3) []*Edge {
	return findHorizon(face, nil, point)
}

func findHorizon(face *Face, edge *Edge, point *glm.Vec3) []*Edge {
	fmt.Println("face, edge", fmt.Sprintf("%p", face), fmt.Sprintf("%p", edge), " visited ", face.Visited)
	if edge == nil {
		e := face.Edges[0]
		fmt.Println("nil 1")
		edges := findHorizon(e.Twin.Face, e.Twin, point)
		e = e.Next
		fmt.Println("nil 2")
		edges = append(edges, findHorizon(e.Twin.Face, e.Twin, point)...)
		e = e.Next
		fmt.Println("nil 3")
		edges = append(edges, findHorizon(e.Twin.Face, e.Twin, point)...)
		return edges
	}
	if face.Visited {
		fmt.Println("visited")
		return nil
	}

	if !face.canSee(point) {
		fmt.Printf("add %p\n", edge.Twin)
		return []*Edge{edge.Twin}
	}

	fmt.Println("setting visited")
	face.Visited = true

	e := edge.Next
	fmt.Print("non-nil 1 ")
	edges := findHorizon(e.Twin.Face, e.Twin, point)
	fmt.Println(edges)
	e = e.Next
	fmt.Print("non-nil 2 ")
	edges = append(edges, findHorizon(e.Twin.Face, e.Twin, point)...)
	fmt.Println(edges)

	return edges
}

// NextConflict returns the index of the face and conflict of the conflict with
// the highest distance from it's associated plane. or -1,-1 if nothing else is
// left
func NextConflict(faces []*Face) (int, int) {
	var maxDist float32
	iface, iconflict := -1, -1

	for n, face := range faces {
		for m := range face.Conflicts {
			if face.Conflicts[m].Distance > maxDist {
				maxDist = face.Conflicts[m].Distance
				iface = n
				iconflict = m
			}
		}
	}

	return iface, iconflict
}

// FindExtremums returns the 6 indices and 6 vec3 of the extremums for each axis
// fomatted [minx, miny, minz, maxx, maxy, maxz]
func FindExtremums(points []glm.Vec3) (extremumIndices [6]int, extremums [6]glm.Vec3) {
	extremums = [6]glm.Vec3{
		{math.MaxFloat32, 0, 0}, {0, math.MaxFloat32, 0}, {0, 0, math.MaxFloat32},
		{-math.MaxFloat32, 0, 0}, {0, -math.MaxFloat32, 0}, {0, 0, -math.MaxFloat32},
	}
	for i := range points {
		for n := 0; n < 3; n++ {
			if extremums[n][n] > points[i][n] {
				extremums[n] = points[i]
				extremumIndices[n] = i
			}
			if extremums[n][n] > points[i][n] {
				extremums[n] = points[i]
				extremumIndices[n] = i
			}
			if extremums[n][n] > points[i][n] {
				extremums[n] = points[i]
				extremumIndices[n] = i
			}

			if extremums[3+n][n] < points[i][n] {
				extremums[3+n] = points[i]
				extremumIndices[3+n] = i
			}
			if extremums[3+n][n] < points[i][n] {
				extremums[3+n] = points[i]
				extremumIndices[3+n] = i
			}
			if extremums[3+n][n] < points[i][n] {
				extremums[3+n] = points[i]
				extremumIndices[3+n] = i
			}
		}
	}
	return
}

// CalculateEpsilon calculates the epsilon the algorithm should use given the
// extremums of the point cloud [minx, miny, minz, maxx, maxy, maxz]
func CalculateEpsilon(extremums [6]glm.Vec3) float32 {
	var maxima float32
	for n := 0; n < 3; n++ {
		maxima += math.Max(math.Abs(extremums[n][n]), math.Abs(extremums[3+n][n]))
	}
	return epsilonbase * maxima * 3
}

// BuildInitialTetrahedron builds the initial tetrahedron from the given 4 indices
func BuildInitialTetrahedron(a, b, c, d int, points []glm.Vec3) []*Face {
	ab := points[b].Sub(&points[a])
	ac := points[c].Sub(&points[a])
	cd := points[d].Sub(&points[c])
	ca := points[a].Sub(&points[c])
	ba := points[a].Sub(&points[b])
	bd := points[d].Sub(&points[b])
	dc := points[c].Sub(&points[d])
	db := points[b].Sub(&points[d])

	f0 := &Face{Vertices: [3]int{a, b, c}, Normal: ac.Cross(&ab), Point: points[a]}
	f1 := &Face{Vertices: [3]int{c, d, a}, Normal: ca.Cross(&cd), Point: points[c]}
	f2 := &Face{Vertices: [3]int{b, a, d}, Normal: bd.Cross(&ba), Point: points[b]}
	f3 := &Face{Vertices: [3]int{d, c, b}, Normal: db.Cross(&dc), Point: points[d]}

	f0.Faces = [3]*Face{f2, f3, f1}
	f1.Faces = [3]*Face{f3, f2, f0}
	f2.Faces = [3]*Face{f0, f1, f3}
	f3.Faces = [3]*Face{f1, f0, f2}

	// edges of f0
	e00 := &Edge{Tail: a, Face: f0}
	e01 := &Edge{Tail: b, Face: f0}
	e02 := &Edge{Tail: c, Face: f0}

	// edges of f1
	e10 := &Edge{Tail: c, Face: f1}
	e11 := &Edge{Tail: d, Face: f1}
	e12 := &Edge{Tail: a, Face: f1}

	// edges of f2
	e20 := &Edge{Tail: b, Face: f2}
	e21 := &Edge{Tail: a, Face: f2}
	e22 := &Edge{Tail: d, Face: f2}

	// edges of f3
	e30 := &Edge{Tail: d, Face: f3}
	e31 := &Edge{Tail: c, Face: f3}
	e32 := &Edge{Tail: b, Face: f3}

	// Connect the faces to the edges
	f0.Edges = [3]*Edge{e00, e01, e02}
	f1.Edges = [3]*Edge{e10, e11, e12}
	f2.Edges = [3]*Edge{e20, e21, e22}
	f3.Edges = [3]*Edge{e30, e31, e32}

	//Setup twin edges
	e00.Twin, e20.Twin = e20, e00
	e01.Twin, e31.Twin = e31, e01
	e02.Twin, e12.Twin = e12, e02
	e10.Twin, e30.Twin = e30, e10
	e11.Twin, e21.Twin = e21, e11
	e22.Twin, e32.Twin = e32, e22

	// Circular connect the edges
	// e0*
	e00.Next, e00.Prev = e01, e02
	e01.Next, e01.Prev = e02, e00
	e02.Next, e02.Prev = e00, e01

	// e1*
	e10.Next, e10.Prev = e11, e12
	e11.Next, e11.Prev = e12, e10
	e12.Next, e12.Prev = e10, e11

	// e2*
	e20.Next, e20.Prev = e21, e22
	e21.Next, e21.Prev = e22, e20
	e22.Next, e22.Prev = e20, e21

	// e3*
	e30.Next, e30.Prev = e31, e32
	e31.Next, e31.Prev = e32, e30
	e32.Next, e32.Prev = e30, e31

	return []*Face{f0, f1, f2, f3}
}
