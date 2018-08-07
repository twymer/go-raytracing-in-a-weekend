package main

import "math"

type Ray struct {
	Origin, Direction Vector
}

func (r Ray) PointAt(t float64) Vector {
	return r.Origin.AddVector(r.Direction.MultiplyFloat(t))
}

func (r Ray) HitSphere(center Vector, radius float64) float64 {
	oc := r.Origin.SubtractVector(center)
	a := Dot(r.Direction, r.Direction)
	b := 2 * Dot(oc, r.Direction)
	c := Dot(oc, oc) - radius*radius
	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return -1
	} else {
		return (-b - math.Sqrt(discriminant)) / (2 * a)
	}
}
