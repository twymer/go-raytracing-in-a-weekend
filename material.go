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
