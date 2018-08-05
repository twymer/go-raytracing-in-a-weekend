package main

type Ray struct {
	Origin, Direction Vector
}

func (r Ray) PointAt(t float64) Vector {
	return r.Origin.Add(r.Direction.Multiply(t))
}

func (r Ray) Color() Vector {
	unitDirection := UnitVector(r.Direction)
	t := .5*unitDirection.Y + 1.0
	return Vector{1, 1, 1}.Multiply(1 - t).Add(Vector{.5, .7, 1}.Multiply(t))
}
