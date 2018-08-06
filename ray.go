package main

type Ray struct {
	Origin, Direction Vector
}

func (r Ray) PointAt(t float64) Vector {
	return r.Origin.Add(r.Direction.Multiply(t))
}

func (r Ray) Color() Vector {
	if r.HitSphere(Vector{0, 0, -1}, .5) {
		return Vector{1, 0, 0}
	}

	unitDirection := UnitVector(r.Direction)
	t := .5*unitDirection.Y + 1.0
	return Vector{1, 1, 1}.Multiply(1 - t).Add(Vector{.5, .7, 1}.Multiply(t))
}

func (r Ray) HitSphere(center Vector, radius float64) bool {
	oc := r.Origin.Subtract(center)
	a := Dot(r.Direction, r.Direction)
	b := 2 * Dot(oc, r.Direction)
	c := Dot(oc, oc) - radius*radius
	discriminant := b*b - 4*a*c
	return discriminant > 0
}
