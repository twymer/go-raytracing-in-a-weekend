package main

import "testing"

func TestPointAt(t *testing.T) {
	v1 := Vector{0, 1, 0}
	v2 := Vector{2, 3, 4}
	ray := Ray{v1, v2}

	result := ray.PointAt(5)

	compareVectors(
		t,
		Vector{10, 16, 20},
		result,
	)
}
