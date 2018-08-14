package main

import (
	"math"
	"math/rand"
)

type Material interface {
	Scatter(r *Ray, record HitRecord) (bool, Vector, *Ray)
}

type Lambertian struct {
	Albedo Vector
}

func NewLambertian(albedo Vector) Lambertian {
	return Lambertian{albedo}
}

func (m Lambertian) Scatter(r *Ray, record HitRecord) (bool, Vector, *Ray) {
	target := record.P.AddVector(record.Normal).AddVector(RandomInUnitSphere())
	scattered := &Ray{record.P, target.SubtractVector(record.P)}

	return true, m.Albedo, scattered
}

type Metal struct {
	Albedo Vector
	Fuzz   float64
}

func NewMetal(albedo Vector, fuzz float64) Metal {
	return Metal{albedo, fuzz}
}

func (m Metal) Scatter(r *Ray, record HitRecord) (bool, Vector, *Ray) {
	reflected := Reflect(r.Direction.UnitVector(), record.Normal)
	scattered := &Ray{record.P, reflected.AddVector(RandomInUnitSphere().MultiplyFloat(m.Fuzz))}

	return (Dot(scattered.Direction, record.Normal) > 0), m.Albedo, scattered
}

type Dielectric struct {
	RefractionIndex float64
}

func NewDielectric(refIndex float64) Dielectric {
	return Dielectric{refIndex}
}

func (m Dielectric) Scatter(r *Ray, record HitRecord) (bool, Vector, *Ray) {
	outwardNormal := Vector{}
	scattered := Ray{}
	reflected := Reflect(r.Direction, record.Normal)
	niOverNt, cosine, reflectProb := 0.0, 0.0, 0.0
	attenuation := Vector{1, 1, 1}

	if Dot(r.Direction, record.Normal) > 0 {
		outwardNormal = record.Normal.MultiplyFloat(-1)
		niOverNt = m.RefractionIndex
		cosine = m.RefractionIndex * Dot(r.Direction, record.Normal) / r.Direction.Length()
	} else {
		outwardNormal = record.Normal
		niOverNt = 1 / m.RefractionIndex
		cosine = -Dot(r.Direction, record.Normal) / r.Direction.Length()
	}

	refractionPossible, refracted := Refract(r.Direction, outwardNormal, niOverNt)

	if refractionPossible {
		reflectProb = schlick(cosine, m.RefractionIndex)
	} else {
		reflectProb = 1.0
		scattered = Ray{record.P, reflected}
	}

	if rand.Float64() < reflectProb {
		scattered = Ray{record.P, reflected}
	} else {
		scattered = Ray{record.P, refracted}
	}

	return true, attenuation, &scattered
}

func schlick(cosine, refractionIndex float64) float64 {
	r0 := (1 - refractionIndex) / (1 + refractionIndex)
	r0 = r0 * r0

	return r0 + (1-r0)*math.Pow((1-cosine), 5)
}
