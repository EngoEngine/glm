// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package glm

import (
	"luxengine.net/math"
	math64 "math"
)

// RotationOrder  is the order in which
// rotations will be transformed for the purposes of AnglesToQuat
type RotationOrder int

const (
	// XYX , shut up golint
	XYX RotationOrder = iota
	// XYZ , shut up golint
	XYZ
	// XZX , shut up golint
	XZX
	// XZY , shut up golint
	XZY
	// YXY , shut up golint
	YXY
	// YXZ , shut up golint
	YXZ
	// YZY , shut up golint
	YZY
	// YZX , shut up golint
	YZX
	// ZYZ , shut up golint
	ZYZ
	// ZYX , shut up golint
	ZYX
	// ZXZ , shut up golint
	ZXZ
	// ZXY , shut up golint
	ZXY
)

// Quat is a Quaternion. A Quaternion is an extension of the imaginary numbers; there's all sorts of
// interesting theory behind it. In 3D graphics we mostly use it as a cheap way of
// representing rotation since quaternions are cheaper to multiply by, and easier to
// interpolate than matrices.
//
// A Quaternion has two parts: W, the so-called scalar component,
// and "V", the vector component. The vector component is considered to
// be the part in 3D space, while W (loosely interpreted) is its 4D coordinate.
type Quat struct {
	W float32
	V Vec3
}

// QuatIdent return the identity quaternion.
// The quaternion identity: W=1; V=(0,0,0).
//
// As with all identities, multiplying any quaternion by this will yield the same
// quaternion you started with.
func QuatIdent() Quat {
	return Quat{1., Vec3{0, 0, 0}}
}

// QuatRotate creates an angle from an axis and an angle relative to that axis.
//
// This is cheaper than HomogRotate3D.
func QuatRotate(angle float32, axis *Vec3) Quat {
	s, c := math.Sin(angle/2), math.Cos(angle/2)
	return Quat{c, axis.Mul(s)}
}

// X is a convenient alias for q.V[0]
func (q1 Quat) X() float32 {
	return q1.V[0]
}

// Y is a convenient alias for q.V[1]
func (q1 Quat) Y() float32 {
	return q1.V[1]
}

// Z is a convenient alias for q.V[2]
func (q1 Quat) Z() float32 {
	return q1.V[2]
}

// Add adds two quaternions. It's no more complicated than
// adding their W and V components.
func (q1 *Quat) Add(q2 *Quat) Quat {
	return Quat{q1.W + q2.W, q1.V.Add(&q2.V)}
}

// AddOf is a memory friendly version of add. q1 = q2 + q3
func (q1 *Quat) AddOf(q2, q3 *Quat) {
	q1.W = q2.W + q3.W
	q1.V.AddOf(&q2.V, &q3.V)
}

// AddWith is a memory friendly version of Add.
// In quaternion cases you COULD use AddOf with q1 twice: q1.AddOf(&q1,&q2).
// This is here just for API consistency.
func (q1 *Quat) AddWith(q2 *Quat) {
	q1.W += q2.W
	q1.V.AddWith(&q2.V)
}

// Sub subtracts two quaternions. It's no more complicated than
// subtracting their W and V components.
func (q1 *Quat) Sub(q2 *Quat) Quat {
	return Quat{q1.W - q2.W, q1.V.Sub(&q2.V)}
}

// SubOf is a memory friendly version of add. q1 = q2 + q3
func (q1 *Quat) SubOf(q2, q3 *Quat) {
	q1.W = q2.W + q3.W
	q1.V.SubOf(&q2.V, &q3.V)
}

// SubWith is a memory friendly version of Sub.
// In quaternion cases you COULD use SubOf with q1 twice: q1.SubOf(&q1,&q2).
// This is here just for API consistency.
func (q1 *Quat) SubWith(q2 *Quat) {
	q1.W += q2.W
	q1.V.SubWith(&q2.V)
}

