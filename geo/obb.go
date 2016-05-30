package geo

import (
	"github.com/engoengine/glm"
	"github.com/luxengine/math"
)

// OBB is a Oriented Bounding Box.
type OBB struct {
	// The center of the OBB
	Center glm.Vec3
	// The orientation of the OBB, these need to be orthonormal.
	Orientation [3]glm.Vec3
	// The half extends of the OBB.
	HalfExtend glm.Vec3
}

// ClosestPointOBBPoint returns the point in or on the OBB closest to p
func ClosestPointOBBPoint(o *OBB, p *glm.Vec3) glm.Vec3 {
	var closestPoint glm.Vec3

	d := p.Sub(&o.Center)

	// Start result at center of box; make steps from there

	// For each OBB axis...
	for i := 0; i < len(o.HalfExtend); i++ {
		// ...project d onto that axis and get the distance along the axis of d
		// from the box center
		dist := d.Dot(&o.Orientation[i])

		// If distance farther than the box extents, clamp to the box
		if dist > o.HalfExtend[i] {
			dist = o.HalfExtend[i]
		} else if dist < -o.HalfExtend[i] {
			dist = -o.HalfExtend[i]
		}

		closestPoint.AddScaledVec(dist, &o.Orientation[i])
	}
	return closestPoint
}

// SqDistOBBPoint returns the square distance of p to the OBB.
func SqDistOBBPoint(o *OBB, p *glm.Vec3) float32 {
	v := p.Sub(&o.Center)

	var sqDist float32

	for i := 0; i < len(o.Orientation); i++ {
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

// TestOBBOBB returns true if these OBB overlap.
func TestOBBOBB(a, b *OBB) bool {
	// TODO(hydroflame): find a good value for that said epsilon
	const (
		epsilon = 0.0001
	)

	var R glm.Mat3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			R[j*3+i] = a.Orientation[i].Dot(&b.Orientation[i])
		}
	}

	// Compute translation vector t
	t := b.Center.Sub(&a.Center)

	// Bring translation into a's coordinate frame
	t = glm.Vec3{t.Dot(&a.Orientation[0]), t.Dot(&a.Orientation[1]), t.Dot(&a.Orientation[2])}

	var AbsR glm.Mat3
	// Compute common subexpressions. Add in an epsilon term to counteract
	// arithmetic errors when two edges are parallel and their cross product is
	// (near) zero.
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			AbsR[j*3+i] = math.Abs(R[j*3+i]) + epsilon
		}
	}

	var ra, rb float32
	// Test axes L = A0, L = A1, L = A2
	for i := 0; i < 3; i++ {
		ra = a.HalfExtend[i]
		rb = b.HalfExtend[0]*AbsR[0*3+i] + b.HalfExtend[1]*AbsR[1*3+i] + b.HalfExtend[2]*AbsR[2*3+i]
		if math.Abs(t[i]) > ra+rb {
			return false
		}
	}

	// Test axes L = B0, L = B1, L = B2
	for i := 0; i < 3; i++ {
		ra = a.HalfExtend[0]*AbsR[i*3+0] + a.HalfExtend[1]*AbsR[i*3+1] + a.HalfExtend[2]*AbsR[i*3+2]
		rb = b.HalfExtend[i]
		if math.Abs(t[0]*R[i*3+0]+t[1]*R[i*3+1]+t[2]*R[i*3+2]) > ra+rb {
			return false
		}
	}

	// Test axis L = A0 x B0
	ra = a.HalfExtend[1]*AbsR[2] + a.HalfExtend[2]*AbsR[1]
	rb = b.HalfExtend[1]*AbsR[6] + b.HalfExtend[2]*AbsR[3]
	if math.Abs(t[2]*R[1]-t[1]*R[2]) > ra+rb {
		return false
	}

	// Test axis L = A0 x B1
	ra = a.HalfExtend[1]*AbsR[5] + a.HalfExtend[2]*AbsR[4]
	rb = b.HalfExtend[0]*AbsR[6] + b.HalfExtend[2]*AbsR[0]
	if math.Abs(t[2]*R[4]-t[1]*R[5]) > ra+rb {
		return false
	}

	// Test axis L = A0 x B2
	ra = a.HalfExtend[1]*AbsR[8] + a.HalfExtend[2]*AbsR[7]
	rb = b.HalfExtend[0]*AbsR[3] + b.HalfExtend[1]*AbsR[0]
	if math.Abs(t[2]*R[7]-t[1]*R[8]) > ra+rb {
		return false
	}

	// Test axis L = A1 x B0
	ra = a.HalfExtend[0]*AbsR[2] + a.HalfExtend[2]*AbsR[0]
	rb = b.HalfExtend[1]*AbsR[7] + b.HalfExtend[2]*AbsR[4]
	if math.Abs(t[0]*R[2]-t[2]*R[0]) > ra+rb {
		return false
	}

	// Test axis L = A1 x B1
	ra = a.HalfExtend[0]*AbsR[5] + a.HalfExtend[2]*AbsR[3]
	rb = b.HalfExtend[0]*AbsR[7] + b.HalfExtend[2]*AbsR[1]
	if math.Abs(t[0]*R[5]-t[2]*R[3]) > ra+rb {
		return false
	}

	// Test axis L = A1 x B2
	ra = a.HalfExtend[0]*AbsR[8] + a.HalfExtend[2]*AbsR[6]
	rb = b.HalfExtend[0]*AbsR[4] + b.HalfExtend[1]*AbsR[1]
	if math.Abs(t[0]*R[8]-t[2]*R[6]) > ra+rb {
		return false
	}

	// Test axis L = A2 x B0
	ra = a.HalfExtend[0]*AbsR[1] + a.HalfExtend[1]*AbsR[0]
	rb = b.HalfExtend[1]*AbsR[8] + b.HalfExtend[2]*AbsR[5]
	if math.Abs(t[1]*R[0]-t[0]*R[1]) > ra+rb {
		return false
	}

	// Test axis L = A2 x B1
	ra = a.HalfExtend[0]*AbsR[4] + a.HalfExtend[1]*AbsR[3]
	rb = b.HalfExtend[0]*AbsR[8] + b.HalfExtend[2]*AbsR[2]
	if math.Abs(t[1]*R[3]-t[0]*R[4]) > ra+rb {
		return false
	}

	// Test axis L = A2 x B2
	ra = a.HalfExtend[0]*AbsR[7] + a.HalfExtend[1]*AbsR[6]
	rb = b.HalfExtend[0]*AbsR[5] + b.HalfExtend[1]*AbsR[2]
	if math.Abs(t[1]*R[6]-t[0]*R[7]) > ra+rb {
		return false
	}

	return true
}
