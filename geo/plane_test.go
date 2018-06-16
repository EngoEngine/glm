package geo

import (
	"github.com/engoengine/glm"
	"testing"
)

func TestPlane_PlaneFromPoints(t *testing.T) {
	tests := []struct {
		points [3]glm.Vec3
		plane  Plane
	}{
		/*{ // 0. up, above 0,0,0
			points: [3]glm.Vec3{glm.Vec3{0, 1, 0}, glm.Vec3{0, 1, 1}, glm.Vec3{1, 1, 0}},
			plane: Plane{
				N: glm.Vec3{0, 1, 0},
				D: 1,
			},
		},*/
	}

	for i, test := range tests {
		got := PlaneFromPoints(&test.points[0], &test.points[1], &test.points[2])
		if !got.P.ApproxEqualThreshold(&test.plane.P, 1e-4) {
			t.Errorf("[%d] P = %v, want %v", i, got.P, test.plane.P)
		}
		if !got.N.ApproxEqualThreshold(&test.plane.N, 1e-4) {
			t.Errorf("[%d] N = %v, want %v", i, got.N, test.plane.N)
		}
	}
}

func TestPlane_DistanceToPlane(t *testing.T) {
	p := Plane{
		N: glm.Vec3{0, 1, 0},
		P: glm.Vec3{0, 1, 0},
	}
	v := glm.Vec3{5, 2, 5}

	want := float32(1.0)
	if d := DistanceToPlane(&p, &v); !glm.FloatEqualThreshold(d, want, 1e-4) {
		t.Errorf("DistanceToPlane = %f, want %f", d, want)
	}
}
