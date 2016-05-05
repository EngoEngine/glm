package glm

import (
	"github.com/luxengine/math"
)

// Rotate2D returns a rotation Matrix about a angle in 2-D space. Specifically
// about the origin. It is a 2x2 matrix, if you need a 3x3 for Homogeneous math
// (e.g. composition with a Translation matrix) see HomogRotate2D.
func Rotate2D(angle float32) Mat2 {
	sin, cos := math.Sincos(angle)
	return Mat2{
		cos, sin,
		-sin, cos,
	}
}

// Rotate3DX returns a 3x3 (non-homogeneous) Matrix that rotates by angle about
// the X-axis.
//
// Where c is cos(angle) and s is sin(angle)
//    [1  0  0]
//    [0  c -s]
//    [0  s  c]
func Rotate3DX(angle float32) Mat3 {
	sin, cos := math.Sincos(angle)
	return Mat3{
		1, 0, 0,
		0, cos, sin,
		0, -sin, cos,
	}
}

// Rotate3DY returns a 3x3 (non-homogeneous) Matrix that rotates by angle about
// the Y-axis.
//
// Where c is cos(angle) and s is sin(angle)
//    [c 0 s]
//    [0 1 0]
//    [s 0 c]
func Rotate3DY(angle float32) Mat3 {
	sin, cos := math.Sincos(angle)
	return Mat3{
		cos, 0, -sin,
		0, 1, 0,
		sin, 0, cos,
	}
}

// Rotate3DZ returns a 3x3 (non-homogeneous) Matrix that rotates by angle about
// the Z-axis.
//
// Where c is cos(angle) and s is sin(angle)
//    [c -s  0]
//    [s  c  0]
//    [0  0  1]
func Rotate3DZ(angle float32) Mat3 {
	sin, cos := math.Sincos(angle)
	return Mat3{
		cos, sin, 0,
		-sin, cos, 0,
		0, 0, 1,
	}
}

// Translate2D returns a homogeneous (3x3 for 2D-space) Translation matrix that moves a point by Tx units in the x-direction and Ty units in the y-direction
//
//    [[1, 0, Tx]]
//    [[0, 1, Ty]]
//    [[0, 0, 1 ]]
func Translate2D(Tx, Ty float32) Mat3 {
	return Mat3{
		1, 0, 0,
		0, 1, 0,
		Tx, Ty, 1,
	}
}

// Translate3D returns a homogeneous (4x4 for 3D-space) Translation matrix that
// moves a point by Tx units in the x-direction, Ty units in the y-direction,
// and Tz units in the z-direction.
//
//    [[1, 0, 0, Tx]]
//    [[0, 1, 0, Ty]]
//    [[0, 0, 1, Tz]]
//    [[0, 0, 0, 1 ]]
func Translate3D(Tx, Ty, Tz float32) Mat4 {
	return Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		Tx, Ty, Tz, 1,
	}
}

// HomogRotate2D is the same as Rotate2D, except homogeneous (3x3 with the extra
// row/col being all zeroes with a one in the bottom right).
func HomogRotate2D(angle float32) Mat3 {
	sin, cos := math.Sincos(angle)
	return Mat3{
		cos, sin, 0,
		-sin, cos, 0,
		0, 0, 1,
	}
}

// HomogRotate3DX is the same as Rotate3DX, except homogeneous (4x4 with the
// extra row/col being all zeroes with a one in the bottom right).
func HomogRotate3DX(angle float32) Mat4 {
	sin, cos := math.Sincos(angle)
	return Mat4{
		1, 0, 0, 0,
		0, cos, sin, 0,
		0, -sin, cos, 0,
		0, 0, 0, 1,
	}
}

// HomogRotate3DY is the same as Rotate3DY, except homogeneous (4x4 with the
// extra row/col being all zeroes with a one in the bottom right).
func HomogRotate3DY(angle float32) Mat4 {
	sin, cos := math.Sincos(angle)
	return Mat4{
		cos, 0, -sin, 0,
		0, 1, 0, 0,
		sin, 0, cos, 0,
		0, 0, 0, 1,
	}
}

