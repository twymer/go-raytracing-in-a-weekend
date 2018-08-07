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
		cam.LowerLeftCorner.Add(cam.Horizontal.Multiply(u).Add(cam.Vertical.Multiply(v)).Subtract(cam.Origin)),
	}
}
