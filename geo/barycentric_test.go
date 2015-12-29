package geo

import (
	"github.com/luxengine/glm"
	"testing"
)

func TestBarycentric(t *testing.T) {
	a, b, c, p := glm.Vec3{1, 2, 3}, glm.Vec3{4, 2, 3}, glm.Vec3{1, 2, 5}, glm.Vec3{2, 3, 4}
	t.Log(Barycentric(&a, &b, &c, &p))
}
