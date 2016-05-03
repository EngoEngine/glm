package glm

import (
	"unsafe"
)

// Transform is a utility type used to aggregate transformations.
type Transform Mat4

// NewTransform returns a new, initialized transform.
func NewTransform() Transform {
	return Transform(Ident4())
}

// Init sets this transform to the identity transform.
func (t *Transform) Init() {
	*t = Transform(Ident4())
}

// Translate3f adds a translation to this transform of {x, y, z}.
func (t *Transform) Translate3f(x, y, z float32) {
	tran := Translate3D(x, y, z)
	((*Mat4)(t)).Mul4With(&tran)
}

// TranslateVec3 adds a translation to this transform of v.
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

// Scale scales this transform by f. Non-uniform scales are prohibited to
// guarantee the existence of the normal matrix.
func (t *Transform) Scale(f float32) {
	m := Scale3D(f, f, f)
	((*Mat4)(t)).Mul4With(&m)
}

// SetScale rotates this transform by q.
func (t *Transform) SetScale(f float32) {
	*t = Transform(Scale3D(f, f, f))
}

// Normal returns the normal matrix of this transform, this is used in most
// light shading algorithms.
func (t *Transform) Normal() Mat3 {
	m := ((*Mat4)(t)).Mat3()
	return m.Inverse()
}

// GLPointer returns the pointer to the first element of the underlying 4x4
// matrix. This is can be passed directly to any OpenGL function.
func (t *Transform) GLPointer() unsafe.Pointer {
	return unsafe.Pointer(t)
}
