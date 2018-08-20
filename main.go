package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
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

func Color(r *Ray, world Hitable, depth int) Vector {
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

func randomScene() []Hitable {
	// Due to loop being nested -11 -> 11 and us adding one manually
	n := 22*22 + 1

	list := make([]Hitable, n)

	list[0] = Sphere{
		Vector{0, -1000, 0},
		1000,
		NewLambertian(Vector{.5, .5, .5}),
	}

	i := 1

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := rand.Float64()
			center := Vector{
				float64(a) + .9*rand.Float64(),
				.2,
				float64(b) + .9*rand.Float64(),
			}

			if (center.SubtractVector(Vector{4, .2, 0}).Length() > .9) {
				if chooseMat < .8 {
					// Diffuse
					list[i] = Sphere{
						center,
						.2,
						NewLambertian(
							Vector{
								rand.Float64() * rand.Float64(),
								rand.Float64() * rand.Float64(),
								rand.Float64() * rand.Float64(),
							},
						),
					}
				} else if chooseMat < .95 {
					// Metal
					list[i] = Sphere{
						center,
						.2,
						NewMetal(
							Vector{
								.5 * (1 + rand.Float64()),
								.5 * (1 + rand.Float64()),
								.5 * (1 + rand.Float64()),
							},
							.5*rand.Float64(),
						),
					}
				} else {
					// Glass
					list[i] = Sphere{
						center,
						.2,
						NewDielectric(1.5),
					}
				}

				i++
			}
		}
	}

	list[i] = Sphere{Vector{0, 1, 0}, 1, NewDielectric(1.5)}
	i++

	list[i] = Sphere{
		Vector{-4, 1, 0},
		1,
		NewLambertian(Vector{.4, .2, .1}),
	}
	i++

	list[i] = Sphere{
		Vector{4, 1, 0},
		1,
		NewMetal(Vector{.7, .6, .5}, 0),
	}

	return list
}

func initializeCamera(nx, ny int) Camera {
	lookFrom := Vector{13, 2, 3}
	lookAt := Vector{0, 0, 0}

	return NewCamera(
		lookFrom,
		lookAt,
		Vector{0, 1, 0},
		20,
		float64(nx)/float64(ny),
		.1,
		10,
	)
}

func computePixelColor(i, j, nx, ny, ns int, cam Camera, world HitableList) [3]int {
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

	return [3]int{
		color.R(),
		color.G(),
		color.B(),
	}
}

var cpuprofile = flag.Bool("cpuprofile", false, "write cpu profile")
var memprofile = flag.Bool("memprofile", true, "write memory profile")
var width = flag.Int("width", 200, "width in pixels")
var height = flag.Int("height", 100, "height in pixels")
var sampling = flag.Int("sampling", 100, "samples per pixels")

func main() {
	flag.Parse()

	if *cpuprofile {
		cf, err := os.Create("cpu.prof")
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(cf); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	// Open file first even though we don't need it yet because if
	// we wait until the end and can't open it .. that's annoying
	f, err := os.Create("out.png")
	check(err)
	defer f.Close()

	nx := *width
	ny := *height
	ns := *sampling

	img := image.NewRGBA(image.Rect(0, 0, nx, ny))

	world := HitableList{randomScene()}

	cam := initializeCamera(nx, ny)

	for j := 0; j < ny; j++ {
		for i := 0; i < nx; i++ {
			pixelColor := computePixelColor(
				i,
				j,
				nx,
				ny,
				ns,
				cam,
				world,
			)

			img.Set(
				i,
				ny-1-j,
				color.RGBA{
					uint8(pixelColor[0]),
					uint8(pixelColor[1]),
					uint8(pixelColor[2]),
					255,
				},
			)
		}
	}

	png.Encode(f, img)

	if *memprofile {
		time.Sleep(500 * time.Millisecond)
		pf, _ := os.Create("mem.prof")
		defer pf.Close()
		runtime.GC()
		pprof.WriteHeapProfile(pf)
	}

	f.Sync()
}
