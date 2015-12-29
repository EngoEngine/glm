package glm

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestQuatMulIdentity(t *testing.T) {
	//t.Parallel()

	i1 := &Quat{1.0, Vec3{0, 0, 0}}
	i2 := QuatIdent()
	i3 := QuatIdent()

	mul := i2.Mul(&i3)

	if !FloatEqual(mul.W, 1.0) {
		t.Errorf("Multiplication of identities does not yield identity")
	}

	if mul.V[0] != i1.V[0] {
		t.Errorf("Multiplication of identities does not yield identity")
	}
	if mul.V[1] != i1.V[1] {
		t.Errorf("Multiplication of identities does not yield identity")
	}
	if mul.V[2] != i1.V[2] {
		t.Errorf("Multiplication of identities does not yield identity")
	}
}

func TestQuatRotateOnAxis(t *testing.T) {
	//t.Parallel()

	var angleDegrees float32 = 30.0
	axis := Vec3{1, 0, 0}

	i1 := QuatRotate(DegToRad(angleDegrees), &axis)

	rotatedAxis := i1.Rotate(&axis)

	if !FloatEqualThreshold(rotatedAxis[0], axis[0], 1e-4) {
		t.Errorf("Rotation of axis does not yield identity")
	}
	if !FloatEqualThreshold(rotatedAxis[1], axis[1], 1e-4) {
		t.Errorf("Rotation of axis does not yield identity")
	}
	if !FloatEqualThreshold(rotatedAxis[2], axis[2], 1e-4) {
		t.Errorf("Rotation of axis does not yield identity")
	}
}

func TestQuatRotateOffAxis(t *testing.T) {
	//t.Parallel()

	angleRads := DegToRad(30.0)
	axis := Vec3{1, 0, 0}

	i1 := QuatRotate(angleRads, &axis)

	vector := Vec3{0, 1, 0}
	rotatedVector := i1.Rotate(&vector)

	s, c := math.Sincos(float64(angleRads))
	answer := Vec3{0, float32(c), float32(s)}

	if !FloatEqualThreshold(rotatedVector[0], answer[0], 1e-4) {
		t.Errorf("Rotation of vector does not yield answer")
	}
	if !FloatEqualThreshold(rotatedVector[1], answer[1], 1e-4) {
		t.Errorf("Rotation of vector does not yield answer")
	}
	if !FloatEqualThreshold(rotatedVector[2], answer[2], 1e-4) {
		t.Errorf("Rotation of vector does not yield answer")
	}
}

func TestQuatIdentityToMatrix(t *testing.T) {
	//t.Parallel()

	quat := QuatIdent()
	matrix := quat.Mat4()
	answer := Ident4()

	if !matrix.ApproxEqual(&answer) {
		t.Errorf("Identity quaternion does not yield identity matrix")
	}
}

func TestQuatRotationToMatrix(t *testing.T) {
	//t.Parallel()

	angle := DegToRad(45.0)

	axis := Vec3{1, 2, 3}
	axis.Normalize()
	quat := QuatRotate(angle, &axis)
	matrix := quat.Mat4()
	answer := HomogRotate3D(angle, &axis)

	if !matrix.ApproxEqualThreshold(&answer, 1e-4) {
		t.Errorf("Rotation quaternion does not yield correct rotation matrix; got: %v expected: %v", matrix, answer)
	}
}

// Taken from the Matlab AnglesToQuat documentation example
func TestAnglesToQuatZYX(t *testing.T) {
	//t.Parallel()

	q := AnglesToQuat(.7854, 0.1, 0, ZYX)

	t.Log("Calculated quaternion: ", q, "\n")

	if !FloatEqualThreshold(q.W, .9227, 1e-3) {
		t.Errorf("Quaternion W incorrect. Got: %f Expected: %f", q.W, .9227)
	}

	if !q.V.ApproxEqualThreshold(&Vec3{-0.0191, 0.0462, 0.3822}, 1e-3) {
		t.Errorf("Quaternion V incorrect. Got: %v, Expected: %v", q.V, Vec3{-0.0191, 0.0462, 0.3822})
	}
}

func TestQuatMatRotateY(t *testing.T) {
	//t.Parallel()

	q := QuatRotate(float32(math.Pi), &Vec3{0, 1, 0})
	q.Normalize()
	v := Vec3{1, 0, 0}

	result := q.Rotate(&v)

	r := Rotate3DY(float32(math.Pi))
	expected := r.Mul3x1(&v)
	t.Logf("Computed from rotation matrix: %v", expected)
	if !result.ApproxEqualThreshold(&expected, 1e-4) {
		t.Errorf("Quaternion rotating vector doesn't match 3D matrix method. Got: %v, Expected: %v", result, expected)
	}
	m := q.Mul(&Quat{0, v})
	c := q.Conjugated()
	mcv := m.Mul(&c).V
	expected = mcv
	t.Logf("Computed from conjugate method: %v", expected)
	if !result.ApproxEqualThreshold(&expected, 1e-4) {
		t.Errorf("Quaternion rotating vector doesn't match slower conjugate method. Got: %v, Expected: %v", result, expected)
	}

	expected = Vec3{-1, 0, 0}
	if !result.ApproxEqualThreshold(&expected, 4e-4) { // The result we get for z is like 8e-8, but a 1e-4 threshold juuuuuust causes it to freak out when compared to 0.0
		t.Errorf("Quaternion rotating vector doesn't match hand-computed result. Got: %v, Expected: %v", result, expected)
	}
}

