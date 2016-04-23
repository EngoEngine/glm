package geo2

import (
	"github.com/luxengine/glm"
	"github.com/luxengine/glm/geo2/internal/qhull"
	"github.com/luxengine/math"
	"reflect"
	"unsafe"
)

// ConvexHull is a shape that doesn't have any edge with an interior angle
// bigger then π radian.
type ConvexHull struct {
	// Vertices are the vertices of the convex hull.
	Vertices []qhull.Vertex

	// Edges are the edges of the convex hull.
	Edges []qhull.Edge

	// bestSupport are the indices in `Vertices` that points to the vertex with
	// the highest dot product with the matching vector in
	// internal/qhull.SupportDirection
	bestSupport [3]int

	// pointer to the memory backing up the convex hull
	// we need to keep this because go doesn't like keeping track of slices that
	// are products of unsafe casting. Linters will say this is unused but
	// whatever.
	mem *byte
}

// Support returns the support of this convex hull in the given direction.
func (c *ConvexHull) Support(dir *glm.Vec2) int {
	// TODO(hydroflame): finish implementation

	// find the index of the support we should start using at first. We could
	// use the dot product for quickhull2.SupportDirection but this is just a
	// couple of load, compare, and jumps so it's way faster.
	var i int
	if dir[0] > 0 {
		if dir[1] > 0 {
			if 2*dir[0] > dir[1] {
				i = 0
			} else {
				i = 1
			}
		} else {
			if 2*dir[0] > -dir[1] {
				i = 0
			} else {
				i = 2
			}
		}
	} else {
		if dir[1] > 0 {
			i = 1
		} else {
			i = 2
		}
	}

	return i

	/*dot := c.Vertices[i].Position.Dot(dir)

	var useNext bool
	if d := c.Vertices[c.Edges[c.Vertices[i].Next].Vertices[1]].Position.Dot(dir); d > dot {
		useNext = true
		i = c.Edges[c.Vertices[i].Next].Vertices[1]
	}

	for {
		var j int
		if useNext {
			j = c.Edges[c.Vertices[i].Next].Vertices[1]

		} else {
			j = c.Edges[c.Vertices[i].Prev].Vertices[0]
		}
		if d := c.Vertices[j].Position.Dot(dir); d > dot {
			i = j
			dot = d
		} else {
			return c.Vertices[i].Position
		}
	}*/
}

// SupportSlow is like Support but the implementation is much simpler
// and probably correct, however it's kinda slow. this is used as a reference
// when testing the fast support.
func (c *ConvexHull) SupportSlow(dir *glm.Vec2) glm.Vec2 {
	var bestDot float32
	var bestIndex int
	for n := range c.Vertices {
		if d := c.Vertices[n].Position.Dot(dir); d > bestDot {
			bestDot = d
			bestIndex = n
		}
	}
	return c.Vertices[bestIndex].Position
}

// TestConvexHullConvexHull return true if the 2 convex hulls intersect. Not
// implemented.
func TestConvexHullConvexHull(a, b *ConvexHull) bool {
	// TODO(hydroflame): implement!
	panic("not implemented")
}

// Quickhull returns the 2D convex hull of the given points using the quickhull
// algorithm.
func Quickhull(points []glm.Vec2) ConvexHull {
	// Make a copy of the Vertices so we know which one are left to check
	pts := make([]glm.Vec2, len(points), len(points))
	copy(pts, points)

	// Find the minimum X and maximum X
	var imin, imax int
	var minx, maxx float32 = math.MaxFloat32, -math.MaxFloat32

	for n := range pts {
		if pts[n][0] < minx {
			minx = pts[n][0]
			imin = n
		}

		if pts[n][0] > maxx {
			maxx = pts[n][0]
			imax = n
		}
	}

	// start the hull with the min and max X
	hull := []glm.Vec2{pts[imin], pts[imax]}

	// remove these 2 points
	pts[imin] = pts[len(pts)-1]
	pts = pts[:len(pts)-1]

	if imax == len(pts) {
		imax = imin
	}

	pts[imax] = pts[len(pts)-1]
	pts = pts[:len(pts)-1]

	for n := 0; n < len(hull); {
		i := PointFarthestFromEdge(&hull[n], &hull[(n+1)%len(hull)], pts)

		if i != -1 {
			hull = append(hull, glm.Vec2{})
			copy(hull[n+2:], hull[n+1:])
			hull[n+1] = pts[i]

			pts[i] = pts[len(pts)-1]
			pts = pts[:len(pts)-1]
		} else {
			n++
		}
	}

	return convexHull2FromPoints(hull)
}

// convexHull2FromPoints packs the memory and builds the convex hull struct.
func convexHull2FromPoints(points []glm.Vec2) ConvexHull {
	vertexArraySize := len(points) * int(qhull.VertexSize)
	edgeArraySize := len(points) * int(qhull.EdgeSize)
	vecArraySize := len(points) * int(unsafe.Sizeof(glm.Vec2{}))

	// make a backing array that holds continuously all the memory of the hull
	backup := make([]byte, vertexArraySize+edgeArraySize+vecArraySize)

	// make the slice header for the vertices, this starts at the very begining
	// of the backing array.
	vertexSliceHeader := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&backup[0])),
		Len:  len(points),
		Cap:  len(points),
	}

	// the edge slice starts at the byte after the end of the memory used by the
	// vertices.
	edgeSliceHeader := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&backup[vertexArraySize])),
		Len:  len(points),
		Cap:  len(points),
	}

	hull := ConvexHull{
		mem:      &backup[0],
		Vertices: *(*[]qhull.Vertex)(unsafe.Pointer(&vertexSliceHeader)),
		Edges:    *(*[]qhull.Edge)(unsafe.Pointer(&edgeSliceHeader)),
	}

	// link all the edges/vertices togheter.
	for n := range points {
		hull.Vertices[n] = qhull.Vertex{
			Position: points[n],
			Prev:     (n - 1 + len(points)) % len(points),
			Next:     (n + 1) % len(points),
		}
		hull.Edges[n] = qhull.Edge{
			Vertices: [2]int{n, (n + 1) % len(points)},
			PrevEdge: (n - 1 + len(points)) % len(points),
			NextEdge: (n + 1) % len(points),
		}

	}

	// calculate an average point to use for the dot product.
	var avg glm.Vec2
	for n := range hull.Vertices {
		avg.AddWith(&hull.Vertices[n].Position)
	}
	avg.MulWith(1 / float32(len(hull.Vertices)))

	// find the indices of the best supports in each of the
	// qhull.SupportDirection
	var bestSupportDot [3]float32
	for n := range hull.Vertices {
		for m := range qhull.SupportDirection {
			tmp := hull.Vertices[n].Position.Sub(&avg)
			if d := qhull.SupportDirection[m].Dot(&tmp); d > bestSupportDot[m] {
				hull.bestSupport[m] = n
				bestSupportDot[m] = d
			}
		}
	}

	return hull
}
