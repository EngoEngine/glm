package geo

import (
	"github.com/luxengine/glm"
)

// OBB2 is a Oriented Bounding Box for 2d
type OBB2 struct {
	Center      glm.Vec2
	Orientation [2]glm.Vec2
	Radius      glm.Vec2
}

// ClosestPointOBB2Point returns the point in or on the OBB closest to 'p'
func ClosestPointOBB2Point(a *OBB2, p *glm.Vec2) glm.Vec2 {
	var closestPoint glm.Vec2

	d := p.Sub(&a.Center)

	// Start result at center of box; make steps from there

	// For each OBB axis...
	for i := 0; i < len(a.Radius); i++ {
		// ...project d onto that axis and get the distance along the axis of d
		// from the box center
		dist := d.Dot(&a.Orientation[i])

		// If distance farther than the box extents, clamp to the box
		if dist > a.Radius[i] {
			dist = a.Radius[i]
		} else if dist < -a.Radius[i] {
			dist = -a.Radius[i]
		}

		closestPoint.AddScaledVec(dist, &a.Orientation[i])
	}
	return closestPoint
}

// SqDistOBB2Point returns the square distance of 'p' to the OBB2
func SqDistOBB2Point(a *OBB2, p *glm.Vec2) float32 {
	v := p.Sub(&a.Center)

	var sqDist float32

	for i := 0; i < len(a.Center); i++ {
		// Project vector from box center to 'p' on each axis, getting the
		// distance of 'p' along that axis, and count any excess distance
		// outside box extents
		var excess float32
		d := v.Dot(&a.Orientation[i])

		if d < -a.Radius[i] {
			excess = d + a.Radius[i]
		} else if d > a.Radius[i] {
			excess = d - a.Radius[i]
		}
		sqDist += excess * excess
	}
	return sqDist
}
