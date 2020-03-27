// eclipse illustrates the eclipse
package main

import (
	"image/color"

	"github.com/ajstarks/fc"
)

func main() {
	width, height := 500, 500
	black := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}

	canvas := fc.NewCanvas("Eclipse", width, height)
	r := 10.0
	y := 50.0
	for x := 10.0; x < 100.0; x += 15 {
		canvas.Circle(x, 50, r+1, white)
		canvas.Circle(x, y, r, black)
		y -= 2
	}
	canvas.EndRun()
}
