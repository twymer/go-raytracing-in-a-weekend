package main

type Ray struct {
	Origin, Destination Vector
}

func (r Ray) PointAt(t float64) Vector {
	return r.Origin.Add(r.Destination.Mul(t))
}
