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
		if test.a.Intersects(&test.b) != test.intersects {
			t.Errorf("[%d] Intersection test failed %v %v", i, test.a, test.b)
		}
	}
}

func TestAABB3_Intersects(t *testing.T) {
	tests := []struct {
		a, b       AABB3
		intersects bool
	}{
		{ // 0
			a: AABB3{
				Center: glm.Vec3{0, 0, 1},
				Radius: glm.Vec3{1, 1, 1},
			},
			b: AABB3{
				Center: glm.Vec3{0.5, 0.5, 1},
				Radius: glm.Vec3{1, 1, 1},
			},
			intersects: true,
		},
		{ // 1
			a: AABB3{
				Center: glm.Vec3{0, 0, 1},
				Radius: glm.Vec3{1, 1, 1},
			},
			b: AABB3{
				Center: glm.Vec3{5, 5, 1},
				Radius: glm.Vec3{1, 1, 1},
			},
			intersects: false,
		},
		{ // 2
			a: AABB3{
				Center: glm.Vec3{5, 0, 1},
				Radius: glm.Vec3{1, 1, 1},
			},
			b: AABB3{
				Center: glm.Vec3{7, 0, 1},
				Radius: glm.Vec3{1, 0, 1},
			},
			intersects: true,
		},
		{ // 3
			a: AABB3{
				Center: glm.Vec3{0, 0, 1},
				Radius: glm.Vec3{1, 1, 1},
			},
			b: AABB3{
				Center: glm.Vec3{0, 0, 1},
				Radius: glm.Vec3{1, 1, 1},
			},
			intersects: true,
		},
		{ // 4
			a: AABB3{
				Center: glm.Vec3{0, 0, 1},
				Radius: glm.Vec3{1, 1, 1},
			},
			b: AABB3{
				Center: glm.Vec3{2, 0, 1},
				Radius: glm.Vec3{1, 1, 1},
			},
			intersects: true,
		},

		{ // 5
			a: AABB3{
				Center: glm.Vec3{0, 0, 1},
				Radius: glm.Vec3{1, 1, 1},
			},
			b: AABB3{
				Center: glm.Vec3{0, 6, 1},
				Radius: glm.Vec3{1, 1, 1},
			},
			intersects: false,
		},

		{ // 6
			a: AABB3{
				Center: glm.Vec3{0, 0, 1},
				Radius: glm.Vec3{1, 1, 1},
			},
			b: AABB3{
				Center: glm.Vec3{0, 0, 7},
				Radius: glm.Vec3{1, 1, 1},
			},
			intersects: false,
		},
	}

	for i, test := range tests {
		if test.a.Intersects(&test.b) != test.intersects {
			t.Errorf("[%d] Intersection test failed %v %v", i, test.a, test.b)
		}
	}
}

func TestBoundingSphere2_Intersects(t *testing.T) {
	tests := []struct {
		a, b       BoundingSphere2
		intersects bool
	}{
		{
			a: BoundingSphere2{
				Center:  glm.Vec2{0, 0},
				Radius:  1,
				Radius2: 1,
			},
			b: BoundingSphere2{
				Center:  glm.Vec2{0, 0},
				Radius:  1,
				Radius2: 1,
			},
			intersects: true,
		},
		{
			a: BoundingSphere2{
				Center:  glm.Vec2{0, 0},
				Radius:  1,
				Radius2: 1,
			},
			b: BoundingSphere2{
				Center:  glm.Vec2{4, 4},
				Radius:  1,
				Radius2: 1,
			},
			intersects: false,
		},

		{
			a: BoundingSphere2{
				Center:  glm.Vec2{0, 0},
				Radius:  1,
				Radius2: 1,
			},
			b: BoundingSphere2{
				Center:  glm.Vec2{2, 0},
				Radius:  1,
				Radius2: 1,
			},
			intersects: true,
		},
	}

	for i, test := range tests {
		if test.a.Intersects(&test.b) != test.intersects {
			t.Errorf("[%d] Intersection test failed %v %v", i, test.a, test.b)
		}
	}
}

