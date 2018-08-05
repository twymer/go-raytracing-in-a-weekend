package main

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

func (v1 Vector) Add(v2 Vector) Vector {
	return Vector{
		v1.X + v2.X,
		v1.Y + v2.Y,
		v1.Z + v2.Z,
	}
}

func (v1 Vector) Sub(v2 Vector) Vector {
	return Vector{
		v1.X - v2.X,
		v1.Y - v2.Y,
		v1.Z - v2.Z,
	}
}

func (v Vector) Mul(t float64) Vector {
	return Vector{
		v.X * t,
		v.Y * t,
		v.Z * t,
	}
}

func (v1 Vector) Cross(v2 Vector) Vector {
	return Vector{
		v1.Y*v2.Z - v1.Z*v2.Y,
		-(v1.X*v2.Z - v1.Z*v2.X),
		v1.X*v2.Y - v1.Y*v2.X,
	}
}

func (v1 Vector) Dot(v2 Vector) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}
