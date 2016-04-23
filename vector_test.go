package glm

import (
	"testing"
)

func TestVec2_String(t *testing.T) {
	v := Vec2{1, 1}
	if s := v.String(); s != "{ 1.000000,  1.000000}" {
		t.Errorf("(%v).String() = %s, want { 1.000000,  1.000000}", v, s)
	}
}

func TestVec3_String(t *testing.T) {
	v := Vec3{1, 1, 1}
	if s := v.String(); s != "{ 1.000000,  1.000000,  1.000000}" {
		t.Errorf("(%v).String() = %s, want { 1.000000,  1.000000,  1.000000}", v, s)
	}
}

func TestVec4_String(t *testing.T) {
	v := Vec4{1, 1, 1, 1}
	if s := v.String(); s != "{ 1.000000,  1.000000,  1.000000,  1.000000}" {
		t.Errorf("(%v).String() = %s, want { 1.000000,  1.000000,  1.000000,  1.000000}", v, s)
	}
}

func TestVec2_Vec3(t *testing.T) {
	tests := []struct {
		vec2 Vec2
		arg  float32
		vec3 Vec3
	}{
		{Vec2{1, 0}, 1, Vec3{1, 0, 1}},
		{Vec2{-4, 6}, 5, Vec3{-4, 6, 5}},
		{Vec2{0, 0}, -9, Vec3{0, 0, -9}},
	}
	for _, test := range tests {
		if v := test.vec2.Vec3(test.arg); !v.ApproxEqual(&test.vec3) {
			t.Errorf("%v.Vec3(%f) = %v, want %v", test.vec2, test.arg, v, test.vec3)
		}
	}
}

func TestVec2_Vec4(t *testing.T) {
	tests := []struct {
		vec2 Vec2
		args [2]float32
		vec4 Vec4
	}{
		{Vec2{1, 0}, [2]float32{1, 1}, Vec4{1, 0, 1, 1}},
		{Vec2{-4, 6}, [2]float32{5, 5}, Vec4{-4, 6, 5, 5}},
		{Vec2{0, 0}, [2]float32{-9, 9}, Vec4{0, 0, -9, 9}},
	}
	for _, test := range tests {
		if v := test.vec2.Vec4(test.args[0], test.args[1]); !v.ApproxEqual(&test.vec4) {
			t.Errorf("%v.Vec4(%f, %f) = %v, want %v", test.vec2, test.args[0], test.args[1], v, test.vec4)
		}
	}
}

func TestVec3_Vec2(t *testing.T) {
	tests := []struct {
		vec3 Vec3
		vec2 Vec2
	}{
		{Vec3{1, 0, 0}, Vec2{1, 0}},
		{Vec3{3, -2, 5}, Vec2{3, -2}},
		{Vec3{0, 7, 6}, Vec2{0, 7}},
	}
	for _, test := range tests {
		if v := test.vec3.Vec2(); !v.ApproxEqual(&test.vec2) {
			t.Errorf("%v.Vec2() = %v, want %v", test.vec3, v, test.vec2)
		}
	}
}

func TestVec3_Vec4(t *testing.T) {
	tests := []struct {
		vec3 Vec3
		arg  float32
		vec4 Vec4
	}{
		{Vec3{1, 0, 7}, 1, Vec4{1, 0, 7, 1}},
		{Vec3{-4, 6, 7}, 5, Vec4{-4, 6, 7, 5}},
		{Vec3{0, 0, 7}, -9, Vec4{0, 0, 7, -9}},
	}
	for _, test := range tests {
		if v := test.vec3.Vec4(test.arg); !v.ApproxEqual(&test.vec4) {
			t.Errorf("%v.Vec4(%f) = %v, want %v", test.vec3, test.arg, v, test.vec4)
		}
	}
}

func TestVec4_Vec3(t *testing.T) {
	tests := []struct {
		vec4 Vec4
		vec3 Vec3
	}{
		{Vec4{1, 0, 0, 7}, Vec3{1, 0, 0}},
		{Vec4{3, -2, 5, 2}, Vec3{3, -2, 5}},
		{Vec4{0, 7, 6, -12.4}, Vec3{0, 7, 6}},
	}
	for _, test := range tests {
		if v := test.vec4.Vec3(); !v.ApproxEqual(&test.vec3) {
			t.Errorf("%v.Vec3() = %v, want %v", test.vec4, v, test.vec3)
		}
	}
}

func TestVec4_Vec2(t *testing.T) {
	tests := []struct {
		vec4 Vec4
		vec2 Vec2
	}{
		{Vec4{1, 0, 0, 7}, Vec2{1, 0}},
		{Vec4{3, -2, 5, 2}, Vec2{3, -2}},
		{Vec4{0, 7, 6, -12.4}, Vec2{0, 7}},
	}
	for _, test := range tests {
		if v := test.vec4.Vec2(); !v.ApproxEqual(&test.vec2) {
			t.Errorf("%v.Vec2() = %v, want %v", test.vec4, v, test.vec2)
		}
	}
}

