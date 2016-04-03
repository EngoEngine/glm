package geo

import (
	"github.com/luxengine/glm"
)

// Capsule3 is a capsule in 3D.
type Capsule3 struct {
	A, B   glm.Vec3
	Radius float32
}

// TestCapsule3Capsule3 returns true if these Capsules overlap.
func TestCapsule3Capsule3(a, b *Capsule3) bool {
	_, _, u, _, _ := ClosestPointSegmentSegment(&a.A, &a.B, &b.A, &b.B)
	// If squared distance is smaller than squared sum of radii, they collide
	radius := a.Radius + b.Radius
	return u <= radius*radius
}

// TestCapsule3Sphere3 returns true if the capsule and the sphere overlap.
func TestCapsule3Sphere3(c *Capsule3, s *Sphere3) bool {
	dist2 := SqDistPointSegment3(&c.A, &c.B, &s.Center)
	radius := s.Radius + c.Radius
	return dist2 <= radius*radius
}
