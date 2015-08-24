package glm

import (
	"math/rand"
	"testing"
)

func TestMat4_Mat3x4(t *testing.T) {
	m4 := Mat4{rand.Float32() * 10, rand.Float32() * 10, rand.Float32() * 10, 0,
		rand.Float32() * 10, rand.Float32() * 10, rand.Float32() * 10, 0,
		rand.Float32() * 10, rand.Float32() * 10, rand.Float32() * 10, 0,
		rand.Float32() * 10, rand.Float32() * 10, rand.Float32() * 10, 1}
	m3x4 := m4.Mat3x4()
	nm4 := m3x4.Mat4()
	if m4 != nm4 {
		t.Errorf("m.Mat3x4().Mat4() =\n%s want\n%sm3x4\n%s", nm4.String(), m4.String(), m3x4.String())
		return
	}
}

func TestMat3x4_Det(t *testing.T) {
	m4 := Mat4{rand.Float32() * 10, rand.Float32() * 10, rand.Float32() * 10, 0,
		rand.Float32() * 10, rand.Float32() * 10, rand.Float32() * 10, 0,
		rand.Float32() * 10, rand.Float32() * 10, rand.Float32() * 10, 0,
		rand.Float32() * 10, rand.Float32() * 10, rand.Float32() * 10, 1}
	m3x4 := m4.Mat3x4()
	if d4, d3 := m4.Det(), m3x4.Det(); d4 != d3 {
		t.Errorf("Det(m) = %f, want %f", d3, d4)
	}
}

func TestMat3x4_Inv(t *testing.T) {
	m4 := Mat4{-1.793091, 5.359944, -5.777826, 0,
		0.408224, -1.586935, 1.855001, 0,
		2.300430, -3.715716, 3.871170, 0,
		6.165272, -18.411301, 19.036600, 1}
	m3x4 := m4.Mat3x4()
	i4 := m4.Inverse()
	i3 := m3x4.Inverse()
	i3_4 := i3.Mat4()
	if !i3_4.ApproxEqualThreshold(&i4, 1e-4) {
		t.Errorf("m.Inv() =\n%s, want\n%s", i3_4.String(), i4.String())
		return
	}
}
