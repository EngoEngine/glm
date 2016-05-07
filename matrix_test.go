package glm

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestMulIdent(t *testing.T) {
	t.Parallel()
	i1 := Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
	i2 := Ident4()
	i3 := Ident4()

	mul := i2.Mul4(&i3)

	if i1 != mul {
		t.Errorf("Multiplication of identities does not yield identity")
	}
}

func TestMat2x3Ident(t *testing.T) {
	i0, i1, i2 := Ident2x3(), Ident2x3(), Ident2x3()
	if im := i0.Mul2x3(&i1); !im.Equal(&i2) {
		t.Errorf("Identity multiplication doesn't yield identity \n%sx\n%s=\n%s", i0.String(), i1.String(), im.String())
	}
}

func TestMat3x4Ident(t *testing.T) {
	i0, i1, i2 := Ident3x4(), Ident3x4(), Ident3x4()
	if im := i0.Mul3x4(&i1); !im.Equal(&i2) {
		t.Errorf("Identity multiplication doesn't yield identity \n%sx\n%s=\n%s", i0.String(), i1.String(), im.String())
	}
}

func TestMatRowsSquare(t *testing.T) {
	t.Parallel()
	v0 := Vec4{1, 2, 3, 4}
	v1 := Vec4{5, 6, 7, 8}
	v2 := Vec4{9, 10, 11, 12}
	v3 := Vec4{13, 14, 15, 16}
	rows := [4]Vec4{v0, v1, v2, v3}
	m1 := Mat4FromRows(&v0, &v1, &v2, &v3)

	t.Logf("4x4 matrix as built from rows: %v", m1)
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			if !FloatEqualThreshold(m1.At(r, c), rows[r][c], 1e-5) {
				t.Errorf("Matrix element at (%d,%d) wrong when built from rows. Got: %f, Expected: %f", r, c, m1.At(r, c), rows[r][c])
			}
		}
	}

	v0, v1, v2, v3 = m1.Rows()
	r2 := [4]Vec4{v0, v1, v2, v3}

	t.Logf("4x4 matrix returned rows: %v", r2)
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			if !FloatEqualThreshold(r2[r][c], rows[r][c], 1e-5) {
				t.Errorf("Matrix element at (%d,%d) wrong when rows are gotten. Got: %f, Expected: %f", r, c, r2[r][c], rows[r][c])
			}
		}
	}
}

func TestMatColsSquare(t *testing.T) {
	t.Parallel()
	cols := [4]Vec4{Vec4{1, 2, 3, 4},
		Vec4{5, 6, 7, 8},
		Vec4{9, 10, 11, 12},
		Vec4{13, 14, 15, 16},
	}
	m1 := Mat4FromCols(&cols[0], &cols[1], &cols[2], &cols[3])

	t.Logf("4x4 matrix as built from cols: %v", m1)
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			if m1.At(r, c) != cols[c][r] {
				t.Errorf("Matrix element at (%d,%d) wrong when built from rows. Got: %f, Expected: %f", r, c, m1.At(r, c), cols[c][r])
			}
		}
	}

	v0, v1, v2, v3 := m1.Cols()
	r2 := [4]Vec4{v0, v1, v2, v3}

	t.Logf("4x4 matrix returned cols: %v", r2)
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			if r2[c][r] != cols[c][r] {
				t.Errorf("Matrix element at (%d,%d) wrong when rows are gotten. Got: %f, Expected: %f", r, c, r2[c][r], cols[c][r])
			}
		}
	}
}