func TestVec2_Elem(t *testing.T) {
	tests := []struct {
		vec2 Vec2
		x, y float32
	}{
		{Vec2{1, 2}, 1, 2},
		{Vec2{-5, -9}, -5, -9},
		{Vec2{0, 0}, 0, 0},
	}
	for _, test := range tests {
		if x, y := test.vec2.Elem(); !(FloatEqual(x, test.x) && FloatEqual(y, test.y)) {
			t.Errorf("%v.Elem() = %f, %f, want %f, %f", test.vec2, x, y, test.x, test.y)
		}
	}
}

func TestVec3_Elem(t *testing.T) {
	tests := []struct {
		vec3    Vec3
		x, y, z float32
	}{
		{Vec3{1, 2, 5}, 1, 2, 5},
		{Vec3{-5, -9, -2}, -5, -9, -2},
		{Vec3{0, 0, 14.5}, 0, 0, 14.5},
	}
	for _, test := range tests {
		if x, y, z := test.vec3.Elem(); !(FloatEqual(x, test.x) && FloatEqual(y, test.y) && FloatEqual(z, test.z)) {
			t.Errorf("%v.Elem() = %f, %f, %f, want %f, %f, %f", test.vec3, x, y, z, test.x, test.y, test.z)
		}
	}
}

func TestVec4_Elem(t *testing.T) {
	tests := []struct {
		vec4       Vec4
		x, y, z, w float32
	}{
		{Vec4{1, 2, 5, -17.4}, 1, 2, 5, -17.4},
		{Vec4{-5, -9, -2, 99.2}, -5, -9, -2, 99.2},
		{Vec4{0, 0, 14.5, -111.11}, 0, 0, 14.5, -111.11},
	}
	for _, test := range tests {
		if x, y, z, w := test.vec4.Elem(); !(FloatEqual(x, test.x) && FloatEqual(y, test.y) && FloatEqual(z, test.z) && FloatEqual(w, test.w)) {
			t.Errorf("%v.Elem() = %f, %f, %f, %f, want %f, %f, %f, %f", test.vec4, x, y, z, w, test.x, test.y, test.z, test.w)
		}
	}
}

var crossTests = []struct {
	v1, v2, v3 Vec3
}{
	{Vec3{1, 0, 0}, Vec3{0, 1, 0}, Vec3{0, 0, 1}},
	{Vec3{0, 1, 0}, Vec3{1, 0, 0}, Vec3{0, 0, -1}},

	{Vec3{1, 2, 3}, Vec3{3, 2, 1}, Vec3{-4, 8, -4}},
	{Vec3{3, 2, 1}, Vec3{1, 2, 3}, Vec3{4, -8, 4}},
}

func TestVec3_Cross(t *testing.T) {
	for _, test := range crossTests {
		if c := test.v1.Cross(&test.v2); !c.ApproxEqual(&test.v3) {
			t.Errorf("%v X %v = %v, want %v", test.v1, test.v2, c, test.v3)
		}
	}
}

func TestVec3_CrossOf(t *testing.T) {
	for _, test := range crossTests {
		var c Vec3
		c.CrossOf(&test.v1, &test.v2)
		if !c.ApproxEqual(&test.v3) {
			t.Errorf("%v X %v = %v, want %v", test.v1, test.v2, c, test.v3)
		}
	}
}

func TestVec3_CrossWith(t *testing.T) {
	for _, test := range crossTests {
		c := test.v1
		c.CrossWith(&test.v2)
		if !c.ApproxEqual(&test.v3) {
			t.Errorf("%v X %v = %v, want %v", test.v1, test.v2, c, test.v3)
		}
	}
}

var vec2Tests = []struct {
	v1, v2, add, sub, mul, component, invert, normal Vec2
	f, dot, len                                      float32
}{
	{
		v1:        Vec2{1, 1},
		v2:        Vec2{1, 1},
		add:       Vec2{2, 2},
		sub:       Vec2{0, 0},
		f:         7,
		mul:       Vec2{7, 7},
		component: Vec2{1, 1},
		dot:       2,
		len:       1.4142135623730950488016887242096980785696718753769480,
		invert:    Vec2{-1, -1},
		normal:    Vec2{0.707106781186, 0.707106781186},
	},
	{
		v1:        Vec2{4, 5},
		v2:        Vec2{-4, -5},
		add:       Vec2{0, 0},
		sub:       Vec2{8, 10},
		f:         1,
		mul:       Vec2{4, 5},
		component: Vec2{-16, -25},
		dot:       -41,
		len:       6.40312,
		invert:    Vec2{-4, -5},
		normal:    Vec2{0.624695, 0.780869},
	},
	{
		v1:        Vec2{-1, -2},
		v2:        Vec2{-4, -9},
		add:       Vec2{-5, -11},
		sub:       Vec2{3, 7},
		f:         -1,
		mul:       Vec2{1, 2},
		component: Vec2{4, 18},
		dot:       22,
		len:       2.23607,
		invert:    Vec2{1, 2},
		normal:    Vec2{-0.447214, -0.894427},
	},
	{
		v1:        Vec2{0.1, 0.9},
		v2:        Vec2{0.9, 0.1},
		add:       Vec2{1, 1},
		sub:       Vec2{-0.8, 0.8},
		f:         0.5,
		mul:       Vec2{0.05, 0.45},
		component: Vec2{0.09, 0.09},
		dot:       0.18,
		len:       0.905539,
		invert:    Vec2{-0.1, -0.9},
		normal:    Vec2{0.110432, 0.993884},
	},
}

