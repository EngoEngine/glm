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

// Intersects returns true if these AABB overlap.
func (a *AABB3) Intersects(b *AABB3) bool {
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

// UpdateAABB updates this AABB by the transformed AABB3 b, a cannot be the same
// as b.
func (a *AABB3) UpdateAABB(b *AABB3, transform glm.Mat3x4) {
	for i := 0; i < 3; i++ {
		b.Center[i] = transform[i+9]
		b.Radius[i] = 0
		for j := 0; j < 3; j++ {
			b.Center[i] += transform[j*3+i] * a.Center[j]
			b.Radius[i] += math.Abs(transform[j*3+i]) * a.Radius[j]
		}
	}
}