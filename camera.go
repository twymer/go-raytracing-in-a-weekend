package main

import "math"

type Camera struct {
	LowerLeftCorner, Horizontal, Vertical, Origin Vector
}

func NewCamera(lookFrom, lookAt, up Vector, vfov, aspect float64) Camera {
	theta := vfov * math.Pi / 180
	halfHeight := math.Tan(theta / 2)
	halfWidth := aspect * halfHeight

	w := lookFrom.SubtractVector(lookAt).UnitVector()
	u := Cross(up, w).UnitVector()
	v := Cross(w, u)

	lowerLeft := lookFrom.SubtractVector(u.MultiplyFloat(halfWidth)).SubtractVector(v.MultiplyFloat(halfHeight)).SubtractVector(w)

	return Camera{
		lowerLeft,
		u.MultiplyFloat(2 * halfWidth),
		v.MultiplyFloat(2 * halfHeight),
		lookFrom,
	}
}

func (cam Camera) GetRay(u, v float64) Ray {
	return Ray{
		cam.Origin,
		cam.LowerLeftCorner.AddVector(cam.Horizontal.MultiplyFloat(u).AddVector(cam.Vertical.MultiplyFloat(v)).SubtractVector(cam.Origin)),
	}
}
