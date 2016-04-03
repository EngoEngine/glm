package geo2

import (
	"github.com/luxengine/glm"
	"testing"
)

func TestPlane_PlaneFromPoints(t *testing.T) {
	tests := []struct {
		points [2]glm.Vec2
		plane  Plane
	}{
	/*{ // 0
		points: [2]glm.Vec2{glm.Vec2{0, 0}, glm.Vec2{1, 0}},
		plane: Plane{
			N: glm.Vec2{0, 1},
			D: 0,
		},
	},*/
	}

	for i, test := range tests {
		got := PlaneFromPoints(&test.points[0], &test.points[1])
		if !got.P.ApproxEqualThreshold(&test.plane.P, 1e-4) {
			t.Errorf("[%d] D = %v, want %v", i, got.P, test.plane.P)
		}
		if !got.N.ApproxEqualThreshold(&test.plane.N, 1e-4) {
			t.Errorf("[%d] N = %v, want %v", i, got.N, test.plane.N)
		}
	}
}

func TestPlane_DistanceToPlane(t *testing.T) {
	p := Plane{
		N: glm.Vec2{0, 1},
		P: glm.Vec2{0, 1},
	}
	v := glm.Vec2{5, 2}

	want := float32(1.0)
	if d := DistanceToPlane(&p, &v); !glm.FloatEqualThreshold(d, want, 1e-4) {
		t.Errorf("DistanceToPlane = %f, want %f", d, want)
	}
}
