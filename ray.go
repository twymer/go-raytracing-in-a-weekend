package main

type Ray struct {
	Origin, Direction Vector
}

func (r Ray) PointAt(t float64) Vector {
	return r.Origin.Add(r.Direction.Multiply(t))
}

// func (r Ray) Color() Vector {
// 	unitDirection := UnitVector(r.Direction)
// }
