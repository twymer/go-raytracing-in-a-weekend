package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
)

// Stolen from gobyexample.com
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func RandomInUnitSphere() Vector {
	p := Vector{}
	for {
		p = Vector{
			rand.Float64(),
			rand.Float64(),
			rand.Float64(),
		}.MultiplyFloat(
			2,
		).SubtractVector(
			Vector{1, 1, 1},
		)

		if p.SquaredLength() < 1.0 {
			return p
		}
	}
}

func Color(r Ray, world Hitable, depth int) Vector {
	hit, record := world.Hit(r, 0.001, math.Inf(0))

	if hit {
		wat, attenuation, scattered := record.Material.Scatter(r, record)

		if depth < 50 && wat {
			return attenuation.MultiplyVector(Color(scattered, world, depth+1))
		} else {
			return Vector{0, 0, 0}
		}
	} else {
		unitDirection := r.Direction.UnitVector()
		t := .5 * (unitDirection.Y + 1)
		return Vector{1, 1, 1}.MultiplyFloat(1 - t).AddVector(
			Vector{.5, .7, 1}.MultiplyFloat(t),
		)
	}

}

func main() {
	f, err := os.Create("output.ppm")
	check(err)
	defer f.Close()

	nx, ny, ns := 400, 200, 100

	f.WriteString("P3\n")
	f.WriteString(fmt.Sprintf("%d %d\n", nx, ny))
	f.WriteString("255\n")

	world := HitableList{
		[]Hitable{
			Sphere{Vector{0, 0, -1}, .5, NewLambertian(Vector{.1, .2, .5})},
			Sphere{Vector{0, -100.5, -1}, 100, NewLambertian(Vector{.8, .8, 0})},
			Sphere{Vector{1, 0, -1}, .5, NewMetal(Vector{.8, .6, .2}, .3)},
			Sphere{Vector{-1, 0, -1}, .5, NewDielectric(1.5)},
			Sphere{Vector{-1, 0, -1}, -.45, NewDielectric(1.5)},
		},
	}

	cam := NewCamera()

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			color := Vector{0, 0, 0}

			for s := 0; s < ns; s++ {
				u := (float64(i) + rand.Float64()) / float64(nx)
				v := (float64(j) + rand.Float64()) / float64(ny)

				r := cam.GetRay(u, v)
				color = color.AddVector(Color(r, world, 0))
			}

			color = color.DivideFloat(float64(ns))
			color = Vector{math.Sqrt(color.X), math.Sqrt(color.Y), math.Sqrt(color.Z)}
			color = color.MultiplyFloat(255.99)

			f.WriteString(
				fmt.Sprintf("%d %d %d\n", color.R(), color.G(), color.B()),
			)
		}
	}

	f.Sync()
}