func BenchmarkQuatRotateOptimized(b *testing.B) {
	b.StopTimer()
	rand := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		q := QuatRotate(rand.Float32(), &Vec3{rand.Float32(), rand.Float32(), rand.Float32()})
		v := &Vec3{rand.Float32(), rand.Float32(), rand.Float32()}
		q.Normalized()
		b.StartTimer()

		s := q.Rotate(v)
		_ = s
	}
}

func BenchmarkQuatRotateConjugate(b *testing.B) {
	b.StopTimer()
	rand := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		q := QuatRotate(rand.Float32(), &Vec3{rand.Float32(), rand.Float32(), rand.Float32()})
		v := Vec3{rand.Float32(), rand.Float32(), rand.Float32()}
		q.Normalized()
		b.StartTimer()

		m := q.Mul(&Quat{0, v})
		c := q.Conjugated()

		v = m.Mul(&c).V
	}
}

func BenchmarkQuatArrayAccess(b *testing.B) {
	b.StopTimer()
	rand := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		q := QuatRotate(rand.Float32(), &Vec3{rand.Float32(), rand.Float32(), rand.Float32()})
		b.StartTimer()

		_ = q.V[0]
	}
}

func BenchmarkQuatFuncElementAccess(b *testing.B) {
	b.StopTimer()
	rand := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		q := QuatRotate(rand.Float32(), &Vec3{rand.Float32(), rand.Float32(), rand.Float32()})
		b.StartTimer()

		_ = q.X()
	}
}

func TestMat4ToQuat(t *testing.T) {
	// http://www.euclideanspace.com/maths/geometry/rotations/conversions/matrixToQuaternion/examples/index.htm

	iden := Ident4()
	qiden := QuatIdent()
	tests := []struct {
		Description string
		Rotation    *Mat4
		Expected    *Quat
	}{
		{
			"forward",
			&iden,
			&qiden,
		},
		{
			"heading 90 degree",
			&Mat4{
				0, 0, -1, 0,
				0, 1, 0, 0,
				1, 0, 0, 0,
				0, 0, 0, 1,
			},
			&Quat{0.7071, Vec3{0, 0.7071, 0}},
		},
		{
			"heading 180 degree",
			&Mat4{
				-1, 0, 0, 0,
				0, 1, 0, 0,
				0, 0, -1, 0,
				0, 0, 0, 1,
			},
			&Quat{0, Vec3{0, 1, 0}},
		},
		{
			"attitude 90 degree",
			&Mat4{
				0, 1, 0, 0,
				-1, 0, 0, 0,
				0, 0, 1, 0,
				0, 0, 0, 1,
			},
			&Quat{0.7071, Vec3{0, 0, 0.7071}},
		},
		{
			"bank 90 degree",
			&Mat4{
				1, 0, 0, 0,
				0, 0, 1, 0,
				0, -1, 0, 0,
				0, 0, 0, 1,
			},
			&Quat{0.7071, Vec3{0.7071, 0, 0}},
		},
	}

	threshold := float32(math.Pow(10, -2))
	for _, c := range tests {
		if r := Mat4ToQuat(c.Rotation); !r.ApproxEqualThreshold(c.Expected, threshold) {
			t.Errorf("%v failed: Mat4ToQuat(%v) != %v (got %v)", c.Description, c.Rotation, c.Expected, r)
		}
	}
}

func TestQuatRotate(t *testing.T) {
	qiden := QuatIdent()
	tests := []struct {
		Description string
		Angle       float32
		Axis        *Vec3
		Expected    *Quat
	}{
		{
			"forward",
			0, &Vec3{0, 0, 0},
			&qiden,
		},
		{
			"heading 90 degree",
			DegToRad(90), &Vec3{0, 1, 0},
			&Quat{0.7071, Vec3{0, 0.7071, 0}},
		},
		{
			"heading 180 degree",
			DegToRad(180), &Vec3{0, 1, 0},
			&Quat{0, Vec3{0, 1, 0}},
		},
		{
			"attitude 90 degree",
			DegToRad(90), &Vec3{0, 0, 1},
			&Quat{0.7071, Vec3{0, 0, 0.7071}},
		},
		{
			"bank 90 degree",
			DegToRad(90), &Vec3{1, 0, 0},
			&Quat{0.7071, Vec3{0.7071, 0, 0}},
		},
	}

	threshold := float32(math.Pow(10, -2))
	for _, c := range tests {
		if r := QuatRotate(c.Angle, c.Axis); !r.OrientationEqualThreshold(c.Expected, threshold) {
			t.Errorf("%v failed: QuatRotate(%v, %v) != %v (got %v)", c.Description, c.Angle, c.Axis, c.Expected, r)
		}
	}
}

func TestQuatLookAtV(t *testing.T) {
	// http://www.euclideanspace.com/maths/algebra/realNormedAlgebra/quaternions/transforms/examples/index.htm
	qiden := QuatIdent()
	tests := []struct {
		Description     string
		Eye, Center, Up *Vec3
		Expected        *Quat
	}{
		{
			"forward",
			&Vec3{0, 0, 0},
			&Vec3{0, 0, -1},
			&Vec3{0, 1, 0},
			&qiden,
		},
		{
			"heading 90 degree",
			&Vec3{0, 0, 0},
			&Vec3{1, 0, 0},
			&Vec3{0, 1, 0},
			&Quat{0.7071, Vec3{0, 0.7071, 0}},
		},
		{
			"heading 180 degree",
			&Vec3{0, 0, 0},
			&Vec3{0, 0, 1},
			&Vec3{0, 1, 0},
			&Quat{0, Vec3{0, 1, 0}},
		},
		{
			"attitude 90 degree",
			&Vec3{0, 0, 0},
			&Vec3{0, 0, -1},
			&Vec3{1, 0, 0},
			&Quat{0.7071, Vec3{0, 0, 0.7071}},
		},
		{
			"bank 90 degree",
			&Vec3{0, 0, 0},
			&Vec3{0, -1, 0},
			&Vec3{0, 0, -1},
			&Quat{0.7071, Vec3{0.7071, 0, 0}},
		},
	}

	threshold := float32(math.Pow(10, -2))
	for _, c := range tests {
		if r := QuatLookAtV(c.Eye, c.Center, c.Up); !r.OrientationEqualThreshold(c.Expected, threshold) {
			t.Errorf("%v failed: QuatLookAtV(%v, %v, %v) != %v (got %v)", c.Description, c.Eye, c.Center, c.Up, c.Expected, r)
		}
	}
}

