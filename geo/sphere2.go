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

// Intersects return true if the spheres overlap.
func (a *Sphere2) Intersects(b *Sphere2) bool {
	d := b.Center.Sub(&a.Center)
	l2 := d.Len2()
	r := a.Radius + b.Radius
	return l2 <= r*r
}

// AABB2 returns the AABB bounding this sphere.
//
// NOTE: If you need to use this function you better start questioning the
// algorithm you're implementing as the sphere is both faster and bounds the
// underlying object better.
func (a *Sphere2) AABB2() AABB2 {
	return AABB2{
		Center: a.Center,
		Radius: glm.Vec2{a.Radius, a.Radius},
	}
}

// OfSphereAndPt updates the bounding sphere to encompass v if needed.
func (a *Sphere2) OfSphereAndPt(v *glm.Vec2) {
	// Compute squared distance between point and sphere center
	d := v.Sub(&a.Center)
	dist2 := d.Len2()
	// Only update s if point p is outside it
	if dist2 > a.Radius*a.Radius {
		dist := math.Sqrt(dist2)
		newRadius := (a.Radius + dist) * 0.5
		k := (newRadius - a.Radius) / dist
		a.Radius = newRadius
		a.Center.AddScaledVec(k, &d)
		a.Radius2 = a.Radius * a.Radius
	}
}

// Ritter is an algorithm to find the bounding sphere of a set of points
func (a *Sphere2) Ritter(points []glm.Vec2) {
	// Get sphere encompassing two approximately most distant points
	a.FromDistantPoints(points)

	// Grow sphere to include all points
	for i := range points {
		a.OfSphereAndPt(&points[i])
	}
}

// FromDistantPoints reshapes the bounding sphere to wrap all the points.
func (a *Sphere2) FromDistantPoints(points []glm.Vec2) {
	// Find the most separated point pair defining the encompassing AABB
	min, max := MostSeparatePointsOnAABB2(points)
	// Set up sphere to just encompass these two points
	a.Center.AddOf(&points[min], &points[max])
	a.Center.MulWith(0.5)

	v := points[max].Sub(&a.Center)
	a.Radius2 = v.Len2()
	a.Radius = math.Sqrt(a.Radius2)
}
