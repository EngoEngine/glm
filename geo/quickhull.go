package geo

import (
	"github.com/luxengine/glm"
	"github.com/luxengine/math"
)

// Quickhull2 returns a 2D Convex hull of the given points.
func Quickhull2(points []glm.Vec2) []glm.Vec2 {

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

	return hull
}

/*
// Quickhull3 returns a 3D Convex hull of the given points.
func Quickhull3(points []glm.Vec3) {

}*/
