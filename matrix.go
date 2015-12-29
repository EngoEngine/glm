package glm

import (
	"bytes"
	"fmt"
	"github.com/luxengine/math"
	"text/tabwriter"
)

// Mat2 represents a 2x2 matrix
type Mat2 [4]float32

// Mat3 represents a 3x3 matrix
type Mat3 [9]float32

// Mat4 represents a 4x4 matrix
type Mat4 [16]float32

// Mat3 returns
//    [m m 0]
//    [m m 0]
//    [0 0 1]
func (m1 *Mat2) Mat3() Mat3 {
	third := Vec3{0, 0, 1}
	col0, col1 := m1.Cols()
	v0 := col0.Vec3(0)
	v1 := col1.Vec3(0)
	return Mat3FromCols(
		&v0,
		&v1,
		&third,
	)
}

// Mat4 returns
//    [m m 0 0]
//    [m m 0 0]
//    [0 0 0 0]
//    [0 0 0 1]
func (m1 *Mat2) Mat4() Mat4 {
	a, b := Vec4{0, 0, 1, 0}, Vec4{0, 0, 0, 1}
	col0, col1 := m1.Cols()
	v0 := col0.Vec4(0, 0)
	v1 := col1.Vec4(0, 0)
	return Mat4FromCols(
		&v0,
		&v1,
		&a,
		&b,
	)
}

// Mat2 returns
//    [m m ?]
//    [m m ?]
//    [? ? ?]
func (m1 *Mat3) Mat2() Mat2 {
	col0, col1, _ := m1.Cols()
	v0 := col0.Vec2()
	v1 := col1.Vec2()
	return Mat2FromCols(
		&v0,
		&v1,
	)
}

// Mat4 returns
//    [m m m 0]
//    [m m m 0]
//    [m m m 0]
//    [0 0 0 1]
func (m1 *Mat3) Mat4() Mat4 {
	a := Vec4{0, 0, 0, 1}
	col0, col1, col2 := m1.Cols()
	v0 := col0.Vec4(0)
	v1 := col1.Vec4(0)
	v2 := col2.Vec4(0)
	return Mat4FromCols(
		&v0,
		&v1,
		&v2,
		&a,
	)
}

// Mat2 returns
//    [m m ? ?]
//    [m m ? ?]
//    [? ? ? ?]
//    [? ? ? ?]
func (m1 *Mat4) Mat2() Mat2 {
	col0, col1, _, _ := m1.Cols()
	v0 := col0.Vec2()
	v1 := col1.Vec2()
	return Mat2FromCols(
		&v0,
		&v1,
	)
}

// Mat3 returns
//    [m m m ?]
//    [m m m ?]
//    [m m m ?]
//    [? ? ? ?]
func (m1 *Mat4) Mat3() Mat3 {
	col0, col1, col2, _ := m1.Cols()
	v0 := col0.Vec3()
	v1 := col1.Vec3()
	v2 := col2.Vec3()
	return Mat3FromCols(
		&v0,
		&v1,
		&v2,
	)
}

// Mat3x4 returns
//    [m m m m]
//    [m m m m]
//    [m m m m]
//    [? ? ? ?]
func (m1 *Mat4) Mat3x4() Mat3x4 {
	return Mat3x4{m1[0], m1[1], m1[2],
		m1[4], m1[5], m1[6],
		m1[8], m1[9], m1[10],
		m1[12], m1[13], m1[14]}
}

// SetCol sets a Column within the Matrix, so it mutates the calling matrix.
func (m1 *Mat2) SetCol(col int, v *Vec2) {
	m1[col*2+0], m1[col*2+1] = v[0], v[1]
}

// SetRow sets a Row within the Matrix, so it mutates the calling matrix.
func (m1 *Mat2) SetRow(row int, v *Vec2) {
	m1[row+0], m1[row+2] = v[0], v[1]
}

// Diag is a basic operation on a square matrix that simply
// returns main diagonal (meaning all elements such that row==col).
func (m1 *Mat2) Diag() Vec2 {
	return Vec2{m1[0], m1[3]}
}

// Ident2 returns the 2x2 identity matrix.
// The identity matrix is a square matrix with the value 1 on its
// diagonals. The characteristic property of the identity matrix is that
// any matrix multiplied by it is itself. (MI = M; IN = N)
func Ident2() Mat2 {
	return Mat2{1, 0, 0, 1}
}

// Diag2 creates a diagonal matrix from the entries of the input vector.
// That is, for each pointer for row==col, vector[row] is the entry. Otherwise it's 0.
//
// Another way to think about it is that the identity is this function where the every vector element is 1.
func Diag2(v *Vec2) Mat2 {
	return Mat2{v[0], 0, 0, v[1]}
}

// Mat2FromRows builds a new matrix from row vectors.
// The resulting matrix will still be in column major order, but this can be
// good for hand-building matrices.
func Mat2FromRows(row0, row1 *Vec2) Mat2 {
	return Mat2{row0[0], row1[0], row0[1], row1[1]}
}

// Mat2FromCols builds a new matrix from column vectors.
func Mat2FromCols(col0, col1 *Vec2) Mat2 {
	return Mat2{col0[0], col0[1], col1[0], col1[1]}
}

// Add performs an element-wise addition of two matrices, this is
// equivalent to iterating over every element of m1 and adding the corresponding value of m2.
func (m1 *Mat2) Add(m2 *Mat2) Mat2 {
	return Mat2{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3]}
}

// AddOf is a memory friendly version of Add.
func (m1 *Mat2) AddOf(m2, m3 *Mat2) {
	m1[0] = m2[0] + m3[0]
	m1[1] = m2[1] + m3[1]
	m1[2] = m2[2] + m3[2]
	m1[3] = m2[3] + m3[3]
}

// AddWith is a memory friendly version of Add.
func (m1 *Mat2) AddWith(m2 *Mat2) {
	m1[0] += m2[0]
	m1[1] += m2[1]
	m1[2] += m2[2]
	m1[3] += m2[3]
}

// Sub performs an element-wise subtraction of two matrices, this is
// equivalent to iterating over every element of m1 and subtracting the corresponding value of m2.
func (m1 *Mat2) Sub(m2 *Mat2) Mat2 {
	return Mat2{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3]}
}

// SubOf is a memory friendly version of Sub.
func (m1 *Mat2) SubOf(m2, m3 *Mat2) {
	m1[0] = m2[0] - m3[0]
	m1[1] = m2[1] - m3[1]
	m1[2] = m2[2] - m3[2]
	m1[3] = m2[3] - m3[3]
}

// SubWith is a memory friendly version of Sub.
func (m1 *Mat2) SubWith(m2 *Mat2) {
	m1[0] -= m2[0]
	m1[1] -= m2[1]
	m1[2] -= m2[2]
	m1[3] -= m2[3]
}

// Mul performs a scalar multiplcation of the matrix. This is equivalent to iterating
// over every element of the matrix and multiply it by c.
func (m1 *Mat2) Mul(c float32) Mat2 {
	return Mat2{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c}
}

// MulOf is a memory friendly version of Mul.
func (m1 *Mat2) MulOf(m2 *Mat2, c float32) {
	m1[0] = m2[0] * c
	m1[1] = m2[1] * c
	m1[2] = m2[2] * c
	m1[3] = m2[3] * c
}

// MulWith is a memory friendly version of Mul.
func (m1 *Mat2) MulWith(c float32) {
	m1[0] *= c
	m1[1] *= c
	m1[2] *= c
	m1[3] *= c
}

// Mul2x1 performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 *Mat2) Mul2x1(m2 *Vec2) Vec2 {
	return Vec2{
		m1[0]*m2[0] + m1[2]*m2[1],
		m1[1]*m2[0] + m1[3]*m2[1],
	}
}

// Mul2 performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 *Mat2) Mul2(m2 *Mat2) Mat2 {
	return Mat2{
		m1[0]*m2[0] + m1[2]*m2[1],
		m1[1]*m2[0] + m1[3]*m2[1],
		m1[0]*m2[2] + m1[2]*m2[3],
		m1[1]*m2[2] + m1[3]*m2[3],
	}
}

// Mul2Of is a memory friendly version of Mul2.
func (m1 *Mat2) Mul2Of(m2, m3 *Mat2) {
	m1[0] = m2[0]*m3[0] + m2[2]*m3[1]
	m1[1] = m2[1]*m3[0] + m2[3]*m3[1]
	m1[2] = m2[0]*m3[2] + m2[2]*m3[3]
	m1[3] = m2[1]*m3[2] + m2[3]*m3[3]
}

// Mul2With is a memory friendly version of Mul2.
func (m1 *Mat2) Mul2With(m2 *Mat2) {
	v0 := m1[0]
	v1 := m1[1]
	v2 := m1[2]
	v3 := m1[3]
	m1[0] = v0*m2[0] + v2*m2[1]
	m1[1] = v1*m2[0] + v3*m2[1]
	m1[2] = v0*m2[2] + v2*m2[3]
	m1[3] = v1*m2[2] + v3*m2[3]
}

// Transposed produces the transpose of this matrix. For any MxN matrix
// the transpose is an NxM matrix with the rows swapped with the columns. For instance
// the transpose of the Mat3x2 is a Mat2x3 like so:
//
//    [[a b]]    [[a c e]]
//    [[c d]] =  [[b d f]]
//    [[e f]]
func (m1 *Mat2) Transposed() Mat2 {
	return Mat2{m1[0], m1[2], m1[1], m1[3]}
}

// Transpose transpose this matrix with itself as destination. For any MxN matrix
// the transpose is an NxM matrix with the rows swapped with the columns. For instance
// the transpose of the Mat3x2 is a Mat2x3 like so:
//
//    [[a b]]    [[a c e]]
//    [[c d]] =  [[b d f]]
//    [[e f]]
func (m1 *Mat2) Transpose() {
	v1 := m1[1]
	m1[1] = m1[2]
	m1[2] = v1
}

//TransposeOf is a memory friendly version of Transposed.
func (m1 *Mat2) TransposeOf(m2 *Mat2) {
	m1[0], m1[1], m1[2], m1[3] = m2[0], m2[2], m2[1], m2[3]
}

// Det returns the determinant of a matrix. The determinant is a measure of a square matrix's
// singularity and invertability, among other things. In this library, the
// determinant is hard coded based on pre-computed cofactor expansion, and uses
// no loops. Of course, the addition and multiplication must still be done.
func (m1 *Mat2) Det() float32 {
	return m1[0]*m1[2] - m1[1]*m1[3]
}

// Inverse computes the inverse of a square matrix. An inverse is a square matrix such that when multiplied by the
// original, yields the identity.
//
// M_inv * M = M * M_inv = I
//
// In this library, the math is precomputed, and uses no loops, though the multiplications, additions, determinant calculation, and scaling
// are still done. This can still be (relatively) expensive for a 4x4.
//
// This function checks the determinant to see if the matrix is invertible.
// If the determinant is 0.0, this function returns the zero matrix. However, due to floating point errors, it is
// entirely plausible to get a false positive or negative.
// In the future, an alternate function may be written which takes in a pre-computed determinant.
func (m1 *Mat2) Inverse() Mat2 {
	det := m1.Det()
	if FloatEqual(det, float32(0.0)) {
		return Mat2{}
	}

	retMat := Mat2{m1[3], -m1[1], -m1[2], m1[0]}

	return retMat.Mul(1 / det)
}

// ApproxEqual performs an element-wise approximate equality test between two matrices,
// as if FloatEqual had been used.
func (m1 *Mat2) ApproxEqual(m2 *Mat2) bool {
	return FloatEqual(m1[0], m2[0]) && FloatEqual(m1[1], m2[1]) && FloatEqual(m1[2], m2[2]) && FloatEqual(m1[3], m2[3])
}

// ApproxEqualThreshold performs an element-wise approximate equality test between two matrices
// with a given epsilon threshold, as if FloatEqualThreshold had been used.
func (m1 *Mat2) ApproxEqualThreshold(m2 *Mat2, threshold float32) bool {
	return FloatEqualThreshold(m1[0], m2[0], threshold) && FloatEqualThreshold(m1[1], m2[1], threshold) && FloatEqualThreshold(m1[2], m2[2], threshold) && FloatEqualThreshold(m1[3], m2[3], threshold)
}

// ApproxFuncEqual performs an element-wise approximate equality test between two matrices
// with a given equality functions, intended to be used with FloatEqualFunc; although and comparison
// function may be used in practice.
func (m1 *Mat2) ApproxFuncEqual(m2 *Mat2, eq func(float32, float32) bool) bool {
	return eq(m1[0], m2[0]) && eq(m1[1], m2[1]) && eq(m1[2], m2[2]) && eq(m1[3], m2[3])
}

// At returns the matrix element at the given row and column.
// This is equivalent to mat[col * numRow + row] where numRow is constant
// (E.G. for a Mat3x2 it's equal to 3)
//
// This method is garbage-in garbage-out. For instance, on a Mat4 asking for
// At(5,0) will work just like At(1,1). Or it may panic if it's out of bounds.
func (m1 *Mat2) At(row, col int) float32 {
	return m1[col*2+row]
}

