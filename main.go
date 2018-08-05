package main

import (
	"fmt"
	"os"
)

// Stolen from gobyexample.com
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Create("output.ppm")
	check(err)
	defer f.Close()

	var nx, ny int = 200, 100

	f.WriteString("P3\n")
	f.WriteString(fmt.Sprintf("%d %d\n", nx, ny))
	f.WriteString("255\n")

	lower_left := Vector{-2, -1, -1}
	horizontal := Vector{4, 0, 0}
	vertical := Vector{0, 2, 0}
	origin := Vector{0, 0, 0}

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			u := float64(i) / float64(nx)
			v := float64(j) / float64(ny)

			r := Ray{
				origin,
				lower_left.Add(
					horizontal.Multiply(u),
				).Add(
					vertical.Multiply(v),
				),
			}

			color := r.Color().Multiply(255.99)

			f.WriteString(
				fmt.Sprintf("%d %d %d\n", color.R(), color.G(), color.B()),
			)
		}
	}

	f.Sync()
}
