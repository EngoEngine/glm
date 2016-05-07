package glm

import (
	"github.com/luxengine/math"
	"testing"
)

func TestHomogRotate3D(t *testing.T) {
	t.Parallel()
	iden := Ident4()
	tests := []struct {
		Description string
		Angle       float32
		Axis        *Vec3
		Expected    *Mat4
	}{
		{
			"forward",
			0, &Vec3{0, 0, 0},
			&iden,
		},
		{
			"heading 90 degree",
			DegToRad(90), &Vec3{0, 1, 0},
			&Mat4{
				0, 0, -1, 0,
				0, 1, 0, 0,
				1, 0, 0, 0,
				0, 0, 0, 1,
			},
		},
		{
			"heading 180 degree",
			DegToRad(180), &Vec3{0, 1, 0},
			&Mat4{
				-1, 0, 0, 0,
				0, 1, 0, 0,
				0, 0, -1, 0,
				0, 0, 0, 1,
			},
		},
		{
			"attitude 90 degree",
			DegToRad(90), &Vec3{0, 0, 1},
			&Mat4{
				0, 1, 0, 0,
				-1, 0, 0, 0,
				0, 0, 1, 0,
				0, 0, 0, 1,
			},
		},
		{
			"bank 90 degree",
			DegToRad(90), &Vec3{1, 0, 0},
			&Mat4{
				1, 0, 0, 0,
				0, 0, 1, 0,
				0, -1, 0, 0,
				0, 0, 0, 1,
			},
		},
	}

	threshold := math.Pow(10, -2)
	for _, c := range tests {
		if r := HomogRotate3D(c.Angle, c.Axis); !r.EqualThreshold(c.Expected, threshold) {
			t.Errorf("%v failed: HomogRotate3D(%v, %v) != %v (got %v)", c.Description, c.Angle, c.Axis, c.Expected, r)
		}
	}
}

func TestExtract3DScale(t *testing.T) {
	t.Parallel()

	t1 := Translate3D(10, 12, -5)
	h := HomogRotate3D(math.Pi/2, &Vec3{1, 0, 0})
	s := Scale3D(2, 3, 4)
	t2 := t1.Mul4(&h)
	t3 := t2.Mul4(&s)

	iden := Ident4()
	tests := []struct {
		M       Mat4
		X, Y, Z float32
	}{
		{
			iden,
			1, 1, 1,
		}, {
			Scale3D(1, 2, 3),
			1, 2, 3,
		}, {
			t3,
			2, 3, 4,
		},
	}

	eq := FloatEqualFunc(1e-6)
	for _, c := range tests {
		if x, y, z := Extract3DScale(&c.M); !eq(x, c.X) || !eq(y, c.Y) || !eq(z, c.Z) {
			t.Errorf("ExtractScale(%v) != %v, %v, %v (got %v, %v, %v)", c.M, c.X, c.Y, c.Z, x, y, z)
		}
	}
}

func TestExtractMaxScale(t *testing.T) {
	t.Parallel()
	t1 := Translate3D(10, 12, -5)
	h := HomogRotate3D(math.Pi/2, &Vec3{1, 0, 0})
	s := Scale3D(2, 3, 4)
	t2 := t1.Mul4(&h)
	t3 := t2.Mul4(&s)

	iden := Ident4()
	tests := []struct {
		M Mat4
		V float32
	}{
		{
			iden,
			1,
		}, {
			Scale3D(1, 2, 3),
			3,
		}, {
			t3,
			4,
		},
	}

	eq := FloatEqualFunc(1e-6)
	for _, c := range tests {
		if r := ExtractMaxScale(&c.M); !eq(r, c.V) {
			t.Errorf("ExtractMaxScale(%v) != %v (got %v)", c.M, c.V, r)
		}
	}
}

func TestTransformCoordinate(t *testing.T) {
	t.Parallel()

	t1 := Translate3D(0, 1, 1)
	s := Scale3D(2, 2, 2)
	t3 := t1.Mul4(&s)

	iden := Ident4()
	tests := [...]struct {
		v *Vec3
		m *Mat4

		out *Vec3
	}{
		{&Vec3{1, 1, 1}, &iden, &Vec3{1, 1, 1}},
		{&Vec3{1, 1, 1}, &t3, &Vec3{2, 3, 3}},
	}

	for _, test := range tests {
		if v := TransformCoordinate(test.v, test.m); !test.out.ApproxEqualThreshold(&v, 1e-4) {
			t.Errorf("TransformCoordinate on vector %v and matrix %v fails to give result %v (got %v)", test.v, test.m, test.out, v)
		}
	}
}

func TestTransformNormal(t *testing.T) {
	t.Parallel()
	t1 := Translate3D(0, 1, 1)
	s := Scale3D(2, 2, 2)
	t3 := t1.Mul4(&s)

	iden := Ident4()
	tests := [...]struct {
		v *Vec3
		m *Mat4

		out *Vec3
	}{
		{&Vec3{1, 1, 1}, &iden, &Vec3{1, 1, 1}},
		{&Vec3{1, 1, 1}, &t3, &Vec3{2, 2, 2}},
	}

	for _, test := range tests {
		if v := TransformNormal(test.v, test.m); !test.out.ApproxEqualThreshold(&v, 1e-4) {
			t.Errorf("TransformNormal on vector %v and matrix %v fails to give result %v (got %v)", test.v, test.m, test.out, v)
		}
	}
}

