package main

import "math"

type Sphere struct {
	Center   Vector
	Radius   float64
	Material Material
}

func (s Sphere) Hit(r *Ray, tMin, tMax float64) (bool, *HitRecord) {
	oc := r.Origin.SubtractVector(s.Center)
	a := Dot(r.Direction, r.Direction)
	b := Dot(oc, r.Direction)
	c := Dot(oc, oc) - s.Radius*s.Radius
	discriminant := b*b - a*c

	if discriminant > 0 {
		temp := (-b - math.Sqrt(b*b-a*c)) / a

		if temp < tMax && temp > tMin {
			p := r.PointAt(temp)
			return true, &HitRecord{
				T:        temp,
				P:        p,
				Normal:   p.SubtractVector(s.Center).DivideFloat(s.Radius),
				Material: s.Material,
			}
		}

		temp = (-b + math.Sqrt(b*b-a*c)) / a

		if temp < tMax && temp > tMin {
			p := r.PointAt(temp)
			return true, &HitRecord{
				T:        temp,
				P:        p,
				Normal:   p.SubtractVector(s.Center).DivideFloat(s.Radius),
				Material: s.Material,
			}
		}
	}

	return false, nil
}
