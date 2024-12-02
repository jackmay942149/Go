package main

import (
	"chapter9/math"
	"testing"
)

type testPair struct {
	values []float64
	min    float64
	max    float64
}

var tests = []testPair{
	{[]float64{1, 2, 3}, 1, 3},
	{[]float64{1, 1, 1, 2, 6}, 1, 6},
	{[]float64{-1.3, 2, 3}, -1.3, 3},
}

func TestMinMax(t *testing.T) {
	for _, pair := range tests {
		minimum := math.Min(pair.values)
		maximum := math.Max(pair.values)
		if maximum != pair.max || minimum != pair.min {
			t.Error(
				"For", pair.values,
				"Expected Max", pair.max,
				"Expected Min", pair.min,
				"Got Max", maximum,
				"Got Min", minimum,
			)
		}
	}
}
