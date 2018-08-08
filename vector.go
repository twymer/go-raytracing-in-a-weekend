package main

import "math"

type Vector struct {
	X, Y, Z float64
}

func (v Vector) R() int {
	return int(v.X)
}

func (v Vector) G() int {
	return int(v.Y)
}

func (v Vector) B() int {
	return int(v.Z)
}

func (v Vector) Array() [3]float64 {
	return [3]float64{
		v.X,
		v.Y,
		v.Z,
	}
}

func (v1 Vector) AddVector(v2 Vector) Vector {
	return Vector{
		v1.X + v2.X,
		v1.Y + v2.Y,
		v1.Z + v2.Z,
	}
}

func (v Vector) AddFloat(t float64) Vector {
	return Vector{
		v.X + t,
		v.Y + t,
		v.Z + t,
	}
}

func (v1 Vector) SubtractVector(v2 Vector) Vector {
	return Vector{
		v1.X - v2.X,
		v1.Y - v2.Y,
		v1.Z - v2.Z,
	}
}

func (v Vector) SubtractFloat(t float64) Vector {
	return Vector{
		v.X - t,
		v.Y - t,
		v.Z - t,
	}
}

func (v Vector) MultiplyFloat(t float64) Vector {
	return Vector{
		v.X * t,
		v.Y * t,
		v.Z * t,
	}
}

func (v1 Vector) MultiplyVector(v2 Vector) Vector {
	return Vector{
		v1.X * v2.X,
		v1.Y * v2.Y,
		v1.Z * v2.Z,
	}
}

func (v Vector) DivideFloat(t float64) Vector {
	return Vector{
		v.X / t,
		v.Y / t,
		v.Z / t,
	}
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vector) SquaredLength() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v1 Vector) Cross(v2 Vector) Vector {
	return Vector{
		v1.Y*v2.Z - v1.Z*v2.Y,
		-(v1.X*v2.Z - v1.Z*v2.X),
		v1.X*v2.Y - v1.Y*v2.X,
	}
}

func Dot(v1, v2 Vector) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func (v Vector) UnitVector() Vector {
	return v.DivideFloat(v.Length())
}

func Reflect(v, n Vector) Vector {
	return v.SubtractVector(n.MultiplyFloat(2 * Dot(v, n)))
}

func Refract(v, n Vector, niOverNt float64) (bool, Vector) {
	uv := v.UnitVector()
	dt := Dot(uv, n)
	discriminant := 1 - niOverNt*niOverNt*(1-dt*dt)

	if discriminant > 0 {
		return true, uv.SubtractVector(n.MultiplyFloat(dt)).MultiplyFloat(niOverNt).SubtractVector(n.MultiplyFloat(math.Sqrt(discriminant)))
	} else {
		return false, Vector{}
	}
}