// Mul multiplies two quaternions. This can be seen as a rotation. Note that
// Multiplication is NOT commutative, meaning q1.Mul(q2) does not necessarily
// equal q2.Mul(q1).
func (q1 *Quat) Mul(q2 *Quat) Quat {
	var c Vec3
	c.CrossOf(&q1.V, &q2.V)
	var m1 Vec3
	m1.MulOf(q1.W, &q2.V)
	var m2 Vec3
	m2.MulOf(q2.W, &q1.V)
	var a1 Vec3
	a1.AddOf(&c, &m1)
	//q1.V.Cross(&q2.V).Add(q2.V.Mul(q1.W)).Add(q1.V.Mul(q2.W))
	return Quat{q1.W*q2.W - q1.V.Dot(&q2.V), a1.Add(&m2)}
}

// MulOf is a memory friendly version of Mul.
func (q1 *Quat) MulOf(q2, q3 *Quat) {
	q1.W = q2.W*q3.W - q2.V.Dot(&q3.V)

	q1.V.CrossOf(&q2.V, &q3.V)
	q1.V.AddScaledVec(q2.W, &q3.V)
	q1.V.AddScaledVec(q3.W, &q2.V)
}

// MulWith is a memory friendly version of Mul.
// Use this when you want q1 both as dest and arg.
func (q1 *Quat) MulWith(q2 *Quat) {
	w := q1.W
	v := q1.V
	q1.W = w*q2.W - v.Dot(&q2.V)

	q1.V.CrossOf(&v, &q2.V)
	q1.V.AddScaledVec(w, &q2.V)
	q1.V.AddScaledVec(q2.W, &v)
}

// Scale scales every element of the quaternion by some constant factor.
func (q1 *Quat) Scale(c float32) Quat {
	return Quat{q1.W * c, Vec3{q1.V[0] * c, q1.V[1] * c, q1.V[2] * c}}
}

// ScaleOf scales every element of the quaternion by some constant factor.
func (q1 *Quat) ScaleOf(c float32, q2 *Quat) {
	q1.W = c * q2.W
	q1.V.MulOf(c, &q2.V)
}

// ScaleWith scales every element of the quaternion by some constant factor.
func (q1 *Quat) ScaleWith(c float32) {
	q1.W *= c
	q1.V.MulWith(c)
}

// Conjugated returns the conjugate of a quaternion. Equivalent to
// Quat{q1.W, q1.V.Mul(-1)}
func (q1 *Quat) Conjugated() Quat {
	return Quat{q1.W, q1.V.Mul(-1)}
}

// ConjugateOf is a memory friendly version of Conjugated. q1 = conjugate(q2)
func (q1 *Quat) ConjugateOf(q2 *Quat) {
	q1.W = q2.W
	q1.V.MulOf(-1, &q2.V)
}

// Conjugate is a memory friendly version of Conjugated. q1 = conjugate(q1)
func (q1 *Quat) Conjugate() {
	q1.V.MulWith(-1)
}

// Len returns the Length of the quaternion, also known as its Norm. This is the same thing as
// the Len of a Vec4
func (q1 *Quat) Len() float32 {
	return math.Sqrt(q1.W*q1.W + q1.V[0]*q1.V[0] + q1.V[1]*q1.V[1] + q1.V[2]*q1.V[2])
}

// Norm is an alias for Len() since both are very common terms.
func (q1 *Quat) Norm() float32 {
	return q1.Len()
}

// Normalized Normalizes the quaternion, returning its versor (unit quaternion).
//
// This is the same as normalizing it as a Vec4.
func (q1 *Quat) Normalized() Quat {
	length := q1.Len()

	if FloatEqual(1, length) {
		return *q1
	}
	if length == 0 {
		return QuatIdent()
	}
	if length == InfPos {
		length = MaxValue
	}

	il := 1.0 / length

	return Quat{q1.W * il, q1.V.Mul(il)}
}

