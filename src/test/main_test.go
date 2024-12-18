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