func TestTransposeSquare(t *testing.T) {
	t.Parallel()
	v := [4]Vec4{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	m := Mat4FromCols(&v[0], &v[1], &v[2], &v[3])

	transpose := m.Transposed()

	correct := Mat4FromRows(&v[0], &v[1], &v[2], &v[3])

	if correct != transpose {
		t.Errorf("Transpose not correct. Got: %v, expected: %v", transpose, correct)
	}
}

func TestAtSet(t *testing.T) {
	t.Parallel()
	m := Mat3{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}

	if v := m.At(0, 2); !FloatEqualThreshold(v, 7, 1e-4) {
		t.Errorf("Incorrect value gotten by At: %v, expected %v", v, 3)
	}

	m.Set(0, 2, 9001)

	if v := m.At(0, 2); !FloatEqualThreshold(v, 9001, 1e-4) {
		t.Errorf("Value set by Set not gotten by At: %v, expected %v", v, 9001)
	}

	correctMat := Mat3{
		1, 2, 3,
		4, 5, 6,
		9001, 8, 9,
	}

	if !correctMat.EqualThreshold(&m, 1e-4) {
		t.Errorf("After set, not equal to matrix that should be identical. Got: %v, expected: %v", m, correctMat)
	}
}

func TestDiagTrace(t *testing.T) {
	t.Parallel()
	m := Diag4(&Vec4{1, 2, 3, 4})

	tr := m.Trace()

	if !FloatEqualThreshold(tr, 10, 1e-4) {
		t.Errorf("Trace of matrix seeded with diagonal vector {1,2,3,4} not equal to 10. Got %v", tr)
	}
}

func TestMatAbs(t *testing.T) {
	t.Parallel()
	m := Mat4{1, -3, 4, 5, -6, 8, -9, 10, 0, 1, 6, 2, 357, 3, 436}
	result := Mat4{1, 3, 4, 5, 6, 8, 9, 10, 0, 1, 6, 2, 357, 3, 436}

	m.AbsSelf()

	if !result.EqualThreshold(&m, 1e-6) {
		t.Errorf("Matrix absolute value does not work properly. Got: %v, Expected: %v", m, result)
	}
}

func TestString(t *testing.T) {
	t.Parallel()
	m := Ident4()

	str := fmt.Sprintf(` %[2]f %[1]f %[1]f %[1]f
 %[1]f %[2]f %[1]f %[1]f
 %[1]f %[1]f %[2]f %[1]f
 %[1]f %[1]f %[1]f %[2]f
`, 0.0, 1.0)

	if str != m.String() {
		t.Errorf("Mat string conversion not working got \n%q expected \n%q", m.String(), str)
	}
}

func TestMat2Conv(t *testing.T) {
	t.Parallel()
	m2 := Ident2()

	if m3 := m2.Mat3(); m3 != Ident3() {
		t.Errorf("Mat2 \n%sMat3 \n%s", m2.String(), m3.String())
	}

	if m4 := m2.Mat4(); m4 != Ident4() {
		t.Errorf("Mat2 \n%sMat4 \n%s", m2.String(), m4.String())
	}
}

func TestMat3Conv(t *testing.T) {
	t.Parallel()
	m3 := Ident3()

	if m2 := m3.Mat2(); m2 != Ident2() {
		t.Errorf("Mat3 \n%sMat2 \n%s", m3.String(), m2.String())
	}

	if m4 := m3.Mat4(); m4 != Ident4() {
		t.Errorf("Mat3 \n%sMat4 \n%s", m3.String(), m4.String())
	}

	if m2x3 := m3.Mat2x3(); m2x3 != Ident2x3() {
		t.Errorf("Mat3 \n%sMat2x3 \n%s", m3.String(), m2x3.String())
	}
}

func TestMat4Conv(t *testing.T) {
	t.Parallel()
	m4 := Ident4()

	if m2 := m4.Mat2(); m2 != Ident2() {
		t.Errorf("Mat4 \n%sMat2 \n%s", m4.String(), m2.String())
	}

	if m3 := m4.Mat3(); m3 != Ident3() {
		t.Errorf("Mat4 \n%sMat3 \n%s", m4.String(), m3.String())
	}
}

func TestMat2x3Conv(t *testing.T) {
	t.Parallel()
	m2x3 := Mat2x3{
		1, 2,
		3, 4,
		5, 6,
	}
	m3want := Mat3{
		1, 2, 0,
		3, 4, 0,
		5, 6, 1,
	}
	m2want := Mat2{
		1, 2,
		3, 4,
	}

	var m2 Mat2
	m2x3.Mat2In(&m2)

	if m2 != m2want {
		t.Errorf("Mat2x3 \n%sMat2 \n%s", m2x3.String(), m2.String())
	}

	if m2 := m2x3.Mat2(); m2 != m2want {
		t.Errorf("Mat2x3 \n%sMat2 \n%s", m2x3.String(), m2.String())
	}

	var m3 Mat3
	m2x3.Mat3In(&m3)

	if m3 != m3want {
		t.Errorf("Mat2x3 \n%sMat3 \n%s", m2x3.String(), m3.String())
	}

	if m3 := m2x3.Mat3(); m3 != m3want {
		t.Errorf("Mat2x3 \n%sMat3 \n%s", m2x3.String(), m3.String())
	}
}

func TestMat2SetCol(t *testing.T) {
	t.Parallel()
	m2 := Ident2()
	m2.SetCol(0, &Vec2{2, 2})
	expected := Mat2{2, 2, 0, 1}
	if m2 != expected {
		t.Errorf("unexpected result matrix from Mat2.SetCol, %+v, %+v", m2.String(), expected.String())
	}
}

func TestMat3SetCol(t *testing.T) {
	t.Parallel()
	m3 := Ident3()
	m3.SetCol(0, &Vec3{2, 2, 2})
	expected := Mat3{2, 2, 2, 0, 1, 0, 0, 0, 1}
	if m3 != expected {
		t.Errorf("unexpected result matrix from Mat3.SetCol, %+v, %+v", m3.String(), expected.String())
	}
}

func TestMat4SetCol(t *testing.T) {
	t.Parallel()
	m4 := Ident4()
	m4.SetCol(0, &Vec4{2, 2, 2, 2})
	expected := Mat4{2, 2, 2, 2, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}
	if m4 != expected {
		t.Errorf("unexpected result matrix from Mat4.SetCol, %s, %+v", m4.String(), expected.String())
	}
}

func TestMat2SetRow(t *testing.T) {
	t.Parallel()
	m2 := Ident2()
	m2.SetRow(0, &Vec2{2, 2})
	expected := Mat2{2, 0, 2, 1}
	if m2 != expected {
		t.Errorf("unexpected result matrix from Mat2.SetCol, %+v, %+v", m2.String(), expected.String())
	}
}

func TestMat3SetRow(t *testing.T) {
	t.Parallel()
	m3 := Ident3()
	m3.SetRow(0, &Vec3{2, 2, 2})
	expected := Mat3{2, 0, 0, 2, 1, 0, 2, 0, 1}
	if m3 != expected {
		t.Errorf("unexpected result matrix from Mat3.SetCol, %+v, %+v", m3.String(), expected.String())
	}
}

func TestMat4SetRow(t *testing.T) {
	t.Parallel()
	m4 := Ident4()
	m4.SetRow(0, &Vec4{2, 2, 2, 2})
	expected := Mat4{2, 0, 0, 0, 2, 1, 0, 0, 2, 0, 1, 0, 2, 0, 0, 1}
	if m4 != expected {
		t.Errorf("unexpected result matrix from Mat4.SetCol, %s, %+v", m4.String(), expected.String())
	}
}
func TestMat2Diag2(t *testing.T) {
	t.Parallel()
	m := Ident2()
	diag := m.Diag()
	expected := Vec2{1, 1}
	if diag != expected {
		t.Errorf("Unexpected diagonal %+v,%+v", diag, expected)
	}
}
func TestMat3Diag3(t *testing.T) {
	t.Parallel()
	m := Ident3()
	diag := m.Diag()
	expected := Vec3{1, 1, 1}
	if diag != expected {
		t.Errorf("Unexpected diagonal %+v,%+v", diag, expected)
	}
}
func TestMat4Diag4(t *testing.T) {
	t.Parallel()
	m := Ident4()
	diag := m.Diag()
	expected := Vec4{1, 1, 1, 1}
	if diag != expected {
		t.Errorf("Unexpected diagonal %+v,%+v", diag, expected)
	}
}
func TestMat2Ident(t *testing.T) {
	t.Parallel()
	expected := Mat2{1, 0, 0, 1}
	iden := Ident2()
	if expected != iden {
		t.Errorf("Unexpected identity Mat2 %+v, %+v", iden, expected)
	}
}
func TestMat3Ident(t *testing.T) {
	t.Parallel()
	expected := Mat3{1, 0, 0, 0, 1, 0, 0, 0, 1}
	iden := Ident3()
	if expected != iden {
		t.Errorf("Unexpected identity Mat3 %+v, %+v", iden, expected)
	}
}
func TestMat4Ident(t *testing.T) {
	t.Parallel()
	expected := Mat4{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}
	iden := Ident4()
	if expected != iden {
		t.Errorf("Unexpected identity Mat4 %+v, %+v", iden, expected)
	}
}

func TestDiag2(t *testing.T) {
	t.Parallel()
	vec := &Vec2{1, 1}
	m := Diag2(vec)
	if Ident2() != m {
		t.Errorf("Unexpected Mat2 from Diag2 %+v", m)
	}
}
func TestDiag3(t *testing.T) {
	t.Parallel()
	vec := &Vec3{1, 1, 1}
	m := Diag3(vec)
	if Ident3() != m {
		t.Errorf("Unexpected Mat3 from Diag3 %+v", m)
	}
}
func TestDiag4(t *testing.T) {
	t.Parallel()
	vec := &Vec4{1, 1, 1, 1}
	m := Diag4(vec)
	if Ident4() != m {
		t.Errorf("Unexpected Mat4 from Diag4 %+v", m)
	}
}

func TestMat2FromRow(t *testing.T) {
	t.Parallel()
	m := Mat2FromRows(&Vec2{1, 0}, &Vec2{0, 1})
	if m != Ident2() {
		t.Errorf("Unexpected result from Mat2FromRow %+v", m)
	}
}

func TestMat2FromCols(t *testing.T) {
	from := Mat2FromCols(&Vec2{1, 2}, &Vec2{3, 4})
	m2 := Mat2{
		1, 2,
		3, 4,
	}
	if m2 != from {
		t.Errorf("Mat2FromCols unexpected result \n%swant\n%s", from, m2)
	}
}

var mat2tests = []struct {
	m0, m1, add, sub, cmat, mul, abs, tran, inv Mat2
	c, det, trace                               float32
}{
	{
		m0:    Mat2{0, 0, 0, 0},
		m1:    Mat2{0, 0, 0, 0},
		add:   Mat2{0, 0, 0, 0},
		sub:   Mat2{0, 0, 0, 0},
		cmat:  Mat2{0, 0, 0, 0},
		mul:   Mat2{0, 0, 0, 0},
		abs:   Mat2{0, 0, 0, 0},
		tran:  Mat2{0, 0, 0, 0},
		inv:   Mat2{0, 0, 0, 0},
		c:     1,
		det:   0,
		trace: 0,
	},
	{
		m0:    Mat2{-1, 2, 3, 4},
		m1:    Mat2{4, 3, 2, 1},
		add:   Mat2{3, 5, 5, 5},
		sub:   Mat2{-5, -1, 1, 3},
		cmat:  Mat2{-0.5, 1, 1.5, 2},
		mul:   Mat2{5, 20, 1, 8},
		abs:   Mat2{1, 2, 3, 4},
		tran:  Mat2{-1, 3, 2, 4},
		inv:   Mat2{-2.0 / 5.0, 1.0 / 5.0, 3.0 / 10.0, 1.0 / 10.0},
		c:     0.5,
		det:   -10,
		trace: 3,
	},
	{
		m0:    Mat2{-1, -2, -3, -4},
		m1:    Mat2{4, 3, 2, 1},
		add:   Mat2{3, 1, -1, -3},
		sub:   Mat2{-5, -5, -5, -5},
		cmat:  Mat2{-0.5, -1, -1.5, -2},
		mul:   Mat2{-13, -20, -5, -8},
		abs:   Mat2{1, 2, 3, 4},
		tran:  Mat2{-1, -3, -2, -4},
		inv:   Mat2{2, -1, -1.5, 0.5},
		c:     0.5,
		det:   -2,
		trace: -5,
	},
}

var mat3tests = []struct {
	m0, m1, add, sub, cmat, mul, abs, tran, inv Mat3
	c, det, trace                               float32
}{
	{
		m0:    Mat3{},
		m1:    Mat3{},
		add:   Mat3{},
		sub:   Mat3{},
		cmat:  Mat3{},
		mul:   Mat3{},
		abs:   Mat3{},
		tran:  Mat3{},
		inv:   Mat3{},
		c:     1,
		det:   0,
		trace: 0,
	},
	{
		m0:    Mat3{-1, -2, -3, -4, -5, -6, -7, -8, -9},
		m1:    Mat3{9, 8, 7, 6, 5, 4, 3, 2, 1},
		add:   Mat3{8, 6, 4, 2, 0, -2, -4, -6, -8},
		sub:   Mat3{-10, -10, -10, -10, -10, -10, -10, -10, -10},
		cmat:  Mat3{-1.0 * 0.7, -2.0 * 0.7, -3.0 * 0.7, -4.0 * 0.7, -5.0 * 0.7, -6.0 * 0.7, -7.0 * 0.7, -8.0 * 0.7, -9.0 * 0.7},
		mul:   Mat3{-90, -114, -138, -54, -69, -84, -18, -24, -30},
		abs:   Mat3{1, 2, 3, 4, 5, 6, 7, 8, 9},
		tran:  Mat3{-1, -4, -7, -2, -5, -8, -3, -6, -9},
		inv:   Mat3{},
		c:     0.7,
		det:   0,
		trace: -15,
	},
	{
		m0:    Mat3{-1, -2, -3, 4, -5, -6, -7, -8, -9},
		m1:    Mat3{9, 8, 7, 6, 5, 4, 3, 2, 1},
		add:   Mat3{8, 6, 4, 10, 0, -2, -4, -6, -8},
		sub:   Mat3{-10, -10, -10, -2, -10, -10, -10, -10, -10},
		cmat:  Mat3{-1.0 * 0.7, -2.0 * 0.7, -3.0 * 0.7, 4.0 * 0.7, -5.0 * 0.7, -6.0 * 0.7, -7.0 * 0.7, -8.0 * 0.7, -9.0 * 0.7},
		mul:   Mat3{-26, -114, -138, -14, -69, -84, -2, -24, -30},
		abs:   Mat3{1, 2, 3, 4, 5, 6, 7, 8, 9},
		tran:  Mat3{-1, 4, -7, -2, -5, -8, -3, -6, -9},
		inv:   Mat3{-1.0 / 16.0, 1.0 / 8.0, -1.0 / 16.0, 13.0 / 8.0, -1.0 / 4.0, -3.0 / 8.0, -67.0 / 48.0, 1.0 / 8.0, 13.0 / 48.0},
		c:     0.7,
		det:   48,
		trace: -15,
	},
}

var mat4tests = []struct {
	m0, m1, add, sub, cmat, mul, abs, tran, inv Mat4
	c, det, trace                               float32
}{
	{
		m0:    Mat4{},
		m1:    Mat4{},
		add:   Mat4{},
		sub:   Mat4{},
		cmat:  Mat4{},
		mul:   Mat4{},
		abs:   Mat4{},
		tran:  Mat4{},
		inv:   Mat4{},
		c:     1,
		det:   0,
		trace: 0,
	},
	{
		m0:    Mat4{-0.3, -2, -3, -4, -5.2, -6, -7, -1, -0.1, -10, -11, -12, -1.1, -14, -15, -16},
		m1:    Mat4{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		add:   Mat4{0.7, 0, 0, 0, -0.2, 0, 0, 7, 8.9, 0, 0, 0, 11.9, 0, 0, 0},
		sub:   Mat4{-1.3, -4, -6, -8, -10.2, -12, -14, -9, -9.1, -20, -22, -24, -14.1, -28, -30, -32},
		cmat:  Mat4{-0.3 * 0.3, -2 * 0.3, -3 * 0.3, -4 * 0.3, -5.2 * 0.3, -6 * 0.3, -7 * 0.3, -1 * 0.3, -0.1 * 0.3, -10 * 0.3, -11 * 0.3, -12 * 0.3, -1.1 * 0.3, -14 * 0.3, -15 * 0.3, -16 * 0.3},
		mul:   Mat4{-15.4, -100, -110, -106, -42.2, -228, -254, -238, -69, -356, -398, -370, -95.8, -484, -542, -502},
		abs:   Mat4{0.3, 2, 3, 4, 5.2, 6, 7, 1, 0.1, 10, 11, 12, 1.1, 14, 15, 16},
		tran:  Mat4{-0.3, -5.2, -0.1, -1.1, -2, -6, -10, -14, -3, -7, -11, -15, -4, -1, -12, -16},
		inv:   Mat4{-0.454545, 8.32667e-17, 1.36364, -0.909091, 0.808442, 0.142857, 1.03896, -0.99026, -0.298701, -0.285714, -2.03247, 1.61688, -0.396104, 0.142857, 0.902597, -0.649351},
		c:     0.3,
		det:   61.6,
		trace: -33.3,
	},
}

func TestMat2_Add(t *testing.T) {
	for i, test := range mat2tests {
		if add := test.m0.Add(&test.m1); !add.Equal(&test.add) {
			t.Errorf("[%d] Add() = \n%swant\n%s", i, add.String(), test.add.String())
		}
		add := test.m0
		add.AddWith(&test.m1)
		if !add.Equal(&test.add) {
			t.Errorf("[%d] AddWith = \n%swant\n%s", i, add.String(), test.add.String())
		}
		add = Mat2{}
		add.AddOf(&test.m0, &test.m1)
		if !add.Equal(&test.add) {
			t.Errorf("[%d] AddOf = \n%swant\n%s", i, add.String(), test.add.String())
		}
	}
}

func TestMat2_Sub(t *testing.T) {
	for i, test := range mat2tests {
		if sub := test.m0.Sub(&test.m1); !sub.Equal(&test.sub) {
			t.Errorf("[%d] Sub() = \n%swant\n%s", i, sub.String(), test.sub.String())
		}
		sub := test.m0
		sub.SubWith(&test.m1)
		if !sub.Equal(&test.sub) {
			t.Errorf("[%d] SubWith = \n%swant\n%s", i, sub.String(), test.sub.String())
		}
		sub = Mat2{}
		sub.SubOf(&test.m0, &test.m1)
		if !sub.Equal(&test.sub) {
			t.Errorf("[%d] SubOf = \n%swant\n%s", i, sub.String(), test.sub.String())
		}
	}
}

func TestMat2_Mul(t *testing.T) {
	for i, test := range mat2tests {
		if cmat := test.m0.Mul(test.c); !cmat.Equal(&test.cmat) {
			t.Errorf("[%d] Mul() = \n%swant\n%s", i, cmat.String(), test.cmat.String())
		}
		cmat := test.m0
		cmat.MulWith(test.c)
		if !cmat.Equal(&test.cmat) {
			t.Errorf("[%d] MulWith = \n%swant\n%s", i, cmat.String(), test.cmat.String())
		}
		cmat = Mat2{}
		cmat.MulOf(&test.m0, test.c)
		if !cmat.Equal(&test.cmat) {
			t.Errorf("[%d] MulOf = \n%swant\n%s", i, cmat.String(), test.cmat.String())
		}
	}
}

func TestMat2_Mul2(t *testing.T) {
	for i, test := range mat2tests {
		if mul := test.m0.Mul2(&test.m1); !mul.Equal(&test.mul) {
			t.Errorf("[%d] Mul2() = \n%swant\n%s", i, mul.String(), test.mul.String())
		}
		mul := test.m0
		mul.Mul2With(&test.m1)
		if !mul.Equal(&test.mul) {
			t.Errorf("[%d] Mul2With = \n%swant\n%s", i, mul.String(), test.mul.String())
		}
		mul = Mat2{}
		mul.Mul2Of(&test.m0, &test.m1)
		if !mul.Equal(&test.mul) {
			t.Errorf("[%d] Mul2Of = \n%swant\n%s", i, mul.String(), test.mul.String())
		}
	}
}

func TestMat2_Transpose(t *testing.T) {
	for i, test := range mat2tests {
		if tr := test.m0.Transposed(); !tr.Equal(&test.tran) {
			t.Errorf("[%d] Transposed() = \n%swant\n%s", i, tr.String(), test.tran.String())
		}
		tr := test.m0
		tr.Transpose()
		if !tr.Equal(&test.tran) {
			t.Errorf("[%d] Transpose() = \n%swant\n%s", i, tr.String(), test.tran.String())
		}
		tr = Mat2{}
		tr.TransposeOf(&test.m0)
		if !tr.Equal(&test.tran) {
			t.Errorf("[%d] Transpose() = \n%swant\n%s", i, tr.String(), test.tran.String())
		}
	}
}

func TestMat2_Det(t *testing.T) {
	for i, test := range mat2tests {
		if det := test.m0.Det(); !FloatEqual(det, test.det) {
			t.Errorf("[%d] Det() = %f, want %f", i, det, test.det)
		}
	}
}

func TestMat2_Inverse(t *testing.T) {
	for i, test := range mat2tests {
		if inv := test.m0.Inverse(); !inv.Equal(&test.inv) {
			t.Errorf("[%d] Inverse() = \n%swant\n%s", i, inv.String(), test.inv.String())
		}
		inv := test.m0
		inv.Invert()
		if !inv.Equal(&test.inv) {
			t.Errorf("[%d] Invert() = \n%swant\n%s", i, inv.String(), test.inv.String())
		}
		inv = Mat2{}
		inv.InverseOf(&test.m0)
		if !inv.Equal(&test.inv) {
			t.Errorf("[%d] InverseOf() = \n%swant\n%s", i, inv.String(), test.inv.String())
		}
	}
}

func TestMat2_Equal(t *testing.T) {
	for i, test := range mat2tests {
		if !test.m0.Equal(&test.m0) || !test.m0.EqualThreshold(&test.m0, 0) {
			t.Errorf("[%d] not equal", i)
		}
	}
}

func TestMat2_AtSet(t *testing.T) {
	const a = float32(1729)
	var mat Mat2
	for r := 0; r < mat.RowLen(); r++ {
		for c := 0; c < mat.ColLen(); c++ {
			mat.Set(r, c, a)
			if v := mat.At(r, c); !FloatEqual(a, v) {
				t.Errorf("At(%d, %d) = %f, want %f", r, c, v, a)
			}
		}
	}
}

func TestMat2_Index(t *testing.T) {
	var index int
	var mat Mat2
	for c := 0; c < mat.ColLen(); c++ {
		for r := 0; r < mat.RowLen(); r++ {
			if i := mat.Index(r, c); i != index {
				t.Errorf("Index(%d, %d) = %d, want %d", r, c, i, index)
			}
			index++
		}
	}
}

func TestMat2_Row(t *testing.T) {
	for i, test := range mat2tests {
		for r := 0; r < test.m0.RowLen(); r++ {
			var row Vec2
			for c := 0; c < test.m0.ColLen(); c++ {
				row[c] = test.m0.At(r, c)
			}
			if mr := test.m0.Row(r); mr != row {
				t.Errorf("[%d] Row(%d) = %s, want %s", i, r, mr.String(), row.String())
			}
		}
	}
}

func TestMat2_Rows(t *testing.T) {
	for i, test := range mat2tests {
		var rows [2]Vec2
		for r := 0; r < test.m0.RowLen(); r++ {
			for c := 0; c < test.m0.ColLen(); c++ {
				rows[r][c] = test.m0.At(r, c)
			}
		}
		r0, r1 := test.m0.Rows()
		mrows := [2]Vec2{r0, r1}
		if rows != mrows {
			t.Errorf("[%d] Rows unexpected result", i)
		}
	}
}

func TestMat2_Col(t *testing.T) {
	for i, test := range mat2tests {
		for c := 0; c < test.m0.ColLen(); c++ {
			var col Vec2
			for r := 0; r < test.m0.RowLen(); r++ {
				col[r] = test.m0.At(r, c)
			}
			if mc := test.m0.Col(c); mc != col {
				t.Errorf("[%d] Col(%d) = %s, want %s", i, c, mc.String(), col.String())
			}
		}
	}
}

func TestMat2_Cols(t *testing.T) {
	for i, test := range mat2tests {
		var cols [2]Vec2
		for c := 0; c < test.m0.RowLen(); c++ {
			for r := 0; r < test.m0.RowLen(); r++ {
				cols[c][r] = test.m0.At(r, c)
			}
		}
		c0, c1 := test.m0.Cols()
		mcols := [2]Vec2{c0, c1}
		if cols != mcols {
			t.Errorf("[%d] Cols unexpected result", i)
		}
	}
}

func TestMat2_Trace(t *testing.T) {
	for i, test := range mat2tests {
		if trace := test.m0.Trace(); !FloatEqual(test.trace, trace) {
			t.Errorf("[%d] Trace() = %f, want %f", i, trace, test.trace)
		}
	}
}

func TestMat2_Abs(t *testing.T) {
	for i, test := range mat2tests {
		if abs := test.m0.Abs(); !abs.Equal(&test.abs) {
			t.Errorf("[%d] Abs() = \n%swant\n%s", i, abs.String(), test.abs.String())
		}
		abs := test.m0
		abs.AbsSelf()
		if !abs.Equal(&test.abs) {
			t.Errorf("[%d] AbsSelf() = \n%swant\n%s", i, abs.String(), test.abs.String())
		}
		abs = Mat2{}
		abs.AbsOf(&test.m0)
		if !abs.Equal(&test.abs) {
			t.Errorf("[%d] AbsOf() = \n%swant\n%s", i, abs.String(), test.abs.String())
		}
	}
}

func TestMat2_Ident(t *testing.T) {
	var i0, i1, i2 Mat2
	i0.Ident()
	i1.Ident()
	i2.Ident()
	if im := i0.Mul2(&i1); !im.Equal(&i2) {
		t.Errorf("Identity multiplication doesn't yield identity \n%sx\n%s=\n%s", i0.String(), i1.String(), im.String())
	}
}

func TestMat2_String(t *testing.T) {
	m := Ident2()
	expect := " 1.000000 0.000000\n 0.000000 1.000000\n"
	if s := m.String(); s != expect {
		t.Errorf("string don't match %q, want %q", s, expect)
	}
}

//////////////////////////////

func TestMat3_Add(t *testing.T) {
	for i, test := range mat3tests {
		if add := test.m0.Add(&test.m1); !add.Equal(&test.add) {
			t.Errorf("[%d] Add() = \n%swant\n%s", i, add.String(), test.add.String())
		}
		add := test.m0
		add.AddWith(&test.m1)
		if !add.Equal(&test.add) {
			t.Errorf("[%d] AddWith = \n%swant\n%s", i, add.String(), test.add.String())
		}
		add = Mat3{}
		add.AddOf(&test.m0, &test.m1)
		if !add.Equal(&test.add) {
			t.Errorf("[%d] AddOf = \n%swant\n%s", i, add.String(), test.add.String())
		}
	}
}

func TestMat3_Sub(t *testing.T) {
	for i, test := range mat3tests {
		if sub := test.m0.Sub(&test.m1); !sub.Equal(&test.sub) {
			t.Errorf("[%d] Sub() = \n%swant\n%s", i, sub.String(), test.sub.String())
		}
		sub := test.m0
		sub.SubWith(&test.m1)
		if !sub.Equal(&test.sub) {
			t.Errorf("[%d] SubWith = \n%swant\n%s", i, sub.String(), test.sub.String())
		}
		sub = Mat3{}
		sub.SubOf(&test.m0, &test.m1)
		if !sub.Equal(&test.sub) {
			t.Errorf("[%d] SubOf = \n%swant\n%s", i, sub.String(), test.sub.String())
		}
	}
}

func TestMat3_Mul(t *testing.T) {
	for i, test := range mat3tests {
		if cmat := test.m0.Mul(test.c); !cmat.EqualThreshold(&test.cmat, 1e-6) {
			t.Errorf("[%d] Mul() = \n%swant\n%s", i, cmat.String(), test.cmat.String())
		}
		cmat := test.m0
		cmat.MulWith(test.c)
		if !cmat.EqualThreshold(&test.cmat, 1e-6) {
			t.Errorf("[%d] MulWith() = \n%swant\n%s", i, cmat.String(), test.cmat.String())
		}
		cmat = Mat3{}
		cmat.MulOf(&test.m0, test.c)
		if !cmat.EqualThreshold(&test.cmat, 1e-6) {
			t.Errorf("[%d] MulOf() = \n%swant\n%s", i, cmat.String(), test.cmat.String())
		}
	}
}

func TestMat3FromRows(t *testing.T) {
	r0, r1, r2 := Vec3{1, 4, 7}, Vec3{2, 5, 8}, Vec3{3, 6, 9}
	m := Mat3FromRows(&r0, &r1, &r2)
	expect := Mat3{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if expect != m {
		t.Errorf("Mat3FromRows dont match %+v %+v", m, expect)
	}
}

func TestMat3FromCols(t *testing.T) {
	c0, c1, c2 := Vec3{1, 2, 3}, Vec3{4, 5, 6}, Vec3{7, 8, 9}
	m := Mat3FromCols(&c0, &c1, &c2)
	expect := Mat3{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if expect != m {
		t.Errorf("Mat3FromRows dont match %+v %+v", m, expect)
	}
}

func TestMat3_Mul3(t *testing.T) {
	for i, test := range mat3tests {
		if mul := test.m0.Mul3(&test.m1); !mul.Equal(&test.mul) {
			t.Errorf("[%d] Mul2() = \n%swant\n%s", i, mul.String(), test.mul.String())
		}
		mul := test.m0
		mul.Mul3With(&test.m1)
		if !mul.Equal(&test.mul) {
			t.Errorf("[%d] Mul2With = \n%swant\n%s", i, mul.String(), test.mul.String())
		}
		mul = Mat3{}
		mul.Mul3Of(&test.m0, &test.m1)
		if !mul.Equal(&test.mul) {
			t.Errorf("[%d] Mul2Of = \n%swant\n%s", i, mul.String(), test.mul.String())
		}
	}
}

func TestMat3_Transpose(t *testing.T) {
	for i, test := range mat3tests {
		tr := test.m0.Transposed()
		if !tr.Equal(&test.tran) {
			t.Errorf("[%d] Transposed() = \n%swant\n%s", i, tr.String(), test.tran.String())
		}
		tr = test.m0
		tr.Transpose()
		if !tr.Equal(&test.tran) {
			t.Errorf("[%d] Transpose() = \n%swant\n%s", i, tr.String(), test.tran.String())
		}
		tr = Mat3{}
		tr.TransposeOf(&test.m0)
		if !tr.Equal(&test.tran) {
			t.Errorf("[%d] TransposeOf() = \n%swant\n%s", i, tr.String(), test.tran.String())
		}
	}
}

func TestMat3_Det(t *testing.T) {
	for i, test := range mat3tests {
		if det := test.m0.Det(); !FloatEqual(det, test.det) {
			t.Errorf("[%d] Det() = %f, want %f", i, det, test.det)
		}
	}
}

func TestMat3_Inverse(t *testing.T) {
	for i, test := range mat3tests {
		if inv := test.m0.Inverse(); !inv.Equal(&test.inv) {
			t.Errorf("[%d] Inverse() = \n%swant\n%s", i, inv.String(), test.inv.String())
		}
		inv := test.m0
		inv.Invert()
		if !inv.Equal(&test.inv) {
			t.Errorf("[%d] Invert() = \n%swant\n%s", i, inv.String(), test.inv.String())
		}
		inv = Mat3{}
		inv.InverseOf(&test.m0)
		if !inv.Equal(&test.inv) {
			t.Errorf("[%d] InverseOf() = \n%swant\n%s", i, inv.String(), test.inv.String())
		}
	}
}

func TestMat3_Equal(t *testing.T) {
	for i, test := range mat3tests {
		if !test.m0.Equal(&test.m0) || !test.m0.EqualThreshold(&test.m0, 0) {
			t.Errorf("[%d] not equal", i)
		}
	}
}

func TestMat3_AtSet(t *testing.T) {
	const a = float32(1729)
	var mat Mat3
	for r := 0; r < mat.RowLen(); r++ {
		for c := 0; c < mat.ColLen(); c++ {
			mat.Set(r, c, a)
			if v := mat.At(r, c); !FloatEqual(a, v) {
				t.Errorf("At(%d, %d) = %f, want %f", r, c, v, a)
			}
		}
	}
}

func TestMat3_Index(t *testing.T) {
	var index int
	var mat Mat3
	for c := 0; c < mat.ColLen(); c++ {
		for r := 0; r < mat.RowLen(); r++ {
			if i := mat.Index(r, c); i != index {
				t.Errorf("Index(%d, %d) = %d, want %d", r, c, i, index)
			}
			index++
		}
	}
}

func TestMat3_Row(t *testing.T) {
	for i, test := range mat3tests {
		for r := 0; r < test.m0.RowLen(); r++ {
			var row Vec3
			for c := 0; c < test.m0.ColLen(); c++ {
				row[c] = test.m0.At(r, c)
			}
			if mr := test.m0.Row(r); mr != row {
				t.Errorf("[%d] Row(%d) = %s, want %s", i, r, mr.String(), row.String())
			}
		}
	}
}

func TestMat3_Rows(t *testing.T) {
	for i, test := range mat3tests {
		var rows [3]Vec3
		for r := 0; r < test.m0.RowLen(); r++ {
			for c := 0; c < test.m0.ColLen(); c++ {
				rows[r][c] = test.m0.At(r, c)
			}
		}
		r0, r1, r2 := test.m0.Rows()
		mrows := [3]Vec3{r0, r1, r2}
		if rows != mrows {
			t.Errorf("[%d] Rows unexpected result", i)
		}
	}
}

func TestMat3_Col(t *testing.T) {
	for i, test := range mat3tests {
		for c := 0; c < test.m0.ColLen(); c++ {
			var col Vec3
			for r := 0; r < test.m0.RowLen(); r++ {
				col[r] = test.m0.At(r, c)
			}
			if mc := test.m0.Col(c); mc != col {
				t.Errorf("[%d] Col(%d) = %s, want %s", i, c, mc.String(), col.String())
			}
		}
	}
}

func TestMat3_Cols(t *testing.T) {
	for i, test := range mat3tests {
		var cols [3]Vec3
		for c := 0; c < test.m0.RowLen(); c++ {
			for r := 0; r < test.m0.RowLen(); r++ {
				cols[c][r] = test.m0.At(r, c)
			}
		}
		c0, c1, c2 := test.m0.Cols()
		mcols := [3]Vec3{c0, c1, c2}
		if cols != mcols {
			t.Errorf("[%d] Cols unexpected result", i)
		}
	}
}

func TestMat3_Trace(t *testing.T) {
	for i, test := range mat3tests {
		if trace := test.m0.Trace(); !FloatEqual(test.trace, trace) {
			t.Errorf("[%d] Trace() = %f, want %f", i, trace, test.trace)
		}
	}
}

func TestMat3_Abs(t *testing.T) {
	for i, test := range mat3tests {
		if abs := test.m0.Abs(); !abs.Equal(&test.abs) {
			t.Errorf("[%d] Abs() = \n%swant\n%s", i, abs.String(), test.abs.String())
		}
		abs := test.m0
		abs.AbsSelf()
		if !abs.Equal(&test.abs) {
			t.Errorf("[%d] AbsSelf() = \n%swant\n%s", i, abs.String(), test.abs.String())
		}
		abs = Mat3{}
		abs.AbsOf(&test.m0)
		if !abs.Equal(&test.abs) {
			t.Errorf("[%d] AbsOf() = \n%swant\n%s", i, abs.String(), test.abs.String())
		}
	}
}

func TestMat3_Ident(t *testing.T) {
	var i0, i1, i2 Mat3
	i0.Ident()
	i1.Ident()
	i2.Ident()
	if im := i0.Mul3(&i1); !im.Equal(&i2) {
		t.Errorf("Identity multiplication doesn't yield identity \n%sx\n%s=\n%s", i0.String(), i1.String(), im.String())
	}
}

func TestMat3_String(t *testing.T) {
	m := Ident3()
	expect := " 1.000000 0.000000 0.000000\n 0.000000 1.000000 0.000000\n 0.000000 0.000000 1.000000\n"
	if s := m.String(); s != expect {
		t.Errorf("string don't match %q, want %q", s, expect)
	}
}

////////////////////////
////////////////////////

func TestMat4_Add(t *testing.T) {
	for i, test := range mat4tests {
		if add := test.m0.Add(&test.m1); !add.EqualThreshold(&test.add, 1e-5) {
			t.Errorf("[%d] Add() = \n%swant\n%s", i, add.String(), test.add.String())
		}
		add := test.m0
		add.AddWith(&test.m1)
		if !add.EqualThreshold(&test.add, 1e-5) {
			t.Errorf("[%d] AddWith = \n%swant\n%s", i, add.String(), test.add.String())
		}
		add = Mat4{}
		add.AddOf(&test.m0, &test.m1)
		if !add.EqualThreshold(&test.add, 1e-5) {
			t.Errorf("[%d] AddOf = \n%swant\n%s", i, add.String(), test.add.String())
		}
	}
}

func TestMat4_Sub(t *testing.T) {
	for i, test := range mat4tests {
		if sub := test.m0.Sub(&test.m1); !sub.Equal(&test.sub) {
			t.Errorf("[%d] Sub() = \n%swant\n%s", i, sub.String(), test.sub.String())
		}
		sub := test.m0
		sub.SubWith(&test.m1)
		if !sub.Equal(&test.sub) {
			t.Errorf("[%d] SubWith = \n%swant\n%s", i, sub.String(), test.sub.String())
		}
		sub = Mat4{}
		sub.SubOf(&test.m0, &test.m1)
		if !sub.Equal(&test.sub) {
			t.Errorf("[%d] SubOf = \n%swant\n%s", i, sub.String(), test.sub.String())
		}
	}
}

func TestMat4_Mul(t *testing.T) {
	for i, test := range mat4tests {
		if cmat := test.m0.Mul(test.c); !cmat.EqualThreshold(&test.cmat, 1e-6) {
			t.Errorf("[%d] Mul() = \n%swant\n%s", i, cmat.String(), test.cmat.String())
		}
		cmat := test.m0
		cmat.MulWith(test.c)
		if !cmat.EqualThreshold(&test.cmat, 1e-6) {
			t.Errorf("[%d] MulWith() = \n%swant\n%s", i, cmat.String(), test.cmat.String())
		}
		cmat = Mat4{}
		cmat.MulOf(&test.m0, test.c)
		if !cmat.EqualThreshold(&test.cmat, 1e-6) {
			t.Errorf("[%d] MulOf() = \n%swant\n%s", i, cmat.String(), test.cmat.String())
		}
	}
}

func TestMat4_Mul4(t *testing.T) {
	for i, test := range mat4tests {
		if mul := test.m0.Mul4(&test.m1); !mul.EqualThreshold(&test.mul, 1e-5) {
			t.Errorf("[%d] Mul2() = \n%swant\n%s", i, mul.String(), test.mul.String())
		}
		mul := test.m0
		mul.Mul4With(&test.m1)
		if !mul.EqualThreshold(&test.mul, 1e-5) {
			t.Errorf("[%d] Mul2With = \n%swant\n%s", i, mul.String(), test.mul.String())
		}
		mul = Mat4{}
		mul.Mul4Of(&test.m0, &test.m1)
		if !mul.EqualThreshold(&test.mul, 1e-5) {
			t.Errorf("[%d] Mul2Of = \n%swant\n%s", i, mul.String(), test.mul.String())
		}
	}
}

func TestMat4_Transpose(t *testing.T) {
	for i, test := range mat4tests {
		tr := test.m0.Transposed()
		if !tr.Equal(&test.tran) {
			t.Errorf("[%d] Transposed() = \n%swant\n%s", i, tr.String(), test.tran.String())
		}
		tr = test.m0
		tr.Transpose()
		if !tr.Equal(&test.tran) {
			t.Errorf("[%d] Transpose() = \n%swant\n%s", i, tr.String(), test.tran.String())
		}
		tr = Mat4{}
		tr.TransposeOf(&test.m0)
		if !tr.Equal(&test.tran) {
			t.Errorf("[%d] TransposeOf() = \n%swant\n%s", i, tr.String(), test.tran.String())
		}
	}
}

func TestMat4_Det(t *testing.T) {
	for i, test := range mat4tests {
		if det := test.m0.Det(); !FloatEqualThreshold(det, test.det, 1e-4) {
			t.Errorf("[%d] Det() = %f, want %f", i, det, test.det)
		}
	}
}

func TestMat4_Inverse(t *testing.T) {
	for i, test := range mat4tests {
		if inv := test.m0.Inverse(); !inv.EqualThreshold(&test.inv, 1e-4) {
			t.Errorf("[%d] Inverse() = \n%swant\n%s", i, inv.String(), test.inv.String())
		}
		inv := test.m0
		inv.Invert()
		if !inv.EqualThreshold(&test.inv, 1e-4) {
			t.Errorf("[%d] Invert() = \n%swant\n%s", i, inv.String(), test.inv.String())
		}
		inv = Mat4{}
		inv.InverseOf(&test.m0)
		if !inv.EqualThreshold(&test.inv, 1e-4) {
			t.Errorf("[%d] InverseOf() = \n%swant\n%s", i, inv.String(), test.inv.String())
		}
	}
}

func TestMat4_Equal(t *testing.T) {
	for i, test := range mat4tests {
		if !test.m0.Equal(&test.m0) || !test.m0.EqualThreshold(&test.m0, 0) {
			t.Errorf("[%d] not equal", i)
		}
	}
}

func TestMat4_AtSet(t *testing.T) {
	const a = float32(1729)
	var mat Mat4
	for r := 0; r < mat.RowLen(); r++ {
		for c := 0; c < mat.ColLen(); c++ {
			mat.Set(r, c, a)
			if v := mat.At(r, c); !FloatEqual(a, v) {
				t.Errorf("At(%d, %d) = %f, want %f", r, c, v, a)
			}
		}
	}
}

func TestMat4_Index(t *testing.T) {
	var index int
	var mat Mat4
	for c := 0; c < mat.ColLen(); c++ {
		for r := 0; r < mat.RowLen(); r++ {
			if i := mat.Index(r, c); i != index {
				t.Errorf("Index(%d, %d) = %d, want %d", r, c, i, index)
			}
			index++
		}
	}
}

func TestMat4_Row(t *testing.T) {
	for i, test := range mat4tests {
		for r := 0; r < test.m0.RowLen(); r++ {
			var row Vec4
			for c := 0; c < test.m0.ColLen(); c++ {
				row[c] = test.m0.At(r, c)
			}
			if mr := test.m0.Row(r); mr != row {
				t.Errorf("[%d] Row(%d) = %s, want %s", i, r, mr.String(), row.String())
			}
		}
	}
}

func TestMat4_Rows(t *testing.T) {
	for i, test := range mat4tests {
		var rows [4]Vec4
		for r := 0; r < test.m0.RowLen(); r++ {
			for c := 0; c < test.m0.ColLen(); c++ {
				rows[r][c] = test.m0.At(r, c)
			}
		}
		r0, r1, r2, r3 := test.m0.Rows()
		mrows := [4]Vec4{r0, r1, r2, r3}
		if rows != mrows {
			t.Errorf("[%d] Rows unexpected result", i)
		}
	}
}

func TestMat4_Col(t *testing.T) {
	for i, test := range mat4tests {
		for c := 0; c < test.m0.ColLen(); c++ {
			var col Vec4
			for r := 0; r < test.m0.RowLen(); r++ {
				col[r] = test.m0.At(r, c)
			}
			if mc := test.m0.Col(c); mc != col {
				t.Errorf("[%d] Col(%d) = %s, want %s", i, c, mc.String(), col.String())
			}
		}
	}
}

func TestMat4_Cols(t *testing.T) {
	for i, test := range mat4tests {
		var cols [4]Vec4
		for c := 0; c < test.m0.RowLen(); c++ {
			for r := 0; r < test.m0.RowLen(); r++ {
				cols[c][r] = test.m0.At(r, c)
			}
		}
		c0, c1, c2, c3 := test.m0.Cols()
		mcols := [4]Vec4{c0, c1, c2, c3}
		if cols != mcols {
			t.Errorf("[%d] Cols unexpected result", i)
		}
	}
}

func TestMat4_Trace(t *testing.T) {
	for i, test := range mat4tests {
		if trace := test.m0.Trace(); !FloatEqual(test.trace, trace) {
			t.Errorf("[%d] Trace() = %f, want %f", i, trace, test.trace)
		}
	}
}

func TestMat4_Abs(t *testing.T) {
	for i, test := range mat4tests {
		if abs := test.m0.Abs(); !abs.Equal(&test.abs) {
			t.Errorf("[%d] Abs() = \n%swant\n%s", i, abs.String(), test.abs.String())
		}
		abs := test.m0
		abs.AbsSelf()
		if !abs.Equal(&test.abs) {
			t.Errorf("[%d] AbsSelf() = \n%swant\n%s", i, abs.String(), test.abs.String())
		}
		abs = Mat4{}
		abs.AbsOf(&test.m0)
		if !abs.Equal(&test.abs) {
			t.Errorf("[%d] AbsOf() = \n%swant\n%s", i, abs.String(), test.abs.String())
		}
	}
}

func TestMat4_Ident(t *testing.T) {
	var i0, i1, i2 Mat4
	i0.Ident()
	i1.Ident()
	i2.Ident()
	if im := i0.Mul4(&i1); !im.Equal(&i2) {
		t.Errorf("Identity multiplication doesn't yield identity \n%sx\n%s=\n%s", i0.String(), i1.String(), im.String())
	}
}

func TestMat4_String(t *testing.T) {
	m := Ident3()
	expect := " 1.000000 0.000000 0.000000\n 0.000000 1.000000 0.000000\n 0.000000 0.000000 1.000000\n"
	if s := m.String(); s != expect {
		t.Errorf("string don't match %q, want %q", s, expect)
	}
}

func BenchmarkMatAdd(b *testing.B) {
	b.StopTimer()
	rand := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	m1 := &Mat4{}
	m2 := &Mat4{}
	m3 := &Mat4{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		for j := 0; j < len(m1); j++ {
			m3[j], m2[j] = rand.Float32(), rand.Float32()
		}
		b.StartTimer()

		m1.AddOf(m2, m3)
	}
}

func BenchmarkMatScale(b *testing.B) {
	b.StopTimer()
	rand := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		m1 := &Mat4{}

		for j := 0; j < len(m1); j++ {
			m1[j] = rand.Float32()
		}
		c := rand.Float32()
		b.StartTimer()

		m1.MulWith(c)
	}
}

func BenchmarkMatMul(b *testing.B) {
	b.StopTimer()
	rand := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		m1 := &Mat4{}
		m2 := &Mat4{}

		for j := 0; j < len(m1); j++ {
			m1[j], m2[j] = rand.Float32(), rand.Float32()
		}
		b.StartTimer()

		m1.Mul4With(m2)
	}
}

func BenchmarkMatTranspose(b *testing.B) {
	b.StopTimer()
	rand := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	m1 := Mat4{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		for j := 0; j < len(m1); j++ {
			m1[j] = rand.Float32()
		}
		b.StartTimer()

		m1.Transpose()
	}
}

func BenchmarkMatDet(b *testing.B) {
	b.StopTimer()
	rand := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		m1 := Mat4{}

		for j := 0; j < len(m1); j++ {
			m1[j] = rand.Float32()
		}
		b.StartTimer()

		_ = m1.Det()
	}
}

func BenchmarkMatInvSelf(b *testing.B) {
	b.StopTimer()
	rand := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		m1 := &Mat4{}

		for j := 0; j < len(m1); j++ {
			m1[j] = rand.Float32()
		}
		b.StartTimer()

		m1.Invert()
	}
}

func BenchmarkMatInvNew(b *testing.B) {
	b.StopTimer()
	rand := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		m1 := &Mat4{}

		for j := 0; j < len(m1); j++ {
			m1[j] = rand.Float32()
		}
		b.StartTimer()
		m1.Invert()
	}
}
