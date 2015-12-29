package geo

import (
	"github.com/luxengine/glm"
	"testing"
)

func TestPlane_PlaneFromPoints(t *testing.T) {
	a, b, c := glm.Vec3{1, 1, 1}, glm.Vec3{0, 1, 1}, glm.Vec3{0, 1, 0}
	p := PlaneFromPoints(&a, &b, &c)
	t.Log(p)
}

func TestPlane_DistanceToPlane(t *testing.T) {
	p := Plane{
		N: glm.Vec3{0, 1, 0},
		D: 1,
	}
	v := glm.Vec3{5, 2, 5}

	want := float32(1.0)
	if d := p.DistanceToPlane(&v); !glm.FloatEqualThreshold(d, want, 1e-4) {
		t.Errorf("DistanceToPlane = %f, want %f", d, want)
	}
}
