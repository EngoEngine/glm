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
