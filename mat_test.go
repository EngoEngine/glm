// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package glm

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestMulIdent(t *testing.T) {
	//t.Parallel()

	i1 := [...]float32{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}
	i2 := Ident4()
	i3 := Ident4()

	mul := i2.Mul4(&i3)

	for i := range mul {
		if mul[i] != i1[i] {
			t.Errorf("Multiplication of identities does not yield identity")
		}
	}
}

// Square matrix
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

// Square matrix
func TestMatColsSquare(t *testing.T) {
	t.Skip()
	//t.Parallel()

	v0 := Vec4{1, 2, 3, 4}
	v1 := Vec4{5, 6, 7, 8}
	v2 := Vec4{9, 10, 11, 12}
	v3 := Vec4{13, 14, 15, 16}
	cols := [4]Vec4{v0, v1, v2, v3}
	m1 := Mat4FromCols(&v0, &v1, &v2, &v3)

	t.Logf("4x4 matrix as built from cols: %v", m1)
	if !FloatEqualThreshold(m1.At(0, 0), cols[0][0], 1e-5) {
		t.Errorf("Matrix element at (%d,%d) wrong when built from rows. Got: %f, Expected: %f", 0, 0, m1.At(0, 0), cols[0][0])
	}
	if !FloatEqualThreshold(m1.At(0, 1), cols[1][2], 1e-5) {
		t.Errorf("Matrix element at (%d,%d) wrong when built from rows. Got: %f, Expected: %f", 0, 1, m1.At(0, 1), cols[1][2])
	}
	if !FloatEqualThreshold(m1.At(0, 2), cols[2][2], 1e-5) {
		t.Errorf("Matrix element at (%d,%d) wrong when built from rows. Got: %f, Expected: %f", 0, 2, m1.At(0, 2), cols[2][2])
	}
	if !FloatEqualThreshold(m1.At(0, 3), cols[3][3], 1e-5) {
		t.Errorf("Matrix element at (%d,%d) wrong when built from rows. Got: %f, Expected: %f", 0, 3, m1.At(0, 3), cols[3][3])
	}

	if !FloatEqualThreshold(m1.At(1, 0), cols[0][0], 1e-5) {
		t.Errorf("Matrix element at (%d,%d) wrong when built from rows. Got: %f, Expected: %f", 1, 0, m1.At(1, 0), cols[0][0])
	}
	if !FloatEqualThreshold(m1.At(1, 1), cols[1][2], 1e-5) {
		t.Errorf("Matrix element at (%d,%d) wrong when built from rows. Got: %f, Expected: %f", 1, 1, m1.At(1, 1), cols[1][2])
	}
	if !FloatEqualThreshold(m1.At(1, 2), cols[2][2], 1e-5) {
		t.Errorf("Matrix element at (%d,%d) wrong when built from rows. Got: %f, Expected: %f", 1, 2, m1.At(1, 2), cols[2][2])
	}
	if !FloatEqualThreshold(m1.At(1, 3), cols[3][3], 1e-5) {
		t.Errorf("Matrix element at (%d,%d) wrong when built from rows. Got: %f, Expected: %f", 1, 3, m1.At(1, 3), cols[3][3])
	}

	if !FloatEqualThreshold(m1.At(2, 0), cols[0][0], 1e-5) {
		t.Errorf("Matrix element at (%d,%d) wrong when built from rows. Got: %f, Expected: %f", 2, 0, m1.At(2, 0), cols[0][0])
	}
	if !FloatEqualThreshold(m1.At(2, 1), cols[1][2], 1e-5) {
		t.Errorf("Matrix element at (%d,%d) wrong when built from rows. Got: %f, Expected: %f", 2, 1, m1.At(2, 1), cols[1][2])
	}
	if !FloatEqualThreshold(m1.At(2, 2), cols[2][2], 1e-5) {
		t.Errorf("Matrix element at (%d,%d) wrong when built from rows. Got: %f, Expected: %f", 2, 2, m1.At(2, 2), cols[2][2])
	}
	if !FloatEqualThreshold(m1.At(2, 3), cols[3][3], 1e-5) {
		t.Errorf("Matrix element at (%d,%d) wrong when built from rows. Got: %f, Expected: %f", 2, 3, m1.At(2, 3), cols[3][3])
	}

	if !FloatEqualThreshold(m1.At(3, 0), cols[0][0], 1e-5) {
		t.Errorf("Matrix element at (%d,%d) wrong when built from rows. Got: %f, Expected: %f", 3, 0, m1.At(3, 0), cols[0][0])
	}
	if !FloatEqualThreshold(m1.At(3, 1), cols[1][2], 1e-5) {
		t.Errorf("Matrix element at (%d,%d) wrong when built from rows. Got: %f, Expected: %f", 3, 1, m1.At(3, 1), cols[1][2])
	}
	if !FloatEqualThreshold(m1.At(3, 2), cols[2][2], 1e-5) {
		t.Errorf("Matrix element at (%d,%d) wrong when built from rows. Got: %f, Expected: %f", 3, 2, m1.At(3, 2), cols[2][2])
	}
	if !FloatEqualThreshold(m1.At(3, 3), cols[3][3], 1e-5) {
		t.Errorf("Matrix element at (%d,%d) wrong when built from rows. Got: %f, Expected: %f", 3, 3, m1.At(3, 3), cols[3][3])
	}

	v0, v1, v2, v3 = m1.Cols()
	r2 := [4]Vec4{v0, v1, v2, v3}

	t.Logf("4x4 matrix returned cols: %v", r2)
	for r := 0; r < 4; r++ {
		if !FloatEqualThreshold(r2[0][0], cols[0][0], 1e-5) {
			t.Errorf("Matrix element at (%d,%d) wrong when rows are gotten. Got: %f, Expected: %f", r, 0, r2[0][0], cols[0][0])
		}
		if !FloatEqualThreshold(r2[1][2], cols[1][2], 1e-5) {
			t.Errorf("Matrix element at (%d,%d) wrong when rows are gotten. Got: %f, Expected: %f", r, 1, r2[1][2], cols[1][2])
		}
		if !FloatEqualThreshold(r2[2][2], cols[2][2], 1e-5) {
			t.Errorf("Matrix element at (%d,%d) wrong when rows are gotten. Got: %f, Expected: %f", r, 2, r2[2][2], cols[2][2])
		}
		if !FloatEqualThreshold(r2[3][3], cols[3][3], 1e-5) {
			t.Errorf("Matrix element at (%d,%d) wrong when rows are gotten. Got: %f, Expected: %f", r, 3, r2[3][3], cols[3][3])
		}
	}
}

