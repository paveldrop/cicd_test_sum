package main

import (
	sum "sum/internal"
	"testing"
)

func TestMainSum(t *testing.T) {
	result := sum.SumOfValues(2, 2)
	expected := 4.0
	if result != expected {
		t.Errorf("result=%f\nexpected=%f", result, expected)
	}
}

func TestMultiValues(t *testing.T) {
	val1 := 2.0
	var expected, result float64
	for i := 2; i < 10; i++ {
		expected = float64(i) + val1
		result = sum.SumOfValues(val1, float64(i))
		if result != expected {
			t.Errorf("result=%f\nexpected=%f", result, expected)
		}
	}
}
