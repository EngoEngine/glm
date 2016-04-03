package geo

import (
	"github.com/luxengine/glm"
	"github.com/luxengine/math"
)

// Sphere2 is a bounding volume for spheres in 2D.
type Sphere2 struct {
	Center          glm.Vec2
	Radius, Radius2 float32
}

// TestSphere2Sphere2 return true if the spheres overlap.
func TestSphere2Sphere2(a, b *Sphere2) bool {
	d := b.Center.Sub(&a.Center)
	l2 := d.Len2()
	r := a.Radius + b.Radius
	return l2 <= r*r
}

// AABB2FromSphere2 returns the AABB bounding this sphere.
//
// NOTE: If you need to use this function you better start questioning the
// algorithm you're implementing as the sphere is both faster and bounds the
// underlying object better.
func AABB2FromSphere2(s *Sphere2) AABB2 {
	return AABB2{
		Center: s.Center,
		Radius: glm.Vec2{s.Radius, s.Radius},
	}
}

// MergeSphere2Point updates the bounding sphere to encompass v if needed.
func MergeSphere2Point(s *Sphere2, v *glm.Vec2) {
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

// RitterSphere is an algorithm to find the bounding sphere of a set of points
func RitterSphere(points []glm.Vec2) Sphere2 {
	// Get sphere encompassing two approximately most distant points
	s := Sphere2FromDistantPoints(points)

	// Grow sphere to include all points
	for i := range points {
		MergeSphere2Point(&s, &points[i])
	}
	return s
}

// Sphere2FromDistantPoints reshapes the bounding sphere to wrap all the points.
func Sphere2FromDistantPoints(points []glm.Vec2) Sphere2 {
	var s Sphere2
	// Find the most separated point pair defining the encompassing AABB
	min, max := MostSeparatePointsOnAABB2(points)
	// Set up sphere to just encompass these two points
	s.Center.AddOf(&points[min], &points[max])
	s.Center.MulWith(0.5)

	v := points[max].Sub(&s.Center)
	s.Radius2 = v.Len2()
	s.Radius = math.Sqrt(s.Radius2)
	return s
}
