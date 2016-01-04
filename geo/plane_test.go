package geo

import (
	"github.com/luxengine/glm"
	"testing"
)

func TestPlane3_Plane3FromPoints(t *testing.T) {
	tests := []struct {
		points [3]glm.Vec3
		plane  Plane3
	}{
		{ // 0. up, above 0,0,0
			points: [3]glm.Vec3{glm.Vec3{0, 1, 0}, glm.Vec3{0, 1, 1}, glm.Vec3{1, 1, 0}},
			plane: Plane3{
				N: glm.Vec3{0, 1, 0},
				D: 1,
			},
		},
		{ // 1. up triangle moved away from origin
			points: [3]glm.Vec3{glm.Vec3{0 + 5, 1, 0 + 5}, glm.Vec3{0 + 5, 1, 1 + 5}, glm.Vec3{1 + 5, 1, 0 + 5}},
			plane: Plane3{
				N: glm.Vec3{0, 1, 0},
				D: 1,
			},
		},
		{ // 2. down, near origin
			points: [3]glm.Vec3{glm.Vec3{0, 1, 1}, glm.Vec3{0, 1, 0}, glm.Vec3{1, 1, 0}},
			plane: Plane3{
				N: glm.Vec3{0, -1, 0},
				D: 1,
			},
		},
		{ // 3. up, at -1
			points: [3]glm.Vec3{glm.Vec3{0, -1, 0}, glm.Vec3{0, -1, 1}, glm.Vec3{1, -1, 0}},
			plane: Plane3{
				N: glm.Vec3{0, 1, 0},
				D: -1,
			},
		},
	}

	for i, test := range tests {
		got := Plane3FromPoints(&test.points[0], &test.points[1], &test.points[2])
		if !glm.FloatEqualThreshold(got.D, test.plane.D, 1e-4) {
			t.Errorf("[%d] D = %f, want %f", i, got.D, test.plane.D)
		}
		if !got.N.ApproxEqualThreshold(&test.plane.N, 1e-4) {
			t.Errorf("[%d] N = %v, want %v", i, got.N, test.plane.N)
		}
	}
}

func TestPlane3_DistanceToPlane(t *testing.T) {
	p := Plane3{
		N: glm.Vec3{0, 1, 0},
		D: 1,
	}
	v := glm.Vec3{5, 2, 5}

	want := float32(1.0)
	if d := p.DistanceToPlane(&v); !glm.FloatEqualThreshold(d, want, 1e-4) {
		t.Errorf("DistanceToPlane = %f, want %f", d, want)
	}
}

func TestPlane2_Plane2FromPoints(t *testing.T) {
	tests := []struct {
		points [2]glm.Vec2
		plane  Plane2
	}{
		{
			points: [2]glm.Vec2{glm.Vec2{0, 0}, glm.Vec2{1, 0}},
			plane: Plane2{
				N: glm.Vec2{0, 1},
				D: 0,
			},
		},
		{
			points: [2]glm.Vec2{glm.Vec2{0, 0}, glm.Vec2{1, 1}},
			plane: Plane2{
				N: glm.Vec2{-0.7072, 0.7072},
				D: 0,
			},
		},
		{
			points: [2]glm.Vec2{glm.Vec2{0, 0}, glm.Vec2{1, -1}},
			plane: Plane2{
				N: glm.Vec2{0.7072, 0.7072},
				D: 0,
			},
		},
		{
			points: [2]glm.Vec2{glm.Vec2{0, 1}, glm.Vec2{1, 0}},
			plane: Plane2{
				N: glm.Vec2{0.7072, 0.7072},
				D: 1,
			},
		},
	}

	for i, test := range tests {
		got := Plane2FromPoints(&test.points[0], &test.points[1])
		if !glm.FloatEqualThreshold(got.D, test.plane.D, 1e-4) {
			t.Errorf("[%d] D = %f, want %f", i, got.D, test.plane.D)
		}
		if !got.N.ApproxEqualThreshold(&test.plane.N, 1e-4) {
			t.Errorf("[%d] N = %v, want %v", i, got.N, test.plane.N)
		}
	}
}

func TestPlane2_DistanceToPlane(t *testing.T) {
	tests := []struct {
	}{}
}
