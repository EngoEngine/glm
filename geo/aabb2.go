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
	if math.Abs(a.Center[0]-b.Center[0]) > (a.Radius[0] + b.Radius[0]) {
		return false
	}

	if math.Abs(a.Center[1]-b.Center[1]) > (a.Radius[1] + b.Radius[1]) {
		return false
	}

	return true
}

// UpdateAABB updates this AABB by the transformed AABB3 b, a cannot be the same
// as b.
func (a *AABB2) UpdateAABB(b *AABB2, transform glm.Mat2x3) {
	for i := 0; i < 2; i++ {
		b.Center[i] = transform[i+4]
		b.Radius[i] = 0
		for j := 0; j < 2; j++ {
			b.Center[i] += transform[j*2+i] * a.Center[j]
			b.Radius[i] += math.Abs(transform[j*2+i]) * a.Radius[j]
		}
	}
}