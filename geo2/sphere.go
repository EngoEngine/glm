package geo2

import (
	"github.com/luxengine/glm"
	"github.com/luxengine/math"
)

// Sphere is a bounding volume for spheres.
type Sphere struct {
	Center          glm.Vec2
	Radius, Radius2 float32
}

// TestSphereSphere return true if the spheres overlap.
func TestSphereSphere(a, b *Sphere) bool {
	d := b.Center.Sub(&a.Center)
	l2 := d.Len2()
	r := a.Radius + b.Radius
	return l2 <= r*r
}

// AABBFromSphere returns the AABB bounding this sphere.
//
// NOTE: If you need to use this function you better start questioning the
// algorithm you're implementing as the sphere is both faster and bounds the
// underlying object better.
func AABBFromSphere(s *Sphere) AABB {
	return AABB{
		Center:     s.Center,
		HalfExtend: glm.Vec2{s.Radius, s.Radius},
	}
}

// MergePoint updates the bounding sphere to encompass v if needed.
func (s *Sphere) MergePoint(v *glm.Vec2) {
	// Compute squared distance between point and sphere center
	d := v.Sub(&s.Center)
	dist2 := d.Len2()
	// Only update s if point p is outside it
	if dist2 > s.Radius*s.Radius {
		dist := math.Sqrt(dist2)
		newRadius := (s.Radius + dist) * 0.5
		k := (newRadius - s.Radius) / dist
		s.Radius = newRadius
		s.Center.AddScaledVec(k, &d)
		s.Radius2 = s.Radius * s.Radius
	}
}

// RitterSphere is an algorithm to find the bounding sphere of a set of points.
func RitterSphere(points []glm.Vec2) Sphere {
	// Get sphere encompassing two approximately most distant points.
	s := sphereFromDistantPoints(points)

	// Grow sphere to include all points.
	for i := range points {
		s.MergePoint(&points[i])
	}
	return s
}

// sphereFromDistantPoints generates a bounding sphere that wraps MOST of the
// points. Use RitterSphere to have certainty that everythings wrapped up.
func sphereFromDistantPoints(points []glm.Vec2) Sphere {
	var s Sphere
	// Find the most separated point pair defining the encompassing AABB
	min, max := MostSeparatePointsOnAABB(points)
	// Set up sphere to just encompass these two points
	s.Center.AddOf(&points[min], &points[max])
	s.Center.MulWith(0.5)

	v := points[max].Sub(&s.Center)
	s.Radius2 = v.Len2()
	s.Radius = math.Sqrt(s.Radius2)
	return s
}
