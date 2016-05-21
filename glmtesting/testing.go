package glmtesting

import (
	"github.com/luxengine/glm"
	"github.com/luxengine/glm/flops/32/flops"
	"github.com/luxengine/math"
)

// FloatEqual returns true if v0 == v1 for every component. Will also return true
// when the components of both vectors are NaN.
func FloatEqual(v0, v1 float32) bool {
	if !flops.Eq(v0, v1) && !(math.IsNaN(v0) && math.IsNaN(v1)) {
		return false
	}
	return true
}

// Vec2Equal returns true if v0 == v1 for every component. Will also return true
// when the components of both vectors are NaN.
func Vec2Equal(v0, v1 glm.Vec2) bool {
	for n := 0; n < len(v0); n++ {
		if !flops.Eq(v0[n], v1[n]) && !(math.IsNaN(v0[n]) && math.IsNaN(v1[n])) {
			return false
		}
	}
	return true
}

// Vec3Equal returns true if v0 == v1 for every component. Will also return true
// when the components of both vectors are NaN.
func Vec3Equal(v0, v1 glm.Vec3) bool {
	for n := 0; n < len(v0); n++ {
		if !flops.Eq(v0[n], v1[n]) && !(math.IsNaN(v0[n]) && math.IsNaN(v1[n])) {
			return false
		}
	}
	return true
}

// Vec4Equal returns true if v0 == v1 for every component. Will also return true
// when the components of both vectors are NaN.
func Vec4Equal(v0, v1 glm.Vec4) bool {
	for n := 0; n < len(v0); n++ {
		if !flops.Eq(v0[n], v1[n]) && !(math.IsNaN(v0[n]) && math.IsNaN(v1[n])) {
			return false
		}
	}
	return true
}
