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
		{ // 6
			a: AABB{
				Center:     glm.Vec2{0, 0},
				HalfExtend: glm.Vec2{1, 1},
			},
			b: AABB{
				Center:     glm.Vec2{2, 2},
				HalfExtend: glm.Vec2{1, 1},
			},
			intersects: true,
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

func TestUpdateAABB(t *testing.T) {
	tests := []struct {
		base AABB
		t    glm.Mat2x3
		fill AABB
	}{
		{ //0
			base: AABB{
				Center:     glm.Vec2{0, 0},
				HalfExtend: glm.Vec2{1, 1},
			},
			t: glm.Mat2x3{
				1, 2,
				3, 4,
				5, 6,
			},
			fill: AABB{
				Center:     glm.Vec2{10, 130},
				HalfExtend: glm.Vec2{3, 30},
			},
		},
	}

	for i, test := range tests {
		UpdateAABB(&test.base, &test.base, &test.t)
		if test.fill != test.base {
			t.Errorf("[%d] Update expected %v got %v", i, test.fill, test.base)
		}
	}
}

func BenchmarkUpdateAABB(tb *testing.B) {
	a := AABB{
		Center:     glm.Vec2{0, 0},
		HalfExtend: glm.Vec2{1, 1},
	}
	t := glm.Mat2x3{
		1, 2,
		3, 4,
		5, 6,
	}
	for n := 0; n < tb.N; n++ {
		UpdateAABB(&a, &a, &t)
	}
}
