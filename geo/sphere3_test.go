package geo

import (
	"github.com/luxengine/glm"
	"testing"
)

func TestSphere3_Intersects(t *testing.T) {
	tests := []struct {
		a, b       Sphere3
		intersects bool
	}{
		{
			a: Sphere3{
				Center:  glm.Vec3{0, 0, 0},
				Radius:  1,
				Radius2: 1,
			},
			b: Sphere3{
				Center:  glm.Vec3{0, 0, 0},
				Radius:  1,
				Radius2: 1,
			},
			intersects: true,
		},
		{
			a: Sphere3{
				Center:  glm.Vec3{0, 0, 0},
				Radius:  1,
				Radius2: 1,
			},
			b: Sphere3{
				Center:  glm.Vec3{0, 4, 4},
				Radius:  1,
				Radius2: 1,
			},
			intersects: false,
		},

		{
			a: Sphere3{
				Center:  glm.Vec3{0, 0, 0},
				Radius:  1,
				Radius2: 1,
			},
			b: Sphere3{
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

func TestSphere3_AABB3(t *testing.T) {
	tests := []struct {
		a Sphere3
		b AABB3
	}{
		{
			a: Sphere3{
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
			a: Sphere3{
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
			a: Sphere3{
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

func BenchmarkSphere3_Intersects(tb *testing.B) {
	a := Sphere3{
		Center:  glm.Vec3{0, 0, 0},
		Radius:  1,
		Radius2: 1,
	}
	b := Sphere3{
		Center:  glm.Vec3{0, 0, 0},
		Radius:  1,
		Radius2: 1,
	}

	for n := 0; n < tb.N; n++ {
		a.Intersects(&b)
	}
}
