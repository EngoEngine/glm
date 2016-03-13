package geo

import (
	"github.com/luxengine/glm"
)

// OBB2 is a Oriented Bounding Box for 2d
type OBB2 struct {
	Center      glm.Vec2
	Orientation [2]glm.Vec2
	Radius      glm.Vec2
}