func TestCompareLookAt(t *testing.T) {
	type OrigExp [2]*Vec3

	tests := []struct {
		Description     string
		Eye, Center, Up *Vec3
		Pos             []OrigExp
	}{
		{
			"forward, identity rotation",
			// looking from viewer into screen z-, up y+
			&Vec3{0, 0, 0}, &Vec3{0, 0, -1}, &Vec3{0, 1, 0},
			[]OrigExp{
				{&Vec3{1, 2, 3}, &Vec3{1, 2, 3}},
			},
		},
		{
			"heading -90 degree, look right",
			// look x+
			// rotate around y -90 deg
			&Vec3{0, 0, 0}, &Vec3{1, 0, 0}, &Vec3{0, 1, 0},
			[]OrigExp{
				{&Vec3{1, 2, 3}, &Vec3{3, 2, -1}},

				{&Vec3{1, 1, -1}, &Vec3{-1, 1, -1}},
				{&Vec3{1, 1, 1}, &Vec3{1, 1, -1}},
				{&Vec3{1, -1, 1}, &Vec3{1, -1, -1}},
				{&Vec3{1, -1, -1}, &Vec3{-1, -1, -1}},

				{&Vec3{-1, 1, -1}, &Vec3{-1, 1, 1}},
				{&Vec3{-1, 1, 1}, &Vec3{1, 1, 1}},
				{&Vec3{-1, -1, 1}, &Vec3{1, -1, 1}},
				{&Vec3{-1, -1, -1}, &Vec3{-1, -1, 1}},
			},
		},
		{
			"heading 180 degree",
			&Vec3{0, 0, 0}, &Vec3{0, 0, 1}, &Vec3{0, 1, 0},
			[]OrigExp{
				{&Vec3{1, 2, 3}, &Vec3{-1, 2, -3}},
			},
		},
		{
			"attitude 90 degree",
			&Vec3{0, 0, 0}, &Vec3{0, 0, -1}, &Vec3{1, 0, 0},
			[]OrigExp{
				{&Vec3{1, 2, 3}, &Vec3{-2, 1, 3}},
			},
		},
		{
			"bank 90 degree, look down",
			// look y-
			// rotate around x -90 deg
			// up toward z-
			&Vec3{0, 0, 0}, &Vec3{0, -1, 0}, &Vec3{0, 0, -1},
			[]OrigExp{
				{&Vec3{1, 2, 3}, &Vec3{1, -3, 2}},

				{&Vec3{1, 1, -1}, &Vec3{1, 1, 1}},
				{&Vec3{1, 1, 1}, &Vec3{1, -1, 1}},
				{&Vec3{1, -1, 1}, &Vec3{1, -1, -1}},
				{&Vec3{1, -1, -1}, &Vec3{1, 1, -1}},

				{&Vec3{-1, 1, -1}, &Vec3{-1, 1, 1}},
				{&Vec3{-1, 1, 1}, &Vec3{-1, -1, 1}},
				{&Vec3{-1, -1, 1}, &Vec3{-1, -1, -1}},
				{&Vec3{-1, -1, -1}, &Vec3{-1, 1, -1}},
			},
		},
		{
			"half roll",
			// immelmann turn without the half roll
			// looking from screen to viewer z+
			// upside down, y-
			&Vec3{0, 0, 0}, &Vec3{0, 0, 1}, &Vec3{0, -1, 0},
			[]OrigExp{
				{&Vec3{1, 1, -1}, &Vec3{1, -1, 1}},
				{&Vec3{1, 1, 1}, &Vec3{1, -1, -1}},
				{&Vec3{1, -1, 1}, &Vec3{1, 1, -1}},
				{&Vec3{1, -1, -1}, &Vec3{1, 1, 1}},

				{&Vec3{-1, 1, -1}, &Vec3{-1, -1, 1}},
				{&Vec3{-1, 1, 1}, &Vec3{-1, -1, -1}},
				{&Vec3{-1, -1, 1}, &Vec3{-1, 1, -1}},
				{&Vec3{-1, -1, -1}, &Vec3{-1, 1, 1}},
			},
		},
		{
			"roll left",
			// look x-
			// rotate around y 90 deg
			// up toward viewer z+
			&Vec3{0, 0, 0}, &Vec3{-1, 0, 0}, &Vec3{0, 0, 1},
			[]OrigExp{
				{&Vec3{1, 1, -1}, &Vec3{1, -1, 1}},
				{&Vec3{1, 1, 1}, &Vec3{1, 1, 1}},
				{&Vec3{1, -1, 1}, &Vec3{-1, 1, 1}},
				{&Vec3{1, -1, -1}, &Vec3{-1, -1, 1}},

				{&Vec3{-1, 1, -1}, &Vec3{1, -1, -1}},
				{&Vec3{-1, 1, 1}, &Vec3{1, 1, -1}},
				{&Vec3{-1, -1, 1}, &Vec3{-1, 1, -1}},
				{&Vec3{-1, -1, -1}, &Vec3{-1, -1, -1}},
			},
		},
	}

	threshold := float32(math.Pow(10, -2))
	for _, c := range tests {
		m := LookAtV(c.Eye, c.Center, c.Up)
		q := QuatLookAtV(c.Eye, c.Center, c.Up)

		for i, p := range c.Pos {
			t.Log(c.Description, i)
			o, e := p[0], p[1]
			v4 := o.Vec4(0)
			mv4 := m.Mul4x1(&v4)
			rm := mv4.Vec3()
			rq := q.Rotate(o)

			if !rq.ApproxEqualThreshold(&rm, threshold) {
				t.Errorf("%v failed: QuatLookAtV() != LookAtV()", c.Description)
			}

			if !e.ApproxEqualThreshold(&rm, threshold) {
				t.Errorf("%v failed: (%v).Mul4x1(%v) != %v (got %v)", c.Description, m, o, e, rm)
			}

			if !e.ApproxEqualThreshold(&rq, threshold) {
				t.Errorf("%v failed: (%v).Rotate(%v) != %v (got %v)", c.Description, q, o, e, rq)
			}
		}
	}
}

