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
}

func NewMetal(albedo Vector) Metal {
	return Metal{albedo}
}

func (m Metal) Scatter(r Ray, record HitRecord) (bool, Vector, Ray) {
	reflected := Reflect(r.Direction.UnitVector(), record.Normal)
	scattered := Ray{record.P, reflected}

	return (Dot(scattered.Direction, record.Normal) > 0), m.Albedo, scattered
}
