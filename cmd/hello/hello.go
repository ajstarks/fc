package main

import (
	"image/color"

	"github.com/ajstarks/fc"
)

func main() {
	width := 500
	height := 500
	blue := color.RGBA{0, 0, 255, 255}
	white := color.RGBA{255, 255, 255, 255}

	canvas := fc.NewCanvas("hello", width, height)

	canvas.Circle(50, 0, 100, blue)
	canvas.CText(50, 25, 10, "hello, world", white)
	canvas.Image(50, 75, 200, 200, "earth.jpg")

	canvas.EndRun()
}
