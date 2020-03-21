package main

import (
	"image/color"
	"math/rand"

	"github.com/ajstarks/fc"
)

func main() {

	width := 500
	height := 500

	white := color.RGBA{255, 255, 255, 255}
	blue := color.RGBA{44, 77, 232, 255}
	yellow := color.RGBA{255, 248, 231, 255}

	earthsize := 4
	sunsize := earthsize * 109

	w, canvas := fc.Start("Sun+Earth", width, height)
	fc.Circle(canvas, 150, 50, earthsize, blue)
	fc.Circle(canvas, width, height, sunsize, yellow)

	for i := 0; i < width; i++ {
		x, y := rand.Intn(width), rand.Intn(height)
		fc.Line(canvas, x, y, x, y+1, 0.4, white)
	}
	fc.EndRun(w, canvas, width, height)
}
