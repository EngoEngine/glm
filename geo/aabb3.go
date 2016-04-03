package geo

import (
	"github.com/luxengine/glm"
	"github.com/luxengine/math"
)

// AABB3 is a 3D axis-aligned bounding box
type AABB3 struct {
	Center glm.Vec3
	Radius glm.Vec3
}

// TestAABB3AABB3 returns true if these AABB overlap.
func TestAABB3AABB3(a, b *AABB3) bool {
	if math.Abs(a.Center[0]-b.Center[0]) > (a.Radius[0] + b.Radius[0]) {
		return false
	}

	if math.Abs(a.Center[1]-b.Center[1]) > (a.Radius[1] + b.Radius[1]) {
		return false
	}

	if math.Abs(a.Center[2]-b.Center[2]) > (a.Radius[2] + b.Radius[2]) {
		return false
	}

	return true
}

// UpdateAABB3 computes an enclosing AABB to base transformed by t and puts the
// result in base.
func UpdateAABB3(base, fill *AABB3, t *glm.Mat3x4) {
	for i := 0; i < 3; i++ {
		fill.Center[i] = t[i+9]
		fill.Radius[i] = 0
		for j := 0; j < 3; j++ {
			fill.Center[i] += t[j*3+i] * base.Center[j]
			fill.Radius[i] += math.Abs(t[j*3+i]) * base.Radius[j]
		}
	}
}

// ClosestPointAABB3Point returns the point in or on the AABB3 closest to 'p'
func ClosestPointAABB3Point(a *AABB3, p *glm.Vec3) glm.Vec3 {
	return glm.Vec3{
		math.Clamp(p[0], a.Center[0]-a.Radius[0], a.Center[0]+a.Radius[0]),
		math.Clamp(p[1], a.Center[1]-a.Radius[1], a.Center[1]+a.Radius[1]),
		math.Clamp(p[2], a.Center[2]-a.Radius[2], a.Center[2]+a.Radius[2]),
	}
}

// SqDistAABB3Point returns the square distance of 'p' to the AABB3
func SqDistAABB3Point(a *AABB3, p *glm.Vec3) float32 {
	var sqDist float32

	// For each axis count any excess distance outside box extents
	v := p[0]
	min := a.Center[0] - a.Radius[0]
	max := a.Center[0] + a.Radius[0]
	if v < min {
		sqDist += (min - v) * (min - v)
	}
	if v > max {
		sqDist += (v - max) * (v - max)
	}

	v = p[1]
	min = a.Center[1] - a.Radius[1]
	max = a.Center[1] + a.Radius[1]
	if v < min {
		sqDist += (min - v) * (min - v)
	}
	if v > max {
		sqDist += (v - max) * (v - max)
	}

	v = p[2]
	min = a.Center[2] - a.Radius[2]
	max = a.Center[2] + a.Radius[2]
	if v < min {
		sqDist += (min - v) * (min - v)
	}
	if v > max {
		sqDist += (v - max) * (v - max)
	}

	return sqDist
}