var vec3Tests = []struct {
	v1, v2, add, sub, mul, component, invert, normal Vec3
	f, dot, len                                      float32
}{
	{
		v1:        Vec3{1, 1, 1},
		v2:        Vec3{1, 1, 1},
		add:       Vec3{2, 2, 2},
		sub:       Vec3{0, 0, 0},
		f:         7,
		mul:       Vec3{7, 7, 7},
		component: Vec3{1, 1, 1},
		dot:       3,
		len:       1.73205,
		invert:    Vec3{-1, -1, -1},
		normal:    Vec3{0.57735, 0.57735, 0.57735},
	},
	{
		v1:        Vec3{4, 5, 7},
		v2:        Vec3{-4, -5, -7},
		add:       Vec3{0, 0, 0},
		sub:       Vec3{8, 10, 14},
		f:         1,
		mul:       Vec3{4, 5, 7},
		component: Vec3{-16, -25, -49},
		dot:       -90,
		len:       9.48683,
		invert:    Vec3{-4, -5, -7},
		normal:    Vec3{0.421637, 0.527046, 0.737865},
	},
	{
		v1:        Vec3{-1, -2, -3},
		v2:        Vec3{-4, -9, -11},
		add:       Vec3{-5, -11, -14},
		sub:       Vec3{3, 7, 8},
		f:         -1,
		mul:       Vec3{1, 2, 3},
		component: Vec3{4, 18, 33},
		dot:       55,
		len:       3.74166,
		invert:    Vec3{1, 2, 3},
		normal:    Vec3{-0.267261, -0.534522, -0.801784},
	},
	{
		v1:        Vec3{0.1, 0.9, 0.3},
		v2:        Vec3{0.9, 0.1, 0.7},
		add:       Vec3{1, 1, 1},
		sub:       Vec3{-0.8, 0.8, -0.4},
		f:         0.5,
		mul:       Vec3{0.05, 0.45, 0.15},
		component: Vec3{0.09, 0.09, 0.21},
		dot:       0.39,
		len:       0.953939,
		invert:    Vec3{-0.1, -0.9, -0.3},
		normal:    Vec3{0.104828, 0.943456, 0.314485},
	},
}

var vec4Tests = []struct {
	v1, v2, add, sub, mul, component, invert, normal Vec4
	f, dot, len                                      float32
}{
	{
		v1:        Vec4{1, 1, 1, 1},
		v2:        Vec4{1, 1, 1, 1},
		add:       Vec4{2, 2, 2, 2},
		sub:       Vec4{0, 0, 0, 0},
		f:         7,
		mul:       Vec4{7, 7, 7, 7},
		component: Vec4{1, 1, 1, 1},
		dot:       4,
		len:       2,
		invert:    Vec4{-1, -1, -1, -1},
		normal:    Vec4{0.5, 0.5, 0.5, 0.5},
	},
	{
		v1:        Vec4{4, 5, 7, 8},
		v2:        Vec4{-4, -5, -7, -8},
		add:       Vec4{0, 0, 0, 0},
		sub:       Vec4{8, 10, 14, 16},
		f:         1,
		mul:       Vec4{4, 5, 7, 8},
		component: Vec4{-16, -25, -49, -64},
		dot:       -154,
		len:       12.4097,
		invert:    Vec4{-4, -5, -7, -8},
		normal:    Vec4{0.322329, 0.402911, 0.564076, 0.644658},
	},
	{
		v1:        Vec4{-1, -2, -3, -4},
		v2:        Vec4{-4, -9, -11, -15},
		add:       Vec4{-5, -11, -14, -19},
		sub:       Vec4{3, 7, 8, 11},
		f:         -1,
		mul:       Vec4{1, 2, 3, 4},
		component: Vec4{4, 18, 33, 60},
		dot:       115,
		len:       5.47723,
		invert:    Vec4{1, 2, 3, 4},
		normal:    Vec4{-0.182574, -0.365148, -0.547723, -0.730297},
	},
	{
		v1:        Vec4{0.1, 0.9, 0.3, 0.5},
		v2:        Vec4{0.9, 0.1, 0.7, 0.5},
		add:       Vec4{1, 1, 1, 1},
		sub:       Vec4{-0.8, 0.8, -0.4, 0},
		f:         0.5,
		mul:       Vec4{0.05, 0.45, 0.15, 0.25},
		component: Vec4{0.09, 0.09, 0.21, 0.25},
		dot:       0.64,
		len:       1.07703,
		invert:    Vec4{-0.1, -0.9, -0.3, -0.5},
		normal:    Vec4{0.0928477, 0.835629, 0.278543, 0.464238},
	},
}