func TestQuatMatConversion(t *testing.T) {
	tests := []struct {
		Angle float32
		Axis  *Vec3
	}{}

	for a := 0.0; a <= math.Pi*2; a += math.Pi / 4.0 {
		af := float32(a)
		tests = append(tests, []struct {
			Angle float32
			Axis  *Vec3
		}{
			{af, &Vec3{1, 0, 0}},
			{af, &Vec3{0, 1, 0}},
			{af, &Vec3{0, 0, 1}},
		}...)
	}

	for _, c := range tests {
		m1 := HomogRotate3D(c.Angle, c.Axis)
		q1 := Mat4ToQuat(&m1)
		q2 := QuatRotate(c.Angle, c.Axis)

		if !FloatEqualThreshold(Abs(q1.Dot(&q2)), 1, 1e-4) {
			t.Errorf("Quaternions for %v %v do not match:\n%v\n%v", RadToDeg(c.Angle), c.Axis, q1, q2)
		}
	}
}

func TestQuatGetter(t *testing.T) {
	tests := []*Quat{
		&Quat{0, Vec3{0, 0, 0}},
		&Quat{1, Vec3{2, 3, 4}},
		&Quat{-4, Vec3{-3, -2, -1}},
	}

	for _, q := range tests {
		if r := q.X(); !FloatEqualThreshold(r, q.V[0], 1e-4) {
			t.Errorf("Quat(%v).X() != %v (got %v)", q, q.V[0], r)
		}

		if r := q.Y(); !FloatEqualThreshold(r, q.V[1], 1e-4) {
			t.Errorf("Quat(%v).Y() != %v (got %v)", q, q.V[1], r)
		}

		if r := q.Z(); !FloatEqualThreshold(r, q.V[2], 1e-4) {
			t.Errorf("Quat(%v).Z() != %v (got %v)", q, q.V[2], r)
		}
	}
}

func TestQuatEqual(t *testing.T) {
	tests := []struct {
		A, B     *Quat
		Expected bool
	}{
		{&Quat{1, Vec3{0, 0, 0}}, &Quat{1, Vec3{0, 0, 0}}, true},
		{&Quat{1, Vec3{2, 3, 4}}, &Quat{1, Vec3{2, 3, 4}}, true},
		{&Quat{0.0000000000001, Vec3{0, 0, 0}}, &Quat{0, Vec3{0, 0, 0}}, true},
		{&Quat{MaxValue, Vec3{1, 0, 0}}, &Quat{MaxValue, Vec3{1, 0, 0}}, true},
		{&Quat{0, Vec3{0, 1, 0}}, &Quat{1, Vec3{0, 0, 0}}, false},
		{&Quat{1, Vec3{2, 3, 0}}, &Quat{-4, Vec3{5, 6, 0}}, false},
	}

	for _, c := range tests {
		if r := c.A.ApproxEqualThreshold(c.B, 1e-4); r != c.Expected {
			t.Errorf("Quat(%v).ApproxEqualThreshold(Quat(%v), 1e-4) != %v (got %v)", c.A, c.B, c.Expected, r)
		}
	}
}

func TestQuatOrientationEqual(t *testing.T) {
	tests := []struct {
		A, B     *Quat
		Expected bool
	}{
		{&Quat{1, Vec3{0, 0, 0}}, &Quat{1, Vec3{0, 0, 0}}, true},
		{&Quat{0, Vec3{0, 1, 0}}, &Quat{0, Vec3{0, -1, 0}}, true},
		{&Quat{0, Vec3{0, 1, 0}}, &Quat{1, Vec3{0, 0, 0}}, false},
		{&Quat{1, Vec3{2, 3, 0}}, &Quat{-4, Vec3{5, 6, 0}}, false},
	}

	for _, c := range tests {
		if r := c.A.OrientationEqualThreshold(c.B, 1e-4); r != c.Expected {
			t.Errorf("Quat(%v).OrientationEqualThreshold(Quat(%v), 1e-4) != %v (got %v)", c.A, c.B, c.Expected, r)
		}
	}
}