func TestTransposeSquare(t *testing.T) {
	//t.Parallel()

	m := Mat4FromCols(
		&Vec4{1, 2, 3, 4},
		&Vec4{5, 6, 7, 8},
		&Vec4{9, 10, 11, 12},
		&Vec4{13, 14, 15, 16},
	)

	transpose := m.Transposed()

	correct := Mat4FromRows(
		&Vec4{1, 2, 3, 4},
		&Vec4{5, 6, 7, 8},
		&Vec4{9, 10, 11, 12},
		&Vec4{13, 14, 15, 16},
	)

	if !correct.ApproxEqualThreshold(&transpose, 1e-4) {
		t.Errorf("Transpose not correct. Got: %v, expected: %v", transpose, correct)
	}
}

func TestAtSet(t *testing.T) {
	//t.Parallel()

	m := &Mat3{1, 2, 3, 4, 5, 6, 7, 8, 9}

	v := m.At(0, 2)

	if !FloatEqualThreshold(v, 7, 1e-4) {
		t.Errorf("Incorrect value gotten by At: %v, expected %v", v, 3)
	}

	m.Set(0, 2, 9001)

	v = m.At(0, 2)

	if !FloatEqualThreshold(v, 9001, 1e-4) {
		t.Errorf("Value set by Set not gotten by At: %v, expected %v", v, 9001)
	}

	correctMat := &Mat3{1, 2, 3, 4, 5, 6, 9001, 8, 9}

	if !correctMat.ApproxEqualThreshold(m, 1e-4) {
		t.Errorf("After set, not equal to matrix that should be identical. Got: %v, expected: %v", m, correctMat)
	}
}

