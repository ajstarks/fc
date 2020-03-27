// confetti -- random shapes
package main

import (
	"flag"
	"image/color"
	"math/rand"

	"github.com/ajstarks/fc"
)

func rn(n int) float64 {
	return float64(rand.Intn(n))

}

func rn8(n int) uint8 {
	return uint8(rand.Intn(n))
}

func main() {
	var nshapes, width, height int
	var maxsize int
	flag.IntVar(&width, "w", 500, "width")
	flag.IntVar(&height, "h", 500, "height")
	flag.IntVar(&nshapes, "n", 500, "number of shapes")
	flag.IntVar(&maxsize, "size", 10, "max size")

	flag.Parse()

	canvas := fc.NewCanvas("Confetti", width, height)
	canvas.Circle(50, 50, 5, color.RGBA{127, 0, 0, 255})
	for i := 0; i < nshapes; i++ {
		x := rn(100)
		y := rn(100)
		color := color.RGBA{rn8(255), rn8(255), rn8(255), rn8(255)}
		if i%2 == 0 {
			canvas.Rect(x, y, rn(maxsize), rn(maxsize), color)
		} else {
			canvas.Circle(x, y, rn(maxsize), color)
		}
	}
	canvas.EndRun()
}
