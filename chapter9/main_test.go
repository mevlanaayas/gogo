
package main

import "testing"

type testpair struct {
	values []float64
	expected float64

}

var averageTests = []testpair{
	{ []float64{}, 0 },
	{ []float64{1,1,1,1,1}, 1 },
	{ []float64{-1,1}, 0 },
	{ []float64{-1,1,12}, 4 },

}

var maxTests = []testpair{
	{ []float64{1,3,15}, 15},
	{ []float64{12,11,12,20}, 20},
	{ []float64{0,-1,-2,8}, 8},
	{ []float64{}, 0},

}

var minTests = []testpair{
	{ []float64{1,15}, 1},
	{ []float64{21,12,2}, 2},
	{ []float64{}, 0},
	{ []float64{-3,-87}, -87},
}

func TestAverage(t *testing.T) {
	for _, pair := range averageTests {
		v := Average(pair.values)
		if v != pair.expected {
			t.Error(
				"For", pair.values,
				"expected", pair.expected,
				"got", v,
			)
		}
	}

}

func TestMax(t *testing.T) {
	for _, pair := range maxTests {
		v := Max(pair.values)
		if v != pair.expected {
			t.Error(
				"For", pair.values,
				"expected", pair.expected,
				"got", v,
			)
		}
	}

}

func TestMin(t *testing.T) {
	for _, pair := range minTests {
		v := Min(pair.values)
		if v != pair.expected {
			t.Error(
				"For", pair.values,
				"expected", pair.expected,
				"got", v,
			)
		}
	}

}