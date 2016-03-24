package geo

import (
	"github.com/luxengine/glm"
	"github.com/luxengine/glm/geo/internal/quickhull2"
	"github.com/luxengine/math"
	"reflect"
	"unsafe"
)

// ConvexHull2 is a convex hull in 2D
type ConvexHull2 struct {
	memory      *byte // pointer to the memory backing up the convex hull
	vertices    []quickhull2.Vertex
	edges       []quickhull2.Edge
	bestSupport [3]int
}

// Support returns the support of this convex hull in the given direcion.
func (c *ConvexHull2) Support(dir *glm.Vec2) glm.Vec2 {
	// find the index of the support we should start using at first. We could
	// use the dot product for quickhull2.SupportDirection but this is just a
	// couple of load, compare, and jumps so it's way faster
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
	return glm.Vec2{}
	/*dot := c.vertices[i].Position.Dot(dir)

	var useNext bool
	if d := c.vertices[c.edges[c.vertices[i].Next].Vertices[1]].Position.Dot(dir); d > dot {
		useNext = true
		i = c.edges[c.vertices[i].Next].Vertices[1]
	}

	for {
		var j int
		if useNext {
			j = c.edges[c.vertices[i].Next].Vertices[1]

		} else {
			j = c.edges[c.vertices[i].Prev].Vertices[0]
		}
		if d := c.vertices[j].Position.Dot(dir); d > dot {
			i = j
			dot = d
		} else {
			return c.vertices[i].Position
		}
	}*/
}

// SupportSlow is like Support but the implementation is much simpler and
// probably correct, however it's kinda slow.
func (c *ConvexHull2) SupportSlow(dir *glm.Vec2) glm.Vec2 {
	var bestDot float32
	var bestIndex int
	for n := range c.vertices {
		if d := c.vertices[n].Position.Dot(dir); d > bestDot {
			bestDot = d
			bestIndex = n
		}
	}
	return c.vertices[bestIndex].Position
}

func (c *ConvexHull2) Intersects(o *ConvexHull2) bool { return false }

// Quickhull2 returns the 2D convex hull of the given points.
func Quickhull2(points []glm.Vec2) ConvexHull2 {

	// Make a copy of the vertices so we know which one are left to check
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

func convexHull2FromPoints(points []glm.Vec2) ConvexHull2 {
	vertexArraySize := len(points) * int(unsafe.Sizeof(quickhull2.Vertex{}))
	edgeArraySize := len(points) * int(unsafe.Sizeof(quickhull2.Edge{}))
	vecArraySize := len(points) * int(unsafe.Sizeof(glm.Vec2{}))
	backup := make([]byte, vertexArraySize+edgeArraySize+vecArraySize)

	vertexSliceHeader := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&backup[0])),
		Len:  len(points),
		Cap:  len(points),
	}

	edgeSliceHeader := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&backup[vertexArraySize])),
		Len:  len(points),
		Cap:  len(points),
	}

	hull := ConvexHull2{
		memory:   &backup[0],
		vertices: *(*[]quickhull2.Vertex)(unsafe.Pointer(&vertexSliceHeader)),
		edges:    *(*[]quickhull2.Edge)(unsafe.Pointer(&edgeSliceHeader)),
	}

	var bestSupport [3]int
	var bestSupportDot [3]float32

	for n := range points {
		hull.vertices[n] = quickhull2.Vertex{
			Position: points[n],
			Prev:     (n - 1 + len(points)) % len(points),
			Next:     (n + 1) % len(points),
		}
		hull.edges[n] = quickhull2.Edge{
			Vertices: [2]int{n, (n + 1) % len(points)},
			PrevEdge: (n - 1 + len(points)) % len(points),
			NextEdge: (n + 1) % len(points),
		}

	}

	var avg glm.Vec2
	for n := range hull.vertices {
		avg.AddWith(&hull.vertices[n].Position)
	}
	avg.MulWith(1 / float32(len(hull.vertices)))

	for n := range hull.vertices {
		for m := range quickhull2.SupportDirection {
			tmp := hull.vertices[n].Position.Sub(&avg)
			if d := quickhull2.SupportDirection[m].Dot(&tmp); d > bestSupportDot[m] {
				bestSupport[m] = n
				bestSupportDot[m] = d
			}
		}
	}

	hull.bestSupport = bestSupport
	return hull
}

/*
// Quickhull3 returns a 3D Convex hull of the given points.
func Quickhull3(points []glm.Vec3) {

}*/
