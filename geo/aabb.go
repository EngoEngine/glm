package geo

import (
	"github.com/EngoEngine/math"
	"github.com/engoengine/glm"
)

// AABB is an axis-aligned bounding box
type AABB struct {
	// Center represents the center of the bounding box.
	Center glm.Vec3

	// HalfExtend represents the 3 half extends of the bounding box.
	HalfExtend glm.Vec3
}

// TestAABBAABB returns true if these AABB overlap.
func TestAABBAABB(a, b *AABB) bool {
	if math.Abs(a.Center[0]-b.Center[0]) > a.HalfExtend[0]+b.HalfExtend[0] ||
		math.Abs(a.Center[1]-b.Center[1]) > a.HalfExtend[1]+b.HalfExtend[1] ||
		math.Abs(a.Center[2]-b.Center[2]) > a.HalfExtend[2]+b.HalfExtend[2] {
		return false
	}
	return true
}

// UpdateAABB computes an enclosing AABB base transformed by t and puts the
// result in fill. base and fill must not be the same.
func UpdateAABB(base, fill *AABB, t *glm.Mat3x4) {
	for i := 0; i < 3; i++ {
		fill.Center[i] = t[i+9]
		fill.HalfExtend[i] = 0
		for j := 0; j < 3; j++ {
			fill.Center[i] += t[j*3+i] * base.Center[j]
			fill.HalfExtend[i] += math.Abs(t[j*3+i]) * base.HalfExtend[j]
		}
	}
}

// ClosestPointAABBPoint returns the point in or on the AABB closest to p.
func ClosestPointAABBPoint(a *AABB, p *glm.Vec3) glm.Vec3 {
	return glm.Vec3{
		math.Clamp(p[0], a.Center[0]-a.HalfExtend[0], a.Center[0]+a.HalfExtend[0]),
		math.Clamp(p[1], a.Center[1]-a.HalfExtend[1], a.Center[1]+a.HalfExtend[1]),
		math.Clamp(p[2], a.Center[2]-a.HalfExtend[2], a.Center[2]+a.HalfExtend[2]),
	}
}

// SqDistAABBPoint returns the square distance of p to the AABB.
func SqDistAABBPoint(a *AABB, p *glm.Vec3) float32 {
	var sqDist float32

	// For each axis count any excess distance outside box extents
	v := p[0]
	min := a.Center[0] - a.HalfExtend[0]
	max := a.Center[0] + a.HalfExtend[0]
	if v < min {
		sqDist += (min - v) * (min - v)
	}
	if v > max {
		sqDist += (v - max) * (v - max)
	}

	v = p[1]
	min = a.Center[1] - a.HalfExtend[1]
	max = a.Center[1] + a.HalfExtend[1]
	if v < min {
		sqDist += (min - v) * (min - v)
	}
	if v > max {
		sqDist += (v - max) * (v - max)
	}

	v = p[2]
	min = a.Center[2] - a.HalfExtend[2]
	max = a.Center[2] + a.HalfExtend[2]
	if v < min {
		sqDist += (min - v) * (min - v)
	}
	if v > max {
		sqDist += (v - max) * (v - max)
	}

	return sqDist
}
