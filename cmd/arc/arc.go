package main

import (
	"image/color"
	"math/rand"

	"github.com/ajstarks/fc"
)

const fullcircle = 6.28318530717958647692528676655900 // 2 * Pi

func randcolor() color.RGBA {
	var c color.RGBA
	c.A, c.R, c.G, c.B = 255, byte(rand.Intn(255)), byte(rand.Intn(255)), byte(rand.Intn(255))
	return c
}

func main() {
	width, height := 1000, 1000
	canvas := fc.NewCanvas("arc", width, height)
	canvas.Rect(50, 50, 100, 100, color.RGBA{255, 255, 255, 255})
	for i := 0; i < 100; i++ {
		x := rand.Float64() * 100
		y := rand.Float64() * 100
		r := rand.Float64() * 20
		a1 := 0.0
		a2 := fullcircle * rand.Float64()
		acolor := randcolor()
		canvas.ArcLine(x, y, r, a1, a2, 0.2, acolor)
	}
	canvas.EndRun()
}
