// eclipse illustrates the eclipse
package main

import (
	"image/color"

	"github.com/ajstarks/fc"
)

func main() {
	width, height := 500, 500
	h2 := height / 2
	r := width / 20
	black := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}

	w, canvas := fc.Start("Eclipse", width, height)
	for x, y := 50, h2; x < 450; x += 75 {
		fc.Circle(canvas, x, h2, r+2, white)
		fc.Circle(canvas, x, y, r, black)
		y += 10
	}
	fc.EndRun(w, canvas, width, height)
}
