package geo

import (
	"github.com/engoengine/glm"
	"testing"
)

func TestBarycentric(t *testing.T) {
	tests := []struct {
		a, b, c, p glm.Vec3
		u, v, w    float32
	}{
		{
			a: glm.Vec3{-1, 0, 0},
			b: glm.Vec3{0, 1, 0},
			c: glm.Vec3{1, 0, 0},
			p: glm.Vec3{0, 0.5, 0},
			u: 0.41666663,
			v: 0.16666667,
			w: 0.4166667,
		},
		{
			a: glm.Vec3{-1, 0, 0},
			b: glm.Vec3{0, 1, 0},
			c: glm.Vec3{1, 0, 0},
			p: glm.Vec3{-1, 0, 0},
			u: 1,
			v: 0,
			w: 0,
		},
		{
			a: glm.Vec3{-1, 0, 0},
			b: glm.Vec3{0, 1, 0},
			c: glm.Vec3{1, 0, 0},
			p: glm.Vec3{1, 0, 0},
			u: 0,
			v: 0,
			w: 1,
		},
		{
			a: glm.Vec3{-1, 0, 0},
			b: glm.Vec3{0, 1, 0},
			c: glm.Vec3{1, 0, 0},
			p: glm.Vec3{2, 0, 0},
			u: -0.5,
			v: 0,
			w: 1.5,
		},
	}

	for i, test := range tests {
		u, v, w := Barycentric(&test.a, &test.b, &test.c, &test.p)
		if !glm.FloatEqualThreshold(u, test.u, 1e-4) ||
			!glm.FloatEqualThreshold(v, test.v, 1e-4) ||
			!glm.FloatEqualThreshold(w, test.w, 1e-4) {
			t.Errorf("[%d] a, b, c, p = %v, %v, %v, %v\nu, v, w = %f, %f, %f want %f, %f, %f", i, test.a, test.b, test.c, test.p,
				u, v, w, test.u, test.v, test.w)
		}
	}
}

func BenchmarkBarycentric(tb *testing.B) {
	a, b, c, p := glm.Vec3{1, 2, 3}, glm.Vec3{4, 2, 3}, glm.Vec3{1, 2, 5}, glm.Vec3{2, 3, 4}

	for n := 0; n < tb.N; n++ {
		Barycentric(&a, &b, &c, &p)
	}
}

func BenchmarkBarycentricCache(tb *testing.B) {
	a, b, c, p := glm.Vec3{1, 2, 3}, glm.Vec3{4, 2, 3}, glm.Vec3{1, 2, 5}, glm.Vec3{2, 3, 4}

	cache := BarycentricCacheFromTriangle(&a, &b, &c)

	tb.ResetTimer()
	for n := 0; n < tb.N; n++ {
		BarycentricWithCache(&cache, &p)
	}
}

/*
func BenchmarkBarycentric_OptimizeCandidate(tb *testing.B) {
	a, b, c, p := glm.Vec3{1, 2, 3}, glm.Vec3{4, 2, 3}, glm.Vec3{1, 2, 5}, glm.Vec3{2, 3, 4}

	for n := 0; n < tb.N; n++ {
		barycentric2(&a, &b, &c, &p)
	}
}
*/
