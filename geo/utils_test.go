package geo

import (
	"github.com/luxengine/glm"
	"testing"
)

func TestIsConvexQuad(t *testing.T) {
	tests := []struct {
		a, b, c, d glm.Vec3
		isconvex   bool
	}{
		{
			a:        glm.Vec3{0, 0, 0},
			b:        glm.Vec3{0, 1, 0},
			c:        glm.Vec3{1, 1, 0},
			d:        glm.Vec3{1, 0, 0},
			isconvex: true,
		},
		{
			a:        glm.Vec3{0, 0, 0},
			b:        glm.Vec3{1, 1, 0},
			c:        glm.Vec3{0, 1, 0},
			d:        glm.Vec3{1, 0, 0},
			isconvex: false,
		},
		{
			a:        glm.Vec3{0, 0, 0},
			b:        glm.Vec3{0, 0, 1},
			c:        glm.Vec3{1, 0, 4},
			d:        glm.Vec3{1, 0, 0},
			isconvex: true,
		},
		{
			a:        glm.Vec3{0, 0, 0},
			b:        glm.Vec3{0, 1, 0},
			c:        glm.Vec3{0, 4, 1},
			d:        glm.Vec3{0, 0, 1},
			isconvex: true,
		},
		{
			a:        glm.Vec3{0, 0, 0},
			b:        glm.Vec3{1, 0, 0},
			c:        glm.Vec3{4, 0, 1},
			d:        glm.Vec3{0, 0, 1},
			isconvex: true,
		},
	}

	for i, test := range tests {
		if isconvex := IsConvexQuad(&test.a, &test.b, &test.c, &test.d); isconvex != test.isconvex {
			t.Errorf("[%d] a(%v), b(%v), c(%v), d(%v) = %T, want %T", i,
				test.a, test.b, test.c, test.d, isconvex, test.isconvex)
		}
	}
}

func TestExtremePointsAlongDirection(t *testing.T) {
	tests := []struct {
		direction  glm.Vec3
		points     []glm.Vec3
		imin, imax int
	}{
		{
			direction: glm.Vec3{0, 1, 0},
			points: []glm.Vec3{
				{0, 0, 0},
				{4, -9, 0},
				{2, 1, 0},
				{5.4, 7, 0},
				{1, 2, 0},
				{-4, -5, 0},
			},
			imin: 1,
			imax: 3,
		},
		{
			direction: glm.Vec3{0, -1, 0},
			points: []glm.Vec3{
				{0, 0, 0},
				{4, -9, 0},
				{2, 1, 0},
				{5.4, 7, 0},
				{1, 2, 0},
				{-4, -5, 0},
			},
			imin: 3,
			imax: 1,
		},
		{
			direction: glm.Vec3{1, 1, 0},
			points: []glm.Vec3{
				{0, 0, 0},
				{10, 10, 0},
				{0, -10, 0},
				{1, 0, 0},
				{0, 1, 0},
			},
			imin: 2,
			imax: 1,
		},
	}

	for i, test := range tests {
		imin, imax := ExtremePointsAlongDirection(&test.direction, test.points)
		if imin != test.imin || imax != test.imax {
			t.Errorf("[%d] direction = %v, points = %v, imin, imax = %d, %d want %d, %d",
				i, test.direction, test.points, imin, imax, test.imin, test.imax)
		}
	}
}

func BenchmarkIsConvexQuad(b *testing.B) {
	bench := struct {
		a, b, c, d glm.Vec3
		isconvex   bool
	}{
		a:        glm.Vec3{0, 0, 0},
		b:        glm.Vec3{0, 1, 0},
		c:        glm.Vec3{1, 1, 0},
		d:        glm.Vec3{1, 0, 0},
		isconvex: true,
	}
	for n := 0; n < b.N; n++ {
		IsConvexQuad(&bench.a, &bench.b, &bench.c, &bench.d)
	}
}
