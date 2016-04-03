package geo3

import (
	"github.com/luxengine/glm"
	"testing"
)

func TestAABB_Intersects(t *testing.T) {
	tests := []struct {
		a, b       AABB
		intersects bool
	}{
		{ // 0
			a: AABB{
				Center: glm.Vec3{0, 0, 1},
				Radius: glm.Vec3{1, 1, 1},
			},
			b: AABB{
				Center: glm.Vec3{0.5, 0.5, 1},
				Radius: glm.Vec3{1, 1, 1},
			},
			intersects: true,
		},
		{ // 1
			a: AABB{
				Center: glm.Vec3{0, 0, 1},
				Radius: glm.Vec3{1, 1, 1},
			},
			b: AABB{
				Center: glm.Vec3{5, 5, 1},
				Radius: glm.Vec3{1, 1, 1},
			},
			intersects: false,
		},
		{ // 2
			a: AABB{
				Center: glm.Vec3{5, 0, 1},
				Radius: glm.Vec3{1, 1, 1},
			},
			b: AABB{
				Center: glm.Vec3{7, 0, 1},
				Radius: glm.Vec3{1, 0, 1},
			},
			intersects: true,
		},
		{ // 3
			a: AABB{
				Center: glm.Vec3{0, 0, 1},
				Radius: glm.Vec3{1, 1, 1},
			},
			b: AABB{
				Center: glm.Vec3{0, 0, 1},
				Radius: glm.Vec3{1, 1, 1},
			},
			intersects: true,
		},
		{ // 4
			a: AABB{
				Center: glm.Vec3{0, 0, 1},
				Radius: glm.Vec3{1, 1, 1},
			},
			b: AABB{
				Center: glm.Vec3{2, 0, 1},
				Radius: glm.Vec3{1, 1, 1},
			},
			intersects: true,
		},

		{ // 5
			a: AABB{
				Center: glm.Vec3{0, 0, 1},
				Radius: glm.Vec3{1, 1, 1},
			},
			b: AABB{
				Center: glm.Vec3{0, 6, 1},
				Radius: glm.Vec3{1, 1, 1},
			},
			intersects: false,
		},

		{ // 6
			a: AABB{
				Center: glm.Vec3{0, 0, 1},
				Radius: glm.Vec3{1, 1, 1},
			},
			b: AABB{
				Center: glm.Vec3{0, 0, 7},
				Radius: glm.Vec3{1, 1, 1},
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
		Center: glm.Vec3{0, 0, 0},
		Radius: glm.Vec3{1, 1, 1},
	}
	b := AABB{
		Center: glm.Vec3{0, 0, 0},
		Radius: glm.Vec3{1, 1, 1},
	}
	for n := 0; n < tb.N; n++ {
		TestAABBAABB(&a, &b)
	}
}
