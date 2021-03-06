package flops

import (
	"testing"
)

var equaltests = []struct {
	data   [2]float64
	result bool
}{
	{
		data:   [2]float64{1, 1},
		result: true,
	},
	{
		data:   [2]float64{2, 2},
		result: true,
	},
	{
		data:   [2]float64{10000, 10000},
		result: true,
	},
	{
		data:   [2]float64{100000000, 100000001},
		result: true,
	},
	{
		data:   [2]float64{1.1, 1},
		result: false,
	},
	{
		data:   [2]float64{-1, 1},
		result: false,
	},
	{
		data:   [2]float64{1, -1},
		result: false,
	},
	{
		data:   [2]float64{2, 1},
		result: false,
	},
	{
		data:   [2]float64{1, 2},
		result: false,
	},
	{
		data:   [2]float64{-2, -1},
		result: false,
	},
	{
		data:   [2]float64{-1, -2},
		result: false,
	},
	{
		data:   [2]float64{0.2, 0.1},
		result: false,
	},
	{
		data:   [2]float64{0.1, 0.2},
		result: false,
	},
	{
		data:   [2]float64{-0.2, -0.1},
		result: false,
	},
	{
		data:   [2]float64{-0.1, -0.2},
		result: false,
	},
}

var ztests = []struct {
	f      float64
	result bool
}{
	{
		f:      0.1,
		result: false,
	},
	{
		f:      0.00000001,
		result: true,
	},
	{
		f:      -0.1,
		result: false,
	},
	{
		f:      -0.00000001,
		result: true,
	},
	{
		f:      0.0000001,
		result: true,
	},
	{
		f:      0.000001,
		result: true,
	},
	{
		f:      0.00001,
		result: false,
	},
	{
		f:      0.0001,
		result: false,
	},
}

func TestRefequal(t *testing.T) {
	for i, test := range equaltests {
		if test.result != refequal(test.data[0], test.data[1]) {
			t.Errorf("[%d] wrong result from %f == %f", i, test.data[0], test.data[1])
		}
	}
}

func TestEq(t *testing.T) {
	for i, test := range equaltests {
		if test.result != Eq(test.data[0], test.data[1]) {
			t.Errorf("[%d] wrong result from %f == %f", i, test.data[0], test.data[1])
		}

		if test.result == Ne(test.data[0], test.data[1]) {
			t.Errorf("[%d] wrong result from %f != %f", i, test.data[0], test.data[1])
		}
	}
}

func TestRefz(t *testing.T) {
	for i, test := range ztests {
		if test.result != refz(test.f) {
			t.Errorf("[%d] wrong result from %f == 0", i, test.f)
		}
	}
}

func TestZ(t *testing.T) {
	for i, test := range ztests {
		if test.result != Z(test.f) {
			t.Errorf("[%d] wrong result from %f == 0", i, test.f)
		}
	}
}
