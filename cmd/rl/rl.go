// mondrian makes Composition II with Red Blue and Yellow by Piet Mondrian
package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/ajstarks/fc"
)

func rn(n int) float64 {
	return float64(rand.Intn(n))
}

func main() {

	width := 500
	height := 500
	rand.Seed(time.Now().Unix())
	canvas := fc.NewCanvas("Random Lines", width, height)
	for x := 0.0; x < 100; x++ {
		r := uint8(rand.Intn(255))
		c := color.RGBA{r, r, r, 100}
		canvas.Line(x, 0, rn(100), 100, 1, c)
	}
	canvas.EndRun()
}