// SetNormalizedOf Normalizes the quaternion, returning its versor (unit quaternion).
//
// This is the same as normalizing it as a Vec4.
func (q1 *Quat) SetNormalizedOf(q2 *Quat) {
	length := q2.Len()

	if FloatEqual(1, length) {
		q1.W = q2.W
		q1.V = q2.V
		return
	}
	if length == 0 {
		q1.W = 1
		q1.V = Vec3{}
		return
	}
	if length == InfPos {
		length = MaxValue
	}

	il := 1.0 / length

	q1.W = q2.W * il
	q1.V.MulOf(il, &q2.V)
}

// Normalize Normalizes the quaternion in place.
//
// This is the same as normalizing it as a Vec4.
func (q1 *Quat) Normalize() {
	length := q1.Len()

	if FloatEqual(1, length) {
		return
	}
	if length == 0 {
		q1.W = 1
		q1.V = Vec3{}
		return
	}
	if length == InfPos {
		length = MaxValue
	}

	il := 1.0 / length

	q1.W *= il
	q1.V.MulWith(il)
}

// Inverse returns the inverse of a quaternion. The inverse is equivalent
// to the conjugate divided by the square of the length.
//
// This method computes the square norm by directly adding the sum
// of the squares of all terms instead of actually squaring q1.Len(),
// both for performance and percision.
func (q1 *Quat) Inverse() Quat {
	c := q1.Conjugated()
	return c.Scale(1 / q1.Dot(q1))
}

// InverseOf is a memory friendly version of Inverse.
func (q1 *Quat) InverseOf(q2 *Quat) {
	q1.ConjugateOf(q2)
	q1.ScaleWith(1.0 / q2.Dot(q2))
}

// Invert is a memory friendly version of Inverse.
func (q1 *Quat) Invert() {
	q1.Conjugate()
	q1.ScaleWith(1.0 / q1.Dot(q1))
}

// Rotate rotates a vector by the rotation this quaternion represents.
// This will result in a 3D vector. Strictly speaking, this is
// equivalent to q1.v.q* where the "."" is quaternion multiplication and v is interpreted
// as a quaternion with W 0 and V v. In code:
// q1.Mul(Quat{0,v}).Mul(q1.Conjugate()), and
// then retrieving the imaginary (vector) part.
//
// In practice, we hand-compute this in the general case and simplify
// to save a few operations.
func (q1 *Quat) Rotate(v *Vec3) Vec3 {
	var cross Vec3
	cross.CrossOf(&q1.V, v)
	// v + 2q_w * (q_v x v) + 2q_v x (q_v x v)
	//return v.Add(cross.Mul(2 * q1.W)).Add(q1.V.Mul(2).Cross(cross))
	/*

		                                        2 .Cross(cross)
		                2 * q1.W       q1.V.Mul( )
		      cross.Mul(        ) .Add(                        )
		v.Add(                   )
	*/
	var mul2 Vec3
	mul2.MulOf(2, &q1.V)
	mul2.CrossWith(&cross)

	var w2 Vec3
	w2.MulOf(2*q1.W, &cross)
	var out Vec3
	out.AddOf(v, &w2)
	out.AddWith(&mul2)
	return out
}

// Mat4 returns the homogeneous 3D rotation matrix corresponding to the quaternion.
func (q1 *Quat) Mat4() Mat4 {
	w, x, y, z := q1.W, q1.V[0], q1.V[1], q1.V[2]
	return Mat4{
		1 - 2*y*y - 2*z*z, 2*x*y + 2*w*z, 2*x*z - 2*w*y, 0,
		2*x*y - 2*w*z, 1 - 2*x*x - 2*z*z, 2*y*z + 2*w*x, 0,
		2*x*z + 2*w*y, 2*y*z - 2*w*x, 1 - 2*x*x - 2*y*y, 0,
		0, 0, 0, 1,
	}
}