// Set sets the corresponding matrix element at the given row and column.
// This has a pointer receiver because it mutates the matrix.
//
// This method is garbage-in garbage-out. For instance, on a Mat4 asking for
// Set(5,0,val) will work just like Set(1,1,val). Or it may panic if it's out of bounds.
func (m1 *Mat2) Set(row, col int, value float32) {
	m1[col*2+row] = value
}

// Index returns the index of the given row and column, to be used with direct
// access. E.G. Index(0,0) = 0.
//
// This is a garbage-in garbage-out method. For instance, on a Mat4 asking for the index of
// (5,0) will work the same as asking for (1,1). Or it may give you a value that will cause
// a panic if you try to access the array with it if it's truly out of bounds.
func (Mat2) Index(row, col int) int {
	return col*2 + row
}

// Row returns a vector representing the corresponding row (starting at row 0).
// This package makes no distinction between row and column vectors, so it
// will be a normal VecM for a MxN matrix.
func (m1 *Mat2) Row(row int) Vec2 {
	return Vec2{m1[row+0], m1[row+2]}
}

// Rows decomposes a matrix into its corresponding row vectors.
// This is equivalent to calling mat.Row for each row.
func (m1 *Mat2) Rows() (row0, row1 Vec2) {
	return m1.Row(0), m1.Row(1)
}

// Col returns a vector representing the corresponding column (starting at col 0).
// This package makes no distinction between row and column vectors, so it
// will be a normal VecN for a MxN matrix.
func (m1 *Mat2) Col(col int) Vec2 {
	return Vec2{m1[col*2+0], m1[col*2+1]}
}

// Cols decomposes a matrix into its corresponding column vectors.
// This is equivalent to calling mat.Col for each column.
func (m1 *Mat2) Cols() (col0, col1 Vec2) {
	return m1.Col(0), m1.Col(1)
}

// Trace is a basic operation on a square matrix that simply
// sums up all elements on the main diagonal (meaning all elements such that row==col).
func (m1 *Mat2) Trace() float32 {
	return m1[0] + m1[3]
}

// Abs returns the element-wise absolute value of this matrix
func (m1 *Mat2) Abs() Mat2 {
	return Mat2{Abs(m1[0]), Abs(m1[1]), Abs(m1[2]), Abs(m1[3])}
}

// AbsSelf is a memory friendly version of Abs.
func (m1 *Mat2) AbsSelf() {
	m1[0], m1[1], m1[2], m1[3] = math.Abs(m1[0]), math.Abs(m1[1]), math.Abs(m1[2]), math.Abs(m1[3])
}

// AbsOf is a memory friendly version of Abs.
func (m1 *Mat2) AbsOf(m2 *Mat2) {
	m1[0], m1[1], m1[2], m1[3] = math.Abs(m2[0]), math.Abs(m2[1]), math.Abs(m2[2]), math.Abs(m2[3])
}

// Iden sets this matrix to the identity matrix.
func (m1 *Mat2) Iden() {
	m1[0] = 1
	m1[1] = 0
	m1[2] = 0
	m1[3] = 1
}

// String pretty prints the matrix
func (m1 *Mat2) String() string {
	buf := new(bytes.Buffer)
	w := tabwriter.NewWriter(buf, 4, 4, 1, ' ', tabwriter.AlignRight)
	for i := 0; i < 2; i++ {
		row := m1.Row(i)
		for _, col := range []float32{row[0], row[1]} {
			fmt.Fprintf(w, "%f\t", col)
		}

		fmt.Fprintln(w, "")
	}
	w.Flush()

	return buf.String()
}

// SetCol sets a Column within the Matrix, so it mutates the calling matrix.
func (m1 *Mat3) SetCol(col int, v *Vec3) {
	m1[col*3+0], m1[col*3+1], m1[col*3+2] = v[0], v[1], v[2]
}

// SetRow sets a Row within the Matrix, so it mutates the calling matrix.
func (m1 *Mat3) SetRow(row int, v *Vec3) {
	m1[row+0], m1[row+3], m1[row+6] = v[0], v[1], v[2]
}

// Diag is a basic operation on a square matrix that simply
// returns main diagonal (meaning all elements such that row==col).
func (m1 *Mat3) Diag() Vec3 {
	return Vec3{m1[0], m1[4], m1[8]}
}

// Ident3 returns the 3x3 identity matrix.
// The identity matrix is a square matrix with the value 1 on its
// diagonals. The characteristic property of the identity matrix is that
// any matrix multiplied by it is itself. (MI = M; IN = N)
func Ident3() Mat3 {
	return Mat3{1, 0, 0, 0, 1, 0, 0, 0, 1}
}

// Diag3 creates a diagonal matrix from the entries of the input vector.
// That is, for each pointer for row==col, vector[row] is the entry. Otherwise it's 0.
//
// Another way to think about it is that the identity is this function where the every vector element is 1.
func Diag3(v *Vec3) Mat3 {
	return Mat3{v[0], 0, 0, 0, v[1], 0, 0, 0, v[2]}
}

// Mat3FromRows builds a new matrix from row vectors.
// The resulting matrix will still be in column major order, but this can be
// good for hand-building matrices.
func Mat3FromRows(row0, row1, row2 *Vec3) Mat3 {
	return Mat3{row0[0], row1[0], row2[0], row0[1], row1[1], row2[1], row0[2], row1[2], row2[2]}
}

// Mat3FromCols builds a new matrix from column vectors.
func Mat3FromCols(col0, col1, col2 *Vec3) Mat3 {
	return Mat3{col0[0], col0[1], col0[2], col1[0], col1[1], col1[2], col2[0], col2[1], col2[2]}
}

// Add performs an element-wise addition of two matrices, this is
// equivalent to iterating over every element of m1 and adding the corresponding value of m2.
func (m1 *Mat3) Add(m2 *Mat3) Mat3 {
	return Mat3{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5], m1[6] + m2[6], m1[7] + m2[7], m1[8] + m2[8]}
}

// AddOf is a memory friendly version of Add.
func (m1 *Mat3) AddOf(m2, m3 *Mat3) {
	m1[0] = m2[0] + m3[0]
	m1[1] = m2[1] + m3[1]
	m1[2] = m2[2] + m3[2]
	m1[3] = m2[3] + m3[3]
	m1[4] = m2[4] + m3[4]
	m1[5] = m2[5] + m3[5]
	m1[6] = m2[6] + m3[6]
	m1[7] = m2[7] + m3[7]
	m1[8] = m2[8] + m3[8]
}

// AddWith is a memory friendly version of Add.
func (m1 *Mat3) AddWith(m2 *Mat3) {
	m1[0] += m2[0]
	m1[1] += m2[1]
	m1[2] += m2[2]
	m1[3] += m2[3]
	m1[4] += m2[4]
	m1[5] += m2[5]
	m1[6] += m2[6]
	m1[7] += m2[7]
	m1[8] += m2[8]
}

// Sub performs an element-wise subtraction of two matrices, this is
// equivalent to iterating over every element of m1 and subtracting the corresponding value of m2.
func (m1 *Mat3) Sub(m2 *Mat3) Mat3 {
	return Mat3{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7], m1[8] - m2[8]}
}

// SubOf is a memory friendly version of Sub.
func (m1 *Mat3) SubOf(m2, m3 *Mat3) {
	m1[0] = m2[0] - m3[0]
	m1[1] = m2[1] - m3[1]
	m1[2] = m2[2] - m3[2]
	m1[3] = m2[3] - m3[3]
	m1[4] = m2[4] - m3[4]
	m1[5] = m2[5] - m3[5]
	m1[6] = m2[6] - m3[6]
	m1[7] = m2[7] - m3[7]
	m1[8] = m2[8] - m3[8]
}

// SubWith is a memory friendly version of Sub.
func (m1 *Mat3) SubWith(m2 *Mat3) {
	m1[0] -= m2[0]
	m1[1] -= m2[1]
	m1[2] -= m2[2]
	m1[3] -= m2[3]
	m1[4] -= m2[4]
	m1[5] -= m2[5]
	m1[6] -= m2[6]
	m1[7] -= m2[7]
	m1[8] -= m2[8]
}

// Mul performs a scalar multiplcation of the matrix. This is equivalent to iterating
// over every element of the matrix and multiply it by c.
func (m1 *Mat3) Mul(c float32) Mat3 {
	return Mat3{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c, m1[8] * c}
}

// MulOf is a memory friendly version fo Mul.
func (m1 *Mat3) MulOf(m2 *Mat3, c float32) {
	m1[0] = m2[0] * c
	m1[1] = m2[1] * c
	m1[2] = m2[2] * c
	m1[3] = m2[3] * c
	m1[4] = m2[4] * c
	m1[5] = m2[5] * c
	m1[6] = m2[6] * c
	m1[7] = m2[7] * c
	m1[8] = m2[8] * c
}

// MulWith is a memory friendly version fo Mul.
func (m1 *Mat3) MulWith(c float32) {
	m1[0] *= c
	m1[1] *= c
	m1[2] *= c
	m1[3] *= c
	m1[4] *= c
	m1[5] *= c
	m1[6] *= c
	m1[7] *= c
	m1[8] *= c
}

// Mul3x1 performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 *Mat3) Mul3x1(m2 *Vec3) Vec3 {
	return Vec3{
		m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2],
		m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2],
		m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2],
	}
}

// Mul3x1Transpose is the same as Mul3x1 except it uses the inplace transpose of
// this matrix.
func (m1 *Mat3) Mul3x1Transpose(v *Vec3) Vec3 {
	return Vec3{
		m1[0]*v[0] + m1[1]*v[1] + m1[2]*v[2],
		m1[3]*v[0] + m1[4]*v[1] + m1[5]*v[2],
		m1[6]*v[0] + m1[7]*v[1] + m1[8]*v[2],
	}
}

// Mul3x1In is a memory friendly version of Mul3x1
func (m1 *Mat3) Mul3x1In(m2, dst *Vec3) {
	dst[0] = m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2]
	dst[1] = m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2]
	dst[2] = m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2]
}

// Mul3 performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 *Mat3) Mul3(m2 *Mat3) Mat3 {
	return Mat3{
		m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2],
		m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2],
		m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2],
		m1[0]*m2[3] + m1[3]*m2[4] + m1[6]*m2[5],
		m1[1]*m2[3] + m1[4]*m2[4] + m1[7]*m2[5],
		m1[2]*m2[3] + m1[5]*m2[4] + m1[8]*m2[5],
		m1[0]*m2[6] + m1[3]*m2[7] + m1[6]*m2[8],
		m1[1]*m2[6] + m1[4]*m2[7] + m1[7]*m2[8],
		m1[2]*m2[6] + m1[5]*m2[7] + m1[8]*m2[8],
	}
}

// Mul3Of is a memory friendly version of Mul3.
func (m1 *Mat3) Mul3Of(m2, m3 *Mat3) {
	m1[0] = m2[0]*m3[0] + m2[3]*m3[1] + m2[6]*m3[2]
	m1[1] = m2[1]*m3[0] + m2[4]*m3[1] + m2[7]*m3[2]
	m1[2] = m2[2]*m3[0] + m2[5]*m3[1] + m2[8]*m3[2]
	m1[3] = m2[0]*m3[3] + m2[3]*m3[4] + m2[6]*m3[5]
	m1[4] = m2[1]*m3[3] + m2[4]*m3[4] + m2[7]*m3[5]
	m1[5] = m2[2]*m3[3] + m2[5]*m3[4] + m2[8]*m3[5]
	m1[6] = m2[0]*m3[6] + m2[3]*m3[7] + m2[6]*m3[8]
	m1[7] = m2[1]*m3[6] + m2[4]*m3[7] + m2[7]*m3[8]
	m1[8] = m2[2]*m3[6] + m2[5]*m3[7] + m2[8]*m3[8]
}

// Mul3With is a memory friendly version of Mul3.
func (m1 *Mat3) Mul3With(m2 *Mat3) {
	v0 := m1[0]
	v1 := m1[1]
	v2 := m1[2]
	v3 := m1[3]
	v4 := m1[4]
	v5 := m1[5]
	v6 := m1[6]
	v7 := m1[7]
	v8 := m1[8]
	m1[0] = v0*m2[0] + v3*m2[1] + v6*m2[2]
	m1[1] = v1*m2[0] + v4*m2[1] + v7*m2[2]
	m1[2] = v2*m2[0] + v5*m2[1] + v8*m2[2]
	m1[3] = v0*m2[3] + v3*m2[4] + v6*m2[5]
	m1[4] = v1*m2[3] + v4*m2[4] + v7*m2[5]
	m1[5] = v2*m2[3] + v5*m2[4] + v8*m2[5]
	m1[6] = v0*m2[6] + v3*m2[7] + v6*m2[8]
	m1[7] = v1*m2[6] + v4*m2[7] + v7*m2[8]
	m1[8] = v2*m2[6] + v5*m2[7] + v8*m2[8]
}