// HomogRotate3DZ is the same as Rotate3DZ, except homogeneous (4x4 with the
// extra row/col being all zeroes with a one in the bottom right).
func HomogRotate3DZ(angle float32) Mat4 {
	sin, cos := math.Sincos(angle)
	return Mat4{
		cos, sin, 0, 0,
		-sin, cos, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

// Scale3D creates a homogeneous 3D scaling matrix.
// [[ scaleX, 0     , 0     , 0 ]]
// [[ 0     , scaleY, 0     , 0 ]]
// [[ 0     , 0     , scaleZ, 0 ]]
// [[ 0     , 0     , 0     , 1 ]]
func Scale3D(scaleX, scaleY, scaleZ float32) Mat4 {
	return Mat4{
		scaleX, 0, 0, 0,
		0, scaleY, 0, 0,
		0, 0, scaleZ, 0,
		0, 0, 0, 1,
	}
}

// Scale2D creates a homogeneous 2D scaling matrix.
// [[ scaleX, 0     , 0 ]]
// [[ 0     , scaleY, 0 ]]
// [[ 0     , 0     , 1 ]]
func Scale2D(scaleX, scaleY float32) Mat3 {
	return Mat3{
		scaleX, 0, 0,
		0, scaleY, 0,
		0, 0, 1,
	}
}

// ShearX2D creates a homogeneous 2D shear matrix along the X-axis.
func ShearX2D(shear float32) Mat3 {
	return Mat3{1, 0, 0,
		shear, 1, 0,
		0, 0, 1,
	}
}

// ShearY2D creates a homogeneous 2D shear matrix along the Y-axis.
func ShearY2D(shear float32) Mat3 {
	return Mat3{
		1, shear, 0,
		0, 1, 0,
		0, 0, 1,
	}
}

// ShearX3D creates a homogeneous 3D shear matrix along the X-axis.
func ShearX3D(shearY, shearZ float32) Mat4 {
	return Mat4{1, shearY, shearZ, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

// ShearY3D creates a homogeneous 3D shear matrix along the Y-axis
func ShearY3D(shearX, shearZ float32) Mat4 {
	return Mat4{
		1, 0, 0, 0,
		shearX, 1, shearZ, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

// ShearZ3D creates a homogeneous 3D shear matrix along the Z-axis
func ShearZ3D(shearX, shearY float32) Mat4 {
	return Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		shearX, shearY, 1, 0,
		0, 0, 0, 1,
	}
}

// HomogRotate3D creates a 3D rotation Matrix that rotates by (radian) angle
// about some arbitrary axis given by a Vector. It produces a homogeneous
// matrix.
//
// Where c is cos(angle) and s is sin(angle), and x, y, and z are the first,
// second, and third elements of the axis vector (respectively):
//    [[ x^2(1-c)+c, xy(1-c)-zs, xz(1-c)+ys, 0 ]]
//    [[ xy(1-c)+zs, y^2(1-c)+c, yz(1-c)-xs, 0 ]]
//    [[ xz(1-c)-ys, yz(1-c)+xs, z^2(1-c)+c, 0 ]]
//    [[ 0         , 0         , 0         , 1 ]]
func HomogRotate3D(angle float32, axis *Vec3) Mat4 {
	x, y, z := axis[0], axis[1], axis[2]
	s, c := math.Sincos(angle)
	k := 1 - c

	return Mat4{x*x*k + c, x*y*k + z*s, x*z*k - y*s, 0, x*y*k - z*s, y*y*k + c, y*z*k + x*s, 0, x*z*k + y*s, y*z*k - x*s, z*z*k + c, 0, 0, 0, 0, 1}
}

// Extract3DScale extracts the 3d scaling from a homogeneous matrix.
func Extract3DScale(m *Mat4) (x, y, z float32) {
	return math.Sqrt(m[0]*m[0] + m[1]*m[1] + m[2]*m[2]),
		math.Sqrt(m[4]*m[4] + m[5]*m[5] + m[6]*m[6]),
		math.Sqrt(m[8]*m[8] + m[9]*m[9] + m[10]*m[10])
}

// ExtractMaxScale extracts the maximum scaling from a homogeneous matrix.
func ExtractMaxScale(m *Mat4) float32 {
	scaleX := m[0]*m[0] + m[1]*m[1] + m[2]*m[2]
	scaleY := m[4]*m[4] + m[5]*m[5] + m[6]*m[6]
	scaleZ := m[8]*m[8] + m[9]*m[9] + m[10]*m[10]

	return math.Sqrt(math.Max(scaleX, math.Max(scaleY, scaleZ)))
}

// Mat4Normal calculates the Normal of the Matrix (aka the inverse transpose)
func Mat4Normal(m *Mat4) Mat3 {
	n := m.Inverse()
	n.Transpose()
	return n.Mat3()
}

// TransformCoordinate multiplies a 3D vector by a transformation given by the
// homogeneous 4D matrix m, applying any translation. If this transformation is
// non-affine, it will project this vector onto the plane w=1 before returning
// the result.
//
// This is similar to saying you're transforming and projecting a point.
//
// This is effectively equivalent to the GLSL code
//     vec4 r = (m * vec4(v,1.));
//     r = r/r.w;
//     vec3 newV = r[0]yz;
func TransformCoordinate(v *Vec3, m *Mat4) Vec3 {
	t := v.Vec4(1)
	t = m.Mul4x1(&t)
	t.Mul(1 / t[2])

	return t.Vec3()
}

// TransformNormal multiplies a 3D vector by a transformation given by the
// homogeneous 4D matrix m, NOT applying any translations.
//
// This is similar to saying you're applying a transformation to a direction or
// normal. Rotation still applies (as does scaling), but translating a direction
// or normal is meaningless.
//
// This is effectively equivalent to the GLSL code
//    vec4 r = (m * vec4(v,0.));
//    vec3 newV = r[0]yz
func TransformNormal(v *Vec3, m *Mat4) Vec3 {
	t := v.Vec4(0)
	t = m.Mul4x1(&t)

	return t.Vec3()
}

// LocalToWorld applies the transform matrix to the local vector. Its a shortcut
// when using 3x4 matrices.
func LocalToWorld(local *Vec3, transform *Mat3x4) Vec3 {
	return transform.Transform(local)
}

// LocalToWorldIn is the same as LocalToWorld but with destination vector.
func LocalToWorldIn(local *Vec3, transform *Mat3x4, dst *Vec3) {
	transform.TransformIn(local, dst)
}

// WorldToLocal applies the inverse of the transform to the vector.
func WorldToLocal(world *Vec3, transform *Mat3x4) Vec3 {
	return transform.TransformInverse(world)
}

// WorldToLocalIn is a memory friendly version of WorldToLocal.
func WorldToLocalIn(world *Vec3, transform *Mat3x4, dst *Vec3) {
	transform.TransformInverseIn(world, dst)
}

// LocalToWorldDirn transforms this direction by this matrix.
func LocalToWorldDirn(local *Vec3, transform *Mat3x4) Vec3 {
	return transform.TransformDirection(local)
}

// LocalToWorldDirnIn is a memory friendly version of LocalToWorldDirn.
func LocalToWorldDirnIn(local *Vec3, transform *Mat3x4, dst *Vec3) {
	transform.TransformDirectionIn(local, dst)
}

// WorldToLocalDirn inverse transforms this direction by this matrix.
func WorldToLocalDirn(world *Vec3, transform *Mat3x4) Vec3 {
	return transform.TransformInverseDirection(world)
}

// WorldToLocalDirnIn is a memory friendly version of WorldToLocalDirn.
func WorldToLocalDirnIn(world *Vec3, transform *Mat3x4, dst *Vec3) {
	transform.TransformInverseDirectionIn(world, dst)
}
