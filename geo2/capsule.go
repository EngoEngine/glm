package geo2

import (
	"github.com/luxengine/glm"
)

// Capsule is a cylinder with round end or can be used as a swept sphere.
type Capsule struct {
	A, B   glm.Vec2
	Radius float32
}

// TestCapsuleCapsule returns true if the capsules overlap.
func TestCapsuleCapsule(a, b *Capsule) bool {
	// TODO(hydroflame): implement
	panic("not implemented")
}

// TestCapsuleSphere returns true if the capsule and the sphere overlap.
func TestCapsuleSphere(c *Capsule, s *Sphere) bool {
	dist2 := SqDistPointSegment(&c.A, &c.B, &s.Center)
	radius := s.Radius + c.Radius
	return dist2 <= radius*radius
}
