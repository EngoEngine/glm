package geo

import (
	"github.com/luxengine/glm"
)

// Capsule2 is a capsule in 2D.
type Capsule2 struct {
	A, B   glm.Vec2
	Radius float32
}

// TestCapsule2Sphere2 returns true if the capsule and the sphere overlap.
func TestCapsule2Sphere2(c *Capsule2, s *Sphere2) bool {
	dist2 := SqDistPointSegment2(&c.A, &c.B, &s.Center)
	radius := s.Radius + c.Radius
	return dist2 <= radius*radius
}
