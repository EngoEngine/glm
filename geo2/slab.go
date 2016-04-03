package geo2

import (
	"github.com/luxengine/glm"
)

// Slab represent a region R = (x, y) | Near <= a*x + b*y <= Far
type Slab struct {
	// The direction of the slab.
	Normal glm.Vec2

	// The distance from origin along the Normal that the slab starts and end.
	Near, Far float32
}
