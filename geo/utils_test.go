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
		if IsConvexQuad(&test.a, &test.b, &test.c, &test.d) != test.isconvex {
			t.Errorf("[%d] a(%v), b(%v), c(%v), d(%v) = %T, want %T", i,
				test.a, test.b, test.c, test.d, test.isconvex)
		}
	}
}

func TestPointFarthestFromEdge(t *testing.T) {
	tests := []struct {
		a, b   glm.Vec2
		points []glm.Vec2
		index  int
	}{
		{
			a: glm.Vec2{0, 0},
			b: glm.Vec2{1, 0},
			points: []glm.Vec2{
				glm.Vec2{0.5, 0.5},
				glm.Vec2{0.1, 0.6},
				glm.Vec2{0.6, 0.4},
				glm.Vec2{0.7, 0.3},
				glm.Vec2{0.8, 0.2},
			},
			index: 1,
		},
		{
			a: glm.Vec2{1, 0},
			b: glm.Vec2{0, 0},
			points: []glm.Vec2{
				glm.Vec2{0.5, -0.5},
				glm.Vec2{0.1, -0.6},
				glm.Vec2{0.6, -0.4},
				glm.Vec2{0.7, -0.3},
				glm.Vec2{0.8, -0.2},
			},
			index: 1,
		},
		{
			a: glm.Vec2{0, 0},
			b: glm.Vec2{1, 0},
			points: []glm.Vec2{
				glm.Vec2{0.5, 0.5},
				glm.Vec2{0.6, 0.4},
				glm.Vec2{0.7, 0.3},
				glm.Vec2{0.8, 0.2},
				glm.Vec2{0.1, 0.6},
			},
			index: 4,
		},
		{
			a: glm.Vec2{1, 0},
			b: glm.Vec2{0, 0},
			points: []glm.Vec2{
				glm.Vec2{0.5, -0.5},
				glm.Vec2{0.6, -0.4},
				glm.Vec2{0.1, -0.6},
				glm.Vec2{0.7, -0.3},
				glm.Vec2{0.8, -0.2},
			},
			index: 2,
		},
		{
			a: glm.Vec2{0, 0},
			b: glm.Vec2{0, 1},
			points: []glm.Vec2{
				glm.Vec2{-1, 0.3},
				glm.Vec2{-0.9, 0.4},
			},
			index: 0,
		},
		{
			a: glm.Vec2{0, 1},
			b: glm.Vec2{0, 0},
			points: []glm.Vec2{
				glm.Vec2{0.9, 0.4},
				glm.Vec2{0.7, 0.9},
				glm.Vec2{1, 0.3},
			},
			index: 2,
		},
	}

	for i, test := range tests {
		if index := PointFarthestFromEdge(&test.a, &test.b, test.points); index != test.index {
			t.Errorf("[%d] index = %d want %d", i, index, test.index)
		}
	}
}

func TestExtremePointsAlongDirection2(t *testing.T) {
	tests := []struct {
		direction  glm.Vec2
		points     []glm.Vec2
		imin, imax int
	}{
		{
			direction: glm.Vec2{0, 1},
			points: []glm.Vec2{
				glm.Vec2{0, 0},
				glm.Vec2{4, -9},
				glm.Vec2{2, 1},
				glm.Vec2{5.4, 7},
				glm.Vec2{1, 2},
				glm.Vec2{-4, -5},
			},
			imin: 1,
			imax: 3,
		},
		{
			direction: glm.Vec2{0, -1},
			points: []glm.Vec2{
				glm.Vec2{0, 0},
				glm.Vec2{4, -9},
				glm.Vec2{2, 1},
				glm.Vec2{5.4, 7},
				glm.Vec2{1, 2},
				glm.Vec2{-4, -5},
			},
			imin: 3,
			imax: 1,
		},
		{
			direction: glm.Vec2{1, 1},
			points: []glm.Vec2{
				glm.Vec2{0, 0},
				glm.Vec2{10, 10},
				glm.Vec2{0, -10},
				glm.Vec2{1, 0},
				glm.Vec2{0, 1},
			},
			imin: 2,
			imax: 1,
		},
	}

	for i, test := range tests {
		imin, imax := ExtremePointsAlongDirection2(&test.direction, test.points)
		if imin != test.imin || imax != test.imax {
			t.Errorf("[%d] direction = %v, points = %v, imin, imax = %d, %d want %d, %d",
				i, test.direction, test.points, imin, imax, test.imin, test.imax)
		}
	}
}

func TestExtremePointsAlongDirection3(t *testing.T) {
	tests := []struct {
		direction  glm.Vec3
		points     []glm.Vec3
		imin, imax int
	}{
		{
			direction: glm.Vec3{0, 1, 0},
			points: []glm.Vec3{
				glm.Vec3{0, 0, 0},
				glm.Vec3{4, -9, 0},
				glm.Vec3{2, 1, 0},
				glm.Vec3{5.4, 7, 0},
				glm.Vec3{1, 2, 0},
				glm.Vec3{-4, -5, 0},
			},
			imin: 1,
			imax: 3,
		},
		{
			direction: glm.Vec3{0, -1, 0},
			points: []glm.Vec3{
				glm.Vec3{0, 0, 0},
				glm.Vec3{4, -9, 0},
				glm.Vec3{2, 1, 0},
				glm.Vec3{5.4, 7, 0},
				glm.Vec3{1, 2, 0},
				glm.Vec3{-4, -5, 0},
			},
			imin: 3,
			imax: 1,
		},
		{
			direction: glm.Vec3{1, 1, 0},
			points: []glm.Vec3{
				glm.Vec3{0, 0, 0},
				glm.Vec3{10, 10, 0},
				glm.Vec3{0, -10, 0},
				glm.Vec3{1, 0, 0},
				glm.Vec3{0, 1, 0},
			},
			imin: 2,
			imax: 1,
		},
	}

	for i, test := range tests {
		imin, imax := ExtremePointsAlongDirection3(&test.direction, test.points)
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