// Transposed produces the transpose of this matrix. For any MxN matrix
// the transpose is an NxM matrix with the rows swapped with the columns. For instance
// the transpose of the Mat3x2 is a Mat2x3 like so:
//
//    [[a b]]    [[a c e]]
//    [[c d]] =  [[b d f]]
//    [[e f]]
func (m1 *Mat3) Transposed() Mat3 {
	return Mat3{m1[0], m1[3], m1[6], m1[1], m1[4], m1[7], m1[2], m1[5], m1[8]}
}

// Transpose is a memory friendly version of Transposed.
func (m1 *Mat3) Transpose() {
	v1 := m1[1]
	v2 := m1[2]
	v5 := m1[5]
	v3 := m1[3]
	v6 := m1[6]
	v7 := m1[7]
	m1[1] = v3
	m1[2] = v6
	m1[3] = v1
	m1[5] = v7
	m1[6] = v2
	m1[7] = v5
}

// TransposeOf is a memory friendly version of Transposed.
func (m1 *Mat3) TransposeOf(m2 *Mat3) {
	m1[1] = m2[3]
	m1[2] = m2[6]
	m1[3] = m2[1]
	m1[5] = m2[7]
	m1[6] = m2[2]
	m1[7] = m2[5]
}

// Det returns the determinant of a matrix. The determinant is a measure of a square matrix's
// singularity and invertability, among other things. In this library, the
// determinant is hard coded based on pre-computed cofactor expansion, and uses
// no loops. Of course, the addition and multiplication must still be done.
func (m1 *Mat3) Det() float32 {
	return m1[0]*m1[4]*m1[8] + m1[3]*m1[7]*m1[2] + m1[6]*m1[1]*m1[5] - m1[6]*m1[4]*m1[2] - m1[3]*m1[1]*m1[8] - m1[0]*m1[7]*m1[5]
}

// Inverse computes the inverse of a square matrix. An inverse is a square matrix such that when multiplied by the
// original, yields the identity.
//
// M_inv * M = M * M_inv = I
//
// In this library, the math is precomputed, and uses no loops, though the multiplications, additions, determinant calculation, and scaling
// are still done. This can still be (relatively) expensive for a 4x4.
//
// This function checks the determinant to see if the matrix is invertible.
// If the determinant is 0.0, this function returns the zero matrix. However, due to floating point errors, it is
// entirely plausible to get a false positive or negative.
// In the future, an alternate function may be written which takes in a pre-computed determinant.
func (m1 *Mat3) Inverse() Mat3 {
	det := m1.Det()
	if FloatEqual(det, float32(0.0)) {
		return Mat3{}
	}

	retMat := Mat3{
		m1[4]*m1[8] - m1[5]*m1[7],
		m1[2]*m1[7] - m1[1]*m1[8],
		m1[1]*m1[5] - m1[2]*m1[4],
		m1[5]*m1[6] - m1[3]*m1[8],
		m1[0]*m1[8] - m1[2]*m1[6],
		m1[2]*m1[3] - m1[0]*m1[5],
		m1[3]*m1[7] - m1[4]*m1[6],
		m1[1]*m1[6] - m1[0]*m1[7],
		m1[0]*m1[4] - m1[1]*m1[3],
	}

	return retMat.Mul(1 / det)
}

// Invert is a memory friendly version of Inverse.
func (m1 *Mat3) Invert() {
	det := m1.Det()
	if FloatEqual(det, float32(0.0)) {
		m1[0] = 0
		m1[1] = 0
		m1[2] = 0
		m1[3] = 0
		m1[4] = 0
		m1[5] = 0
		m1[6] = 0
		m1[7] = 0
		m1[8] = 0
	}

	v0 := m1[0]
	v1 := m1[1]
	v2 := m1[2]
	v3 := m1[3]
	v4 := m1[4]
	v5 := m1[5]
	v6 := m1[6]
	v7 := m1[7]
	v8 := m1[8]
	m1[0] = v4*v8 - v5*v7
	m1[1] = v2*v7 - v1*v8
	m1[2] = v1*v5 - v2*v4
	m1[3] = v5*v6 - v3*v8
	m1[4] = v0*v8 - v2*v6
	m1[5] = v2*v3 - v0*v5
	m1[6] = v3*v7 - v4*v6
	m1[7] = v1*v6 - v0*v7
	m1[8] = v0*v4 - v1*v3

	m1.MulWith(1.0 / det)
}

// InverseOf is a memory friendly version fo Inverse.
func (m1 *Mat3) InverseOf(m2 *Mat3) {
	det := m2.Det()
	if FloatEqual(det, float32(0.0)) {
		return
	}

	v0 := m2[0]
	v1 := m2[1]
	v2 := m2[2]
	v3 := m2[3]
	v4 := m2[4]
	v5 := m2[5]
	v6 := m2[6]
	v7 := m2[7]
	v8 := m2[8]
	m1[0] = v4*v8 - v5*v7
	m1[1] = v2*v7 - v1*v8
	m1[2] = v1*v5 - v2*v4
	m1[3] = v5*v6 - v3*v8
	m1[4] = v0*v8 - v2*v6
	m1[5] = v2*v3 - v0*v5
	m1[6] = v3*v7 - v4*v6
	m1[7] = v1*v6 - v0*v7
	m1[8] = v0*v4 - v1*v3

	m1.MulWith(1.0 / det)
}

// ApproxEqual performs an element-wise approximate equality test between two matrices,
// as if FloatEqual had been used.
func (m1 *Mat3) ApproxEqual(m2 *Mat3) bool {
	return FloatEqual(m1[0], m2[0]) && FloatEqual(m1[1], m2[1]) && FloatEqual(m1[2], m2[2]) &&
		FloatEqual(m1[3], m2[3]) && FloatEqual(m1[4], m2[4]) && FloatEqual(m1[5], m2[5]) &&
		FloatEqual(m1[6], m2[6]) && FloatEqual(m1[7], m2[7]) && FloatEqual(m1[8], m2[8])

}

// ApproxEqualThreshold performs an element-wise approximate equality test between two matrices
// with a given epsilon threshold, as if FloatEqualThreshold had been used.
func (m1 *Mat3) ApproxEqualThreshold(m2 *Mat3, threshold float32) bool {
	return FloatEqualThreshold(m1[0], m2[0], threshold) && FloatEqualThreshold(m1[1], m2[1], threshold) && FloatEqualThreshold(m1[2], m2[2], threshold) &&
		FloatEqualThreshold(m1[3], m2[3], threshold) && FloatEqualThreshold(m1[4], m2[4], threshold) && FloatEqualThreshold(m1[5], m2[5], threshold) &&
		FloatEqualThreshold(m1[6], m2[6], threshold) && FloatEqualThreshold(m1[7], m2[7], threshold) && FloatEqualThreshold(m1[8], m2[8], threshold)
}

// ApproxFuncEqual performs an element-wise approximate equality test between two matrices
// with a given equality functions, intended to be used with FloatEqualFunc; although and comparison
// function may be used in practice.
func (m1 *Mat3) ApproxFuncEqual(m2 *Mat3, eq func(float32, float32) bool) bool {
	return eq(m1[0], m2[0]) && eq(m1[1], m2[1]) && eq(m1[2], m2[2]) &&
		eq(m1[3], m2[3]) && eq(m1[4], m2[4]) && eq(m1[5], m2[5]) &&
		eq(m1[6], m2[6]) && eq(m1[7], m2[7]) && eq(m1[8], m2[8])
}

// At returns the matrix element at the given row and column.
// This is equivalent to mat[col * numRow + row] where numRow is constant
// (E.G. for a Mat3x2 it's equal to 3)
//
// This method is garbage-in garbage-out. For instance, on a Mat4 asking for
// At(5,0) will work just like At(1,1). Or it may panic if it's out of bounds.
func (m1 *Mat3) At(row, col int) float32 {
	return m1[col*3+row]
}

// Set sets the corresponding matrix element at the given row and column.
// This has a pointer receiver because it mutates the matrix.
//
// This method is garbage-in garbage-out. For instance, on a Mat4 asking for
// Set(5,0,val) will work just like Set(1,1,val). Or it may panic if it's out of bounds.
func (m1 *Mat3) Set(row, col int, value float32) {
	m1[col*3+row] = value
}

// Index returns the index of the given row and column, to be used with direct
// access. E.G. Index(0,0) = 0.
//
// This is a garbage-in garbage-out method. For instance, on a Mat4 asking for the index of
// (5,0) will work the same as asking for (1,1). Or it may give you a value that will cause
// a panic if you try to access the array with it if it's truly out of bounds.
func (Mat3) Index(row, col int) int {
	return col*3 + row
}

// Row returns a vector representing the corresponding row (starting at row 0).
// This package makes no distinction between row and column vectors, so it
// will be a normal VecM for a MxN matrix.
func (m1 *Mat3) Row(row int) Vec3 {
	return Vec3{m1[row+0], m1[row+3], m1[row+6]}
}

// Rows decomposes a matrix into its corresponding row vectors.
// This is equivalent to calling mat.Row for each row.
func (m1 *Mat3) Rows() (row0, row1, row2 Vec3) {
	return m1.Row(0), m1.Row(1), m1.Row(2)
}

// Col returns a vector representing the corresponding column (starting at col 0).
// This package makes no distinction between row and column vectors, so it
// will be a normal VecN for a MxN matrix.
func (m1 *Mat3) Col(col int) Vec3 {
	return Vec3{m1[col*3+0], m1[col*3+1], m1[col*3+2]}
}

// Cols decomposes a matrix into its corresponding column vectors.
// This is equivalent to calling mat.Col for each column.
func (m1 *Mat3) Cols() (col0, col1, col2 Vec3) {
	return m1.Col(0), m1.Col(1), m1.Col(2)
}

// Trace is a basic operation on a square matrix that simply
// sums up all elements on the main diagonal (meaning all elements such that row==col).
func (m1 *Mat3) Trace() float32 {
	return m1[0] + m1[4] + m1[8]
}

// Abs returns the element-wise absolute value of this matrix
func (m1 *Mat3) Abs() Mat3 {
	return Mat3{math.Abs(m1[0]), math.Abs(m1[1]), math.Abs(m1[2]),
		math.Abs(m1[3]), math.Abs(m1[4]), math.Abs(m1[5]),
		math.Abs(m1[6]), math.Abs(m1[7]), math.Abs(m1[8])}
}

// AbsSelf is a memory friendly version of Abs.
func (m1 *Mat3) AbsSelf() {
	m1[0] = math.Abs(m1[0])
	m1[1] = math.Abs(m1[1])
	m1[2] = math.Abs(m1[2])
	m1[3] = math.Abs(m1[3])
	m1[4] = math.Abs(m1[4])
	m1[5] = math.Abs(m1[5])
	m1[6] = math.Abs(m1[6])
	m1[7] = math.Abs(m1[7])
	m1[8] = math.Abs(m1[8])
}

// AbsOf is a memory friendly version of Abs.
func (m1 *Mat3) AbsOf(m2 *Mat3) {
	m1[0] = math.Abs(m2[0])
	m1[1] = math.Abs(m2[1])
	m1[2] = math.Abs(m2[2])
	m1[3] = math.Abs(m2[3])
	m1[4] = math.Abs(m2[4])
	m1[5] = math.Abs(m2[5])
	m1[6] = math.Abs(m2[6])
	m1[7] = math.Abs(m2[7])
	m1[8] = math.Abs(m2[8])
}

// Iden sets this matrix as the identity matrix.
func (m1 *Mat3) Iden() {
	m1[0] = 1
	m1[1] = 0
	m1[2] = 0
	m1[3] = 0
	m1[4] = 1
	m1[5] = 0
	m1[6] = 0
	m1[7] = 0
	m1[8] = 1
}

// SetOrientation sets this matrix to the orientation matrix represented by that quaternion.
func (m1 *Mat3) SetOrientation(q1 *Quat) {
	w, x, y, z := q1.W, q1.V[0], q1.V[1], q1.V[2]
	m1[0] = 1 - 2*y*y - 2*z*z
	m1[1] = 2*x*y + 2*w*z
	m1[2] = 2*x*z - 2*w*y
	m1[3] = 2*x*y - 2*w*z
	m1[4] = 1 - 2*x*x - 2*z*z
	m1[5] = 2*y*z + 2*w*x
	m1[6] = 2*x*z + 2*w*y
	m1[7] = 2*y*z - 2*w*x
	m1[8] = 1 - 2*x*x - 2*y*y
}

// String pretty prints the matrix
func (m1 *Mat3) String() string {
	buf := new(bytes.Buffer)
	w := tabwriter.NewWriter(buf, 4, 4, 1, ' ', tabwriter.AlignRight)
	for i := 0; i < 3; i++ {
		row := m1.Row(i)
		for _, col := range []float32{row[0], row[1], row[2]} {
			fmt.Fprintf(w, "%f\t", col)
		}

		fmt.Fprintln(w, "")
	}
	w.Flush()

	return buf.String()
}

