package glm

import (
	"math"
	"testing"
)

func TestHomogRotate3D(t *testing.T) {
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

	threshold := float32(math.Pow(10, -2))
	for _, c := range tests {
		if r := HomogRotate3D(c.Angle, c.Axis); !r.ApproxEqualThreshold(c.Expected, threshold) {
			t.Errorf("%v failed: HomogRotate3D(%v, %v) != %v (got %v)", c.Description, c.Angle, c.Axis, c.Expected, r)
		}
	}
}

func TestExtract3DScale(t *testing.T) {

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
