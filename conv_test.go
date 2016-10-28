package glm

import (
	"github.com/EngoEngine/math"

	"testing"
)

var (
	nan  = math.NaN()
	infp = math.Inf(1)
	infm = math.Inf(-1)
)

func TestCartesianToSpherical(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in            Vec3
		r, theta, phi float32
	}{
		{ // http://keisan.casio.com/exec/system/1359533867
			in:    Vec3{5, 12, 9},
			r:     15.8114,
			theta: 0.96525166318993,
			phi:   1.1760052070951,
		},
		{
			in:    Vec3{nan, nan, nan},
			r:     nan,
			theta: nan,
			phi:   nan,
		},
		{ // answer from c++ standard math library.
			in:    Vec3{infp, infp, infp},
			r:     infp,
			theta: nan,
			phi:   0.785398,
		},
		{ // answer from c++ standard math library.
			in:    Vec3{infm, infm, infm},
			r:     infp,
			theta: nan,
			phi:   -2.356194,
		},
	}

	for i, test := range tests {
		if r, theta, phi := CartesianToSpherical(test.in); (!FloatEqualThreshold(r, test.r, 1e-4) && !(math.IsNaN(test.r) && math.IsNaN(r))) ||
			(!FloatEqualThreshold(theta, test.theta, 1e-4) && !(math.IsNaN(test.theta) && math.IsNaN(theta))) ||
			(!FloatEqualThreshold(phi, test.phi, 1e-4) && !(math.IsNaN(test.phi) && math.IsNaN(phi))) {
			t.Errorf("[%d] CartesianToSpherical(%s) = %f, %f, %f want %f, %f, %f", i, test.in.String(), r, theta, phi, test.r, test.theta, test.phi)
		}
	}
}

func TestSphericalToCartesian(t *testing.T) {
	t.Parallel()
	tests := []struct {
		out           Vec3
		r, theta, phi float32
	}{
		{ // http://keisan.casio.com/exec/system/1359533867
			out:   Vec3{5, 12, 9},
			r:     15.8114,
			theta: 0.965250852,
			phi:   1.1760046,
		},
		{
			out:   Vec3{nan, nan, nan},
			r:     nan,
			theta: nan,
			phi:   nan,
		},
		{
			out:   Vec3{nan, nan, nan},
			r:     infp,
			theta: infp,
			phi:   infp,
		},
		{
			out:   Vec3{nan, nan, nan},
			r:     infm,
			theta: infm,
			phi:   infm,
		},
	}

	for i, test := range tests {
		if out := SphericalToCartesian(test.r, test.theta, test.phi); (!FloatEqualThreshold(out[0], test.out[0], 1e-4) && !(math.IsNaN(test.out[0]) && math.IsNaN(out[0]))) ||
			(!FloatEqualThreshold(out[1], test.out[1], 1e-4) && !(math.IsNaN(test.out[1]) && math.IsNaN(out[1]))) ||
			(!FloatEqualThreshold(out[2], test.out[2], 1e-4) && !(math.IsNaN(test.out[2]) && math.IsNaN(out[2]))) {
			t.Errorf("[%d] SphericalToCartesian(%f, %f, %f) = %s want %s", i, test.r, test.theta, test.phi, out.String(), test.out.String())
		}
	}
}

func TestCartesianToCylindrical(t *testing.T) {
	tests := []struct {
		in          Vec3
		rho, phi, z float32
	}{
		{
			in:  Vec3{5, 12, 9},
			rho: 13,
			phi: 1.17601,
			z:   9,
		},
		{
			in:  Vec3{nan, nan, nan},
			rho: nan,
			phi: nan,
			z:   nan,
		},
		{
			in:  Vec3{infp, infp, infp},
			rho: infp,
			phi: 0.785398,
			z:   infp,
		},
		{
			in:  Vec3{infm, infm, infm},
			rho: infp,
			phi: -2.356194,
			z:   infm,
		},
	}
	for i, test := range tests {
		if rho, phi, z := CartesianToCylindrical(test.in); (!FloatEqualThreshold(rho, test.rho, 1e-4) && !(math.IsNaN(test.rho) && math.IsNaN(rho))) ||
			(!FloatEqualThreshold(phi, test.phi, 1e-4) && !(math.IsNaN(test.phi) && math.IsNaN(phi))) ||
			(!FloatEqualThreshold(z, test.z, 1e-4) && !(math.IsNaN(test.z) && math.IsNaN(z))) {
			t.Errorf("[%d] CartesianToCylindrical(%s) = %f, %f, %f want %f, %f, %f", i, test.in.String(), rho, phi, z, test.rho, test.phi, test.z)
		}
	}
}

