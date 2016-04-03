package geo

import (
	"github.com/luxengine/glm"
)

// Rect3 is a rectangle in 3D space, it has no 2D equivalent (a line ?)
type Rect3 struct {
	Center      glm.Vec3
	Orientation [2]glm.Vec3
	Radius      glm.Vec2
}

// ClosestPointRect3Point returns the point on the rectangle closest to 'p'
func ClosestPointRect3Point(r *Rect3, p *glm.Vec3) glm.Vec3 {
	d := p.Sub(&r.Center)
	closestPoint := r.Center

	// Start result at center of box; make steps from there

	// For each axis...
	for i := 0; i < len(r.Orientation); i++ {
		// ...project d onto that axis and get the distance along the axis of d
		// from the box center
		dist := d.Dot(&r.Orientation[i])

		// If distance farther than the box extents, clamp to the box
		if dist > r.Radius[i] {
			dist = r.Radius[i]
		} else if dist < -r.Radius[i] {
			dist = -r.Radius[i]
		}

		closestPoint.AddScaledVec(dist, &r.Orientation[i])
	}
	return closestPoint
}

// SqDistRect3Point returns the square distance of 'p' to the OBB3
func SqDistRect3Point(r *Rect3, p *glm.Vec3) float32 {
	v := p.Sub(&r.Center)

	var sqDist float32

	for i := 0; i < len(r.Orientation); i++ {
		// Project vector from box center to 'p' on each axis, getting the
		// distance of 'p' along that axis, and count any excess distance
		// outside box extents
		var excess float32
		d := v.Dot(&r.Orientation[i])

		if d < -r.Radius[i] {
			excess = d + r.Radius[i]
		} else if d > r.Radius[i] {
			excess = d - r.Radius[i]
		}
		sqDist += excess * excess
	}
	return sqDist
}
