package geo

import (
	"github.com/luxengine/glm"
	"github.com/luxengine/math"
)

// OBB3 is a Oriented Bounding Box for 3d
type OBB3 struct {
	Center      glm.Vec3
	Orientation [3]glm.Vec3
	Radius      glm.Vec3
}

// ClosestPoint returns the point in or on the OBB closest to 'p'
func (a *OBB3) ClosestPoint(p *glm.Vec3) glm.Vec3 {
	var closestPoint glm.Vec3

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

// SqDistOfPoint returns the square distance of 'p' to the OBB3
func (a *OBB3) SqDistOfPoint(p *glm.Vec3) float32 {
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

// Intersects returns true if these OBB overlap.
func (a *OBB3) Intersects(b *OBB3) bool {
	// TODO find a good value for that said epsilon
	const (
		epsilon = 0.000
	)
	var ra, rb float32
	var R, AbsR glm.Mat3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			R[j*3+i] = a.Orientation[i].Dot(&b.Orientation[i])
		}
	}

	// Compute translation vector t
	t := b.Center.Sub(&a.Center)

	// Bring translation into a's coordinate frame
	t = glm.Vec3{t.Dot(&a.Orientation[0]), t.Dot(&a.Orientation[1]), t.Dot(&a.Orientation[2])}

	// Compute common subexpressions. Add in an epsilon term to counteract
	// arithmetic errors when two edges are parallel and their cross product is
	// (near) zero.
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			AbsR[j*3+i] = math.Abs(R[j*3+i]) + epsilon
		}
	}

	// Test axes L = A0, L = A1, L = A2
	for i := 0; i < 3; i++ {
		ra = a.Radius[i]
		rb = b.Radius[0]*AbsR[0*3+i] + b.Radius[1]*AbsR[1*3+i] + b.Radius[2]*AbsR[2*3+i]
		if math.Abs(t[i]) > ra+rb {
			return false
		}
	}

	// Test axes L = B0, L = B1, L = B2
	for i := 0; i < 3; i++ {
		ra = a.Radius[0]*AbsR[i*3+0] + a.Radius[1]*AbsR[i*3+1] + a.Radius[2]*AbsR[i*3+2]
		rb = b.Radius[i]
		if math.Abs(t[0]*R[i*3+0]+t[1]*R[i*3+1]+t[2]*R[i*3+2]) > ra+rb {
			return false
		}
	}

	// Test axis L = A0 x B0
	ra = a.Radius[1]*AbsR[2] + a.Radius[2]*AbsR[1]
	rb = b.Radius[1]*AbsR[6] + b.Radius[2]*AbsR[3]
	if math.Abs(t[2]*R[1]-t[1]*R[2]) > ra+rb {
		return false
	}

	// Test axis L = A0 x B1
	ra = a.Radius[1]*AbsR[5] + a.Radius[2]*AbsR[4]
	rb = b.Radius[0]*AbsR[6] + b.Radius[2]*AbsR[0]
	if math.Abs(t[2]*R[4]-t[1]*R[5]) > ra+rb {
		return false
	}

	// Test axis L = A0 x B2
	ra = a.Radius[1]*AbsR[8] + a.Radius[2]*AbsR[7]
	rb = b.Radius[0]*AbsR[3] + b.Radius[1]*AbsR[0]
	if math.Abs(t[2]*R[7]-t[1]*R[8]) > ra+rb {
		return false
	}

	// Test axis L = A1 x B0
	ra = a.Radius[0]*AbsR[2] + a.Radius[2]*AbsR[0]
	rb = b.Radius[1]*AbsR[7] + b.Radius[2]*AbsR[4]
	if math.Abs(t[0]*R[2]-t[2]*R[0]) > ra+rb {
		return false
	}

	// Test axis L = A1 x B1
	ra = a.Radius[0]*AbsR[5] + a.Radius[2]*AbsR[3]
	rb = b.Radius[0]*AbsR[7] + b.Radius[2]*AbsR[1]
	if math.Abs(t[0]*R[5]-t[2]*R[3]) > ra+rb {
		return false
	}

	// Test axis L = A1 x B2
	ra = a.Radius[0]*AbsR[8] + a.Radius[2]*AbsR[6]
	rb = b.Radius[0]*AbsR[4] + b.Radius[1]*AbsR[1]
	if math.Abs(t[0]*R[8]-t[2]*R[6]) > ra+rb {
		return false
	}

	// Test axis L = A2 x B0
	ra = a.Radius[0]*AbsR[1] + a.Radius[1]*AbsR[0]
	rb = b.Radius[1]*AbsR[8] + b.Radius[2]*AbsR[5]
	if math.Abs(t[1]*R[0]-t[0]*R[1]) > ra+rb {
		return false
	}

	// Test axis L = A2 x B1
	ra = a.Radius[0]*AbsR[4] + a.Radius[1]*AbsR[3]
	rb = b.Radius[0]*AbsR[8] + b.Radius[2]*AbsR[2]
	if math.Abs(t[1]*R[3]-t[0]*R[4]) > ra+rb {
		return false
	}

	// Test axis L = A2 x B2
	ra = a.Radius[0]*AbsR[7] + a.Radius[1]*AbsR[6]
	rb = b.Radius[0]*AbsR[5] + b.Radius[1]*AbsR[2]
	if math.Abs(t[1]*R[6]-t[0]*R[7]) > ra+rb {
		return false
	}

	return true
}