func TestQuatAdd(t *testing.T) {
	tests := []struct {
		A, B     *Quat
		Expected *Quat
	}{
		{&Quat{0, Vec3{0, 0, 0}}, &Quat{0, Vec3{0, 0, 0}}, &Quat{0, Vec3{0, 0, 0}}},
		{&Quat{1, Vec3{0, 0, 0}}, &Quat{1, Vec3{0, 0, 0}}, &Quat{2, Vec3{0, 0, 0}}},
		{&Quat{1, Vec3{2, 3, 4}}, &Quat{5, Vec3{6, 7, 8}}, &Quat{6, Vec3{8, 10, 12}}},
	}

	for _, c := range tests {
		if r := c.A.Add(c.B); !r.ApproxEqualThreshold(c.Expected, 1e-4) {
			t.Errorf("Quat(%v).Add(Quat(%v)) != %v (got %v)", c.A, c.B, c.Expected, r)
		}
	}
}

func TestQuatSub(t *testing.T) {
	tests := []struct {
		A, B     *Quat
		Expected *Quat
	}{
		{&Quat{0, Vec3{0, 0, 0}}, &Quat{0, Vec3{0, 0, 0}}, &Quat{0, Vec3{0, 0, 0}}},
		{&Quat{1, Vec3{0, 0, 0}}, &Quat{1, Vec3{0, 0, 0}}, &Quat{0, Vec3{0, 0, 0}}},
		{&Quat{1, Vec3{2, 3, 4}}, &Quat{5, Vec3{6, 7, 8}}, &Quat{-4, Vec3{-4, -4, -4}}},
	}

	for _, c := range tests {
		if r := c.A.Sub(c.B); !r.ApproxEqualThreshold(c.Expected, 1e-4) {
			t.Errorf("Quat(%v).Sub(Quat(%v)) != %v (got %v)", c.A, c.B, c.Expected, r)
		}
	}
}

func TestQuatScale(t *testing.T) {
	tests := []struct {
		Rotation *Quat
		Scalar   float32
		Expected *Quat
	}{
		{&Quat{0, Vec3{0, 0, 0}}, 1, &Quat{0, Vec3{0, 0, 0}}},
		{&Quat{1, Vec3{0, 0, 0}}, 2, &Quat{2, Vec3{0, 0, 0}}},
		{&Quat{1, Vec3{2, 3, 4}}, 3, &Quat{3, Vec3{6, 9, 12}}},
	}

	for _, c := range tests {
		if r := c.Rotation.Scale(c.Scalar); !r.ApproxEqualThreshold(c.Expected, 1e-4) {
			t.Errorf("Quat(%v).Scale(%v) != %v (got %v)", c.Rotation, c.Scalar, c.Expected, r)
		}
	}
}

func TestQuatLen(t *testing.T) {
	tests := []struct {
		Rotation Quat
		Expected float32
	}{
		{Quat{0, Vec3{1, 0, 0}}, 1},
		{Quat{0, Vec3{0.0000000000001, 0, 0}}, 0},
		{Quat{0, Vec3{MaxValue, 1, 0}}, InfPos},
		{Quat{4, Vec3{1, 2, 3}}, float32(math.Sqrt(1*1 + 2*2 + 3*3 + 4*4))},
		{Quat{0, Vec3{3.1, 4.2, 1.3}}, float32(math.Sqrt(3.1*3.1 + 4.2*4.2 + 1.3*1.3))},
	}

	for _, c := range tests {
		if r := c.Rotation.Len(); !FloatEqualThreshold(c.Expected, r, 1e-4) {
			t.Errorf("Quat(%v).Len() != %v (got %v)", c.Rotation, c.Expected, r)
		}

		if !FloatEqualThreshold(c.Rotation.Len(), c.Rotation.Norm(), 1e-4) {
			t.Error("Quat().Len() != Quat().Norm()")
		}
	}
}

func TestQuatNormalize(t *testing.T) {
	tests := []struct {
		Rotation *Quat
		Expected *Quat
	}{
		{&Quat{0, Vec3{0, 0, 0}}, &Quat{1, Vec3{0, 0, 0}}},
		{&Quat{0, Vec3{1, 0, 0}}, &Quat{0, Vec3{1, 0, 0}}},
		{&Quat{0, Vec3{0.0000000000001, 0, 0}}, &Quat{0, Vec3{1, 0, 0}}},
		{&Quat{0, Vec3{MaxValue, 1, 0}}, &Quat{0, Vec3{1, 0, 0}}},
		{&Quat{4, Vec3{1, 2, 3}}, &Quat{4.0 / 5.477, Vec3{1.0 / 5.477, 2.0 / 5.477, 3.0 / 5.477}}},
		{&Quat{0, Vec3{3.1, 4.2, 1.3}}, &Quat{0, Vec3{3.1 / 5.3795, 4.2 / 5.3795, 1.3 / 5.3795}}},
	}

	for _, c := range tests {
		if r := c.Rotation.Normalized(); !r.ApproxEqualThreshold(c.Expected, 1e-4) {
			t.Errorf("Quat(%v).Normalize() != %v (got %v)", c.Rotation, c.Expected, r)
		}
	}
}

func TestQuatInverse(t *testing.T) {
	tests := []struct {
		Rotation *Quat
		Expected *Quat
	}{
		{&Quat{0, Vec3{1, 0, 0}}, &Quat{0, Vec3{-1, 0, 0}}},
		{&Quat{3, Vec3{-1, 4, 3}}, &Quat{3.0 / 35.0, Vec3{1.0 / 35.0, -4.0 / 35.0, -3.0 / 35.0}}},
		{&Quat{1, Vec3{0, 0, 2}}, &Quat{1.0 / 5.0, Vec3{0, 0, -2.0 / 5.0}}},
	}

	for _, c := range tests {
		if r := c.Rotation.Inverse(); !r.ApproxEqualThreshold(c.Expected, 1e-4) {
			t.Errorf("Quat(%v).Inverse() != %v (got %v)", c.Rotation, c.Expected, r)
		}
	}
}