func TestRotate2D(t *testing.T) {
	t.Parallel()
	tests := []struct {
		angle      float32
		start, end Vec2
	}{
		{
			angle: math.Pi / 2,
			start: Vec2{1, 0},
			end:   Vec2{-0, 1}, // - zero because IEEE FP will produce like -4e-8
		},
		{
			angle: -math.Pi / 2,
			start: Vec2{1, 0},
			end:   Vec2{-0, -1}, // - zero because IEEE FP will produce like -4e-8
		},
		{
			angle: math.Pi / 4,
			start: Vec2{1, 0},
			end:   Vec2{0.70710677, 0.70710677},
		},
	}

	for i, test := range tests {
		m := Rotate2D(test.angle)
		if end := m.Mul2x1(&test.start); !end.ApproxEqualThreshold(&test.end, 1e-2) {
			t.Errorf("[%d] m * v = %v, want %v", i, end, test.end)
		}
	}
}

func TestRotate3D(t *testing.T) {
	t.Parallel()
	tests := []struct {
		angle                      float32
		start, endRX, endRY, endRZ Vec3
	}{
		{
			angle: math.Pi / 2,
			start: Vec3{1, 0, 0},
			endRX: Vec3{1, 0, 0},
			endRY: Vec3{-0, 0, -1},
			endRZ: Vec3{0, 1, 0},
		},
	}

	for i, test := range tests {
		x := Rotate3DX(test.angle)
		y := Rotate3DY(test.angle)
		z := Rotate3DZ(test.angle)
		if end := x.Mul3x1(&test.start); !end.ApproxEqualThreshold(&test.endRX, 1e-2) {
			t.Errorf("[%d] Rotate3DX m * v = %v, want %v", i, end, test.endRX)
		}

		if end := y.Mul3x1(&test.start); !end.ApproxEqualThreshold(&test.endRY, 1e-2) {
			t.Errorf("[%d] Rotate3DY m * v = %v, want %v", i, end, test.endRY)
		}

		if end := z.Mul3x1(&test.start); !end.ApproxEqualThreshold(&test.endRZ, 1e-2) {
			t.Errorf("[%d] Rotate3DZ m * v = %v, want %v", i, end, test.endRZ)
		}
	}
}

func TestTranslate2D(t *testing.T) {
	tests := []struct {
		t Vec2
		m Mat3
	}{
		{
			t: Vec2{0, 0},
			m: Ident3(),
		},
		{
			t: Vec2{1, 2},
			m: Mat3{
				1, 0, 0,
				0, 1, 0,
				1, 2, 1,
			},
		},
		{
			t: Vec2{-3, 9},
			m: Mat3{
				1, 0, 0,
				0, 1, 0,
				-3, 9, 1,
			},
		},

		{
			t: Vec2{1e8, -1e4},
			m: Mat3{
				1, 0, 0,
				0, 1, 0,
				1e8, -1e4, 1,
			},
		},
	}

	for i, test := range tests {
		if m := Translate2D(test.t[0], test.t[1]); !m.Equal(&test.m) {
			t.Errorf("[%d] Translate2D(%f, %f) = %s, want %s", i, test.t[0], test.t[1], m.String(), test.m.String())
		}
	}
}

func TestHomogRotate2D(t *testing.T) {
	tests := []struct {
		angle float32
		mat   Mat3
	}{
		{
			angle: 0,
			mat: Mat3{
				1, 0, 0,
				0, 1, 0,
				0, 0, 1,
			},
		},
		{
			angle: math.Pi / 6,
			mat: Mat3{
				0.866025, 0.500000, 0,
				-0.500000, 0.866025, 0,
				0, 0, 1,
			},
		},
		{
			angle: math.Pi * 7 / 8,
			mat: Mat3{
				-0.92388, 0.382683, 0,
				-0.382683, -0.92388, 0,
				0, 0, 1,
			},
		},
		{ // https://www.wolframalpha.com/input/?i=rotation+matrix&rawformassumption=%7B%22F%22,+%22RotationCalculator%22,+%22alpha%22%7D+-%3E%22315%22&rawformassumption=%7B%22FP%22,+%22RotationCalculator%22,+%22dir%22%7D+-%3E+%22plus%22&rawformassumption=%7B%22F%22,+%22RotationCalculator%22,+%22point%22%7D+-%3E%22%7B0,+0%7D%22&rawformassumption=%7B%22C%22,+%22rotation+matrix%22%7D+-%3E+%7B%22Calculator%22%7D&rawformassumption=%7B%22MC%22,%22%22%7D-%3E%7B%22Formula%22%7D
			angle: math.Pi * 7 / 4,
			mat: Mat3{
				0.7071067814, -0.7071067814, 0,
				0.7071067814, 0.7071067814, 0,
				0, 0, 1,
			},
		},
	}
	for i, test := range tests {
		if m := HomogRotate2D(test.angle); !m.EqualThreshold(&test.mat, 1e-4) {
			t.Errorf("[%d] HomogRotate2D(%f) = \n%swant \n%s", i, test.angle, m.String(), test.mat.String())
		}
	}
}

func TestScale2D(t *testing.T) {
	tests := []struct {
		sx, sy float32
		mat    Mat3
	}{
		{
			sx:  1,
			sy:  1,
			mat: Ident3(),
		},
		{
			sx: 3,
			sy: 2,
			mat: Mat3{
				3, 0, 0,
				0, 2, 0,
				0, 0, 1,
			},
		},
		{
			sx: -14.5,
			sy: 0.0001,
			mat: Mat3{
				-14.5, 0, 0,
				0, 0.0001, 0,
				0, 0, 1,
			},
		},
	}

	for i, test := range tests {
		if m := Scale2D(test.sx, test.sy); !m.Equal(&test.mat) {
			t.Errorf("[%d] Scale2D(%f, %f) = \n%swant\n%s", i, test.sx, test.sy, m.String(), test.mat.String())
		}
	}
}
