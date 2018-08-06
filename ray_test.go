package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPointAt(t *testing.T) {
	v1 := Vector{0, 1, 0}
	v2 := Vector{2, 3, 4}
	ray := Ray{v1, v2}

	result := ray.PointAt(5)

	assert.EqualValues(t, Vector{10, 16, 20}, result)
}
