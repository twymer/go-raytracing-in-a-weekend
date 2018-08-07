package main

type Camera struct {
	LowerLeftCorner, Horizontal, Vertical, Origin Vector
}

func NewCamera() Camera {
	return Camera{
		Vector{-2, -1, -1},
		Vector{4, 0, 0},
		Vector{0, 2, 0},
		Vector{0, 0, 0},
	}
}

func (cam Camera) GetRay(u, v float64) Ray {
	return Ray{
		cam.Origin,
		cam.LowerLeftCorner.AddVector(cam.Horizontal.MultiplyFloat(u).AddVector(cam.Vertical.MultiplyFloat(v)).SubtractVector(cam.Origin)),
	}
}
