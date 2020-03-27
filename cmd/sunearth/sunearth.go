package main

import (
	"image/color"
	"math/rand"

	"github.com/ajstarks/fc"
)

func rn(n int) float64 {
	return float64(rand.Intn(n))
}

func main() {

	width := 500
	height := 500

	white := color.RGBA{255, 255, 255, 255}
	blue := color.RGBA{44, 77, 232, 255}
	yellow := color.RGBA{255, 248, 231, 255}

	earthsize := 1.6
	sunsize := earthsize * 109

	canvas := fc.NewCanvas("sun+earth", width, height)
	canvas.Circle(30, 90, earthsize, blue)
	canvas.Circle(100, 0, sunsize, yellow)

	for i := 0; i < width; i++ {
		x, y := rn(100), rn(100)
		canvas.Line(x, y, x+0.5, y, 0.05, white)
	}
	canvas.EndRun()
}
