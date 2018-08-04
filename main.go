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
			var r float32 = float32(i) / float32(nx)
			var g float32 = float32(j) / float32(ny)
			var b float32 = .2

			var ir int = int(255.99 * r)
			var ig int = int(255.99 * g)
			var ib int = int(255.99 * b)

			f.WriteString(fmt.Sprintf("%d %d %d\n", ir, ig, ib))
		}
	}

	f.Sync()
}
