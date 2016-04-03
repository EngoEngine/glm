package geo

import (
	"github.com/luxengine/glm"
	"testing"
)

func TestSphere2_Intersects(t *testing.T) {
	tests := []struct {
		a, b       Sphere2
		intersects bool
	}{
		{
			a: Sphere2{
				Center:  glm.Vec2{0, 0},
				Radius:  1,
				Radius2: 1,
			},
			b: Sphere2{
				Center:  glm.Vec2{0, 0},
				Radius:  1,
				Radius2: 1,
			},
			intersects: true,
		},
		{
			a: Sphere2{
				Center:  glm.Vec2{0, 0},
				Radius:  1,
				Radius2: 1,
			},
			b: Sphere2{
				Center:  glm.Vec2{4, 4},
				Radius:  1,
				Radius2: 1,
			},
			intersects: false,
		},

		{
			a: Sphere2{
				Center:  glm.Vec2{0, 0},
				Radius:  1,
				Radius2: 1,
			},
			b: Sphere2{
				Center:  glm.Vec2{2, 0},
				Radius:  1,
				Radius2: 1,
			},
			intersects: true,
		},
	}

	for i, test := range tests {
		if TestSphere2Sphere2(&test.a, &test.b) != test.intersects {
			t.Errorf("[%d] Intersection test failed %v %v", i, test.a, test.b)
		}
	}
}

func TestSphere2_AABB2(t *testing.T) {
	tests := []struct {
		a Sphere2
		b AABB2
	}{
		{
			a: Sphere2{
				Center:  glm.Vec2{3, 4},
				Radius:  1,
				Radius2: 1,
			},
			b: AABB2{
				Center: glm.Vec2{3, 4},
				Radius: glm.Vec2{1, 1},
			},
		},
		{
			a: Sphere2{
				Center:  glm.Vec2{-4, 2.3},
				Radius:  1,
				Radius2: 1,
			},
			b: AABB2{
				Center: glm.Vec2{-4, 2.3},
				Radius: glm.Vec2{1, 1},
			},
		},

		{
			a: Sphere2{
				Center:  glm.Vec2{1, 4},
				Radius:  1,
				Radius2: 1,
			},
			b: AABB2{
				Center: glm.Vec2{1, 4},
				Radius: glm.Vec2{1, 1},
			},
		},
	}

	for i, test := range tests {
		aabb := AABB2FromSphere2(&test.a)
		if !aabb.Center.ApproxEqualThreshold(&test.b.Center, 1e-4) ||
			!aabb.Radius.ApproxEqualThreshold(&test.b.Radius, 1e-4) {
			t.Errorf("[%d] %v.AABB = %v, want %v", i, test.a, aabb, test.b)
		}
	}
}

func BenchmarkSphere2_Intersects(tb *testing.B) {
	a := Sphere2{
		Center:  glm.Vec2{0, 0},
		Radius:  1,
		Radius2: 1,
	}
	b := Sphere2{
		Center:  glm.Vec2{0, 0},
		Radius:  1,
		Radius2: 1,
	}

	for n := 0; n < tb.N; n++ {
		TestSphere2Sphere2(&a, &b)
	}
}