// SetCol sets a Column within the Matrix, so it mutates the calling matrix.
func (m1 *Mat4) SetCol(col int, v *Vec4) {
	m1[col*4+0], m1[col*4+1], m1[col*4+2], m1[col*4+3] = v[0], v[1], v[2], v[3]
}

// SetRow sets a Row within the Matrix, so it mutates the calling matrix.
func (m1 *Mat4) SetRow(row int, v *Vec4) {
	m1[row+0], m1[row+4], m1[row+8], m1[row+12] = v[0], v[1], v[2], v[3]
}

// Diag is a basic operation on a square matrix that simply
// returns main diagonal (meaning all elements such that row==col).
func (m1 *Mat4) Diag() Vec4 {
	return Vec4{m1[0], m1[5], m1[10], m1[15]}
}

// Ident4 returns the 4x4 identity matrix.
// The identity matrix is a square matrix with the value 1 on its
// diagonals. The characteristic property of the identity matrix is that
// any matrix multiplied by it is itself. (MI = M; IN = N)
func Ident4() Mat4 {
	return Mat4{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}
}

// Diag4 creates a diagonal matrix from the entries of the input vector.
// That is, for each pointer for row==col, vector[row] is the entry. Otherwise it's 0.
//
// Another way to think about it is that the identity is this function where the every vector element is 1.
func Diag4(v *Vec4) Mat4 {
	return Mat4{v[0], 0, 0, 0, 0, v[1], 0, 0, 0, 0, v[2], 0, 0, 0, 0, v[3]}
}

// Mat4FromRows builds a new matrix from row vectors.
// The resulting matrix will still be in column major order, but this can be
// good for hand-building matrices.
func Mat4FromRows(row0, row1, row2, row3 *Vec4) Mat4 {
	return Mat4{row0[0], row1[0], row2[0], row3[0], row0[1], row1[1], row2[1], row3[1], row0[2], row1[2], row2[2], row3[2], row0[3], row1[3], row2[3], row3[3]}
}

// Mat4FromCols builds a new matrix from column vectors.
func Mat4FromCols(col0, col1, col2, col3 *Vec4) Mat4 {
	return Mat4{col0[0], col0[1], col0[2], col0[3], col1[0], col1[1], col1[2], col1[3], col2[0], col2[1], col2[2], col2[3], col3[0], col3[1], col3[2], col3[3]}
}

// Add performs an element-wise addition of two matrices, this is
// equivalent to iterating over every element of m1 and adding the corresponding value of m2.
func (m1 *Mat4) Add(m2 *Mat4) Mat4 {
	return Mat4{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5], m1[6] + m2[6], m1[7] + m2[7], m1[8] + m2[8], m1[9] + m2[9], m1[10] + m2[10], m1[11] + m2[11], m1[12] + m2[12], m1[13] + m2[13], m1[14] + m2[14], m1[15] + m2[15]}
}

// AddOf is a memory friendly version of Add.
func (m1 *Mat4) AddOf(m2, m3 *Mat4) {
	m1[0] = m2[0] + m3[0]
	m1[1] = m2[1] + m3[1]
	m1[2] = m2[2] + m3[2]
	m1[3] = m2[3] + m3[3]
	m1[4] = m2[4] + m3[4]
	m1[5] = m2[5] + m3[5]
	m1[6] = m2[6] + m3[6]
	m1[7] = m2[7] + m3[7]
	m1[8] = m2[8] + m3[8]
	m1[9] = m2[9] + m3[9]
	m1[10] = m2[10] + m3[10]
	m1[11] = m2[11] + m3[11]
	m1[12] = m2[12] + m3[12]
	m1[13] = m2[13] + m3[13]
	m1[14] = m2[14] + m3[14]
	m1[15] = m2[15] + m3[15]
}

// AddWith is a memory friendly version of Add.
func (m1 *Mat4) AddWith(m2 *Mat4) {
	m1[0] += m2[0]
	m1[1] += m2[1]
	m1[2] += m2[2]
	m1[3] += m2[3]
	m1[4] += m2[4]
	m1[5] += m2[5]
	m1[6] += m2[6]
	m1[7] += m2[7]
	m1[8] += m2[8]
	m1[9] += m2[9]
	m1[10] += m2[10]
	m1[11] += m2[11]
	m1[12] += m2[12]
	m1[13] += m2[13]
	m1[14] += m2[14]
	m1[15] += m2[15]
}

// Sub performs an element-wise subtraction of two matrices, this is
// equivalent to iterating over every element of m1 and subtracting the corresponding value of m2.
func (m1 *Mat4) Sub(m2 *Mat4) Mat4 {
	return Mat4{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7], m1[8] - m2[8], m1[9] - m2[9], m1[10] - m2[10], m1[11] - m2[11], m1[12] - m2[12], m1[13] - m2[13], m1[14] - m2[14], m1[15] - m2[15]}
}

// SubOf is a memory friendly version of Sub.
func (m1 *Mat4) SubOf(m2, m3 *Mat4) {
	m1[0] = m2[0] - m3[0]
	m1[1] = m2[1] - m3[1]
	m1[2] = m2[2] - m3[2]
	m1[3] = m2[3] - m3[3]
	m1[4] = m2[4] - m3[4]
	m1[5] = m2[5] - m3[5]
	m1[6] = m2[6] - m3[6]
	m1[7] = m2[7] - m3[7]
	m1[8] = m2[8] - m3[8]
	m1[9] = m2[9] - m3[9]
	m1[10] = m2[10] - m3[10]
	m1[11] = m2[11] - m3[11]
	m1[12] = m2[12] - m3[12]
	m1[13] = m2[13] - m3[13]
	m1[14] = m2[14] - m3[14]
	m1[15] = m2[15] - m3[15]
}

// SubWith is a memory friendly version of Sub.
func (m1 *Mat4) SubWith(m2 *Mat4) {
	m1[0] -= m2[0]
	m1[1] -= m2[1]
	m1[2] -= m2[2]
	m1[3] -= m2[3]
	m1[4] -= m2[4]
	m1[5] -= m2[5]
	m1[6] -= m2[6]
	m1[7] -= m2[7]
	m1[8] -= m2[8]
	m1[9] -= m2[9]
	m1[10] -= m2[10]
	m1[11] -= m2[11]
	m1[12] -= m2[12]
	m1[13] -= m2[13]
	m1[14] -= m2[14]
	m1[15] -= m2[15]
}

// Mul performs a scalar multiplcation of the matrix. This is equivalent to iterating
// over every element of the matrix and multiply it by c.
func (m1 *Mat4) Mul(c float32) Mat4 {
	return Mat4{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c, m1[8] * c, m1[9] * c, m1[10] * c, m1[11] * c, m1[12] * c, m1[13] * c, m1[14] * c, m1[15] * c}
}

// MulOf is a memory friendly version of Mul.
func (m1 *Mat4) MulOf(m2 *Mat4, c float32) {
	m1[0] = m2[0] * c
	m1[1] = m2[1] * c
	m1[2] = m2[2] * c
	m1[3] = m2[3] * c
	m1[4] = m2[4] * c
	m1[5] = m2[5] * c
	m1[6] = m2[6] * c
	m1[7] = m2[7] * c
	m1[8] = m2[8] * c
	m1[9] = m2[9] * c
	m1[10] = m2[10] * c
	m1[11] = m2[11] * c
	m1[12] = m2[12] * c
	m1[13] = m2[13] * c
	m1[14] = m2[14] * c
	m1[15] = m2[15] * c
}

// MulWith is a memory friendly version of Mul.
func (m1 *Mat4) MulWith(c float32) {
	m1[0] *= c
	m1[1] *= c
	m1[2] *= c
	m1[3] *= c
	m1[4] *= c
	m1[5] *= c
	m1[6] *= c
	m1[7] *= c
	m1[8] *= c
	m1[9] *= c
	m1[10] *= c
	m1[11] *= c
	m1[12] *= c
	m1[13] *= c
	m1[14] *= c
	m1[15] *= c
}

// Mul4x1 performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 *Mat4) Mul4x1(m2 *Vec4) Vec4 {
	return Vec4{
		m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2] + m1[12]*m2[3],
		m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2] + m1[13]*m2[3],
		m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2] + m1[14]*m2[3],
		m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2] + m1[15]*m2[3],
	}
}

// Mul4 performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 *Mat4) Mul4(m2 *Mat4) Mat4 {
	return Mat4{
		m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2] + m1[12]*m2[3],
		m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2] + m1[13]*m2[3],
		m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2] + m1[14]*m2[3],
		m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2] + m1[15]*m2[3],
		m1[0]*m2[4] + m1[4]*m2[5] + m1[8]*m2[6] + m1[12]*m2[7],
		m1[1]*m2[4] + m1[5]*m2[5] + m1[9]*m2[6] + m1[13]*m2[7],
		m1[2]*m2[4] + m1[6]*m2[5] + m1[10]*m2[6] + m1[14]*m2[7],
		m1[3]*m2[4] + m1[7]*m2[5] + m1[11]*m2[6] + m1[15]*m2[7],
		m1[0]*m2[8] + m1[4]*m2[9] + m1[8]*m2[10] + m1[12]*m2[11],
		m1[1]*m2[8] + m1[5]*m2[9] + m1[9]*m2[10] + m1[13]*m2[11],
		m1[2]*m2[8] + m1[6]*m2[9] + m1[10]*m2[10] + m1[14]*m2[11],
		m1[3]*m2[8] + m1[7]*m2[9] + m1[11]*m2[10] + m1[15]*m2[11],
		m1[0]*m2[12] + m1[4]*m2[13] + m1[8]*m2[14] + m1[12]*m2[15],
		m1[1]*m2[12] + m1[5]*m2[13] + m1[9]*m2[14] + m1[13]*m2[15],
		m1[2]*m2[12] + m1[6]*m2[13] + m1[10]*m2[14] + m1[14]*m2[15],
		m1[3]*m2[12] + m1[7]*m2[13] + m1[11]*m2[14] + m1[15]*m2[15],
	}
}

// Mul4Of is a memory friendly version fo Mul4.
func (m1 *Mat4) Mul4Of(m2, m3 *Mat4) {
	m1[0] = m2[0]*m3[0] + m2[4]*m3[1] + m2[8]*m3[2] + m2[12]*m3[3]
	m1[1] = m2[1]*m3[0] + m2[5]*m3[1] + m2[9]*m3[2] + m2[13]*m3[3]
	m1[2] = m2[2]*m3[0] + m2[6]*m3[1] + m2[10]*m3[2] + m2[14]*m3[3]
	m1[3] = m2[3]*m3[0] + m2[7]*m3[1] + m2[11]*m3[2] + m2[15]*m3[3]
	m1[4] = m2[0]*m3[4] + m2[4]*m3[5] + m2[8]*m3[6] + m2[12]*m3[7]
	m1[5] = m2[1]*m3[4] + m2[5]*m3[5] + m2[9]*m3[6] + m2[13]*m3[7]
	m1[6] = m2[2]*m3[4] + m2[6]*m3[5] + m2[10]*m3[6] + m2[14]*m3[7]
	m1[7] = m2[3]*m3[4] + m2[7]*m3[5] + m2[11]*m3[6] + m2[15]*m3[7]
	m1[8] = m2[0]*m3[8] + m2[4]*m3[9] + m2[8]*m3[10] + m2[12]*m3[11]
	m1[9] = m2[1]*m3[8] + m2[5]*m3[9] + m2[9]*m3[10] + m2[13]*m3[11]
	m1[10] = m2[2]*m3[8] + m2[6]*m3[9] + m2[10]*m3[10] + m2[14]*m3[11]
	m1[11] = m2[3]*m3[8] + m2[7]*m3[9] + m2[11]*m3[10] + m2[15]*m3[11]
	m1[12] = m2[0]*m3[12] + m2[4]*m3[13] + m2[8]*m3[14] + m2[12]*m3[15]
	m1[13] = m2[1]*m3[12] + m2[5]*m3[13] + m2[9]*m3[14] + m2[13]*m3[15]
	m1[14] = m2[2]*m3[12] + m2[6]*m3[13] + m2[10]*m3[14] + m2[14]*m3[15]
	m1[15] = m2[3]*m3[12] + m2[7]*m3[13] + m2[11]*m3[14] + m2[15]*m3[15]
}

