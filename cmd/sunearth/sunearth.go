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

	earthsize := 6
	sunsize := earthsize * 109

	w, content := fc.Start("Sun+Earth", width, height)
	earth := fc.Circle(150, 50, earthsize, blue)
	sun := fc.Circle(width, height, sunsize, yellow)
	content.AddObject(earth)
	content.AddObject(sun)
	for i := 0; i < width; i++ {
		x, y := rand.Intn(width), rand.Intn(height)
		content.AddObject(fc.Line(x, y, x, y+1, 0.4, white))
	}
	fc.EndRun(w, content, width, height)
}