func TestVec2_Add(t *testing.T) {
	for _, test := range vec2Tests {
		c := test.v1.Add(&test.v2)
		if !c.ApproxEqualThreshold(&test.add, 1e-4) {
			t.Errorf("%v + %v = %v, want %v", test.v1, test.v2, c, test.add)
		}
	}
}

func TestVec2_AddOf(t *testing.T) {
	for _, test := range vec2Tests {
		var c Vec2
		c.AddOf(&test.v1, &test.v2)
		if !c.ApproxEqualThreshold(&test.add, 1e-4) {
			t.Errorf("%v + %v = %v, want %v", test.v1, test.v2, c, test.add)
		}
	}
}

func TestVec2_AddWith(t *testing.T) {
	for _, test := range vec2Tests {
		c := test.v1
		c.AddWith(&test.v2)
		if !c.ApproxEqualThreshold(&test.add, 1e-4) {
			t.Errorf("%v + %v = %v, want %v", test.v1, test.v2, c, test.add)
		}
	}
}

func TestVec3_Add(t *testing.T) {
	for _, test := range vec3Tests {
		c := test.v1.Add(&test.v2)
		if !c.ApproxEqualThreshold(&test.add, 1e-4) {
			t.Errorf("%v + %v = %v, want %v", test.v1, test.v2, c, test.add)
		}
	}
}

func TestVec3_AddOf(t *testing.T) {
	for _, test := range vec3Tests {
		var c Vec3
		c.AddOf(&test.v1, &test.v2)
		if !c.ApproxEqualThreshold(&test.add, 1e-4) {
			t.Errorf("%v + %v = %v, want %v", test.v1, test.v2, c, test.add)
		}
	}
}

func TestVec3_AddWith(t *testing.T) {
	for _, test := range vec3Tests {
		c := test.v1
		c.AddWith(&test.v2)
		if !c.ApproxEqualThreshold(&test.add, 1e-4) {
			t.Errorf("%v + %v = %v, want %v", test.v1, test.v2, c, test.add)
		}
	}
}

func TestVec4_Add(t *testing.T) {
	for _, test := range vec4Tests {
		c := test.v1.Add(&test.v2)
		if !c.ApproxEqualThreshold(&test.add, 1e-4) {
			t.Errorf("%v + %v = %v, want %v", test.v1, test.v2, c, test.add)
		}
	}
}

func TestVec4_AddOf(t *testing.T) {
	for _, test := range vec4Tests {
		var c Vec4
		c.AddOf(&test.v1, &test.v2)
		if !c.ApproxEqualThreshold(&test.add, 1e-4) {
			t.Errorf("%v + %v = %v, want %v", test.v1, test.v2, c, test.add)
		}
	}
}

func TestVec4_AddWith(t *testing.T) {
	for _, test := range vec4Tests {
		c := test.v1
		c.AddWith(&test.v2)
		if !c.ApproxEqualThreshold(&test.add, 1e-4) {
			t.Errorf("%v + %v = %v, want %v", test.v1, test.v2, c, test.add)
		}
	}
}

func TestVec2_Sub(t *testing.T) {
	for _, test := range vec2Tests {
		c := test.v1.Sub(&test.v2)
		if !c.ApproxEqualThreshold(&test.sub, 1e-4) {
			t.Errorf("%v - %v = %v, want %v", test.v1, test.v2, c, test.sub)
		}
	}
}

func TestVec2_SubOf(t *testing.T) {
	for _, test := range vec2Tests {
		var c Vec2
		c.SubOf(&test.v1, &test.v2)
		if !c.ApproxEqualThreshold(&test.sub, 1e-4) {
			t.Errorf("%v - %v = %v, want %v", test.v1, test.v2, c, test.sub)
		}
	}
}

func TestVec2_SubWith(t *testing.T) {
	for _, test := range vec2Tests {
		c := test.v1
		c.SubWith(&test.v2)
		if !c.ApproxEqualThreshold(&test.sub, 1e-4) {
			t.Errorf("%v - %v = %v, want %v", test.v1, test.v2, c, test.sub)
		}
	}
}

func TestVec3_Sub(t *testing.T) {
	for _, test := range vec3Tests {
		c := test.v1.Sub(&test.v2)
		if !c.ApproxEqualThreshold(&test.sub, 1e-4) {
			t.Errorf("%v - %v = %v, want %v", test.v1, test.v2, c, test.sub)
		}
	}
}

func TestVec3_SubOf(t *testing.T) {
	for _, test := range vec3Tests {
		var c Vec3
		c.SubOf(&test.v1, &test.v2)
		if !c.ApproxEqualThreshold(&test.sub, 1e-4) {
			t.Errorf("%v - %v = %v, want %v", test.v1, test.v2, c, test.sub)
		}
	}
}

