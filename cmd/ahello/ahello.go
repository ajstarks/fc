package main

import (
	"image/color"

	"github.com/ajstarks/fc"
)

func main() {
	width := 500
	height := 500

	white := color.RGBA{255, 255, 255, 255}
	blue := color.RGBA{44, 77, 232, 255}
	midx := width / 2
	iy := height / 5
	ty := 3 * height / 4

	w, canvas := fc.AbsStart("hello", width, height)
	fc.AbsCircle(canvas, midx, height, midx, blue)
	fc.AbsTextMid(canvas, midx, ty, "hello, world", width/10, white)
	fc.AbsImage(canvas, midx, iy, 200, 200, "earth.jpg")
	fc.AbsEndRun(w, canvas, width, height)
}
