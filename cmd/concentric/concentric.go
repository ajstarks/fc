// concentric gray scale circles
package main

import (
	"image/color"

	"github.com/ajstarks/fc"
)

func main() {
	width, height := 500, 500
	canvas := fc.NewCanvas("Concentric", width, height)
	r := 60.0
	g := uint8(5)
	for i := 0; i < 6; i++ {
		canvas.Circle(50, 50, r, color.RGBA{128, g, g, 255})
		r -= 10
		g += 25
	}
	canvas.EndRun()
}