func TestVec3_SubWith(t *testing.T) {
	for _, test := range vec3Tests {
		c := test.v1
		c.SubWith(&test.v2)
		if !c.ApproxEqualThreshold(&test.sub, 1e-4) {
			t.Errorf("%v - %v = %v, want %v", test.v1, test.v2, c, test.sub)
		}
	}
}

func TestVec4_Sub(t *testing.T) {
	for _, test := range vec4Tests {
		c := test.v1.Sub(&test.v2)
		if !c.ApproxEqualThreshold(&test.sub, 1e-4) {
			t.Errorf("%v - %v = %v, want %v", test.v1, test.v2, c, test.sub)
		}
	}
}

func TestVec4_SubOf(t *testing.T) {
	for _, test := range vec4Tests {
		var c Vec4
		c.SubOf(&test.v1, &test.v2)
		if !c.ApproxEqualThreshold(&test.sub, 1e-4) {
			t.Errorf("%v - %v = %v, want %v", test.v1, test.v2, c, test.sub)
		}
	}
}

func TestVec4_SubWith(t *testing.T) {
	for _, test := range vec4Tests {
		c := test.v1
		c.SubWith(&test.v2)
		if !c.ApproxEqualThreshold(&test.sub, 1e-4) {
			t.Errorf("%v - %v = %v, want %v", test.v1, test.v2, c, test.sub)
		}
	}
}

func TestVec2_Mul(t *testing.T) {
	for _, test := range vec2Tests {
		c := test.v1.Mul(test.f)
		if !c.ApproxEqualThreshold(&test.mul, 1e-4) {
			t.Errorf("%v * %f = %v, want %v", test.v1, test.f, c, test.mul)
		}
	}
}

func TestVec2_MulOf(t *testing.T) {
	for _, test := range vec2Tests {
		var c Vec2
		c.MulOf(test.f, &test.v1)
		if !c.ApproxEqualThreshold(&test.mul, 1e-4) {
			t.Errorf("%v * %f = %v, want %v", test.v1, test.f, c, test.mul)
		}
	}
}

func TestVec2_MulWith(t *testing.T) {
	for _, test := range vec2Tests {
		c := test.v1
		c.MulWith(test.f)
		if !c.ApproxEqualThreshold(&test.mul, 1e-4) {
			t.Errorf("%v * %f = %v, want %v", test.v1, test.f, c, test.mul)
		}
	}
}

func TestVec3_Mul(t *testing.T) {
	for _, test := range vec3Tests {
		c := test.v1.Mul(test.f)
		if !c.ApproxEqualThreshold(&test.mul, 1e-4) {
			t.Errorf("%v * %f = %v, want %v", test.v1, test.f, c, test.mul)
		}
	}
}

func TestVec3_MulOf(t *testing.T) {
	for _, test := range vec3Tests {
		var c Vec3
		c.MulOf(test.f, &test.v1)
		if !c.ApproxEqualThreshold(&test.mul, 1e-4) {
			t.Errorf("%v * %f = %v, want %v", test.v1, test.f, c, test.mul)
		}
	}
}

func TestVec3_MulWith(t *testing.T) {
	for _, test := range vec3Tests {
		c := test.v1
		c.MulWith(test.f)
		if !c.ApproxEqualThreshold(&test.mul, 1e-4) {
			t.Errorf("%v * %f = %v, want %v", test.v1, test.f, c, test.mul)
		}
	}
}

func TestVec4_Mul(t *testing.T) {
	for _, test := range vec4Tests {
		c := test.v1.Mul(test.f)
		if !c.ApproxEqualThreshold(&test.mul, 1e-4) {
			t.Errorf("%v * %f = %v, want %v", test.v1, test.f, c, test.mul)
		}
	}
}

func TestVec4_MulOf(t *testing.T) {
	for _, test := range vec4Tests {
		var c Vec4
		c.MulOf(test.f, &test.v1)
		if !c.ApproxEqualThreshold(&test.mul, 1e-4) {
			t.Errorf("%v * %f = %v, want %v", test.v1, test.f, c, test.mul)
		}
	}
}

func TestVec4_MulWith(t *testing.T) {
	for _, test := range vec4Tests {
		c := test.v1
		c.MulWith(test.f)
		if !c.ApproxEqualThreshold(&test.mul, 1e-4) {
			t.Errorf("%v * %f = %v, want %v", test.v1, test.f, c, test.mul)
		}
	}
}

func TestVec2_ComponentProduct(t *testing.T) {
	for _, test := range vec2Tests {
		c := test.v1.ComponentProduct(&test.v2)
		if !c.ApproxEqualThreshold(&test.component, 1e-4) {
			t.Errorf("%v * %v = %v, want %v", test.v1, test.v2, c, test.component)
		}
	}
}

