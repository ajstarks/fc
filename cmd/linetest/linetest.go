package main

import (
	"fmt"
	"image/color"

	"github.com/ajstarks/fc"
)

func main() {
	width, height := 1000, 1000
	blue := color.RGBA{0, 0, 255, 255}
	black := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}

	canvas := fc.NewCanvas("linetest", width, height)

	x := 10.0
	canvas.CornerRect(0, 100, 100, 100, white)
	for stroke := 0.1; stroke <= 1.8; stroke += 0.1 {
		canvas.Line(x, 10, x, 90, stroke, blue)
		canvas.CText(x, 5, 1.5, fmt.Sprintf("%.2f", stroke), black)
		x += 5
	}

	canvas.EndRun()
}