// Mul4With is a memory friendly version fo Mul4.
func (m1 *Mat4) Mul4With(m2 *Mat4) {
	v0 := m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2] + m1[12]*m2[3]
	v1 := m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2] + m1[13]*m2[3]
	v2 := m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2] + m1[14]*m2[3]
	v3 := m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2] + m1[15]*m2[3]
	v4 := m1[0]*m2[4] + m1[4]*m2[5] + m1[8]*m2[6] + m1[12]*m2[7]
	v5 := m1[1]*m2[4] + m1[5]*m2[5] + m1[9]*m2[6] + m1[13]*m2[7]
	v6 := m1[2]*m2[4] + m1[6]*m2[5] + m1[10]*m2[6] + m1[14]*m2[7]
	v7 := m1[3]*m2[4] + m1[7]*m2[5] + m1[11]*m2[6] + m1[15]*m2[7]
	v8 := m1[0]*m2[8] + m1[4]*m2[9] + m1[8]*m2[10] + m1[12]*m2[11]
	v9 := m1[1]*m2[8] + m1[5]*m2[9] + m1[9]*m2[10] + m1[13]*m2[11]
	v10 := m1[2]*m2[8] + m1[6]*m2[9] + m1[10]*m2[10] + m1[14]*m2[11]
	v11 := m1[3]*m2[8] + m1[7]*m2[9] + m1[11]*m2[10] + m1[15]*m2[11]
	v12 := m1[0]*m2[12] + m1[4]*m2[13] + m1[8]*m2[14] + m1[12]*m2[15]
	v13 := m1[1]*m2[12] + m1[5]*m2[13] + m1[9]*m2[14] + m1[13]*m2[15]
	v14 := m1[2]*m2[12] + m1[6]*m2[13] + m1[10]*m2[14] + m1[14]*m2[15]
	v15 := m1[3]*m2[12] + m1[7]*m2[13] + m1[11]*m2[14] + m1[15]*m2[15]

	m1[0] = v0
	m1[1] = v1
	m1[2] = v2
	m1[3] = v3
	m1[4] = v4
	m1[5] = v5
	m1[6] = v6
	m1[7] = v7
	m1[8] = v8
	m1[9] = v9
	m1[10] = v10
	m1[11] = v11
	m1[12] = v12
	m1[13] = v13
	m1[14] = v14
	m1[15] = v15
}

// Transposed produces the transpose of this matrix. For any MxN matrix
// the transpose is an NxM matrix with the rows swapped with the columns. For instance
// the transpose of the Mat3x2 is a Mat2x3 like so:
//
//    [[a b]]    [[a c e]]
//    [[c d]] =  [[b d f]]
//    [[e f]]
func (m1 *Mat4) Transposed() Mat4 {
	return Mat4{m1[0], m1[4], m1[8], m1[12],
		m1[1], m1[5], m1[9], m1[13],
		m1[2], m1[6], m1[10], m1[14],
		m1[3], m1[7], m1[11], m1[15]}
}

// TransposeOf is a memory friendly version of Transposed.
func (m1 *Mat4) TransposeOf(m2 *Mat4) {
	m1[0] = m2[0]
	m1[1] = m2[4]
	m1[2] = m2[8]
	m1[3] = m2[12]
	m1[4] = m2[1]
	m1[5] = m2[5]
	m1[6] = m2[9]
	m1[7] = m2[13]
	m1[8] = m2[2]
	m1[9] = m2[6]
	m1[10] = m2[10]
	m1[11] = m2[14]
	m1[12] = m2[3]
	m1[13] = m2[7]
	m1[14] = m2[11]
	m1[15] = m2[15]
}

// Transpose is a memory friendly version of Transposed.
func (m1 *Mat4) Transpose() {
	//m1[0] = m1[0]
	v1 := m1[1]
	m1[1] = m1[4]
	v2 := m1[2]
	m1[2] = m1[8]
	v3 := m1[3]
	m1[3] = m1[12]
	m1[4] = v1
	//m1[5] = m1[5]
	v6 := m1[6]
	m1[6] = m1[9]
	v7 := m1[7]
	m1[7] = m1[13]
	m1[8] = v2
	m1[9] = v6
	//m1[10] = m1[10]
	v11 := m1[11]
	m1[11] = m1[14]
	m1[12] = v3
	m1[13] = v7
	m1[14] = v11
	//m1[15] = m1[15]
}

// Det returns the determinant of a matrix. The determinant is a measure of a square matrix's
// singularity and invertability, among other things. In this library, the
// determinant is hard coded based on pre-computed cofactor expansion, and uses
// no loops. Of course, the addition and multiplication must still be done.
func (m1 *Mat4) Det() float32 {
	return m1[0]*m1[5]*m1[10]*m1[15] - m1[0]*m1[5]*m1[11]*m1[14] - m1[0]*m1[6]*m1[9]*m1[15] + m1[0]*m1[6]*m1[11]*m1[13] + m1[0]*m1[7]*m1[9]*m1[14] - m1[0]*m1[7]*m1[10]*m1[13] - m1[1]*m1[4]*m1[10]*m1[15] + m1[1]*m1[4]*m1[11]*m1[14] + m1[1]*m1[6]*m1[8]*m1[15] - m1[1]*m1[6]*m1[11]*m1[12] - m1[1]*m1[7]*m1[8]*m1[14] + m1[1]*m1[7]*m1[10]*m1[12] + m1[2]*m1[4]*m1[9]*m1[15] - m1[2]*m1[4]*m1[11]*m1[13] - m1[2]*m1[5]*m1[8]*m1[15] + m1[2]*m1[5]*m1[11]*m1[12] + m1[2]*m1[7]*m1[8]*m1[13] - m1[2]*m1[7]*m1[9]*m1[12] - m1[3]*m1[4]*m1[9]*m1[14] + m1[3]*m1[4]*m1[10]*m1[13] + m1[3]*m1[5]*m1[8]*m1[14] - m1[3]*m1[5]*m1[10]*m1[12] - m1[3]*m1[6]*m1[8]*m1[13] + m1[3]*m1[6]*m1[9]*m1[12]
}

// Inverse computes the inverse of a square matrix. An inverse is a square matrix such that when multiplied by the
// original, yields the identity.
//
// M_inv * M = M * M_inv = I
//
// In this library, the math is precomputed, and uses no loops, though the multiplications, additions, determinant calculation, and scaling
// are still done. This can still be (relatively) expensive for a 4x4.
//
// This function checks the determinant to see if the matrix is invertible.
// If the determinant is 0.0, this function returns the zero matrix. However, due to floating point errors, it is
// entirely plausible to get a false positive or negative.
// In the future, an alternate function may be written which takes in a pre-computed determinant.
func (m1 *Mat4) Inverse() Mat4 {
	det := m1.Det()
	if FloatEqual(det, float32(0.0)) {
		return Mat4{}
	}

	v0 := m1[0]
	v1 := m1[1]
	v2 := m1[2]
	v3 := m1[3]
	v4 := m1[4]
	v5 := m1[5]
	v6 := m1[6]
	v7 := m1[7]
	v8 := m1[8]
	v9 := m1[9]
	v10 := m1[10]
	v11 := m1[11]
	v12 := m1[12]
	v13 := m1[13]
	v14 := m1[14]
	v15 := m1[15]

	//precalculate the most common products
	v7v10 := v7 * v10
	v6v11 := v6 * v11
	v7v9 := v7 * v9
	v5v11 := v5 * v11
	v6v9 := v6 * v9
	v5v10 := v5 * v10
	v1v4 := v1 * v4
	v4v9 := v4 * v9
	v6v8 := v6 * v8
	v5v8 := v5 * v8
	v7v8 := v7 * v8
	v1v12 := v1 * v12
	v2v12 := v2 * v12
	v2v13 := v2 * v13
	v2v15 := v2 * v15
	v3v12 := v3 * v12
	v3v13 := v3 * v13
	v3v14 := v3 * v14
	v4v10 := v4 * v10
	v4v11 := v4 * v11
	v10v15 := v10 * v15
	v7v14 := v7 * v14
	v6v15 := v6 * v15
	v0v11 := v0 * v11
	v1v8 := v1 * v8
	v0v9 := v0 * v9
	v0v5 := v0 * v5
	v0v13 := v0 * v13

	retMat := Mat4{
		-v7v10*v13 + v6v11*v13 + v7v9*v14 - v5v11*v14 - v6v9*v15 + v5v10*v15,
		v3v13*v10 - v2v13*v11 - v3v14*v9 + v1*v11*v14 + v2v15*v9 - v1*v10v15,
		-v3v13*v6 + v2v13*v7 + v3v14*v5 - v1*v7v14 - v2v15*v5 + v1*v6v15,
		v3*v6v9 - v2*v7v9 - v3*v5v10 + v1*v7v10 + v2*v5v11 - v1*v6v11,
		v7v10*v12 - v6v11*v12 - v7v8*v14 + v4v11*v14 + v6v8*v15 - v4v10*v15,
		-v3v12*v10 + v2v12*v11 + v3v14*v8 - v0v11*v14 - v2v15*v8 + v0*v10v15,
		v3v12*v6 - v2v12*v7 - v3v14*v4 + v0*v7v14 + v2v15*v4 - v0*v6v15,
		-v3*v6v8 + v2*v7v8 + v3*v4v10 - v0*v7v10 - v2*v4v11 + v0*v6v11,
		-v7v9*v12 + v5v11*v12 + v7v8*v13 - v4v11*v13 - v5v8*v15 + v4v9*v15,
		v3v12*v9 - v1v12*v11 - v3v13*v8 + v0v11*v13 + v1v8*v15 - v0v9*v15,
		-v3v12*v5 + v1v12*v7 + v3v13*v4 - v0v13*v7 - v1v4*v15 + v0v5*v15,
		v3*v5v8 - v1*v7v8 - v3*v4v9 + v0*v7v9 + v1v4*v11 - v0*v5v11,
		v6v9*v12 - v5v10*v12 - v6v8*v13 + v4v10*v13 + v5v8*v14 - v4v9*v14,
		-v2v12*v9 + v1v12*v10 + v2v13*v8 - v0v13*v10 - v1v8*v14 + v0v9*v14,
		v2v12*v5 - v1v12*v6 - v2v13*v4 + v0v13*v6 + v1v4*v14 - v0v5*v14,
		-v2*v5v8 + v1*v6v8 + v2*v4v9 - v0*v6v9 - v1v4*v10 + v0*v5v10,
	}
	//v2v4, v8v13 v8v14, v10v13, v4v10, v1v7, v4v11, v11v14

	return retMat.Mul(1.0 / det)
}

// Invert is a memory friendly version of Inverse.
func (m1 *Mat4) Invert() {
	det := m1.Det()
	if FloatEqual(det, float32(0.0)) {
		m1[0] = 0
		m1[1] = 0
		m1[2] = 0
		m1[3] = 0
		m1[4] = 0
		m1[5] = 0
		m1[6] = 0
		m1[7] = 0
		m1[8] = 0
		m1[9] = 0
		m1[10] = 0
		m1[11] = 0
		m1[12] = 0
		m1[13] = 0
		m1[14] = 0
		m1[15] = 0
		return
	}

	//m1ake a copy to not override original while reading
	v0 := m1[0]
	v1 := m1[1]
	v2 := m1[2]
	v3 := m1[3]
	v4 := m1[4]
	v5 := m1[5]
	v6 := m1[6]
	v7 := m1[7]
	v8 := m1[8]
	v9 := m1[9]
	v10 := m1[10]
	v11 := m1[11]
	v12 := m1[12]
	v13 := m1[13]
	v14 := m1[14]
	v15 := m1[15]

	//precalculate the most common products
	v7v10 := v7 * v10
	v6v11 := v6 * v11
	v7v9 := v7 * v9
	v5v11 := v5 * v11
	v6v9 := v6 * v9
	v5v10 := v5 * v10
	v1v4 := v1 * v4
	v4v9 := v4 * v9
	v6v8 := v6 * v8
	v5v8 := v5 * v8
	v7v8 := v7 * v8
	v1v12 := v1 * v12
	v2v12 := v2 * v12
	v2v13 := v2 * v13
	v2v15 := v2 * v15
	v3v12 := v3 * v12
	v3v13 := v3 * v13
	v3v14 := v3 * v14
	v4v10 := v4 * v10
	v4v11 := v4 * v11
	v10v15 := v10 * v15
	v7v14 := v7 * v14
	v6v15 := v6 * v15
	v0v11 := v0 * v11
	v1v8 := v1 * v8
	v0v9 := v0 * v9
	v0v5 := v0 * v5
	v0v13 := v0 * v13

	m1[0] = -v7v10*v13 + v6v11*v13 + v7v9*v14 - v5v11*v14 - v6v9*v15 + v5v10*v15
	m1[1] = v3v13*v10 - v2v13*v11 - v3v14*v9 + v1*v11*v14 + v2v15*v9 - v1*v10v15
	m1[2] = -v3v13*v6 + v2v13*v7 + v3v14*v5 - v1*v7v14 - v2v15*v5 + v1*v6v15
	m1[3] = v3*v6v9 - v2*v7v9 - v3*v5v10 + v1*v7v10 + v2*v5v11 - v1*v6v11
	m1[4] = v7v10*v12 - v6v11*v12 - v7v8*v14 + v4v11*v14 + v6v8*v15 - v4v10*v15
	m1[5] = -v3v12*v10 + v2v12*v11 + v3v14*v8 - v0v11*v14 - v2v15*v8 + v0*v10v15
	m1[6] = v3v12*v6 - v2v12*v7 - v3v14*v4 + v0*v7v14 + v2v15*v4 - v0*v6v15
	m1[7] = -v3*v6v8 + v2*v7v8 + v3*v4v10 - v0*v7v10 - v2*v4v11 + v0*v6v11
	m1[8] = -v7v9*v12 + v5v11*v12 + v7v8*v13 - v4v11*v13 - v5v8*v15 + v4v9*v15
	m1[9] = v3v12*v9 - v1v12*v11 - v3v13*v8 + v0v11*v13 + v1v8*v15 - v0v9*v15
	m1[10] = -v3v12*v5 + v1v12*v7 + v3v13*v4 - v0v13*v7 - v1v4*v15 + v0v5*v15
	m1[11] = v3*v5v8 - v1*v7v8 - v3*v4v9 + v0*v7v9 + v1v4*v11 - v0*v5v11
	m1[12] = v6v9*v12 - v5v10*v12 - v6v8*v13 + v4v10*v13 + v5v8*v14 - v4v9*v14
	m1[13] = -v2v12*v9 + v1v12*v10 + v2v13*v8 - v0v13*v10 - v1v8*v14 + v0v9*v14
	m1[14] = v2v12*v5 - v1v12*v6 - v2v13*v4 + v0v13*v6 + v1v4*v14 - v0v5*v14
	m1[15] = -v2*v5v8 + v1*v6v8 + v2*v4v9 - v0*v6v9 - v1v4*v10 + v0*v5v10
	m1.MulWith(1.0 / det)
}