func TestVec2_ComponentProductOf(t *testing.T) {
	for _, test := range vec2Tests {
		var c Vec2
		c.ComponentProductOf(&test.v1, &test.v2)
		if !c.ApproxEqualThreshold(&test.component, 1e-4) {
			t.Errorf("%v * %v = %v, want %v", test.v1, test.v2, c, test.component)
		}
	}
}

func TestVec2_ComponentProductWith(t *testing.T) {
	for _, test := range vec2Tests {
		c := test.v1
		c.ComponentProductWith(&test.v2)
		if !c.ApproxEqualThreshold(&test.component, 1e-4) {
			t.Errorf("%v * %v = %v, want %v", test.v1, test.v2, c, test.component)
		}
	}
}

func TestVec3_ComponentProduct(t *testing.T) {
	for _, test := range vec3Tests {
		c := test.v1.ComponentProduct(&test.v2)
		if !c.ApproxEqualThreshold(&test.component, 1e-4) {
			t.Errorf("%v * %v = %v, want %v", test.v1, test.v2, c, test.component)
		}
	}
}

func TestVec3_ComponentProductOf(t *testing.T) {
	for _, test := range vec3Tests {
		var c Vec3
		c.ComponentProductOf(&test.v1, &test.v2)
		if !c.ApproxEqualThreshold(&test.component, 1e-4) {
			t.Errorf("%v * %v = %v, want %v", test.v1, test.v2, c, test.component)
		}
	}
}

func TestVec3_ComponentProductWith(t *testing.T) {
	for _, test := range vec3Tests {
		c := test.v1
		c.ComponentProductWith(&test.v2)
		if !c.ApproxEqualThreshold(&test.component, 1e-4) {
			t.Errorf("%v * %v = %v, want %v", test.v1, test.v2, c, test.component)
		}
	}
}

func TestVec4_ComponentProduct(t *testing.T) {
	for _, test := range vec4Tests {
		c := test.v1.ComponentProduct(&test.v2)
		if !c.ApproxEqualThreshold(&test.component, 1e-4) {
			t.Errorf("%v * %v = %v, want %v", test.v1, test.v2, c, test.component)
		}
	}
}

func TestVec4_ComponentProductOf(t *testing.T) {
	for _, test := range vec4Tests {
		var c Vec4
		c.ComponentProductOf(&test.v1, &test.v2)
		if !c.ApproxEqualThreshold(&test.component, 1e-4) {
			t.Errorf("%v * %v = %v, want %v", test.v1, test.v2, c, test.component)
		}
	}
}

func TestVec4_ComponentProductWith(t *testing.T) {
	for _, test := range vec4Tests {
		c := test.v1
		c.ComponentProductWith(&test.v2)
		if !c.ApproxEqualThreshold(&test.component, 1e-4) {
			t.Errorf("%v * %v = %v, want %v", test.v1, test.v2, c, test.component)
		}
	}
}

////////////////////////////////////////////////////////////////////////////////////

func TestVec2_Dot(t *testing.T) {
	for _, test := range vec2Tests {
		c := test.v1.Dot(&test.v2)
		if !FloatEqualThreshold(c, test.dot, 1e-4) {
			t.Errorf("%v . %v = %f, want %f", test.v1, test.v2, c, test.dot)
		}
	}
}

func TestVec3_Dot(t *testing.T) {
	for _, test := range vec3Tests {
		c := test.v1.Dot(&test.v2)
		if !FloatEqualThreshold(c, test.dot, 1e-4) {
			t.Errorf("%v . %v = %f, want %f", test.v1, test.v2, c, test.dot)
		}
	}
}

func TestVec4_Dot(t *testing.T) {
	for _, test := range vec4Tests {
		c := test.v1.Dot(&test.v2)
		if !FloatEqualThreshold(c, test.dot, 1e-4) {
			t.Errorf("%v . %v = %f, want %f", test.v1, test.v2, c, test.dot)
		}
	}
}

func TestVec2_Dotf(t *testing.T) {
	for _, test := range vec2Tests {
		c := test.v1.Dotf(test.v2[0], test.v2[1])
		if !FloatEqualThreshold(c, test.dot, 1e-4) {
			t.Errorf("%v . %v = %f, want %f", test.v1, test.v2, c, test.dot)
		}
	}
}

func TestVec3_Dotf(t *testing.T) {
	for _, test := range vec3Tests {
		c := test.v1.Dotf(test.v2[0], test.v2[1], test.v2[2])
		if !FloatEqualThreshold(c, test.dot, 1e-4) {
			t.Errorf("%v . %v = %f, want %f", test.v1, test.v2, c, test.dot)
		}
	}
}

func TestVec4_Dotf(t *testing.T) {
	for _, test := range vec4Tests {
		c := test.v1.Dotf(test.v2[0], test.v2[1], test.v2[2], test.v2[3])
		if !FloatEqualThreshold(c, test.dot, 1e-4) {
			t.Errorf("%v . %v = %f, want %f", test.v1, test.v2, c, test.dot)
		}
	}
}

