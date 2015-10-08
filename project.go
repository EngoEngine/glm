// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package glm

import (
	"errors"
	"github.com/luxengine/math"
)

// Ortho returns a Mat4 that represents a orthographic projection from the given arguments.
func Ortho(left, right, bottom, top, near, far float32) Mat4 {
	rml, tmb, fmn := (right - left), (top - bottom), (far - near)

	return Mat4{float32(2. / rml), 0, 0, 0, 0, float32(2. / tmb), 0, 0, 0, 0, float32(-2. / fmn), 0, float32(-(right + left) / rml), float32(-(top + bottom) / tmb), float32(-(far + near) / fmn), 1}
}

// Ortho2D is equivalent to Ortho with the near and far planes being -1 and 1, respectively
func Ortho2D(left, right, bottom, top float32) Mat4 {
	return Ortho(left, right, bottom, top, -1, 1)
}

// Perspective returns a Mat4 representing a perspective projection of the given arguments.
func Perspective(fovy, aspect, near, far float32) *Mat4 {
	// fovy = (fovy * math.Pi) / 180.0 // convert from degrees to radians
	nmf, f := near-far, 1./math.Tan(fovy/2.0)

	return &Mat4{float32(f / aspect), 0, 0, 0, 0, float32(f), 0, 0, 0, 0, float32((near + far) / nmf), -1, 0, 0, float32((2. * far * near) / nmf), 0}
}

// Frustum returns a Mat4 representing a frustrum transform (squared pyramid with the top cut off)
func Frustum(left, right, bottom, top, near, far float32) Mat4 {
	rml, tmb, fmn := (right - left), (top - bottom), (far - near)
	A, B, C, D := (right+left)/rml, (top+bottom)/tmb, -(far+near)/fmn, -(2*far*near)/fmn

	return Mat4{float32((2. * near) / rml), 0, 0, 0, 0, float32((2. * near) / tmb), 0, 0, float32(A), float32(B), float32(C), -1, 0, 0, float32(D), 0}
}

// LookAt returns a Mat4 that represents a camera transform from the given arguments.
func LookAt(eyeX, eyeY, eyeZ, centerX, centerY, centerZ, upX, upY, upZ float32) *Mat4 {
	return LookAtV(&Vec3{eyeX, eyeY, eyeZ}, &Vec3{centerX, centerY, centerZ}, &Vec3{upX, upY, upZ})
}

// LookAtV generates a transform matrix from world space into the specific eye space
func LookAtV(eye, center, up *Vec3) *Mat4 {
	var f Vec3
	f.SubOf(center, eye)
	f.Normalize()
	var nup Vec3
	nup.SetNormalizeOf(up)
	var s Vec3
	s.CrossOf(&f, &nup)
	s.Normalize()
	var u Vec3
	u.CrossOf(&s, &f)

	M := Mat4{
		s.X(), u.X(), -f.X(), 0,
		s.Y(), u.Y(), -f.Y(), 0,
		s.Z(), u.Z(), -f.Z(), 0,
		0, 0, 0, 1,
	}

	ret := M.Mul4(Translate3D(float32(-eye.X()), float32(-eye.Y()), float32(-eye.Z())))
	return &ret
}

// Project transforms a set of coordinates from object space (in obj) to window coordinates (with depth)
//
// Window coordinates are continuous, not discrete (well, as continuous as an IEEE Floating Point can be), so you won't get exact pixel locations
// without rounding or similar
func Project(obj *Vec3, modelview, projection *Mat4, initialX, initialY, width, height int) (win *Vec3) {
	obj4 := obj.Vec4(1)

	pm := projection.Mul4(modelview)
	vpp := pm.Mul4x1(&obj4)
	win = &Vec3{}
	win[0] = float32(initialX) + (float32(width)*(vpp[0]+1))/2
	win[1] = float32(initialY) + (float32(height)*(vpp[1]+1))/2
	win[2] = (vpp[2] + 1) / 2

	return win
}

// UnProject transforms a set of window coordinates to object space. If your MVP (projection.Mul(modelview) matrix is not invertible, this will return an error
//
// Note that the projection may not be perfect if you use strict pixel locations rather than the exact values given by Projectf.
// (It's still unlikely to be perfect due to precision errors, but it will be closer)
func UnProject(win *Vec3, modelview, projection *Mat4, initialX, initialY, width, height int) (obj Vec3, err error) {
	pm := projection.Mul4(modelview)
	inv := pm.Inverse()
	if inv == (Mat4{}) {
		return Vec3{}, errors.New("Could not find matrix inverse (projection times modelview is probably non-singular)")
	}

	obj4 := inv.Mul4x1(&Vec4{
		(2 * (win[0] - float32(initialX)) / float32(width)) - 1,
		(2 * (win[1] - float32(initialY)) / float32(height)) - 1,
		2*win[2] - 1,
		1.0,
	})
	obj = obj4.Vec3()

	//if obj4[3] > MinValue {}
	obj[0] /= obj4[3]
	obj[1] /= obj4[3]
	obj[2] /= obj4[3]

	return obj, nil
}
