package glmtesting

import (
	"github.com/engoengine/glm"
	"github.com/EngoEngine/math"
	"testing"
)

func TestFloatEqual(t *testing.T) {
	tests := []struct {
		a, b   float32
		result bool
	}{
		{0, 0, true},
		{math.NaN(), 0, false},
		{0, math.NaN(), false},
		{math.NaN(), math.NaN(), true},
	}
	for i, test := range tests {
		if res := FloatEqual(test.a, test.b); res != test.result {
			t.Errorf("[%d] result = %t, want %t", i, res, test.result)
		}
	}
}

func TestVec2Equal(t *testing.T) {
	tests := []struct {
		a, b   glm.Vec2
		result bool
	}{
		{glm.Vec2{0, 0}, glm.Vec2{0, 0}, true},
		{glm.Vec2{math.NaN(), 0}, glm.Vec2{0, 0}, false},
		{glm.Vec2{0, math.NaN()}, glm.Vec2{0, 0}, false},
		{glm.Vec2{math.NaN(), math.NaN()}, glm.Vec2{math.NaN(), math.NaN()}, true},
	}
	for i, test := range tests {
		if res := Vec2Equal(test.a, test.b); res != test.result {
			t.Errorf("[%d] result = %t, want %t", i, res, test.result)
		}
	}
}

func TestVec3Equal(t *testing.T) {
	tests := []struct {
		a, b   glm.Vec3
		result bool
	}{
		{glm.Vec3{0, 0, 0}, glm.Vec3{0, 0, 0}, true},
		{glm.Vec3{math.NaN(), 0, 0}, glm.Vec3{0, 0, 0}, false},
		{glm.Vec3{0, math.NaN(), 0}, glm.Vec3{0, 0, 0}, false},
		{glm.Vec3{0, 0, math.NaN()}, glm.Vec3{0, 0, 0}, false},
		{glm.Vec3{math.NaN(), math.NaN(), math.NaN()}, glm.Vec3{math.NaN(), math.NaN(), math.NaN()}, true},
	}
	for i, test := range tests {
		if res := Vec3Equal(test.a, test.b); res != test.result {
			t.Errorf("[%d] result = %t, want %t", i, res, test.result)
		}
	}
}

func TestVec4Equal(t *testing.T) {
	tests := []struct {
		a, b   glm.Vec4
		result bool
	}{
		{glm.Vec4{0, 0, 0, 0}, glm.Vec4{0, 0, 0, 0}, true},
		{glm.Vec4{math.NaN(), 0, 0, 0}, glm.Vec4{0, 0, 0, 0}, false},
		{glm.Vec4{0, math.NaN(), 0, 0}, glm.Vec4{0, 0, 0, 0}, false},
		{glm.Vec4{0, 0, math.NaN(), 0}, glm.Vec4{0, 0, 0, 0}, false},
		{glm.Vec4{0, 0, 0, math.NaN()}, glm.Vec4{0, 0, 0, 0}, false},
		{glm.Vec4{math.NaN(), math.NaN(), math.NaN(), math.NaN()}, glm.Vec4{math.NaN(), math.NaN(), math.NaN(), math.NaN()}, true},
	}
	for i, test := range tests {
		if res := Vec4Equal(test.a, test.b); res != test.result {
			t.Errorf("[%d] result = %t, want %t", i, res, test.result)
		}
	}
}