func TestQuatSlerp(t *testing.T) {
	tests := []struct {
		A, B     *Quat
		Scalar   float32
		Expected *Quat
	}{
		{&Quat{0, Vec3{0, 0, 0}}, &Quat{0, Vec3{0, 0, 0}}, 0, &Quat{1, Vec3{0, 0, 0}}},
		{&Quat{0, Vec3{1, 0, 0}}, &Quat{0, Vec3{1, 0, 0}}, 0.5, &Quat{0, Vec3{1, 0, 0}}},
		{&Quat{1, Vec3{0, 0, 0}}, &Quat{0, Vec3{1, 0, 0}}, 0.5, &Quat{0.7071067811865475, Vec3{0.7071067811865475, 0, 0}}},
		{&Quat{0.5, Vec3{-0.5, -0.5, 0.5}}, &Quat{0.996, Vec3{-0.080, -0.080, 0}}, 1, &Quat{0.996, Vec3{-0.080, -0.080, 0}}},
		{&Quat{0.5, Vec3{-0.5, -0.5, 0.5}}, &Quat{0.996, Vec3{-0.080, -0.080, 0}}, 0, &Quat{0.5, Vec3{-0.5, -0.5, 0.5}}},
		{&Quat{0.5, Vec3{-0.5, -0.5, 0.5}}, &Quat{0.996, Vec3{-0.080, -0.080, 0}}, 0.2, &Quat{0.6553097459373098, Vec3{-0.44231939784548874, -0.44231939784548874, 0.4237176207195655}}},
		{&Quat{0.996, Vec3{-0.080, -0.080, 0}}, &Quat{0.5, Vec3{-0.5, -0.5, 0.5}}, 0.8, &Quat{0.6553097459373098, Vec3{-0.44231939784548874, -0.44231939784548874, 0.4237176207195655}}},
		{&Quat{1, Vec3{0, 0, 0}}, &Quat{-0.9999999, Vec3{0, 0, 0}}, 0, &Quat{1, Vec3{0, 0, 0}}},
	}

	for _, c := range tests {
		if r := QuatSlerp(c.A, c.B, c.Scalar); !r.ApproxEqualThreshold(c.Expected, 1e-2) {
			t.Errorf("QuatSlerp(%v, %v, %v) != %v (got %v)", c.A, c.B, c.Scalar, c.Expected, r)
		}
	}
}

func TestQuatDot(t *testing.T) {
	tests := []struct {
		A, B     *Quat
		Expected float32
	}{
		{&Quat{0, Vec3{0, 0, 0}}, &Quat{0, Vec3{0, 0, 0}}, 0},
		{&Quat{0, Vec3{1, 2, 3}}, &Quat{0, Vec3{4, 5, 6}}, 32},
		{&Quat{4, Vec3{1, 2, 3}}, &Quat{8, Vec3{5, 6, 7}}, 70},
	}

	for _, c := range tests {
		if r := c.A.Dot(c.B); !FloatEqualThreshold(r, c.Expected, 1e-4) {
			t.Errorf("Quat(%v).Dot(Quat(%v)) != %v (got %v)", c.A, c.B, c.Expected, r)
		}
	}
}

func TestApproxEqual(t *testing.T) {
	q1 := Quat{1, Vec3{2, 3, 4}}
	q2 := Quat{1, Vec3{2, 3, 4}}
	if !q1.ApproxEqual(&q2) {
		t.Errorf("quaternion should be equal %+v, %+v", q1, q2)
	}
	q2 = Quat{2, Vec3{6, 2, 5}}
	if q1.ApproxEqual(&q2) {
		t.Errorf("quaternion shouldnt be equal %+v, %+v", q1, q2)
	}
}

func TestQuatBetweenVector3(t *testing.T) {
	v1 := Vec3{1, 0, 0}
	v2 := Vec3{-1, 0, 0}
	QuatBetweenVectors(&v1, &v2)
}

func TestQuatLerp(t *testing.T) {
	tests := []struct {
		A, B     Quat
		Amount   float32
		Expected Quat
	}{
		{Quat{0, Vec3{0, 0, 0}}, Quat{0, Vec3{0, 0, 0}}, 0, Quat{0, Vec3{0, 0, 0}}},
		{Quat{0, Vec3{1, 2, 3}}, Quat{0, Vec3{4, 5, 6}}, 0.5, Quat{0, Vec3{2.5, 3.5, 4.5}}},
		{Quat{4, Vec3{1, 2, 3}}, Quat{8, Vec3{5, 6, 7}}, 0.75, Quat{7, Vec3{4, 5, 6}}},
	}

	for _, c := range tests {
		if r := QuatLerp(&c.A, &c.B, c.Amount); !r.ApproxEqualThreshold(&c.Expected, 1e-4) {
			t.Errorf("QuatLerp(Quat(%v), (Quat(%v))) != %v (got %v)", c.A, c.B, c.Expected, r)
		}
	}
}

func TestQuatBetweenVectors(t *testing.T) {
	tests := []struct {
		A, B     Vec3
		Expected Quat
	}{
		{Vec3{0, 0, 1}, Vec3{1, 1, 0}, Quat{0.70710677, Vec3{-0.49999997, 0.49999997, 0}}},
		{Vec3{1, 2, 3}, Vec3{4, 5, 6}, Quat{0.9936377, Vec3{-0.04597839, 0.09195679, -0.045978405}}},
		{Vec3{1, 2, 3}, Vec3{5, 6, 7}, Quat{0.99205077, Vec3{-0.051373072, 0.10274618, -0.0513731}}},
	}

	for _, c := range tests {
		if r := QuatBetweenVectors(&c.A, &c.B); !r.ApproxEqualThreshold(&c.Expected, 1e-4) {
			t.Errorf("QuatBetweenVectors(Vec3(%v), (Vec3(%v))) != %v (got %v)", c.A, c.B, c.Expected, r)
		}
	}
}