// InverseOf is a memory friendly version of Inverse.
func (m1 *Mat4) InverseOf(m2 *Mat4) {
	det := m2.Det()
	if FloatEqual(det, float32(0.0)) {
		m1[0] = 0
		m1[1] = 0
		m1[2] = 0
		m1[3] = 0
		m1[4] = 0
		m1[5] = 0
		m1[6] = 0
		m1[7] = 0
		m1[8] = 0
		m1[9] = 0
		m1[10] = 0
		m1[11] = 0
		m1[12] = 0
		m1[13] = 0
		m1[14] = 0
		m1[15] = 0
		return
	}

	//m1ake a copy to not override original while reading
	v0 := m2[0]
	v1 := m2[1]
	v2 := m2[2]
	v3 := m2[3]
	v4 := m2[4]
	v5 := m2[5]
	v6 := m2[6]
	v7 := m2[7]
	v8 := m2[8]
	v9 := m2[9]
	v10 := m2[10]
	v11 := m2[11]
	v12 := m2[12]
	v13 := m2[13]
	v14 := m2[14]
	v15 := m2[15]

	//precalculate the most common products
	v7v10 := v7 * v10
	v6v11 := v6 * v11
	v7v9 := v7 * v9
	v5v11 := v5 * v11
	v6v9 := v6 * v9
	v5v10 := v5 * v10
	v1v4 := v1 * v4
	v4v9 := v4 * v9
	v6v8 := v6 * v8
	v5v8 := v5 * v8
	v7v8 := v7 * v8
	v1v12 := v1 * v12
	v2v12 := v2 * v12
	v2v13 := v2 * v13
	v2v15 := v2 * v15
	v3v12 := v3 * v12
	v3v13 := v3 * v13
	v3v14 := v3 * v14
	v4v10 := v4 * v10
	v4v11 := v4 * v11
	v10v15 := v10 * v15
	v7v14 := v7 * v14
	v6v15 := v6 * v15
	v0v11 := v0 * v11
	v1v8 := v1 * v8
	v0v9 := v0 * v9
	v0v5 := v0 * v5
	v0v13 := v0 * v13

	m1[0] = -v7v10*v13 + v6v11*v13 + v7v9*v14 - v5v11*v14 - v6v9*v15 + v5v10*v15
	m1[1] = v3v13*v10 - v2v13*v11 - v3v14*v9 + v1*v11*v14 + v2v15*v9 - v1*v10v15
	m1[2] = -v3v13*v6 + v2v13*v7 + v3v14*v5 - v1*v7v14 - v2v15*v5 + v1*v6v15
	m1[3] = v3*v6v9 - v2*v7v9 - v3*v5v10 + v1*v7v10 + v2*v5v11 - v1*v6v11
	m1[4] = v7v10*v12 - v6v11*v12 - v7v8*v14 + v4v11*v14 + v6v8*v15 - v4v10*v15
	m1[5] = -v3v12*v10 + v2v12*v11 + v3v14*v8 - v0v11*v14 - v2v15*v8 + v0*v10v15
	m1[6] = v3v12*v6 - v2v12*v7 - v3v14*v4 + v0*v7v14 + v2v15*v4 - v0*v6v15
	m1[7] = -v3*v6v8 + v2*v7v8 + v3*v4v10 - v0*v7v10 - v2*v4v11 + v0*v6v11
	m1[8] = -v7v9*v12 + v5v11*v12 + v7v8*v13 - v4v11*v13 - v5v8*v15 + v4v9*v15
	m1[9] = v3v12*v9 - v1v12*v11 - v3v13*v8 + v0v11*v13 + v1v8*v15 - v0v9*v15
	m1[10] = -v3v12*v5 + v1v12*v7 + v3v13*v4 - v0v13*v7 - v1v4*v15 + v0v5*v15
	m1[11] = v3*v5v8 - v1*v7v8 - v3*v4v9 + v0*v7v9 + v1v4*v11 - v0*v5v11
	m1[12] = v6v9*v12 - v5v10*v12 - v6v8*v13 + v4v10*v13 + v5v8*v14 - v4v9*v14
	m1[13] = -v2v12*v9 + v1v12*v10 + v2v13*v8 - v0v13*v10 - v1v8*v14 + v0v9*v14
	m1[14] = v2v12*v5 - v1v12*v6 - v2v13*v4 + v0v13*v6 + v1v4*v14 - v0v5*v14
	m1[15] = -v2*v5v8 + v1*v6v8 + v2*v4v9 - v0*v6v9 - v1v4*v10 + v0*v5v10
	m1.MulWith(1.0 / det)
}

// ApproxEqual performs an element-wise approximate equality test between two matrices,
// as if FloatEqual had been used.
func (m1 *Mat4) ApproxEqual(m2 *Mat4) bool {
	return FloatEqual(m1[0], m2[0]) && FloatEqual(m1[1], m2[1]) && FloatEqual(m1[2], m2[2]) && FloatEqual(m1[3], m2[3]) &&
		FloatEqual(m1[4], m2[4]) && FloatEqual(m1[5], m2[5]) && FloatEqual(m1[6], m2[6]) && FloatEqual(m1[7], m2[7]) &&
		FloatEqual(m1[8], m2[8]) && FloatEqual(m1[9], m2[9]) && FloatEqual(m1[10], m2[10]) && FloatEqual(m1[11], m2[11]) &&
		FloatEqual(m1[12], m2[12]) && FloatEqual(m1[13], m2[13]) && FloatEqual(m1[14], m2[14]) && FloatEqual(m1[15], m2[15])
}

// ApproxEqualThreshold performs an element-wise approximate equality test between two matrices
// with a given epsilon threshold, as if FloatEqualThreshold had been used.
func (m1 *Mat4) ApproxEqualThreshold(m2 *Mat4, threshold float32) bool {
	return FloatEqualThreshold(m1[0], m2[0], threshold) && FloatEqualThreshold(m1[1], m2[1], threshold) && FloatEqualThreshold(m1[2], m2[2], threshold) && FloatEqualThreshold(m1[3], m2[3], threshold) &&
		FloatEqualThreshold(m1[4], m2[4], threshold) && FloatEqualThreshold(m1[5], m2[5], threshold) && FloatEqualThreshold(m1[6], m2[6], threshold) && FloatEqualThreshold(m1[7], m2[7], threshold) &&
		FloatEqualThreshold(m1[8], m2[8], threshold) && FloatEqualThreshold(m1[9], m2[9], threshold) && FloatEqualThreshold(m1[10], m2[10], threshold) && FloatEqualThreshold(m1[11], m2[11], threshold) &&
		FloatEqualThreshold(m1[12], m2[12], threshold) && FloatEqualThreshold(m1[13], m2[13], threshold) && FloatEqualThreshold(m1[14], m2[14], threshold) && FloatEqualThreshold(m1[15], m2[15], threshold)
}

// ApproxFuncEqual performs an element-wise approximate equality test between two matrices
// with a given equality functions, intended to be used with FloatEqualFunc; although and comparison
// function may be used in practice.
func (m1 *Mat4) ApproxFuncEqual(m2 *Mat4, eq func(float32, float32) bool) bool {
	return eq(m1[0], m2[0]) && eq(m1[1], m2[1]) && eq(m1[2], m2[2]) && eq(m1[3], m2[3]) &&
		eq(m1[4], m2[4]) && eq(m1[5], m2[5]) && eq(m1[6], m2[6]) && eq(m1[7], m2[7]) &&
		eq(m1[8], m2[8]) && eq(m1[9], m2[9]) && eq(m1[10], m2[10]) && eq(m1[11], m2[11]) &&
		eq(m1[12], m2[12]) && eq(m1[13], m2[13]) && eq(m1[14], m2[14]) && eq(m1[15], m2[15])
}

// At returns the matrix element at the given row and column.
// This is equivalent to mat[col * numRow + row] where numRow is constant
// (E.G. for a Mat3x2 it's equal to 3)
//
// This method is garbage-in garbage-out. For instance, on a Mat4 asking for
// At(5,0) will work just like At(1,1). Or it may panic if it's out of bounds.
func (m1 *Mat4) At(row, col int) float32 {
	return m1[col*4+row]
}

// Set sets the corresponding matrix element at the given row and column.
// This has a pointer receiver because it mutates the matrix.
//
// This method is garbage-in garbage-out. For instance, on a Mat4 asking for
// Set(5,0,val) will work just like Set(1,1,val). Or it may panic if it's out of bounds.
func (m1 *Mat4) Set(row, col int, value float32) {
	m1[col*4+row] = value
}

// Index returns the index of the given row and column, to be used with direct
// access. E.G. Index(0,0) = 0.
//
// This is a garbage-in garbage-out method. For instance, on a Mat4 asking for the index of
// (5,0) will work the same as asking for (1,1). Or it may give you a value that will cause
// a panic if you try to access the array with it if it's truly out of bounds.
func (Mat4) Index(row, col int) int {
	return col*4 + row
}

// Row returns a vector representing the corresponding row (starting at row 0).
// This package makes no distinction between row and column vectors, so it
// will be a normal VecM for a MxN matrix.
func (m1 *Mat4) Row(row int) Vec4 {
	return Vec4{m1[row+0], m1[row+4], m1[row+8], m1[row+12]}
}

// Rows decomposes a matrix into its corresponding row vectors.
// This is equivalent to calling mat.Row for each row.
func (m1 *Mat4) Rows() (row0, row1, row2, row3 Vec4) {
	return m1.Row(0), m1.Row(1), m1.Row(2), m1.Row(3)
}

// Col returns a vector representing the corresponding column (starting at col 0).
// This package makes no distinction between row and column vectors, so it
// will be a normal VecN for a MxN matrix.
func (m1 *Mat4) Col(col int) Vec4 {
	return Vec4{m1[col*4+0], m1[col*4+1], m1[col*4+2], m1[col*4+3]}
}

// Cols decomposes a matrix into its corresponding column vectors.
// This is equivalent to calling mat.Col for each column.
func (m1 *Mat4) Cols() (col0, col1, col2, col3 Vec4) {
	return m1.Col(0), m1.Col(1), m1.Col(2), m1.Col(3)
}

// Trace is a basic operation on a square matrix that simply
// sums up all elements on the main diagonal (meaning all elements such that row==col).
func (m1 *Mat4) Trace() float32 {
	return m1[0] + m1[5] + m1[10] + m1[15]
}

// Abs returns the element-wise absolute value of this matrix
func (m1 *Mat4) Abs() Mat4 {
	return Mat4{math.Abs(m1[0]), math.Abs(m1[1]), math.Abs(m1[2]), math.Abs(m1[3]), math.Abs(m1[4]), math.Abs(m1[5]), math.Abs(m1[6]), math.Abs(m1[7]), math.Abs(m1[8]), math.Abs(m1[9]), math.Abs(m1[10]), math.Abs(m1[11]), math.Abs(m1[12]), math.Abs(m1[13]), math.Abs(m1[14]), math.Abs(m1[15])}
}

// AbsOf is a memory friendly version of Abs.
func (m1 *Mat4) AbsOf(m2 *Mat4) {
	m1[0] = math.Abs(m2[0])
	m1[1] = math.Abs(m2[1])
	m1[2] = math.Abs(m2[2])
	m1[3] = math.Abs(m2[3])
	m1[4] = math.Abs(m2[4])
	m1[5] = math.Abs(m2[5])
	m1[6] = math.Abs(m2[6])
	m1[7] = math.Abs(m2[7])
	m1[8] = math.Abs(m2[8])
	m1[9] = math.Abs(m2[9])
	m1[10] = math.Abs(m2[10])
	m1[11] = math.Abs(m2[11])
	m1[12] = math.Abs(m2[12])
	m1[13] = math.Abs(m2[13])
	m1[14] = math.Abs(m2[14])
	m1[15] = math.Abs(m2[15])

}