func TestVec2_Len(t *testing.T) {
	for _, test := range vec2Tests {
		c := test.v1.Len()
		if !FloatEqualThreshold(c, test.len, 1e-4) {
			t.Errorf("%v.Len() = %f, want %f", test.v1, c, test.len)
		}
	}
}

func TestVec3_Len(t *testing.T) {
	for _, test := range vec3Tests {
		c := test.v1.Len()
		if !FloatEqualThreshold(c, test.len, 1e-4) {
			t.Errorf("%v.Len() = %f, want %f", test.v1, c, test.len)
		}
	}
}

func TestVec4_Len(t *testing.T) {
	for _, test := range vec4Tests {
		c := test.v1.Len()
		if !FloatEqualThreshold(c, test.len, 1e-4) {
			t.Errorf("%v.Len() = %f, want %f", test.v1, c, test.len)
		}
	}
}

func TestVec2_Len2(t *testing.T) {
	for _, test := range vec2Tests {
		c := test.v1.Len2()
		if !FloatEqualThreshold(c, test.len*test.len, 1e-4) {
			t.Errorf("%v.Len2() = %f, want %f", test.v1, c, test.len*test.len)
		}
	}
}

func TestVec3_Len2(t *testing.T) {
	for _, test := range vec3Tests {
		c := test.v1.Len2()
		if !FloatEqualThreshold(c, test.len*test.len, 1e-4) {
			t.Errorf("%v.Len2() = %f, want %f", test.v1, c, test.len*test.len)
		}
	}
}

func TestVec4_Len2(t *testing.T) {
	for _, test := range vec4Tests {
		c := test.v1.Len2()
		if !FloatEqualThreshold(c, test.len*test.len, 1e-4) {
			t.Errorf("%v.Len2() = %f, want %f", test.v1, c, test.len*test.len)
		}
	}
}

func TestVec2_Inverse(t *testing.T) {
	for _, test := range vec2Tests {
		c := test.v1.Inverse()
		if !c.ApproxEqualThreshold(&test.invert, 1e-4) {
			t.Errorf("%v.Inverse() = %v, want %v", test.v1, c, test.invert)
		}
	}
}

func TestVec3_Inverse(t *testing.T) {
	for _, test := range vec3Tests {
		c := test.v1.Inverse()
		if !c.ApproxEqualThreshold(&test.invert, 1e-4) {
			t.Errorf("%v.Inverse() = %v, want %v", test.v1, c, test.invert)
		}
	}
}

func TestVec4_Inverse(t *testing.T) {
	for _, test := range vec4Tests {
		c := test.v1.Inverse()
		if !c.ApproxEqualThreshold(&test.invert, 1e-4) {
			t.Errorf("%v.Inverse() = %v, want %v", test.v1, c, test.invert)
		}
	}
}

func TestVec2_Invert(t *testing.T) {
	for _, test := range vec2Tests {
		c := test.v1
		c.Invert()
		if !c.ApproxEqualThreshold(&test.invert, 1e-4) {
			t.Errorf("%v.Inverse() = %v, want %v", test.v1, c, test.invert)
		}
	}
}

func TestVec3_Invert(t *testing.T) {
	for _, test := range vec3Tests {
		c := test.v1
		c.Invert()
		if !c.ApproxEqualThreshold(&test.invert, 1e-4) {
			t.Errorf("%v.Inverse() = %v, want %v", test.v1, c, test.invert)
		}
	}
}

func TestVec4_Invert(t *testing.T) {
	for _, test := range vec4Tests {
		c := test.v1
		c.Invert()
		if !c.ApproxEqualThreshold(&test.invert, 1e-4) {
			t.Errorf("%v.Inverse() = %v, want %v", test.v1, c, test.invert)
		}
	}
}

func TestVec2_Zero(t *testing.T) {
	for _, test := range vec2Tests {
		c := test.v1
		c.Zero()
		if !c.ApproxEqualThreshold(&Vec2{}, 1e-4) {
			t.Errorf("%v.Zero() = %v, want %v", test.v1, c, Vec2{})
		}
	}
}

func TestVec3_Zero(t *testing.T) {
	for _, test := range vec3Tests {
		c := test.v1
		c.Zero()
		if !c.ApproxEqualThreshold(&Vec3{}, 1e-4) {
			t.Errorf("%v.Zero() = %v, want %v", test.v1, c, Vec3{})
		}
	}
}

func TestVec4_Zero(t *testing.T) {
	for _, test := range vec4Tests {
		c := test.v1
		c.Zero()
		if !c.ApproxEqualThreshold(&Vec4{}, 1e-4) {
			t.Errorf("%v.Zero() = %v, want %v", test.v1, c, Vec4{})
		}
	}
}

func TestVec2_Normalized(t *testing.T) {
	for _, test := range vec2Tests {
		c := test.v1.Normalized()
		if !c.ApproxEqualThreshold(&test.normal, 1e-4) {
			t.Errorf("%v.Normal() = %v, want %v", test.v1, c, test.normal)
		}
	}
}