func TestQuat_Ident(t *testing.T) {
	ident := Quat{W: 1, V: Vec3{0, 0, 0}}
	if ident != QuatIdent() {
		t.Errorf("QuatIdent = %v, want %v", QuatIdent(), ident)
	}

	var q Quat
	q.Iden()

	if ident != q {
		t.Errorf("q.Iden = %v, want %v", q, ident)
	}
}

var quatTests = []struct {
	q1, q2, add, sub, mul, scale, conj, normal, inv, svec Quat
	f                                                     float32
	v1                                                    Vec3
	mat3                                                  Mat3
}{
	{
		q1:     Quat{W: 1, V: Vec3{2, 3, 4}},
		q2:     Quat{W: 1, V: Vec3{2, 3, 4}},
		add:    Quat{W: 2, V: Vec3{4, 6, 8}},
		sub:    Quat{W: 0, V: Vec3{0, 0, 0}},
		mul:    Quat{W: -28, V: Vec3{4, 6, 8}},
		scale:  Quat{W: 2, V: Vec3{4, 6, 8}},
		conj:   Quat{W: 1, V: Vec3{-2, -3, -4}},
		normal: Quat{W: float32(1.0 / math.Sqrt(30.0)), V: Vec3{float32(math.Sqrt(2.0 / 15.0)), float32(math.Sqrt(3.0 / 10.0)), float32(2.0 * math.Sqrt(2.0/15.0))}},
		inv:    Quat{W: 1.0 / 30.0, V: Vec3{-1.0 / 15.0, -1.0 / 10.0, -2.0 / 15.0}},
		svec:   Quat{W: -15, V: Vec3{10, -5, 10}},
		v1:     Vec3{3, 2, 1},
		mat3: Mat3{-2.0 / 3.0, 2.0 / 3.0, 1.0 / 3.0,
			2.0 / 15.0, -1.0 / 3.0, 14.0 / 15.0,
			11.0 / 15.0, 2.0 / 3.0, 2.0 / 15.0},
		f: 2,
	},
	{
		q1:     Quat{W: 5, V: Vec3{6, 7, 8}},
		q2:     Quat{W: 3, V: Vec3{4, 5, 6}},
		add:    Quat{W: 8, V: Vec3{10, 12, 14}},
		sub:    Quat{W: 2, V: Vec3{2, 2, 2}},
		mul:    Quat{W: -92, V: Vec3{40, 42, 56}},
		scale:  Quat{W: 2.5, V: Vec3{3, 3.5, 4}},
		conj:   Quat{W: 5, V: Vec3{-6, -7, -8}},
		normal: Quat{W: float32(5.0 / math.Sqrt(174.0)), V: Vec3{float32(math.Sqrt(6.0 / 29.0)), float32(7.0 / math.Sqrt(174.0)), float32(4.0 * math.Sqrt(2.0/87.0))}},
		inv:    Quat{W: 5.0 / 174.0, V: Vec3{-1.0 / 29.0, -7.0 / 174.0, -4.0 / 87.0}},
		svec:   Quat{W: -41.75, V: Vec3{22.5, 10.25, 22}},
		mat3: Mat3{-26.0 / 87.0, 82.0 / 87.0, 13.0 / 87.0,
			2.0 / 87.0, -13.0 / 87.0, 86.0 / 87.0,
			83.0 / 87.0, 26.0 / 87.0, 2.0 / 87.0},
		f:  0.5,
		v1: Vec3{10, 9, 8},
	},
}

func TestQuat_AddOf(t *testing.T) {
	for i, test := range quatTests {
		var q Quat
		q.AddOf(&test.q1, &test.q2)
		if !q.ApproxEqualThreshold(&test.add, 1e-4) {
			t.Errorf("[%d] q1 + q2 = %v, want %v", i, q, test.add)
		}
	}
}

func TestQuat_AddWith(t *testing.T) {
	for i, test := range quatTests {
		q := test.q1
		q.AddWith(&test.q2)
		if !q.ApproxEqualThreshold(&test.add, 1e-4) {
			t.Errorf("[%d] q1 + q2 = %v, want %v", i, q, test.add)
		}
	}
}

func TestQuat_SubOf(t *testing.T) {
	for i, test := range quatTests {
		var q Quat
		q.SubOf(&test.q1, &test.q2)
		if !q.ApproxEqualThreshold(&test.sub, 1e-4) {
			t.Errorf("[%d] q1 - q2 = %v, want %v", i, q, test.sub)
		}
	}
}

func TestQuat_SubWith(t *testing.T) {
	for i, test := range quatTests {
		q := test.q1
		q.SubWith(&test.q2)
		if !q.ApproxEqualThreshold(&test.sub, 1e-4) {
			t.Errorf("[%d] q1 - q2 = %v, want %v", i, q, test.sub)
		}
	}
}

func TestQuat_Mul(t *testing.T) {
	for i, test := range quatTests {
		q := test.q1.Mul(&test.q2)
		if !q.ApproxEqualThreshold(&test.mul, 1e-4) {
			t.Errorf("[%d] q1 * q2 = %v, want %v", i, q, test.mul)
		}
	}
}

