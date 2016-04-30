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
		base   AABB
		t      glm.Mat2x3
		expect AABB
	}{
		{ // 0
			base: AABB{
				Center:     glm.Vec2{5, 4},
				HalfExtend: glm.Vec2{1, 3},
			},
			t: glm.Mat2x3{
				1, 0,
				0, 1,
				0, 0,
			},
			expect: AABB{
				Center:     glm.Vec2{5, 4},
				HalfExtend: glm.Vec2{1, 3},
			},
		},
		{ // 1
			base: AABB{
				Center:     glm.Vec2{5, 4},
				HalfExtend: glm.Vec2{1, 3},
			},
			t: glm.Mat2x3{
				1, 0,
				0, 1,
				4, 4,
			},
			expect: AABB{
				Center:     glm.Vec2{9, 8},
				HalfExtend: glm.Vec2{1, 3},
			},
		},
		{ // 2
			base: AABB{
				Center:     glm.Vec2{5, 4},
				HalfExtend: glm.Vec2{1, 3},
			},
			t: glm.Mat2x3{
				0, 1,
				1, 0,
				0, 0,
			},
			expect: AABB{
				Center:     glm.Vec2{4, 5},
				HalfExtend: glm.Vec2{3, 1},
			},
		},
	}

	for i, test := range tests {
		var aabb AABB
		UpdateAABB(&test.base, &aabb, &test.t)
		if !test.expect.Center.ApproxEqual(&aabb.Center) || !test.expect.HalfExtend.ApproxEqual(&aabb.HalfExtend) {
			t.Errorf("[%d] UpdateAABB expected %v got %v", i, test.expect, aabb)
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
