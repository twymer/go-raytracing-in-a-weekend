package main

import "math"

type Sphere struct {
	Center Vector
	Radius float64
}

func (s Sphere) Hit(r Ray, tMin, tMax float64) (bool, HitRecord) {
	oc := r.Origin.Subtract(s.Center)
	a := Dot(r.Direction, r.Direction)
	b := Dot(oc, r.Direction)
	c := Dot(oc, oc) - s.Radius*s.Radius
	discriminant := b*b - a*c

	record := HitRecord{}

	if discriminant > 0 {
		temp := (-b - math.Sqrt(b*b-a*c)) / a

		if temp < tMax && temp > tMin {
			record.T = temp
			record.P = r.PointAt(temp)
			record.Normal = record.P.Subtract(s.Center).Divide(s.Radius)
			return true, record
		}

		temp = (-b + math.Sqrt(b*b-a*c)) / a

		if temp < tMax && temp > tMin {
			record.T = temp
			record.P = r.PointAt(temp)
			record.Normal = record.P.Subtract(s.Center).Divide(s.Radius)
			return true, record
		}
	}

	return false, record
}
