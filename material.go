package main

type Material interface {
	Scatter(r Ray, record HitRecord) (bool, Vector, Ray)
}

type Lambertian struct {
	Albedo Vector
}

func NewLambertian(albedo Vector) Lambertian {
	return Lambertian{albedo}
}

func (m Lambertian) Scatter(r Ray, record HitRecord) (bool, Vector, Ray) {
	target := record.P.AddVector(record.Normal).AddVector(RandomInUnitSphere())
	scattered := Ray{record.P, target.SubtractVector(record.P)}

	return true, m.Albedo, scattered
}

type Metal struct {
	Albedo Vector
	Fuzz   float64
}

func NewMetal(albedo Vector, fuzz float64) Metal {
	return Metal{albedo, fuzz}
}

func (m Metal) Scatter(r Ray, record HitRecord) (bool, Vector, Ray) {
	reflected := Reflect(r.Direction.UnitVector(), record.Normal)
	scattered := Ray{record.P, reflected.AddVector(RandomInUnitSphere().MultiplyFloat(m.Fuzz))}

	return (Dot(scattered.Direction, record.Normal) > 0), m.Albedo, scattered
}

type Dielectric struct {
	RefractionIndex float64
}

func NewDielectric(refIndex float64) Dielectric {
	return Dielectric{refIndex}
}

func (m Dielectric) Scatter(r Ray, record HitRecord) (bool, Vector, Ray) {
	reflected := Reflect(r.Direction, record.Normal)
	outwardNormal := Vector{}
	niOverNt := 0.0
	attenuation := Vector{1, 1, 1}

	if Dot(r.Direction, record.Normal) > 0 {
		outwardNormal = record.Normal.MultiplyFloat(-1)
		niOverNt = m.RefractionIndex
	} else {
		outwardNormal = record.Normal
		niOverNt = 1 / m.RefractionIndex
	}

	refractionPossible, refracted := Refract(r.Direction, outwardNormal, niOverNt)

	if refractionPossible {
		return true, attenuation, Ray{record.P, refracted}
	} else {
		return false, attenuation, Ray{record.P, reflected}
	}
}