func TestQuat_MulOf(t *testing.T) {
	for i, test := range quatTests {
		var q Quat
		q.MulOf(&test.q1, &test.q2)
		if !q.ApproxEqualThreshold(&test.mul, 1e-4) {
			t.Errorf("[%d] q1 * q2 = %v, want %v", i, q, test.mul)
		}
	}
}

func TestQuat_MulWith(t *testing.T) {
	for i, test := range quatTests {
		q := test.q1
		q.MulWith(&test.q2)
		if !q.ApproxEqualThreshold(&test.mul, 1e-4) {
			t.Errorf("[%d] q1 * q2 = %v, want %v", i, q, test.mul)
		}
	}
}

func TestQuat_Scale(t *testing.T) {
	for i, test := range quatTests {
		q := test.q1.Scale(test.f)
		if !q.ApproxEqualThreshold(&test.scale, 1e-4) {
			t.Errorf("[%d] q1 * f = %v, want %v", i, q, test.scale)
		}
	}
}

func TestQuat_ScaleOf(t *testing.T) {
	for i, test := range quatTests {
		var q Quat
		q.ScaleOf(test.f, &test.q1)
		if !q.ApproxEqualThreshold(&test.scale, 1e-4) {
			t.Errorf("[%d] q1 * f = %v, want %v", i, q, test.scale)
		}
	}
}

func TestQuat_ScaleWith(t *testing.T) {
	for i, test := range quatTests {
		q := test.q1
		q.ScaleWith(test.f)
		if !q.ApproxEqualThreshold(&test.scale, 1e-4) {
			t.Errorf("[%d] q1 * f = %v, want %v", i, q, test.scale)
		}
	}
}

func TestQuat_Conjugated(t *testing.T) {
	for i, test := range quatTests {
		q := test.q1.Conjugated()
		if !q.ApproxEqualThreshold(&test.conj, 1e-4) {
			t.Errorf("[%d] conj(q1) = %v, want %v", i, q, test.conj)
		}
	}
}

func TestQuat_Conjugate(t *testing.T) {
	for i, test := range quatTests {
		q := test.q1
		q.Conjugate()
		if !q.ApproxEqualThreshold(&test.conj, 1e-4) {
			t.Errorf("[%d] conj(q1) = %v, want %v", i, q, test.conj)
		}
	}
}

func TestQuat_ConjugateOf(t *testing.T) {
	for i, test := range quatTests {
		var q Quat
		q.ConjugateOf(&test.q1)
		if !q.ApproxEqualThreshold(&test.conj, 1e-4) {
			t.Errorf("[%d] conj(q1) = %v, want %v", i, q, test.conj)
		}
	}
}

func TestQuat_Normalized(t *testing.T) {
	for i, test := range quatTests {
		q := test.q1.Normalized()
		if !q.ApproxEqualThreshold(&test.normal, 1e-4) {
			t.Errorf("[%d] unit(q1) = %v, want %v", i, q, test.normal)
		}
	}
}

func TestQuat_Normalize(t *testing.T) {
	for i, test := range quatTests {
		q := test.q1
		q.Normalize()
		if !q.ApproxEqualThreshold(&test.normal, 1e-4) {
			t.Errorf("[%d] unit(q1) = %v, want %v", i, q, test.normal)
		}
	}
}

func TestQuat_SetNormalizeOf(t *testing.T) {
	for i, test := range quatTests {
		var q Quat
		q.SetNormalizedOf(&test.q1)
		if !q.ApproxEqualThreshold(&test.normal, 1e-4) {
			t.Errorf("[%d] unit(q1) = %v, want %v", i, q, test.normal)
		}
	}
}

func TestQuat_Inverse(t *testing.T) {
	for i, test := range quatTests {
		q := test.q1.Inverse()
		if !q.ApproxEqualThreshold(&test.inv, 1e-4) {
			t.Errorf("[%d] inv(q1) = %v, want %v", i, q, test.inv)
		}
	}
}

func TestQuat_Invert(t *testing.T) {
	for i, test := range quatTests {
		q := test.q1
		q.Invert()
		if !q.ApproxEqualThreshold(&test.inv, 1e-4) {
			t.Errorf("[%d] inv(q1) = %v, want %v", i, q, test.inv)
		}
	}
}

func TestQuat_InverseOf(t *testing.T) {
	for i, test := range quatTests {
		var q Quat
		q.InverseOf(&test.q1)
		if !q.ApproxEqualThreshold(&test.inv, 1e-4) {
			t.Errorf("[%d] inv(q1) = %v, want %v", i, q, test.inv)
		}
	}
}

func TestQuat_AddScaledVec(t *testing.T) {
	for i, test := range quatTests {
		q := test.q1
		q.AddScaledVec(test.f, &test.v1)
		if !q.ApproxEqualThreshold(&test.svec, 1e-4) {
			t.Errorf("[%d] addscaledvec(q1) = %v, want %v", i, q, test.svec)
		}
	}
}

func TestQuat_Mat3(t *testing.T) {
	for i, test := range quatTests {
		q := test.q1
		q.Normalize()
		m := q.Mat3()
		if !m.ApproxEqualThreshold(&test.mat3, 1e-4) {
			t.Errorf("[%d] mat3(q1) = %v, want %v", i, m, test.mat3)
		}
	}
}

func TestQuat_Mat4(t *testing.T) {
	for i, test := range quatTests {
		q := test.q1
		q.Normalize()
		tmp := q.Mat4()
		m := tmp.Mat3()
		if !m.ApproxEqualThreshold(&test.mat3, 1e-4) {
			t.Errorf("[%d] mat3(q1) = %v, want %v", i, m, test.mat3)
		}
	}
}
