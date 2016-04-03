package geo3

import (
	"github.com/luxengine/glm"
	"testing"
)

func TestQuickhull(t *testing.T) {
	t.Skip("Quickhull 3D is not implemented yet")
	points := []glm.Vec3{{0, 0, 0}, {1, 1, 1},
		{2, 0, 0}, {0, 2, 0}, {0, 0, 2},
		{-1, 0, 0}, {0, -1, 0}, {0, 0, -1}, {0.1, 0.1, 0.1},
		{0, 1.9, 1.9}} // that last one should see 2 faces :\
	Quickhull(points)
}