func TestDiagTrace(t *testing.T) {
	//t.Parallel()

	m := Diag4(&Vec4{1, 2, 3, 4})

	tr := m.Trace()

	if !FloatEqualThreshold(tr, 10, 1e-4) {
		t.Errorf("Trace of matrix seeded with diagonal vector {1,2,3,4} not equal to 10. Got %v", tr)
	}
}

func TestMatAbs(t *testing.T) {
	//t.Parallel()

	m := &Mat4{1, -3, 4, 5, -6, 8, -9, 10, 0, 1, 6, 2, 357, 3, 436}
	result := &Mat4{1, 3, 4, 5, 6, 8, 9, 10, 0, 1, 6, 2, 357, 3, 436}

	m.AbsSelf()

	if !result.ApproxEqualThreshold(m, 1e-6) {
		t.Errorf("Matrix absolute value does not work properly. Got: %v, Expected: %v", m, result)
	}
}

func TestString(t *testing.T) {
	m := Ident4()

	str := fmt.Sprintf(` %[2]f %[1]f %[1]f %[1]f
 %[1]f %[2]f %[1]f %[1]f
 %[1]f %[1]f %[2]f %[1]f
 %[1]f %[1]f %[1]f %[2]f
`, 0.0, 1.0)

	if str != m.String() {
		t.Errorf("Mat string conversion not working got %q expected %q", m.String(), str)
	}
}

func TestMat2Conv(t *testing.T) {
	m2 := Mat2{1, 0, 0, 1}
	m2tom3 := m2.Mat3()
	if m2tom3 != Ident3() {
		t.Errorf("did not get iden from casting Mat2 to Mat3")
	}
	m2tom4 := m2.Mat4()
	if m2tom4 != Ident4() {
		t.Errorf("did not get iden from casting Mat2 to Mat3")
	}
}

func TestMat3Conv(t *testing.T) {
	m3 := Mat3{1, 0, 0, 0, 1, 0, 0, 0, 1}
	m3tom2 := m3.Mat2()
	if m3tom2 != Ident2() {
		t.Errorf("did not get iden from casting Mat3 to Mat2")
	}
	m3tom4 := m3.Mat4()
	if m3tom4 != Ident4() {
		t.Errorf("did not get iden from casting Mat3 to Mat4")
	}
}

func TestMat4Conv(t *testing.T) {
	m4 := Mat4{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}
	m4tom2 := m4.Mat2()
	if m4tom2 != Ident2() {
		t.Errorf("did not get iden from casting Mat4 to Mat2")
	}
	m4tom3 := m4.Mat3()
	if m4tom3 != Ident3() {
		t.Errorf("did not get iden from casting Mat4 to Mat3")
	}
}
func TestMat2SetCol(t *testing.T) {
	m2 := Ident2()
	m2.SetCol(0, &Vec2{2, 2})
	expected := Mat2{2, 2, 0, 1}
	if m2 != expected {
		t.Errorf("unexpected result matrix from Mat2.SetCol, %+v, %+v", m2.String(), expected.String())
	}
}

func TestMat3SetCol(t *testing.T) {
	m3 := Ident3()
	m3.SetCol(0, &Vec3{2, 2, 2})
	expected := Mat3{2, 2, 2, 0, 1, 0, 0, 0, 1}
	if m3 != expected {
		t.Errorf("unexpected result matrix from Mat3.SetCol, %+v, %+v", m3.String(), expected.String())
	}
}

