package geo2

import (
	"github.com/luxengine/glm"
	"testing"
)

func TestTestAABBAABB(t *testing.T) {
	tests := []struct {
		a, b       AABB
		intersects bool
	}{
		{ // 0
			a: AABB{
				Center:     glm.Vec2{0, 0},
				HalfExtend: glm.Vec2{1, 1},
			},
			b: AABB{
				Center:     glm.Vec2{0.5, 0.5},
				HalfExtend: glm.Vec2{1, 1},
			},
			intersects: true,
		},
		{ // 1
			a: AABB{
				Center:     glm.Vec2{0, 0},
				HalfExtend: glm.Vec2{1, 1},
			},
			b: AABB{
				Center:     glm.Vec2{5, 5},
				HalfExtend: glm.Vec2{1, 1},
			},
			intersects: false,
		},
		{ // 2
			a: AABB{
				Center:     glm.Vec2{5, 0},
				HalfExtend: glm.Vec2{1, 1},
			},
			b: AABB{
				Center:     glm.Vec2{7, 0},
				HalfExtend: glm.Vec2{1, 0},
			},
			intersects: true,
		},
		{ // 3
			a: AABB{
				Center:     glm.Vec2{0, 0},
				HalfExtend: glm.Vec2{1, 1},
			},
			b: AABB{
				Center:     glm.Vec2{0, 0},
				HalfExtend: glm.Vec2{1, 1},
			},
			intersects: true,
		},
		{ // 4
			a: AABB{
				Center:     glm.Vec2{0, 0},
				HalfExtend: glm.Vec2{1, 1},
			},
			b: AABB{
				Center:     glm.Vec2{2, 0},
				HalfExtend: glm.Vec2{1, 1},
			},
			intersects: true,
		},
		{ // 5
			a: AABB{
				Center:     glm.Vec2{0, 0},
				HalfExtend: glm.Vec2{1, 1},
			},
			b: AABB{
				Center:     glm.Vec2{0, 6},
				HalfExtend: glm.Vec2{1, 1},
			},
			intersects: false,
		},
	}

	for i, test := range tests {
		if TestAABBAABB(&test.a, &test.b) != test.intersects {
			t.Errorf("[%d] Intersection test failed %v %v", i, test.a, test.b)
		}
	}
}

func BenchmarkTestAABBAABB(tb *testing.B) {
	a := AABB{
		Center:     glm.Vec2{0, 0},
		HalfExtend: glm.Vec2{1, 1},
	}
	b := AABB{
		Center:     glm.Vec2{0, 0},
		HalfExtend: glm.Vec2{1, 1},
	}
	for n := 0; n < tb.N; n++ {
		TestAABBAABB(&a, &b)
	}
}
