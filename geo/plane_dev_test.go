// +build dev
package geo

import (
	"github.com/luxengine/glm"
	"testing"
)

func TestPlane2_Dev(t *testing.T) {
	a, b := glm.Vec2{1, 2}, glm.Vec2{1, 3}
	t.Log(Plane2FromPoints(&a, &b))
}