func TestCylindricalToCartesian(t *testing.T) {
	t.Parallel()
	tests := []struct {
		out         Vec3
		rho, phi, z float32
	}{
		{
			out: Vec3{5, 12, 9},
			rho: 13,
			phi: 1.17601,
			z:   9,
		},
		{
			out: Vec3{nan, nan, nan},
			rho: nan,
			phi: nan,
			z:   nan,
		},
		{
			out: Vec3{nan, nan, infp},
			rho: infp,
			phi: infp,
			z:   infp,
		},
		{
			out: Vec3{nan, nan, infm},
			rho: infm,
			phi: infm,
			z:   infm,
		},
	}

	for i, test := range tests {
		if out := CylindricalToCartesian(test.rho, test.phi, test.z); (!FloatEqualThreshold(out[0], test.out[0], 1e-4) && !(math.IsNaN(test.out[0]) && math.IsNaN(out[0]))) ||
			(!FloatEqualThreshold(out[1], test.out[1], 1e-4) && !(math.IsNaN(test.out[1]) && math.IsNaN(out[1]))) ||
			(!FloatEqualThreshold(out[2], test.out[2], 1e-4) && !(math.IsNaN(test.out[2]) && math.IsNaN(out[2]))) {
			t.Errorf("[%d] CylindricalToCartesian(%f, %f, %f) = %s want %s", i, test.rho, test.phi, test.z, out.String(), test.out.String())
		}
	}
}

func TestSphericalToCylindrical(t *testing.T) {
	t.Parallel()
	tests := []struct {
		out           Vec3
		r, theta, phi float32
	}{
		{
			out:   Vec3{13, 1.17601, 9},
			r:     15.8114,
			theta: 0.965250852,
			phi:   1.1760046,
		},
		{
			out:   Vec3{nan, nan, nan},
			r:     nan,
			theta: nan,
			phi:   nan,
		},
		{
			out:   Vec3{nan, infp, nan},
			r:     infp,
			theta: infp,
			phi:   infp,
		},
		{
			out:   Vec3{nan, infm, nan},
			r:     infm,
			theta: infm,
			phi:   infm,
		},
	}

	for i, test := range tests {
		if rho, phi2, z := SphericalToCylindrical(test.r, test.theta, test.phi); (!FloatEqualThreshold(rho, test.out[0], 1e-4) && !(math.IsNaN(test.out[0]) && math.IsNaN(rho))) ||
			(!FloatEqualThreshold(phi2, test.out[1], 1e-4) && !(math.IsNaN(test.out[1]) && math.IsNaN(phi2))) ||
			(!FloatEqualThreshold(z, test.out[2], 1e-4) && !(math.IsNaN(test.out[2]) && math.IsNaN(z))) {
			t.Errorf("[%d] SphericalToCylindrical(%f, %f, %f) = [%f, %f, %f] want %s", i, test.r, test.theta, test.phi, rho, phi2, z, test.out.String())
		}
	}
}

// work

func TestCylindricalToSpherical(t *testing.T) {
	t.Parallel()
	tests := []struct {
		out         Vec3
		rho, phi, z float32
	}{
		{
			out: Vec3{15.8114, 0.965250852, 1.1760046},
			rho: 13,
			phi: 1.17601,
			z:   9,
		},
		{
			out: Vec3{nan, nan, nan},
			rho: nan,
			phi: nan,
			z:   nan,
		},
		{
			out: Vec3{infp, 0.785398, infp},
			rho: infp,
			phi: infp,
			z:   infp,
		},
		{
			out: Vec3{infp, -2.356194, infm},
			rho: infm,
			phi: infm,
			z:   infm,
		},
	}

	for i, test := range tests {
		if r, theta, phi2 := CylindricalToSpherical(test.rho, test.phi, test.z); (!FloatEqualThreshold(r, test.out[0], 1e-4) && !(math.IsNaN(test.out[0]) && math.IsNaN(r))) ||
			(!FloatEqualThreshold(theta, test.out[1], 1e-4) && !(math.IsNaN(test.out[1]) && math.IsNaN(theta))) ||
			(!FloatEqualThreshold(phi2, test.out[2], 1e-4) && !(math.IsNaN(test.out[2]) && math.IsNaN(phi2))) {
			t.Errorf("[%d] CylindricalToSpherical(%f, %f, %f) = [%f, %f, %f] want %s", i, test.rho, test.phi, test.z, r, theta, phi2, test.out.String())
		}
	}
}

var deg2rad = []struct {
	Deg, Rad float32
}{
	{0, 0},
	{90, math.Pi / 2},
	{180, math.Pi},
	{270, math.Pi + math.Pi/2},
	{360, math.Pi * 2},
	{-90, -math.Pi / 2},
	{-360, -math.Pi * 2},
	{nan, nan},
	{infp, infp},
	{infm, infm},
}

func TestDegToRad(t *testing.T) {
	t.Parallel()
	for i, c := range deg2rad {
		if r := DegToRad(c.Deg); !FloatEqualThreshold(r, c.Rad, 1e-4) && !(math.IsNaN(r) && math.IsNaN(c.Rad)) {
			t.Errorf("[%d] DegToRad(%v) != %v (got %v)", i, c.Deg, c.Rad, r)
		}
	}
}

func TestRadToDeg(t *testing.T) {
	t.Parallel()
	for i, c := range deg2rad {
		if r := RadToDeg(c.Rad); !FloatEqualThreshold(r, c.Deg, 1e-4) && !(math.IsNaN(r) && math.IsNaN(c.Deg)) {
			t.Errorf("[%d] RadToDeg(%v) != %v (got %v)", i, c.Rad, c.Deg, r)
		}
	}
}
