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

func Color(r Ray, world Hitable) Vector {
	hit, record := world.Hit(r, 0.0, math.Inf(0))

	if hit {
		v := Vector{
			record.Normal.X,
			record.Normal.Y,
			record.Normal.Z,
		}
		return v.AddFloat(1).Multiply(.5)
	} else {
		unitDirection := r.Direction.UnitVector()
		t := .5 * (unitDirection.Y + 1)
		return Vector{1, 1, 1}.Multiply(1 - t).Add(
			Vector{.5, .7, 1}.Multiply(t),
		)
	}

}

func main() {
	f, err := os.Create("output.ppm")
	check(err)
	defer f.Close()

	nx, ny, ns := 200, 100, 100

	f.WriteString("P3\n")
	f.WriteString(fmt.Sprintf("%d %d\n", nx, ny))
	f.WriteString("255\n")

	world := HitableList{
		[]Hitable{
			Sphere{Vector{0, 0, -1}, .5},
			Sphere{Vector{0, -100.5, -1}, 100},
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
				color = color.Add(Color(r, world))
			}

			color = color.Divide(float64(ns)).Multiply(255.99)

			f.WriteString(
				fmt.Sprintf("%d %d %d\n", color.R(), color.G(), color.B()),
			)
		}
	}

	f.Sync()
}
