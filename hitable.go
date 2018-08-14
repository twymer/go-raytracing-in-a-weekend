package main

type HitRecord struct {
	T         float64
	P, Normal Vector
	Material  Material
}

type Hitable interface {
	Hit(r *Ray, tMin, tMax float64) (bool, HitRecord)
}

type HitableList struct {
	List []Hitable
}

func (hl HitableList) Hit(r *Ray, tMin, tMax float64) (bool, HitRecord) {
	hitAnything := false
	closest := tMax
	closestPoint := HitRecord{}

	for _, hitable := range hl.List {
		hit, hitRecord := hitable.Hit(r, tMin, closest)

		if hit {
			hitAnything = true
			closest = hitRecord.T
			closestPoint = hitRecord
		}
	}

	return hitAnything, closestPoint
}