// Dot returns the dot product between two quaternions, equivalent to if this was a Vec4
func (q1 *Quat) Dot(q2 *Quat) float32 {
	return q1.W*q2.W + q1.V[0]*q2.V[0] + q1.V[1]*q2.V[1] + q1.V[2]*q2.V[2]
}

// ApproxEqual returns whether the quaternions are approximately equal, as if
// FloatEqual was called on each matching element
func (q1 *Quat) ApproxEqual(q2 *Quat) bool {
	return FloatEqual(q1.W, q2.W) && q1.V.ApproxEqual(&q2.V)
}

// ApproxEqualThreshold returns whether the quaternions are approximately equal with a given tolerence, as if
// FloatEqualThreshold was called on each matching element with the given epsilon
func (q1 *Quat) ApproxEqualThreshold(q2 *Quat, epsilon float32) bool {
	return FloatEqualThreshold(q1.W, q2.W, epsilon) && q1.V.ApproxEqualThreshold(&q2.V, epsilon)
}

// ApproxEqualFunc returns whether the quaternions are approximately equal using the given comparison function, as if
// the function had been called on each individual element
func (q1 *Quat) ApproxEqualFunc(q2 *Quat, f func(float32, float32) bool) bool {
	return f(q1.W, q2.W) && q1.V.ApproxFuncEqual(&q2.V, f)
}

// OrientationEqual returns whether the quaternions represents the same orientation
//
// Different values can represent the same orientation (q == -q) because quaternions avoid singularities
// and discontinuities involved with rotation in 3 dimensions by adding extra dimensions
func (q1 *Quat) OrientationEqual(q2 *Quat) bool {
	return q1.OrientationEqualThreshold(q2, Epsilon)
}

// OrientationEqualThreshold returns whether the quaternions represents the same orientation with a given tolerence
func (q1 *Quat) OrientationEqualThreshold(q2 *Quat, epsilon float32) bool {
	n1 := q1.Normalized()
	n2 := q2.Normalized()
	return math.Abs(n1.Dot(&n2)) > 1-epsilon
}

// QuatSlerp is *S*pherical *L*inear Int*erp*olation, a method of interpolating
// between two quaternions. This always takes the straightest path on the sphere between
// the two quaternions, and maintains constant velocity.
//
// However, it's expensive and QuatSlerp(q1,q2) is not the same as QuatSlerp(q2,q1)
func QuatSlerp(q1, q2 *Quat, amount float32) Quat {
	n1, n2 := q1.Normalized(), q2.Normalized()
	dot := n1.Dot(&n2)

	// If the inputs are too close for comfort, linearly interpolate and normalize the result.
	if dot > 0.9995 {
		return QuatNlerp(&n1, &n2, amount)
	}

	// This is here for precision errors, I'm perfectly aware the *technically* the dot is bound [-1,1], but since Acos will freak out if it's not (even if it's just a liiiiitle bit over due to normal error) we need to clamp it
	dot = Clamp(dot, -1, 1)

	theta := math.Acos(dot) * amount
	c, s := math.Cos(theta), math.Sin(theta)

	n1dot := n1.Scale(dot)
	n2n1dot := n2.Sub(&n1dot)
	rel := n2n1dot.Normalized()
	rel = rel.Scale(s)
	n1c := n1.Scale(c)
	return n1c.Add(&rel)

	/*
		n1, n2 = n1.Normalize(), n2.Normalize()
		dot := n1.Dot(n2)

		// If the inputs are too close for comfort, linearly interpolate and normalize the result.
		if dot > 0.9995 {
			return QuatNlerp(n1, n2, amount)
		}

		// This is here for precision errors, I'm perfectly aware the *technically* the dot is bound [-1,1], but since Acos will freak out if it's not (even if it's just a liiiiitle bit over due to normal error) we need to clamp it
		dot = Clamp(dot, -1, 1)

		theta := math.Acos(dot) * amount
		c, s := math.Cos(theta), math.Sin(theta)
		rel := n2.Sub(n1.Scale(dot)).Normalize()

		return n1.Scale(c).Add(rel.Scale(s))
	*/
}

