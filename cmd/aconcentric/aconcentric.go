// concentric gray scale circles
package main

import (
	"image/color"

	"github.com/ajstarks/fc"
)

func main() {
	width, height := 500, 500
	w, canvas := fc.AbsStart("Concentric", width, height)
	r := height / 2
	for g := uint8(0); g < 250; g += 50 {
		fc.AbsCircle(canvas, width/2, height/2, r, color.RGBA{g, g, g, 255})
		r -= 50
	}
	fc.AbsEndRun(w, canvas, width, height)
}
