package geo

import (
	"github.com/engoengine/glm"
)

// Rect is a rectangle in 3D space, they are a simpler version of OBBs.
type Rect struct {
	// the center of the rectangle
	Center glm.Vec3
	// the orientation of the rectangle in space.
	Orientation [2]glm.Vec3
	// the half extends of the rectangle.
	HalfExtend glm.Vec2
}

// ClosestPointRectPoint returns the point on the rectangle closest to p
func ClosestPointRectPoint(r *Rect, p *glm.Vec3) glm.Vec3 {
	d := p.Sub(&r.Center)
	closestPoint := r.Center

	// Start result at center of box; make steps from there

	// For each axis...
	for i := 0; i < len(r.Orientation); i++ {
		// ...project d onto that axis and get the distance along the axis of d
		// from the box center
		dist := d.Dot(&r.Orientation[i])

		// If distance farther than the box extents, clamp to the box
		if dist > r.HalfExtend[i] {
			dist = r.HalfExtend[i]
		} else if dist < -r.HalfExtend[i] {
			dist = -r.HalfExtend[i]
		}

		closestPoint.AddScaledVec(dist, &r.Orientation[i])
	}
	return closestPoint
}

// SqDistRectPoint returns the square distance of p to the rectangle.
func SqDistRectPoint(r *Rect, p *glm.Vec3) float32 {
	v := p.Sub(&r.Center)

	var sqDist float32

	for i := 0; i < len(r.Orientation); i++ {
		// Project vector from box center to p on each axis, getting the
		// distance of p along that axis, and count any excess distance
		// outside box extents.
		var excess float32
		d := v.Dot(&r.Orientation[i])

		if d < -r.HalfExtend[i] {
			excess = d + r.HalfExtend[i]
		} else if d > r.HalfExtend[i] {
			excess = d - r.HalfExtend[i]
		}
		sqDist += excess * excess
	}
	return sqDist
}
