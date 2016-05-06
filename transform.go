package glm

import (
	"unsafe"
)

// Transform is a utility type used to aggregate transformations. Transform
// concatenation, like matrix multiplication, is not commutative.
type Transform Mat4

// NewTransform returns a new, initialized transform.
func NewTransform() Transform {
	return Transform(Ident4())
}

// Iden sets this transform to the identity transform. You NEED to call this
// EXCEPT IF:
// - You get your transform from NewTransform
// - You're gonna call Set* BEFORE Translate* or Rotate*
func (t *Transform) Iden() {
	*t = Transform(Ident4())
}

// Translate3f concatenates a translation to this transform of {x, y, z}.
func (t *Transform) Translate3f(x, y, z float32) {
	tran := Translate3D(x, y, z)
	((*Mat4)(t)).Mul4With(&tran)
}

// TranslateVec3 concatenates a translation to this transform of v.
func (t *Transform) TranslateVec3(v *Vec3) {
	tran := Translate3D(v[0], v[1], v[2])
	((*Mat4)(t)).Mul4With(&tran)
}

// SetTranslate3f sets the transform to a translate transform of {x, y, z}.
func (t *Transform) SetTranslate3f(x, y, z float32) {
	*t = Transform(Translate3D(x, y, z))
}

// SetTranslateVec3 sets the transform to a translate transform of v.
func (t *Transform) SetTranslateVec3(v *Vec3) {
	*t = Transform(Translate3D(v[0], v[1], v[2]))
}

// RotateQuat rotates this transform by q.
func (t *Transform) RotateQuat(q *Quat) {
	m := q.Mat4()
	((*Mat4)(t)).Mul4With(&m)
}

// SetRotateQuat rotates this transform by q.
func (t *Transform) SetRotateQuat(q *Quat) {
	*t = Transform(q.Mat4())
}

// Concatenate Transform t2 into t.
func (t *Transform) Concatenate(t2 *Transform) {
	((*Mat4)(t)).Mul4With((*Mat4)(t2))
}

// LocalToWorld transform a given point and returns the world point that this
// transform generates.
func (t *Transform) LocalToWorld(v *Vec3) Vec3 {
	v4 := v.Vec4(1)
	v4 = (*Mat4)(t).Mul4x1(&v4)
	return v4.Vec3()
}

// WorldToLocal transform a given point and returns the local point that this
// transform generates.
func (t *Transform) WorldToLocal(v *Vec3) Vec3 {
	// BUG(hydroflame): the current implementation currently inverse the matrix
	// on every call ... that may not be the most efficient.
	inv := (*Mat4)(t).Inverse()
	v4 := v.Vec4(1)
	v4 = inv.Mul4x1(&v4)
	return v4.Vec3()
}

// Normal returns the normal matrix of this transform, this is used in most
// light shading algorithms.
func (t *Transform) Normal() Mat3 {
	// Since we prevent scaling we are guaranteed that the upper 3x3 matrix is
	// orthogonal and (TODO(hydroflame): find the word for when a matrix has all
	// unit vectors), we can just throw it back and it's the correct transform
	// matrix.
	return ((*Mat4)(t)).Mat3()
}

// Mat4 simply returns the Mat4 associated with this Transform. This effectively
// makes a copy.
func (t *Transform) Mat4() Mat4 {
	return *((*Mat4)(t))
}

// Pointer returns the pointer to the first element of the underlying 4x4
// matrix. This is can be passed directly to OpenGL function.
func (t *Transform) Pointer() unsafe.Pointer {
	return unsafe.Pointer(t)
}

// String return a string that represents this transform (a mat4).
func (t *Transform) String() string {
	return (*Mat4)(t).String()
}

// Transform2D is a utility type used to aggregate transformations. Transform
// concatenation, like matrix multiplication, is not commutative.
type Transform2D Mat3

// NewTransform2D returns a new, initialized transform.
func NewTransform2D() Transform2D {
	return Transform2D(Ident3())
}

// Iden sets this transform to the identity transform. You NEED to call this
// EXCEPT IF:
// - You get your transform from NewTransform
// - You're gonna call Set* BEFORE Translate* or Rotate*
func (t *Transform2D) Iden() {
	*t = Transform2D(Ident3())
}

// Translate2f concatenates a translation to this transform of {x, y, z}.
func (t *Transform2D) Translate2f(x, y float32) {
	tran := Translate2D(x, y)
	((*Mat3)(t)).Mul3With(&tran)
}

// TranslateVec2 concatenates a translation to this transform of v.
func (t *Transform2D) TranslateVec2(v *Vec2) {
	tran := Translate2D(v[0], v[1])
	((*Mat3)(t)).Mul3With(&tran)
}

// SetTranslate2f sets the transform to a translate transform of {x, y, z}.
func (t *Transform2D) SetTranslate2f(x, y float32) {
	*t = Transform2D(Translate2D(x, y))
}

// SetTranslateVec2 sets the transform to a translate transform of v.
func (t *Transform2D) SetTranslateVec2(v *Vec2) {
	*t = Transform2D(Translate2D(v[0], v[1]))
}

// Rotate concatenates a rotation of angle (radian).
func (t *Transform2D) Rotate(angle float32) {
	rot := HomogRotate2D(angle)
	((*Mat3)(t)).Mul3With(&rot)
}

// SetRotate sets the transform to a rotate transform of angle (radian).
func (t *Transform2D) SetRotate(angle float32) {
	*t = Transform2D(HomogRotate2D(angle))
}

// LocalToWorld transform a given point and returns the world point that this
// transform generates.
func (t *Transform2D) LocalToWorld(v *Vec2) Vec2 {
	v3 := v.Vec3(1)
	v3 = (*Mat3)(t).Mul3x1(&v3)
	return v3.Vec2()
}

// WorldToLocal transform a given point and returns the local point that this
// transform generates.
func (t *Transform2D) WorldToLocal(v *Vec2) Vec2 {
	// BUG(hydroflame): the current implementation currently inverse the matrix
	// on every call ... that may not be the most efficient.
	inv := (*Mat3)(t).Inverse()
	v3 := v.Vec3(1)
	v3 = inv.Mul3x1(&v3)
	return v3.Vec2()
}

// Concatenate Transform t2 into t.
func (t *Transform2D) Concatenate(t2 *Transform2D) {
	((*Mat3)(t)).Mul3With((*Mat3)(t2))
}

// Mat3 simply returns the Mat3 associated with this Transform. This effectively
// makes a copy.
func (t *Transform2D) Mat3() Mat3 {
	return *((*Mat3)(t))
}

// Pointer returns the pointer to the first element of the underlying 4x4
// matrix. This is can be passed directly to OpenGL function.
func (t *Transform2D) Pointer() unsafe.Pointer {
	return unsafe.Pointer(t)
}

// String return a string that represents this transform (a mat4).
func (t *Transform2D) String() string {
	return (*Mat3)(t).String()
}
