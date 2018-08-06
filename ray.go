package main

import "math"

type Ray struct {
	Origin, Direction Vector
}

func (r Ray) PointAt(t float64) Vector {
	return r.Origin.Add(r.Direction.Multiply(t))
}

func (r Ray) Color() Vector {
	t := r.HitSphere(Vector{0, 0, -1}, .5)

	if t > 0 {
		n := UnitVector(
			r.PointAt(t).Subtract(Vector{0, 0, -1}),
		)

		return n.AddFloat(1).Multiply(.5)
	}

	unitDirection := UnitVector(r.Direction)
	t = .5*unitDirection.Y + 1.0
	return Vector{1, 1, 1}.Multiply(1 - t).Add(Vector{.5, .7, 1}.Multiply(t))
}

func (r Ray) HitSphere(center Vector, radius float64) float64 {
	oc := r.Origin.Subtract(center)
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
