package geo

import (
	"github.com/luxengine/glm"
)

// Slab represent a region R = (x, y, z) | Near <= a*x + b*y + c*z <= Far
type Slab struct {
	// The direction of the slab.
	Normal glm.Vec3

	// The distance from origin along the Normal that the slab starts and end.
	Near, Far float32
}
