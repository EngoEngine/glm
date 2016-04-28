package geo2

import (
	"github.com/luxengine/glm"
	"testing"
)

func TestSphere_Intersects(t *testing.T) {
	tests := []struct {
		a, b       Sphere
		intersects bool
	}{
		{
			a: Sphere{
				Center:  glm.Vec2{0, 0},
				Radius:  1,
				Radius2: 1,
			},
			b: Sphere{
				Center:  glm.Vec2{0, 0},
				Radius:  1,
				Radius2: 1,
			},
			intersects: true,
		},
		{
			a: Sphere{
				Center:  glm.Vec2{0, 0},
				Radius:  1,
				Radius2: 1,
			},
			b: Sphere{
				Center:  glm.Vec2{4, 4},
				Radius:  1,
				Radius2: 1,
			},
			intersects: false,
		},

		{
			a: Sphere{
				Center:  glm.Vec2{0, 0},
				Radius:  1,
				Radius2: 1,
			},
			b: Sphere{
				Center:  glm.Vec2{2, 0},
				Radius:  1,
				Radius2: 1,
			},
			intersects: true,
		},
	}

	for i, test := range tests {
		if TestSphereSphere(&test.a, &test.b) != test.intersects {
			t.Errorf("[%d] Intersection test failed %v %v", i, test.a, test.b)
		}
	}
}

func TestSphere_AABB(t *testing.T) {
	tests := []struct {
		a Sphere
		b AABB
	}{
		{
			a: Sphere{
				Center:  glm.Vec2{3, 4},
				Radius:  1,
				Radius2: 1,
			},
			b: AABB{
				Center:     glm.Vec2{3, 4},
				HalfExtend: glm.Vec2{1, 1},
			},
		},
		{
			a: Sphere{
				Center:  glm.Vec2{-4, 2.3},
				Radius:  1,
				Radius2: 1,
			},
			b: AABB{
				Center:     glm.Vec2{-4, 2.3},
				HalfExtend: glm.Vec2{1, 1},
			},
		},

		{
			a: Sphere{
				Center:  glm.Vec2{1, 4},
				Radius:  1,
				Radius2: 1,
			},
			b: AABB{
				Center:     glm.Vec2{1, 4},
				HalfExtend: glm.Vec2{1, 1},
			},
		},
	}

	for i, test := range tests {
		aabb := AABBFromSphere(&test.a)
		if !aabb.Center.ApproxEqualThreshold(&test.b.Center, 1e-4) ||
			!aabb.HalfExtend.ApproxEqualThreshold(&test.b.HalfExtend, 1e-4) {
			t.Errorf("[%d] %v.AABB = %v, want %v", i, test.a, aabb, test.b)
		}
	}
}

func BenchmarkSphere_Intersects(tb *testing.B) {
	a := Sphere{
		Center:  glm.Vec2{0, 0},
		Radius:  1,
		Radius2: 1,
	}
	b := Sphere{
		Center:  glm.Vec2{0, 0},
		Radius:  1,
		Radius2: 1,
	}

	for n := 0; n < tb.N; n++ {
		TestSphereSphere(&a, &b)
	}
}