// AbsSelf is a memory friendly version of Abs.
func (m1 *Mat4) AbsSelf() {
	m1[0] = math.Abs(m1[0])
	m1[1] = math.Abs(m1[1])
	m1[2] = math.Abs(m1[2])
	m1[3] = math.Abs(m1[3])
	m1[4] = math.Abs(m1[4])
	m1[5] = math.Abs(m1[5])
	m1[6] = math.Abs(m1[6])
	m1[7] = math.Abs(m1[7])
	m1[8] = math.Abs(m1[8])
	m1[9] = math.Abs(m1[9])
	m1[10] = math.Abs(m1[10])
	m1[11] = math.Abs(m1[11])
	m1[12] = math.Abs(m1[12])
	m1[13] = math.Abs(m1[13])
	m1[14] = math.Abs(m1[14])
	m1[15] = math.Abs(m1[15])
}

// Iden sets this matrix to the identity matrix
func (m1 *Mat4) Iden() {
	m1[0] = 1
	m1[1] = 0
	m1[2] = 0
	m1[3] = 0
	m1[4] = 0
	m1[5] = 1
	m1[6] = 0
	m1[7] = 0
	m1[8] = 0
	m1[9] = 0
	m1[10] = 1
	m1[11] = 0
	m1[12] = 0
	m1[13] = 0
	m1[14] = 0
	m1[15] = 1
}

// String pretty prints the matrix
func (m1 *Mat4) String() string {
	buf := new(bytes.Buffer)
	w := tabwriter.NewWriter(buf, 4, 4, 1, ' ', tabwriter.AlignRight)
	for i := 0; i < 4; i++ {
		row := m1.Row(i)
		for _, col := range []float32{row[0], row[1], row[2], row[3]} {
			fmt.Fprintf(w, "%f\t", col)
		}

		fmt.Fprintln(w, "")
	}
	w.Flush()

	return buf.String()
}

// Mat3x4 stuff

//Mat3x4 is a 3 row 4 column matrix.
type Mat3x4 [12]float32

// Ident3x4 returns the cheating matrix 3x4 with its diagonal as [1,1,1]
func Ident3x4() Mat3x4 {
	return Mat3x4{1, 0, 0,
		0, 1, 0,
		0, 0, 1,
		0, 0, 0}
}

// Mat4 returns a mat4 with the last row as [0 0 0 1].
func (m1 *Mat3x4) Mat4() Mat4 {
	return Mat4{m1[0], m1[1], m1[2], 0,
		m1[3], m1[4], m1[5], 0,
		m1[6], m1[7], m1[8], 0,
		m1[9], m1[10], m1[11], 1}
}

// Mat4In is a memory friendly version of Mat4.
func (m1 *Mat3x4) Mat4In(m2 *Mat4) {
	m2[0], m2[4], m2[8], m2[12] = m1[0], m1[3], m1[6], m1[9]
	m2[1], m2[5], m2[9], m2[13] = m1[1], m1[4], m1[7], m1[10]
	m2[2], m2[6], m2[10], m2[14] = m1[2], m1[5], m1[8], m1[11]
	m2[3], m2[7], m2[11], m2[15] = 0, 0, 0, 1

}

// SetCol sets a Column within the Matrix, so it mutates the calling matrix.
func (m1 *Mat3x4) SetCol(col int, v *Vec3) {
	m1[col*3+0], m1[col*3+1], m1[col*3+2] = v[0], v[1], v[2]
}

// SetRow sets a Row within the Matrix, so it mutates the calling matrix.
func (m1 *Mat3x4) SetRow(row int, v *Vec4) {
	m1[row+0], m1[row+3], m1[row+6], m1[row+9] = v[0], v[1], v[2], v[3]
}

// Mat3x4FromRows builds a new matrix from row vectors.
// The resulting matrix will still be in column major order, but this can be
// good for hand-building matrices.
func Mat3x4FromRows(row0, row1, row2 *Vec4) Mat3x4 {
	return Mat3x4{row0[0], row1[0], row2[0], row0[1], row1[1], row2[1], row0[2], row1[2], row2[2], row0[3], row1[3], row2[3]}
}

// Mat3x4FromCols builds a new matrix from column vectors.
func Mat3x4FromCols(col0, col1, col2, col3 *Vec3) Mat3x4 {
	return Mat3x4{col0[0], col0[1], col0[2], col1[0], col1[1], col1[2], col2[0], col2[1], col2[2], col3[0], col3[1], col3[2]}
}

// Add performs an element-wise addition of two matrices, this is
// equivalent to iterating over every element of m1 and adding the corresponding value of m2.
func (m1 *Mat3x4) Add(m2 *Mat3x4) Mat3x4 {
	return Mat3x4{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5], m1[6] + m2[6], m1[7] + m2[7], m1[8] + m2[8], m1[9] + m2[9], m1[10] + m2[10], m1[11] + m2[11]}
}

// Sub performs an element-wise subtraction of two matrices, this is
// equivalent to iterating over every element of m1 and subtracting the corresponding value of m2.
func (m1 *Mat3x4) Sub(m2 *Mat3x4) Mat3x4 {
	return Mat3x4{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7], m1[8] - m2[8], m1[9] - m2[9], m1[10] - m2[10], m1[11] - m2[11]}
}

// Mul performs a scalar multiplcation of the matrix. This is equivalent to iterating
// over every element of the matrix and multiply it by c.
func (m1 *Mat3x4) Mul(c float32) Mat3x4 {
	return Mat3x4{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c, m1[8] * c, m1[9] * c, m1[10] * c, m1[11] * c}
}

// Mul4x1 performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 *Mat3x4) Mul4x1(v1 *Vec4) Vec3 {
	return Vec3{
		m1[0]*v1[0] + m1[3]*v1[1] + m1[6]*v1[2] + m1[9]*v1[3],
		m1[1]*v1[0] + m1[4]*v1[1] + m1[7]*v1[2] + m1[10]*v1[3],
		m1[2]*v1[0] + m1[5]*v1[1] + m1[8]*v1[2] + m1[11]*v1[3],
	}
}

// Mul3x1 is a cheat function that assumes the last row is [0 0 0 1] and the vectors last coordinate is 1. It's used in the physics engine to transform coordinates.
func (m1 *Mat3x4) Mul3x1(v1 *Vec3) Vec3 {
	return Vec3{
		m1[0]*v1[0] + m1[3]*v1[1] + m1[6]*v1[2] + m1[9],
		m1[1]*v1[0] + m1[4]*v1[1] + m1[7]*v1[2] + m1[10],
		m1[2]*v1[0] + m1[5]*v1[1] + m1[8]*v1[2] + m1[11],
	}
}

// Mul3x1In is a memory friendly version of Mul3x1, its declaration differs from the rest of the memory utility function to keep the api clean.
func (m1 *Mat3x4) Mul3x1In(v1, dst *Vec3) {
	dst[0] = m1[0]*v1[0] + m1[3]*v1[1] + m1[6]*v1[2] + m1[9]
	dst[1] = m1[1]*v1[0] + m1[4]*v1[1] + m1[7]*v1[2] + m1[10]
	dst[2] = m1[2]*v1[0] + m1[5]*v1[1] + m1[8]*v1[2] + m1[11]
}

// Mul3x4 is a cheat function that assumes the last row of both matrices
// is [0 0 0 1] and performs a 4x4 matrix multiplication.
func (m1 *Mat3x4) Mul3x4(m2 *Mat3x4) Mat3x4 {
	return Mat3x4{
		m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2],
		m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2],
		m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2],

		m1[0]*m2[3] + m1[3]*m2[4] + m1[6]*m2[5],
		m1[1]*m2[3] + m1[4]*m2[4] + m1[7]*m2[5],
		m1[2]*m2[3] + m1[5]*m2[4] + m1[8]*m2[5],

		m1[0]*m2[6] + m1[3]*m2[7] + m1[6]*m2[8],
		m1[1]*m2[6] + m1[4]*m2[7] + m1[7]*m2[8],
		m1[2]*m2[6] + m1[5]*m2[7] + m1[8]*m2[8],

		m1[0]*m2[9] + m1[3]*m2[10] + m1[6]*m2[11] + m1[9],
		m1[1]*m2[9] + m1[4]*m2[10] + m1[7]*m2[11] + m1[10],
		m1[2]*m2[9] + m1[5]*m2[10] + m1[8]*m2[11] + m1[11],
	}
}

// Mul3x4Of is a memory friendly version of Mul3x4.
func (m1 *Mat3x4) Mul3x4Of(m2, m3 *Mat3x4) {
	m1[0] = m2[0]*m3[0] + m2[3]*m3[1] + m2[6]*m3[2]
	m1[1] = m2[1]*m3[0] + m2[4]*m3[1] + m2[7]*m3[2]
	m1[2] = m2[2]*m3[0] + m2[5]*m3[1] + m2[8]*m3[2]

	m1[3] = m2[0]*m3[3] + m2[3]*m3[4] + m2[6]*m3[5]
	m1[4] = m2[1]*m3[3] + m2[4]*m3[4] + m2[7]*m3[5]
	m1[5] = m2[2]*m3[3] + m2[5]*m3[4] + m2[8]*m3[5]

	m1[6] = m2[0]*m3[6] + m2[3]*m3[7] + m2[6]*m3[8]
	m1[7] = m2[1]*m3[6] + m2[4]*m3[7] + m2[7]*m3[8]
	m1[8] = m2[2]*m3[6] + m2[5]*m3[7] + m2[8]*m3[8]

	m1[9] = m2[0]*m3[9] + m2[3]*m3[10] + m2[6]*m3[11] + m2[9]
	m1[10] = m2[1]*m3[9] + m2[4]*m3[10] + m2[7]*m3[11] + m2[10]
	m1[11] = m2[2]*m3[9] + m2[5]*m3[10] + m2[8]*m3[11] + m2[11]
}

// Mul3x4With is a memory friendly version of Mul3x4.
func (m1 *Mat3x4) Mul3x4With(m2 *Mat3x4) {
	v0 := m1[0]
	v1 := m1[1]
	v2 := m1[2]
	v3 := m1[3]
	v4 := m1[4]
	v5 := m1[5]
	v6 := m1[6]
	v7 := m1[7]
	v8 := m1[8]
	v9 := m1[9]
	v10 := m1[10]
	v11 := m1[11]
	m1[0] = v0*m2[0] + v3*m2[1] + v6*m2[2]
	m1[1] = v1*m2[0] + v4*m2[1] + v7*m2[2]
	m1[2] = v2*m2[0] + v5*m2[1] + v8*m2[2]

	m1[3] = v0*m2[3] + v3*m2[4] + v6*m2[5]
	m1[4] = v1*m2[3] + v4*m2[4] + v7*m2[5]
	m1[5] = v2*m2[3] + v5*m2[4] + v8*m2[5]

	m1[6] = v0*m2[6] + v3*m2[7] + v6*m2[8]
	m1[7] = v1*m2[6] + v4*m2[7] + v7*m2[8]
	m1[8] = v2*m2[6] + v5*m2[7] + v8*m2[8]

	m1[9] = v0*m2[9] + v3*m2[10] + v6*m2[11] + v9
	m1[10] = v1*m2[9] + v4*m2[10] + v7*m2[11] + v10
	m1[11] = v2*m2[9] + v5*m2[10] + v8*m2[11] + v11
}

// Mul4 performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 *Mat3x4) Mul4(m2 *Mat4) Mat3x4 {
	return Mat3x4{
		m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2] + m1[9]*m2[3],
		m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2] + m1[10]*m2[3],
		m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2] + m1[11]*m2[3],
		m1[0]*m2[4] + m1[3]*m2[5] + m1[6]*m2[6] + m1[9]*m2[7],
		m1[1]*m2[4] + m1[4]*m2[5] + m1[7]*m2[6] + m1[10]*m2[7],
		m1[2]*m2[4] + m1[5]*m2[5] + m1[8]*m2[6] + m1[11]*m2[7],
		m1[0]*m2[8] + m1[3]*m2[9] + m1[6]*m2[10] + m1[9]*m2[11],
		m1[1]*m2[8] + m1[4]*m2[9] + m1[7]*m2[10] + m1[10]*m2[11],
		m1[2]*m2[8] + m1[5]*m2[9] + m1[8]*m2[10] + m1[11]*m2[11],
		m1[0]*m2[12] + m1[3]*m2[13] + m1[6]*m2[14] + m1[9]*m2[15],
		m1[1]*m2[12] + m1[4]*m2[13] + m1[7]*m2[14] + m1[10]*m2[15],
		m1[2]*m2[12] + m1[5]*m2[13] + m1[8]*m2[14] + m1[11]*m2[15],
	}
}

