package geo

import (
	"github.com/luxengine/glm"
	"github.com/luxengine/math"
)

// AABB2 is a 2D axis-aligned bounding box
type AABB2 struct {
	Center glm.Vec2
	Radius glm.Vec2
}

// Intersects returns true if these AABB overlap.
func (a *AABB2) Intersects(b *AABB2) bool {
	if math.Abs(a.Center.X()-b.Center.X()) > (a.Radius.X() + b.Radius.X()) {
		return false
	}

	if math.Abs(a.Center.Y()-b.Center.Y()) > (a.Radius.Y() + b.Radius.Y()) {
		return false
	}

	return true
}

// AABB3 is a 3D axis-aligned bounding box
type AABB3 struct {
	Center glm.Vec3
	Radius glm.Vec3
}

// Intersects returns true if these AABB overlap.
func (a *AABB3) Intersects(b *AABB3) bool {
	if math.Abs(a.Center.X()-b.Center.X()) > (a.Radius.X() + b.Radius.X()) {
		return false
	}

	if math.Abs(a.Center.Y()-b.Center.Y()) > (a.Radius.Y() + b.Radius.Y()) {
		return false
	}

	if math.Abs(a.Center.Z()-b.Center.Z()) > (a.Radius.Z() + b.Radius.Z()) {
		return false
	}

	return true
}

// BoundingSphere2 is a bounding volume for spheres in 2D.
type BoundingSphere2 struct {
	Center          glm.Vec2
	Radius, Radius2 float32
}

// Intersects return true if the spheres overlap.
func (a *BoundingSphere2) Intersects(b *BoundingSphere2) bool {
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
func (a *BoundingSphere2) AABB2() AABB2 {
	return AABB2{
		Center: a.Center,
		Radius: glm.Vec2{a.Radius, a.Radius},
	}
}

// BoundingSphere3 is a bounding volume for spheres in 3D.
type BoundingSphere3 struct {
	Center          glm.Vec3
	Radius, Radius2 float32
}

// Intersects return true if the spheres overlap.
func (a *BoundingSphere3) Intersects(b *BoundingSphere3) bool {
	d := b.Center.Sub(&a.Center)
	l2 := d.Len2()
	r := a.Radius + b.Radius
	return l2 <= r*r
}

// AABB3 returns the AABB bounding this sphere.
//
// NOTE: If you need to use this function you better start questioning the
// algorithm you're implementing as the sphere is both faster and bounds the
// underlying object better.
func (a *BoundingSphere3) AABB3() AABB3 {
	return AABB3{
		Center: a.Center,
		Radius: glm.Vec3{a.Radius, a.Radius, a.Radius},
	}
}
