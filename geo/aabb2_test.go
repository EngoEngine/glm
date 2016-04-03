package geo

import (
	"github.com/luxengine/glm"
	"testing"
)

func TestAABB2_Intersects(t *testing.T) {
	tests := []struct {
		a, b       AABB2
		intersects bool
	}{
		{ // 0
			a: AABB2{
				Center: glm.Vec2{0, 0},
				Radius: glm.Vec2{1, 1},
			},
			b: AABB2{
				Center: glm.Vec2{0.5, 0.5},
				Radius: glm.Vec2{1, 1},
			},
			intersects: true,
		},
		{ // 1
			a: AABB2{
				Center: glm.Vec2{0, 0},
				Radius: glm.Vec2{1, 1},
			},
			b: AABB2{
				Center: glm.Vec2{5, 5},
				Radius: glm.Vec2{1, 1},
			},
			intersects: false,
		},
		{ // 2
			a: AABB2{
				Center: glm.Vec2{5, 0},
				Radius: glm.Vec2{1, 1},
			},
			b: AABB2{
				Center: glm.Vec2{7, 0},
				Radius: glm.Vec2{1, 0},
			},
			intersects: true,
		},
		{ // 3
			a: AABB2{
				Center: glm.Vec2{0, 0},
				Radius: glm.Vec2{1, 1},
			},
			b: AABB2{
				Center: glm.Vec2{0, 0},
				Radius: glm.Vec2{1, 1},
			},
			intersects: true,
		},
		{ // 4
			a: AABB2{
				Center: glm.Vec2{0, 0},
				Radius: glm.Vec2{1, 1},
			},
			b: AABB2{
				Center: glm.Vec2{2, 0},
				Radius: glm.Vec2{1, 1},
			},
			intersects: true,
		},
		{ // 5
			a: AABB2{
				Center: glm.Vec2{0, 0},
				Radius: glm.Vec2{1, 1},
			},
			b: AABB2{
				Center: glm.Vec2{0, 6},
				Radius: glm.Vec2{1, 1},
			},
			intersects: false,
		},
	}

	for i, test := range tests {
		if TestAABB2AABB2(&test.a, &test.b) != test.intersects {
			t.Errorf("[%d] Intersection test failed %v %v", i, test.a, test.b)
		}
	}
}

func BenchmarkAABB2_Intersects(tb *testing.B) {
	a := AABB2{
		Center: glm.Vec2{0, 0},
		Radius: glm.Vec2{1, 1},
	}
	b := AABB2{
		Center: glm.Vec2{0, 0},
		Radius: glm.Vec2{1, 1},
	}
	for n := 0; n < tb.N; n++ {
		TestAABB2AABB2(&a, &b)
	}
}
