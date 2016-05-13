package glm

import (
	"testing"
)

func TestTransform_Iden(t *testing.T) {
	t.Parallel()
	var i Transform
	i.Iden()
	iden4 := Ident4()
	if m := i.Mat4(); m != iden4 {
		t.Error("Transform.Iden does not yield an identity Mat4")
	}
}

func TestTransform_Translate(t *testing.T) {
	t.Parallel()
	tests := []Vec3{
		{1, 2, 3},
		{3, 2, 1},
		{4, 6, 8},
		{1, 5, 9},
		{9, 8, 0},
		{-9, 8, -0},
	}

	for i, test := range tests {
		expect := Translate3D(test[0], test[1], test[2])
		var tr [2]Transform
		tr[0].Iden()
		tr[0].Translate3f(test[0], test[1], test[2])
		tr[1].Iden()
		tr[1].TranslateVec3(&test)
		trm := [2]Mat4{tr[0].Mat4(), tr[1].Mat4()}
		if !expect.EqualThreshold(&trm[0], 1e-4) {
			t.Errorf("[%d] Translate3f\n%snot equal to\n%s", i, expect.String(), trm[0].String())
		}

		if !expect.EqualThreshold(&trm[1], 1e-4) {
			t.Errorf("[%d] Translate3f\n%snot equal to\n%s", i, expect.String(), trm[1].String())
		}
	}
}

func TestTransform_SetTranslate(t *testing.T) {
	t.Parallel()
	tests := []Vec3{
		{1, 2, 3},
		{3, 2, 1},
		{4, 6, 8},
		{1, 5, 9},
		{9, 8, 0},
		{-9, 8, -0},
	}

	for i, test := range tests {
		expect := Translate3D(test[0], test[1], test[2])
		var tr [2]Transform
		tr[0].Iden()
		tr[1].Iden()

		tr[0].Translate3f(test[0], test[1], test[2])
		tr[1].TranslateVec3(&test)
		trm := [2]Mat4{tr[0].Mat4(), tr[1].Mat4()}
		if !expect.EqualThreshold(&trm[0], 1e-4) {
			t.Errorf("[%d] Translate3f\n%snot equal to\n%s", i, expect.String(), trm[0].String())
		}

		if !expect.EqualThreshold(&trm[1], 1e-4) {
			t.Errorf("[%d] Translate3f\n%snot equal to\n%s", i, expect.String(), trm[1].String())
		}
	}
}

/*
func TestSpecial(t *testing.T) {
	var tr Transform
	tr.Iden()

	v := Vec3{1, 0, 0}
	v.Normalize()
	q := QuatRotate(math.Pi, &v)
	q.Normalize()
	tr.RotateQuat(&q)

	tr.Translate3f(0, 5, 0)

	t.Errorf("\n%s", tr.String())
	inv := (*Mat4)(&tr).Inverse()
	t.Errorf("\n%s", inv.String())

	local := Vec3{0, 1, 0}
	world := tr.LocalToWorld(&local)
	t.Errorf("%s", world.String())

	local = tr.WorldToLocal(&world)
	t.Errorf("%s", local.String())
}
*/
