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

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			v := Vector{
				float64(i) / float64(nx),
				float64(j) / float64(ny),
				0.2,
			}

			v = v.Mul(255.99)

			f.WriteString(fmt.Sprintf("%d %d %d\n", v.R(), v.G(), v.B()))
		}
	}

	f.Sync()
}