// QuatLerp is *L*inear Int*erp*olation between two Quaternions, cheap and simple.
//
// Not excessively useful, but uses can be found.
func QuatLerp(q1, q2 *Quat, amount float32) Quat {
	//q1.Add(                        )
	//       q2.Sub(  )
	//              q1 .Scale(amount)
	var1 := q2.Sub(q1)
	var2 := var1.Scale(amount)
	return q1.Add(&var2)
}

// QuatNlerp is *N*ormalized *L*inear Int*erp*olation between two Quaternions. Cheaper than Slerp
// and usually just as good. This is literally Lerp with Normalize() called on it.
//
// Unlike Slerp, constant velocity isn't maintained, but it's much faster and
// Nlerp(q1,q2) and Nlerp(q2,q1) return the same path. You should probably
// use this more often unless you're suffering from choppiness due to the
// non-constant velocity problem.
func QuatNlerp(q1, q2 *Quat, amount float32) Quat {
	l := QuatLerp(q1, q2, amount)
	return l.Normalized()
}

// AnglesToQuat performs a rotation in the specified order. If the order is not
// a valid RotationOrder, this function will panic
//
// The rotation "order" is more of an axis descriptor. For instance XZX would
// tell the function to interpret angle1 as a rotation about the X axis, angle2 about
// the Z axis, and angle3 about the X axis again.
//
// Based off the code for the Matlab function "angle2quat", though this implementation
// only supports 3 single angles as opposed to multiple angles.
func AnglesToQuat(angle1, angle2, angle3 float32, order RotationOrder) Quat {
	s := [3]float64{}
	c := [3]float64{}

	s[0], c[0] = math64.Sincos(float64(angle1 / 2))
	s[1], c[1] = math64.Sincos(float64(angle2 / 2))
	s[2], c[2] = math64.Sincos(float64(angle3 / 2))

	ret := Quat{}
	switch order {
	case ZYX:
		ret.W = float32(c[0]*c[1]*c[2] + s[0]*s[1]*s[2])
		ret.V = Vec3{float32(c[0]*c[1]*s[2] - s[0]*s[1]*c[2]),
			float32(c[0]*s[1]*c[2] + s[0]*c[1]*s[2]),
			float32(s[0]*c[1]*c[2] - c[0]*s[1]*s[2]),
		}
	case ZYZ:
		ret.W = float32(c[0]*c[1]*c[2] - s[0]*c[1]*s[2])
		ret.V = Vec3{float32(c[0]*s[1]*s[2] - s[0]*s[1]*c[2]),
			float32(c[0]*s[1]*c[2] + s[0]*s[1]*s[2]),
			float32(s[0]*c[1]*c[2] + c[0]*c[1]*s[2]),
		}
	case ZXY:
		ret.W = float32(c[0]*c[1]*c[2] - s[0]*s[1]*s[2])
		ret.V = Vec3{float32(c[0]*s[1]*c[2] - s[0]*c[1]*s[2]),
			float32(c[0]*c[1]*s[2] + s[0]*s[1]*c[2]),
			float32(c[0]*s[1]*s[2] + s[0]*c[1]*c[2]),
		}
	case ZXZ:
		ret.W = float32(c[0]*c[1]*c[2] - s[0]*c[1]*s[2])
		ret.V = Vec3{float32(c[0]*s[1]*c[2] + s[0]*s[1]*s[2]),
			float32(s[0]*s[1]*c[2] - c[0]*s[1]*s[2]),
			float32(c[0]*c[1]*s[2] + s[0]*c[1]*c[2]),
		}
	case YXZ:
		ret.W = float32(c[0]*c[1]*c[2] + s[0]*s[1]*s[2])
		ret.V = Vec3{float32(c[0]*s[1]*c[2] + s[0]*c[1]*s[2]),
			float32(s[0]*c[1]*c[2] - c[0]*s[1]*s[2]),
			float32(c[0]*c[1]*s[2] - s[0]*s[1]*c[2]),
		}
	case YXY:
		ret.W = float32(c[0]*c[1]*c[2] - s[0]*c[1]*s[2])
		ret.V = Vec3{float32(c[0]*s[1]*c[2] + s[0]*s[1]*s[2]),
			float32(s[0]*c[1]*c[2] + c[0]*c[1]*s[2]),
			float32(c[0]*s[1]*s[2] - s[0]*s[1]*c[2]),
		}
	case YZX:
		ret.W = float32(c[0]*c[1]*c[2] - s[0]*s[1]*s[2])
		ret.V = Vec3{float32(c[0]*c[1]*s[2] + s[0]*s[1]*c[2]),
			float32(c[0]*s[1]*s[2] + s[0]*c[1]*c[2]),
			float32(c[0]*s[1]*c[2] - s[0]*c[1]*s[2]),
		}
	case YZY:
		ret.W = float32(c[0]*c[1]*c[2] - s[0]*c[1]*s[2])
		ret.V = Vec3{float32(s[0]*s[1]*c[2] - c[0]*s[1]*s[2]),
			float32(c[0]*c[1]*s[2] + s[0]*c[1]*c[2]),
			float32(c[0]*s[1]*c[2] + s[0]*s[1]*s[2]),
		}
	case XYZ:
		ret.W = float32(c[0]*c[1]*c[2] - s[0]*s[1]*s[2])
		ret.V = Vec3{float32(c[0]*s[1]*s[2] + s[0]*c[1]*c[2]),
			float32(c[0]*s[1]*c[2] - s[0]*c[1]*s[2]),
			float32(c[0]*c[1]*s[2] + s[0]*s[1]*c[2]),
		}
	case XYX:
		ret.W = float32(c[0]*c[1]*c[2] - s[0]*c[1]*s[2])
		ret.V = Vec3{float32(c[0]*c[1]*s[2] + s[0]*c[1]*c[2]),
			float32(c[0]*s[1]*c[2] + s[0]*s[1]*s[2]),
			float32(s[0]*s[1]*c[2] - c[0]*s[1]*s[2]),
		}
	case XZY:
		ret.W = float32(c[0]*c[1]*c[2] + s[0]*s[1]*s[2])
		ret.V = Vec3{float32(s[0]*c[1]*c[2] - c[0]*s[1]*s[2]),
			float32(c[0]*c[1]*s[2] - s[0]*s[1]*c[2]),
			float32(c[0]*s[1]*c[2] + s[0]*c[1]*s[2]),
		}
	case XZX:
		ret.W = float32(c[0]*c[1]*c[2] - s[0]*c[1]*s[2])
		ret.V = Vec3{float32(c[0]*c[1]*s[2] + s[0]*c[1]*c[2]),
			float32(c[0]*s[1]*s[2] - s[0]*s[1]*c[2]),
			float32(c[0]*s[1]*c[2] + s[0]*s[1]*s[2]),
		}
	default:
		panic("Unsupported rotation order")
	}
	return ret
}

