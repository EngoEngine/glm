package debug

import (
	"github.com/luxengine/glm"
	"testing"
)

func TestDraw(t *testing.T) {
	Draw("draw.svg", glm.Vec2{200, 200}, glm.Vec2{100, 100}, glm.Vec2{0, 100}, glm.Vec2{-10, 150})
}