// Det on 3x4 matrix is a cheat, it assumes the last row is [0 0 0 1].
//    [a d g j]
//    [b e h k]
//    [c f i l]
//    [0 0 0 1]
// aei - afh - bdi + bfg + cdh - ceg
func (m1 *Mat3x4) Det() float32 {
	return m1[0]*m1[4]*m1[8] - m1[0]*m1[5]*m1[7] - m1[1]*m1[3]*m1[8] + m1[1]*m1[5]*m1[6] + m1[2]*m1[3]*m1[7] - m1[2]*m1[4]*m1[6]
}

// Inverse is a cheat function that returns the inverse of this matrix as if it was a 4x4 matrix.
func (m1 *Mat3x4) Inverse() Mat3x4 {

	det := m1.Det()
	if FloatEqual(det, float32(0.0)) {
		return Mat3x4{}
	}

	retMat := Mat3x4{
		m1[4]*m1[8] - m1[5]*m1[7],
		m1[2]*m1[7] - m1[1]*m1[8],
		m1[1]*m1[5] - m1[2]*m1[4],
		m1[5]*m1[6] - m1[3]*m1[8],
		m1[0]*m1[8] - m1[2]*m1[6],
		m1[2]*m1[3] - m1[0]*m1[5],
		m1[3]*m1[7] - m1[4]*m1[6],
		m1[1]*m1[6] - m1[0]*m1[7],
		m1[0]*m1[4] - m1[1]*m1[3],
		m1[5]*m1[7]*m1[9] - m1[4]*m1[8]*m1[9] - m1[5]*m1[6]*m1[10] + m1[3]*m1[8]*m1[10] + m1[4]*m1[6]*m1[11] - m1[3]*m1[7]*m1[11],
		-m1[2]*m1[9]*m1[7] + m1[1]*m1[9]*m1[8] + m1[2]*m1[10]*m1[6] - m1[0]*m1[10]*m1[8] - m1[1]*m1[6]*m1[11] + m1[0]*m1[7]*m1[11],
		m1[2]*m1[9]*m1[4] - m1[1]*m1[9]*m1[5] - m1[2]*m1[10]*m1[3] + m1[0]*m1[10]*m1[5] + m1[1]*m1[3]*m1[11] - m1[0]*m1[4]*m1[11],
	}
	return retMat.Mul(1.0 / det)
}

// ApproxEqual performs an element-wise approximate equality test between two matrices,
// as if FloatEqual had been used.
func (m1 *Mat3x4) ApproxEqual(m2 *Mat3x4) bool {
	return FloatEqual(m1[0], m2[0]) && FloatEqual(m1[1], m2[1]) && FloatEqual(m1[2], m2[2]) && FloatEqual(m1[3], m2[3]) &&
		FloatEqual(m1[4], m2[4]) && FloatEqual(m1[5], m2[5]) && FloatEqual(m1[6], m2[6]) && FloatEqual(m1[7], m2[7]) &&
		FloatEqual(m1[8], m2[8]) && FloatEqual(m1[9], m2[9]) && FloatEqual(m1[10], m2[10]) && FloatEqual(m1[11], m2[11])
}

// ApproxEqualThreshold performs an element-wise approximate equality test between two matrices
// with a given epsilon threshold, as if FloatEqualThreshold had been used.
func (m1 *Mat3x4) ApproxEqualThreshold(m2 *Mat3x4, threshold float32) bool {
	return FloatEqualThreshold(m1[0], m2[0], threshold) && FloatEqualThreshold(m1[1], m2[1], threshold) && FloatEqualThreshold(m1[2], m2[2], threshold) && FloatEqualThreshold(m1[3], m2[3], threshold) &&
		FloatEqualThreshold(m1[4], m2[4], threshold) && FloatEqualThreshold(m1[5], m2[5], threshold) && FloatEqualThreshold(m1[6], m2[6], threshold) && FloatEqualThreshold(m1[7], m2[7], threshold) &&
		FloatEqualThreshold(m1[8], m2[8], threshold) && FloatEqualThreshold(m1[9], m2[9], threshold) && FloatEqualThreshold(m1[10], m2[10], threshold) && FloatEqualThreshold(m1[11], m2[11], threshold)
}

// ApproxFuncEqual performs an element-wise approximate equality test between two matrices
// with a given equality functions, intended to be used with FloatEqualFunc; although and comparison
// function may be used in practice.
func (m1 *Mat3x4) ApproxFuncEqual(m2 *Mat3x4, eq func(float32, float32) bool) bool {
	return eq(m1[0], m2[0]) && eq(m1[1], m2[1]) && eq(m1[2], m2[2]) && eq(m1[3], m2[3]) &&
		eq(m1[4], m2[4]) && eq(m1[5], m2[5]) && eq(m1[6], m2[6]) && eq(m1[7], m2[7]) &&
		eq(m1[8], m2[8]) && eq(m1[9], m2[9]) && eq(m1[10], m2[10]) && eq(m1[11], m2[11])
}

// At returns the matrix element at the given row and column.
// This is equivalent to mat[col * numRow + row] where numRow is constant
// (E.G. for a Mat3x2 it's equal to 3)
//
// This method is garbage-in garbage-out. For instance, on a Mat4 asking for
// At(5,0) will work just like At(1,1). Or it may panic if it's out of bounds.
func (m1 *Mat3x4) At(row, col int) float32 {
	return m1[col*3+row]
}

// Set sets the corresponding matrix element at the given row and column.
// This has a pointer receiver because it mutates the matrix.
//
// This method is garbage-in garbage-out. For instance, on a Mat4 asking for
// Set(5,0,val) will work just like Set(1,1,val). Or it may panic if it's out of bounds.
func (m1 *Mat3x4) Set(row, col int, value float32) {
	m1[col*3+row] = value
}

// Index returns the index of the given row and column, to be used with direct
// access. E.G. Index(0,0) = 0.
//
// This is a garbage-in garbage-out method. For instance, on a Mat4 asking for the index of
// (5,0) will work the same as asking for (1,1). Or it may give you a value that will cause
// a panic if you try to access the array with it if it's truly out of bounds.
func (Mat3x4) Index(row, col int) int {
	return col*3 + row
}

// Row returns a vector representing the corresponding row (starting at row 0).
// This package makes no distinction between row and column vectors, so it
// will be a normal VecM for a MxN matrix.
func (m1 *Mat3x4) Row(row int) Vec4 {
	return Vec4{m1[row+0], m1[row+3], m1[row+6], m1[row+9]}
}

// Rows decomposes a matrix into its corresponding row vectors.
// This is equivalent to calling mat.Row for each row.
func (m1 *Mat3x4) Rows() (row0, row1, row2 Vec4) {
	return m1.Row(0), m1.Row(1), m1.Row(2)
}

// Col returns a vector representing the corresponding column (starting at col 0).
// This package makes no distinction between row and column vectors, so it
// will be a normal VecN for a MxN matrix.
func (m1 *Mat3x4) Col(col int) Vec3 {
	return Vec3{m1[col*3+0], m1[col*3+1], m1[col*3+2]}
}

// Cols decomposes a matrix into its corresponding column vectors.
// This is equivalent to calling mat.Col for each column.
func (m1 *Mat3x4) Cols() (col0, col1, col2, col3 Vec3) {
	return m1.Col(0), m1.Col(1), m1.Col(2), m1.Col(3)
}

// Abs returns the element-wise absolute value of this matrix
func (m1 *Mat3x4) Abs() Mat3x4 {
	return Mat3x4{math.Abs(m1[0]), math.Abs(m1[1]), math.Abs(m1[2]), math.Abs(m1[3]), math.Abs(m1[4]), math.Abs(m1[5]), math.Abs(m1[6]), math.Abs(m1[7]), math.Abs(m1[8]), math.Abs(m1[9]), math.Abs(m1[10]), math.Abs(m1[11])}
}

// SetOrientationAndPos sets this matrix to represent this quaternion's orientation and this vector's position.
func (m1 *Mat3x4) SetOrientationAndPos(q1 *Quat, v1 *Vec3) {
	w, x, y, z := q1.W, q1.V[0], q1.V[1], q1.V[2]
	m1[0] = 1 - 2*y*y - 2*z*z
	m1[1] = 2*x*y + 2*w*z
	m1[2] = 2*x*z - 2*w*y
	m1[3] = 2*x*y - 2*w*z
	m1[4] = 1 - 2*x*x - 2*z*z
	m1[5] = 2*y*z + 2*w*x
	m1[6] = 2*x*z + 2*w*y
	m1[7] = 2*y*z - 2*w*x
	m1[8] = 1 - 2*x*x - 2*y*y
	m1[9] = v1[0]
	m1[10] = v1[1]
	m1[11] = v1[2]
}

// Transform is really just calling Mul3x1 but for the physics engine we'll redeclare it that way.
func (m1 *Mat3x4) Transform(v1 *Vec3) Vec3 {
	return m1.Mul3x1(v1)
}

// TransformIn is really just calling Mul3x1In but for the physics engine we'll redeclare it that way.
func (m1 *Mat3x4) TransformIn(v1, dst *Vec3) {
	m1.Mul3x1In(v1, dst)
}

// TransformInverse will transform v1 by using shortcut. Like assuming that the 4th
// column is a translation and that the inner 3x3 matrix is a rotation matrix (meaning
// that we can use the transpose.
func (m1 *Mat3x4) TransformInverse(v1 *Vec3) Vec3 {
	x := v1[0] - m1[9]
	y := v1[1] - m1[10]
	z := v1[2] - m1[11]
	return Vec3{
		x*m1[0] + y*m1[1] + z*m1[2],
		x*m1[3] + y*m1[4] + z*m1[5],
		x*m1[6] + y*m1[7] + z*m1[8],
	}
}

// TransformInverseIn is a memory friendly version of TransformInverse.
func (m1 *Mat3x4) TransformInverseIn(v1, dst *Vec3) {
	x := v1[0] - m1[9]
	y := v1[1] - m1[10]
	z := v1[2] - m1[11]

	dst[0] = x*m1[0] + y*m1[1] + z*m1[2]
	dst[1] = x*m1[3] + y*m1[4] + z*m1[5]
	dst[2] = x*m1[6] + y*m1[7] + z*m1[8]
}

// TransformDirection transforms the given direction by this inner rotation matrix.
func (m1 *Mat3x4) TransformDirection(v1 *Vec3) Vec3 {
	return Vec3{v1[0]*m1[0] + v1[1]*m1[3] + v1[2]*m1[6],
		v1[0]*m1[1] + v1[1]*m1[4] + v1[2]*m1[7],
		v1[0]*m1[2] + v1[1]*m1[5] + v1[2]*m1[8]}
}

// TransformDirectionIn is a memory friendly version of TransformDirection.
func (m1 *Mat3x4) TransformDirectionIn(v1, dst *Vec3) {
	dst[0] = v1[0]*m1[0] + v1[1]*m1[3] + v1[2]*m1[6]
	dst[1] = v1[0]*m1[1] + v1[1]*m1[4] + v1[2]*m1[7]
	dst[2] = v1[0]*m1[2] + v1[1]*m1[5] + v1[2]*m1[8]
}

// TransformInverseDirection uses the fact that the inner 3x3 matrix is a
// rotation matrix to use the transpose to do the inverse of TransformDirection.
func (m1 *Mat3x4) TransformInverseDirection(v1 *Vec3) Vec3 {
	return Vec3{v1[0]*m1[0] + v1[1]*m1[1] + v1[2]*m1[2],
		v1[0]*m1[3] + v1[1]*m1[4] + v1[2]*m1[5],
		v1[0]*m1[6] + v1[1]*m1[7] + v1[2]*m1[8]}
}

// TransformInverseDirectionIn is a memory friendly version of TransformInverseDirection.
func (m1 *Mat3x4) TransformInverseDirectionIn(v1, dst *Vec3) {
	dst[0] = v1[0]*m1[0] + v1[1]*m1[1] + v1[2]*m1[2]
	dst[1] = v1[0]*m1[3] + v1[1]*m1[4] + v1[2]*m1[5]
	dst[2] = v1[0]*m1[6] + v1[1]*m1[7] + v1[2]*m1[8]
}

// GetAxis return one of the axis of the matrix. i needs to be between 0 and 3
// or else this will panic
func (m1 *Mat3x4) GetAxis(i int) Vec3 {
	return Vec3{
		m1[i*3+0],
		m1[i*3+1],
		m1[i*3+2],
	}
}

// String pretty prints the matrix
func (m1 *Mat3x4) String() string {
	buf := new(bytes.Buffer)
	w := tabwriter.NewWriter(buf, 4, 4, 1, ' ', tabwriter.AlignRight)
	for i := 0; i < 3; i++ {
		for _, col := range m1.Row(i) {
			fmt.Fprintf(w, "%f\t", col)
		}

		fmt.Fprintln(w, "")
	}
	w.Flush()

	return buf.String()
}
