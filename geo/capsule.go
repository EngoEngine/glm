package geo

import (
	"github.com/engoengine/glm"
)

// Capsule is a cylinder with round end or can be used as a swept sphere.
type Capsule struct {
	A, B   glm.Vec3
	Radius float32
}

// TestCapsuleCapsule returns true if these Capsules overlap.
func TestCapsuleCapsule(a, b *Capsule) bool {
	_, _, u, _, _ := ClosestPointSegmentSegment(&a.A, &a.B, &b.A, &b.B)
	// If squared distance is smaller than squared sum of radii, they collide
	r := a.Radius + b.Radius
	return u <= r*r
}

// TestCapsuleSphere returns true if the capsule and the sphere overlap.
func TestCapsuleSphere(c *Capsule, s *Sphere) bool {
	dist2 := SqDistPointSegment(&c.A, &c.B, &s.Center)
	r := s.Radius + c.Radius
	return dist2 <= r*r
}