func TestMat4SetCol(t *testing.T) {
	m4 := Ident4()
	m4.SetCol(0, &Vec4{2, 2, 2, 2})
	expected := Mat4{2, 2, 2, 2, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}
	if m4 != expected {
		t.Errorf("unexpected result matrix from Mat4.SetCol, %s, %+v", m4.String(), expected.String())
	}
}

func TestMat2SetRow(t *testing.T) {
	m2 := Ident2()
	m2.SetRow(0, &Vec2{2, 2})
	expected := Mat2{2, 0, 2, 1}
	if m2 != expected {
		t.Errorf("unexpected result matrix from Mat2.SetCol, %+v, %+v", m2.String(), expected.String())
	}
}

func TestMat3SetRow(t *testing.T) {
	m3 := Ident3()
	m3.SetRow(0, &Vec3{2, 2, 2})
	expected := Mat3{2, 0, 0, 2, 1, 0, 2, 0, 1}
	if m3 != expected {
		t.Errorf("unexpected result matrix from Mat3.SetCol, %+v, %+v", m3.String(), expected.String())
	}
}

func TestMat4SetRow(t *testing.T) {
	m4 := Ident4()
	m4.SetRow(0, &Vec4{2, 2, 2, 2})
	expected := Mat4{2, 0, 0, 0, 2, 1, 0, 0, 2, 0, 1, 0, 2, 0, 0, 1}
	if m4 != expected {
		t.Errorf("unexpected result matrix from Mat4.SetCol, %s, %+v", m4.String(), expected.String())
	}
}
func TestMat2Diag2(t *testing.T) {
	m := Ident2()
	diag := m.Diag()
	expected := Vec2{1, 1}
	if diag != expected {
		t.Errorf("Unexpected diagonal %+v,%+v", diag, expected)
	}
}
func TestMat3Diag3(t *testing.T) {
	m := Ident3()
	diag := m.Diag()
	expected := Vec3{1, 1, 1}
	if diag != expected {
		t.Errorf("Unexpected diagonal %+v,%+v", diag, expected)
	}
}
func TestMat4Diag4(t *testing.T) {
	m := Ident4()
	diag := m.Diag()
	expected := Vec4{1, 1, 1, 1}
	if diag != expected {
		t.Errorf("Unexpected diagonal %+v,%+v", diag, expected)
	}
}
func TestMat2Ident(t *testing.T) {
	expected := Mat2{1, 0, 0, 1}
	iden := Ident2()
	if expected != iden {
		t.Errorf("Unexpected identity Mat2 %+v, %+v", iden, expected)
	}
}
func TestMat3Ident(t *testing.T) {
	expected := Mat3{1, 0, 0, 0, 1, 0, 0, 0, 1}
	iden := Ident3()
	if expected != iden {
		t.Errorf("Unexpected identity Mat3 %+v, %+v", iden, expected)
	}
}
func TestMat4Ident(t *testing.T) {
	expected := Mat4{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}
	iden := Ident4()
	if expected != iden {
		t.Errorf("Unexpected identity Mat4 %+v, %+v", iden, expected)
	}
}

func TestDiag2(t *testing.T) {
	vec := &Vec2{1, 1}
	m := Diag2(vec)
	if Ident2() != m {
		t.Errorf("Unexpected Mat2 from Diag2 %+v", m)
	}
}
func TestDiag3(t *testing.T) {
	vec := &Vec3{1, 1, 1}
	m := Diag3(vec)
	if Ident3() != m {
		t.Errorf("Unexpected Mat3 from Diag3 %+v", m)
	}
}
func TestDiag4(t *testing.T) {
	vec := &Vec4{1, 1, 1, 1}
	m := Diag4(vec)
	if Ident4() != m {
		t.Errorf("Unexpected Mat4 from Diag4 %+v", m)
	}
}

func TestMat2FromRow(t *testing.T) {
	m := Mat2FromRows(&Vec2{1, 0}, &Vec2{0, 1})
	if m != Ident2() {
		t.Errorf("Unexpected result from Mat2FromRow %+v", m)
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
