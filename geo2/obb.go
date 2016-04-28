package geo2

import (
	"github.com/luxengine/glm"
)

// OBB is a Oriented Bounding Box.
type OBB struct {
	Center      glm.Vec2
	Orientation [2]glm.Vec2
	HalfExtend  glm.Vec2
}

// ClosestPointOBBPoint returns the point in or on the OBB closest to p
func ClosestPointOBBPoint(a *OBB, p *glm.Vec2) glm.Vec2 {
	var closestPoint glm.Vec2

	d := p.Sub(&a.Center)

	// Start result at center of box; make steps from there

	// For each OBB axis...
	for i := 0; i < len(a.HalfExtend); i++ {
		// ...project d onto that axis and get the distance along the axis of d
		// from the box center
		dist := d.Dot(&a.Orientation[i])

		// If distance farther than the box extents, clamp to the box
		if dist > a.HalfExtend[i] {
			dist = a.HalfExtend[i]
		} else if dist < -a.HalfExtend[i] {
			dist = -a.HalfExtend[i]
		}

		closestPoint.AddScaledVec(dist, &a.Orientation[i])
	}
	return closestPoint
}

// SqDistOBBPoint returns the square distance of 'p' to the OBB
func SqDistOBBPoint(o *OBB, p *glm.Vec2) float32 {
	v := p.Sub(&o.Center)

	var sqDist float32

	for i := 0; i < len(o.Center); i++ {
		// Project vector from box center to 'p' on each axis, getting the
		// distance of 'p' along that axis, and count any excess distance
		// outside box extents
		var excess float32
		d := v.Dot(&o.Orientation[i])

		if d < -o.HalfExtend[i] {
			excess = d + o.HalfExtend[i]
		} else if d > o.HalfExtend[i] {
			excess = d - o.HalfExtend[i]
		}
		sqDist += excess * excess
	}
	return sqDist
}
