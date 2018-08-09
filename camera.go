package main

import (
	"math"
	"math/rand"
)

type Camera struct {
	LowerLeftCorner, Horizontal, Vertical, Origin, U, V, W Vector
	LensRadius                                             float64
}

func NewCamera(lookFrom, lookAt, up Vector, vfov, aspect, aperature, focusDist float64) Camera {
	lensRadius := aperature / 2
	theta := vfov * math.Pi / 180
	halfHeight := math.Tan(theta / 2)
	halfWidth := aspect * halfHeight

	w := lookFrom.SubtractVector(lookAt).UnitVector()
	u := Cross(up, w).UnitVector()
	v := Cross(w, u)

	lowerLeft := lookFrom.SubtractVector(
		u.MultiplyFloat(halfWidth * focusDist),
	).SubtractVector(
		v.MultiplyFloat(halfHeight * focusDist),
	).SubtractVector(w.MultiplyFloat(focusDist))

	return Camera{
		lowerLeft,
		u.MultiplyFloat(2 * halfWidth * focusDist),
		v.MultiplyFloat(2 * halfHeight * focusDist),
		lookFrom,
		u,
		v,
		w,
		lensRadius,
	}
}

func (cam Camera) GetRay(s, t float64) Ray {
	rd := randomInUnitDisk().MultiplyFloat(cam.LensRadius)
	offset := cam.U.MultiplyFloat(rd.X).AddVector(cam.V.MultiplyFloat(rd.Y))

	return Ray{
		cam.Origin.AddVector(offset),
		cam.LowerLeftCorner.AddVector(cam.Horizontal.MultiplyFloat(s)).AddVector(cam.Vertical.MultiplyFloat(t)).SubtractVector(cam.Origin).SubtractVector(offset),
	}
}

func randomInUnitDisk() Vector {
	p := Vector{}
	for {
		p = Vector{rand.Float64(), rand.Float64(), 0}.MultiplyFloat(2).SubtractVector(Vector{1, 1, 0})
		if Dot(p, p) < 1.0 {
			return p
		}
	}
}
