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

	c := fc.NewCanvas("hello", width, height)

	c.Circle(50, 0, 100, blue)
	c.CText(50, 25, 10, "hello, world", white)
	c.Image(50, 75, 100, 100, "earth.jpg")

	c.EndRun()
}