func TestBoundingSphere3_Intersects(t *testing.T) {
	tests := []struct {
		a, b       BoundingSphere3
		intersects bool
	}{
		{
			a: BoundingSphere3{
				Center:  glm.Vec3{0, 0, 0},
				Radius:  1,
				Radius2: 1,
			},
			b: BoundingSphere3{
				Center:  glm.Vec3{0, 0, 0},
				Radius:  1,
				Radius2: 1,
			},
			intersects: true,
		},
		{
			a: BoundingSphere3{
				Center:  glm.Vec3{0, 0, 0},
				Radius:  1,
				Radius2: 1,
			},
			b: BoundingSphere3{
				Center:  glm.Vec3{0, 4, 4},
				Radius:  1,
				Radius2: 1,
			},
			intersects: false,
		},

		{
			a: BoundingSphere3{
				Center:  glm.Vec3{0, 0, 0},
				Radius:  1,
				Radius2: 1,
			},
			b: BoundingSphere3{
				Center:  glm.Vec3{0, 2, 0},
				Radius:  1,
				Radius2: 1,
			},
			intersects: true,
		},
	}

	for i, test := range tests {
		if test.a.Intersects(&test.b) != test.intersects {
			t.Errorf("[%d] Intersection test failed %v %v", i, test.a, test.b)
		}
	}
}

func TestBoundingSphere2_AABB2(t *testing.T) {
	tests := []struct {
		a BoundingSphere2
		b AABB2
	}{
		{
			a: BoundingSphere2{
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
			a: BoundingSphere2{
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
			a: BoundingSphere2{
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
		aabb := test.a.AABB2()
		if !aabb.Center.ApproxEqualThreshold(&test.b.Center, 1e-4) ||
			!aabb.Radius.ApproxEqualThreshold(&test.b.Radius, 1e-4) {
			t.Errorf("[%d] %v.AABB = %v, want %v", i, test.a, aabb, test.b)
		}
	}
}

func TestBoundingSphere3_AABB3(t *testing.T) {
	tests := []struct {
		a BoundingSphere3
		b AABB3
	}{
		{
			a: BoundingSphere3{
				Center:  glm.Vec3{3, 4, 7},
				Radius:  1,
				Radius2: 1,
			},
			b: AABB3{
				Center: glm.Vec3{3, 4, 7},
				Radius: glm.Vec3{1, 1, 1},
			},
		},
		{
			a: BoundingSphere3{
				Center:  glm.Vec3{-4, 2.3, 9.9},
				Radius:  1,
				Radius2: 1,
			},
			b: AABB3{
				Center: glm.Vec3{-4, 2.3, 9.9},
				Radius: glm.Vec3{1, 1, 1},
			},
		},

		{
			a: BoundingSphere3{
				Center:  glm.Vec3{1, 4, 5},
				Radius:  1,
				Radius2: 1,
			},
			b: AABB3{
				Center: glm.Vec3{1, 4, 5},
				Radius: glm.Vec3{1, 1, 1},
			},
		},
	}

	for i, test := range tests {
		aabb := test.a.AABB3()
		if !aabb.Center.ApproxEqualThreshold(&test.b.Center, 1e-4) ||
			!aabb.Radius.ApproxEqualThreshold(&test.b.Radius, 1e-4) {
			t.Errorf("[%d] %v.AABB = %v, want %v", i, test.a, aabb, test.b)
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
		a.Intersects(&b)
	}
}

func BenchmarkAABB3_Intersects(tb *testing.B) {
	a := AABB3{
		Center: glm.Vec3{0, 0, 0},
		Radius: glm.Vec3{1, 1, 1},
	}
	b := AABB3{
		Center: glm.Vec3{0, 0, 0},
		Radius: glm.Vec3{1, 1, 1},
	}
	for n := 0; n < tb.N; n++ {
		a.Intersects(&b)
	}
}

func BenchmarkBoundingSphere2_Intersects(tb *testing.B) {
	a := BoundingSphere2{
		Center:  glm.Vec2{0, 0},
		Radius:  1,
		Radius2: 1,
	}
	b := BoundingSphere2{
		Center:  glm.Vec2{0, 0},
		Radius:  1,
		Radius2: 1,
	}

	for n := 0; n < tb.N; n++ {
		a.Intersects(&b)
	}
}

func BenchmarkBoundingSphere3_Intersects(tb *testing.B) {
	a := BoundingSphere3{
		Center:  glm.Vec3{0, 0, 0},
		Radius:  1,
		Radius2: 1,
	}
	b := BoundingSphere3{
		Center:  glm.Vec3{0, 0, 0},
		Radius:  1,
		Radius2: 1,
	}

	for n := 0; n < tb.N; n++ {
		a.Intersects(&b)
	}
}
