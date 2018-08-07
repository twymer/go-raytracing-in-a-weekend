package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddition(t *testing.T) {
	v1 := Vector{1, 1, 1}
	v2 := Vector{0, 1, 2}

	result := v1.AddVector(v2)

	assert.EqualValues(t, Vector{1, 2, 3}, result)
}

func TestSubtraction(t *testing.T) {
	v1 := Vector{1, 1, 1}
	v2 := Vector{0, 1, 2}

	result := v1.SubtractVector(v2)

	assert.EqualValues(t, Vector{1, 0, -1}, result)
}

func TestMultiplication(t *testing.T) {
	v := Vector{0, 1, 2}

	result := v.MultiplyFloat(4)

	assert.EqualValues(t, Vector{0, 4, 8}, result)
}

func TestDivide(t *testing.T) {
	v := Vector{0, 5, 10}

	result := v.DivideFloat(2)

	assert.EqualValues(t, Vector{0, 2.5, 5}, result)
}

func TestCrossProduct(t *testing.T) {
	v1 := Vector{1, 2, 3}
	v2 := Vector{3, 4, 5}

	result := v1.Cross(v2)

	assert.EqualValues(t, Vector{-2, 4, -2}, result)
}

func TestDotProduct(t *testing.T) {
	v1 := Vector{1, 2, 3}
	v2 := Vector{3, 4, 5}

	result := Dot(v1, v2)

	assert.Equal(t, 26.0, result)
}

func TestLength(t *testing.T) {
	v := Vector{3, 4, 5}

	result := v.Length()

	assert.InDelta(t, 7.07106, result, .00001)
}

func TestUnitVector(t *testing.T) {
	t.Skipped()
}
