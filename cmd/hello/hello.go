package main

import (
	"image/color"

	"github.com/ajstarks/fc"
)

func main() {
	width := 500
	height := 500

	white := color.RGBA{255, 255, 255, 255}
	black := color.RGBA{0, 0, 0, 255}
	blue := color.RGBA{44, 77, 232, 255}

	w, canvas := fc.Start("hello", width, height)

	fc.Rect(canvas, width/2, height/2, width, height, black)
	fc.Circle(canvas, width/2, height, width/2, blue)
	fc.TextMid(canvas, width/2, height/2, "hello, world", width/10, white)
	fc.Image(canvas, width/2, height/5, 200, 200, "earth.jpg")

	fc.EndRun(w, canvas, width, height)
}
