package geo

import (
	"github.com/luxengine/glm"
	"math/rand"
	"testing"
)

func TestIsConvexQuad(t *testing.T) {
	t.Parallel()
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
			t.Errorf("[%d] a(%v), b(%v), c(%v), d(%v) = %t, want %t", i,
				test.a, test.b, test.c, test.d, isconvex, test.isconvex)
		}
	}
}

func TestExtremePointsAlongDirection(t *testing.T) {
	t.Parallel()
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
			t.Errorf("[%d] direction(%v), points(%v) = %d, %d want %d, %d",
				i, test.direction, test.points, imin, imax, test.imin, test.imax)
		}
	}
}

func TestVariance(t *testing.T) {
	tests := []struct {
		slice    []float32
		variance float32
	}{
		{
			[]float32{1, 1},
			0,
		},
		{
			[]float32{1, 2, 3},
			2.0 / 3.0,
		},
		{
			[]float32{1, 3, 5},
			8.0 / 3.0,
		},
		{
			[]float32{600, 470, 170, 430, 300},
			21704,
		},
	}
	for i, test := range tests {
		if variance := Variance(test.slice); !glm.FloatEqual(variance, test.variance) {
			t.Errorf("[%d] Variance(%v) = %f, want %f", i, test.slice, variance, test.variance)
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

func BenchmarkExtremePointsAlongDirection1000(b *testing.B) {
	r := rand.New(rand.NewSource(999))
	dir := glm.Vec3{1, 0, 0}
	points := make([]glm.Vec3, 1000)
	for n := 0; n < 1000; n++ {
		points[n] = glm.Vec3{r.Float32(), r.Float32(), r.Float32()}
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		ExtremePointsAlongDirection(&dir, points)
	}
}

func BenchmarkVariance1000(b *testing.B) {
	r := rand.New(rand.NewSource(999))
	data := make([]float32, 1000)
	for n := 0; n < 1000; n++ {
		data[n] = r.Float32()
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		Variance(data)
	}
}
