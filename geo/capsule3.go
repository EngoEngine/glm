package geo

import (
	"github.com/luxengine/glm"
)

// Capsule3 is a capsule in 3D.
type Capsule3 struct {
	A, B   glm.Vec3
	Radius float32
}

// Intersects returns true if these Capsules overlap.
func (c *Capsule3) Intersects(o *Capsule3) bool {
	_, _, u, _, _ := ClosestPointSegmentSegment(&c.A, &c.B, &o.A, &o.B)
	// If squared distance is smaller than squared sum of radii, they collide
	radius := c.Radius + o.Radius
	return u <= radius*radius
}

// IntersectsSphere returns true if the capsule and the sphere overlap.
func (c *Capsule3) IntersectsSphere(s *Sphere3) bool {
	dist2 := SqDistPointSegment3(&c.A, &c.B, &s.Center)
	radius := s.Radius + c.Radius
	return dist2 <= radius*radius
}
