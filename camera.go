package main

import "math"

type Camera struct {
	LowerLeftCorner, Horizontal, Vertical, Origin Vector
}

func NewCamera(vfov, aspect float64) Camera {
	theta := vfov * math.Pi / 180
	halfHeight := math.Tan(theta / 2)
	halfWidth := aspect * halfHeight

	return Camera{
		Vector{-halfWidth, -halfHeight, -1},
		Vector{2 * halfWidth, 0, 0},
		Vector{0, 2 * halfHeight, 0},
		Vector{0, 0, 0},
	}
}

func (cam Camera) GetRay(u, v float64) Ray {
	return Ray{
		cam.Origin,
		cam.LowerLeftCorner.AddVector(cam.Horizontal.MultiplyFloat(u).AddVector(cam.Vertical.MultiplyFloat(v)).SubtractVector(cam.Origin)),
	}
}
