package geo2

import (
	"github.com/luxengine/glm"
	"testing"
)

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