// Mat4ToQuat converts a pure rotation matrix into a quaternion
func Mat4ToQuat(m *Mat4) Quat {
	// http://www.euclideanspace.com/maths/geometry/rotations/conversions/matrixToQuaternion/index.htm

	if tr := m[0] + m[5] + m[10]; tr > 0 {
		s := 0.5 / math.Sqrt(tr+1.0)
		return Quat{
			0.25 / s,
			Vec3{
				(m[6] - m[9]) * s,
				(m[8] - m[2]) * s,
				(m[1] - m[4]) * s,
			},
		}
	}

	if (m[0] > m[5]) && (m[0] > m[10]) {
		s := 2.0 * math.Sqrt(1.0+m[0]-m[5]-m[10])
		return Quat{
			(m[6] - m[9]) / s,
			Vec3{
				0.25 * s,
				(m[4] + m[1]) / s,
				(m[8] + m[2]) / s,
			},
		}
	}

	if m[5] > m[10] {
		s := 2.0 * math.Sqrt(1.0+m[5]-m[0]-m[10])
		return Quat{
			(m[8] - m[2]) / s,
			Vec3{
				(m[4] + m[1]) / s,
				0.25 * s,
				(m[9] + m[6]) / s,
			},
		}

	}

	s := 2.0 * math.Sqrt(1.0+m[10]-m[0]-m[5])
	return Quat{
		(m[1] - m[4]) / s,
		Vec3{
			(m[8] + m[2]) / s,
			(m[9] + m[6]) / s,
			0.25 * s,
		},
	}
}

