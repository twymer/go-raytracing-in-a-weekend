package main

import "testing"

func compareVectors(t *testing.T, expected Vector, got Vector) {
	if expected.Array() != got.Array() {
		t.Errorf("Failure: Expected %v, got %v", expected, got)
	}
}

func compareFloats(t *testing.T, expected float64, got float64) {
	if expected != got {
		t.Errorf("Failure: Expected %f, got %f", expected, got)
	}
}

func TestAddition(t *testing.T) {
	v1 := Vector{1, 1, 1}
	v2 := Vector{0, 1, 2}

	result := v1.Add(v2)

	compareVectors(
		t,
		Vector{1, 2, 3},
		result,
	)
}

func TestSubtraction(t *testing.T) {
	v1 := Vector{1, 1, 1}
	v2 := Vector{0, 1, 2}

	result := v1.Sub(v2)

	compareVectors(
		t,
		Vector{1, 0, -1},
		result,
	)
}

func TestMultiplication(t *testing.T) {
	v := Vector{0, 1, 2}

	result := v.Mul(4)

	compareVectors(
		t,
		Vector{0, 4, 8},
		result,
	)
}

func TestCrossProduct(t *testing.T) {
	v1 := Vector{1, 2, 3}
	v2 := Vector{3, 4, 5}

	result := v1.Cross(v2)

	compareVectors(
		t,
		Vector{-2, 4, -2},
		result,
	)
}

func TestDotProduct(t *testing.T) {
	v1 := Vector{1, 2, 3}
	v2 := Vector{3, 4, 5}

	result := v1.Dot(v2)

	compareFloats(
		t,
		26,
		result,
	)
}