func TestVec2_Normalize(t *testing.T) {
	for _, test := range vec2Tests {
		c := test.v1
		c.Normalize()
		if !c.ApproxEqualThreshold(&test.normal, 1e-4) {
			t.Errorf("%v.Normal() = %v, want %v", test.v1, c, test.normal)
		}
	}
}

func TestVec3_Normalized(t *testing.T) {
	for _, test := range vec3Tests {
		c := test.v1.Normalized()
		if !c.ApproxEqualThreshold(&test.normal, 1e-4) {
			t.Errorf("%v.Normal() = %v, want %v", test.v1, c, test.normal)
		}
	}
}

func TestVec3_Normalize(t *testing.T) {
	for _, test := range vec3Tests {
		c := test.v1
		c.Normalize()
		if !c.ApproxEqualThreshold(&test.normal, 1e-4) {
			t.Errorf("%v.Normal() = %v, want %v", test.v1, c, test.normal)
		}
	}
}

func TestVec4_Normalized(t *testing.T) {
	for _, test := range vec4Tests {
		c := test.v1.Normalized()
		if !c.ApproxEqualThreshold(&test.normal, 1e-4) {
			t.Errorf("%v.Normal() = %v, want %v", test.v1, c, test.normal)
		}
	}
}

func TestVec4_Normalize(t *testing.T) {
	for _, test := range vec4Tests {
		c := test.v1
		c.Normalize()
		if !c.ApproxEqualThreshold(&test.normal, 1e-4) {
			t.Errorf("%v.Normal() = %v, want %v", test.v1, c, test.normal)
		}
	}
}

func TestVec2_AddScaled(t *testing.T) {
	for _, test := range vec2Tests {
		tmp := test.v2.Mul(test.f)
		tmp.AddWith(&test.v1)
		c := test.v1
		c.AddScaledVec(test.f, &test.v2)
		if !c.ApproxEqualThreshold(&tmp, 1e-4) {
			t.Errorf("v1.AddScaled(f, v2) = %v, want %v", c, tmp)
		}
	}
}

func TestVec3_AddScaled(t *testing.T) {
	for _, test := range vec3Tests {
		tmp := test.v2.Mul(test.f)
		tmp.AddWith(&test.v1)
		c := test.v1
		c.AddScaledVec(test.f, &test.v2)
		if !c.ApproxEqualThreshold(&tmp, 1e-4) {
			t.Errorf("v1.AddScaled(f, v2) = %v, want %v", c, tmp)
		}
	}
}

func TestVec4_AddScaled(t *testing.T) {
	for _, test := range vec4Tests {
		tmp := test.v2.Mul(test.f)
		tmp.AddWith(&test.v1)
		c := test.v1
		c.AddScaledVec(test.f, &test.v2)
		if !c.ApproxEqualThreshold(&tmp, 1e-4) {
			t.Errorf("v1.AddScaled(f, v2) = %v, want %v", c, tmp)
		}
	}
}

func TestNormalizeVec2(t *testing.T) {
	for _, test := range vec2Tests {
		c := NormalizeVec2(test.v1)
		if !c.ApproxEqualThreshold(&test.normal, 1e-4) {
			t.Errorf("NormalizeVec2(c) = %v, want %v", c, test.normal)
		}
	}
}

func TestNormalizeVec3(t *testing.T) {
	for _, test := range vec3Tests {
		c := NormalizeVec3(test.v1)
		if !c.ApproxEqualThreshold(&test.normal, 1e-4) {
			t.Errorf("NormalizeVec3(c) = %v, want %v", c, test.normal)
		}
	}
}

func TestNormalizeVec4(t *testing.T) {
	for _, test := range vec4Tests {
		c := NormalizeVec4(test.v1)
		if !c.ApproxEqualThreshold(&test.normal, 1e-4) {
			t.Errorf("NormalizeVec4(c) = %v, want %v", c, test.normal)
		}
	}
}

func TestVec2_SetNormalizeOf(t *testing.T) {
	for _, test := range vec2Tests {
		var c Vec2
		c.SetNormalizeOf(&test.v1)
		if !c.ApproxEqualThreshold(&test.normal, 1e-4) {
			t.Errorf("c.SetNormalizeOf(v1) = %v, want %v", c, test.normal)
		}
	}
}

func TestVec3_SetNormalizeOf(t *testing.T) {
	for _, test := range vec3Tests {
		var c Vec3
		c.SetNormalizeOf(&test.v1)
		if !c.ApproxEqualThreshold(&test.normal, 1e-4) {
			t.Errorf("c.SetNormalizeOf(v1) = %v, want %v", c, test.normal)
		}
	}
}

func TestVec4_SetNormalizeOf(t *testing.T) {
	for _, test := range vec4Tests {
		var c Vec4
		c.SetNormalizeOf(&test.v1)
		if !c.ApproxEqualThreshold(&test.normal, 1e-4) {
			t.Errorf("c.SetNormalizeOf(v1) = %v, want %v", c, test.normal)
		}
	}
}