// QuatLookAtV creates a rotation from an eye vector to a center vector
//
// It assumes the front of the rotated object at Z- and up at Y+
func QuatLookAtV(eye, center, up *Vec3) Quat {
	// http://www.opengl-tutorial.org/intermediate-tutorials/tutorial-17-quaternions/#I_need_an_equivalent_of_gluLookAt__How_do_I_orient_an_object_towards_a_point__
	// https://bitbucket.org/sinbad/ogre/src/d2ef494c4a2f5d6e2f0f17d3bfb9fd936d5423bb/OgreMain/src/OgreCamera.cpp?at=default#cl-161

	cme := center.Sub(eye)
	direction := cme.Normalized()

	// Find the rotation between the front of the object (that we assume towards Z-,
	// but this depends on your model) and the desired direction
	min1 := Vec3{0, 0, -1}
	rotDir := QuatBetweenVectors(&min1, &direction)

	// Recompute up so that it's perpendicular to the direction
	// You can skip that part if you really want to force up
	//right := direction.Cross(up)
	//up = right.Cross(direction)

	// Because of the 1rst rotation, the up is probably completely screwed up.
	// Find the rotation between the "up" of the rotated object, and the desired up
	dup := Vec3{0, 1, 0}
	upCur := rotDir.Rotate(&dup)
	rotUp := QuatBetweenVectors(&upCur, up)

	rotTarget := rotUp.Mul(&rotDir) // remember, in reverse order.
	return rotTarget.Inverse()      // camera rotation should be inversed!
}

// QuatBetweenVectors calculates the rotation between two vectors
func QuatBetweenVectors(start, dest *Vec3) Quat {
	// http://www.opengl-tutorial.org/intermediate-tutorials/tutorial-17-quaternions/#I_need_an_equivalent_of_gluLookAt__How_do_I_orient_an_object_towards_a_point__
	// https://github.com/g-truc/glm/blob/0.9.5/glm/gtx/quaternion.inl#L225
	// https://bitbucket.org/sinbad/ogre/src/d2ef494c4a2f5d6e2f0f17d3bfb9fd936d5423bb/OgreMain/include/OgreVector3.h?at=default#cl-654

	sn := start.Normalized()
	dn := dest.Normalized()
	epsilon := float32(0.001)

	cosTheta := sn.Dot(&dn)
	if cosTheta < -1.0+epsilon {
		// special case when vectors in opposite directions:
		// there is no "ideal" rotation axis
		// So guess one; any will do as long as it's perpendicular to start
		up := Vec3{1, 0, 0}
		axis := up.Cross(start)
		if axis.Dot(&axis) < epsilon {
			// bad luck, they were parallel, try again!
			up = Vec3{0, 1, 0}
			axis = up.Cross(start)
		}

		naxis := axis.Normalized()
		return QuatRotate(math.Pi, &naxis)
	}

	axis := sn.Cross(&dn)
	s := math.Sqrt((1.0 + cosTheta) * 2.0)

	return Quat{
		s * 0.5,
		axis.Mul(1.0 / s),
	}
}
