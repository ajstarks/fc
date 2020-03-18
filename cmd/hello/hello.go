package main

import (
	"fmt"
	"image/color"

	"github.com/ajstarks/fc"
)

func main() {

	white := color.RGBA{255, 255, 255, 255}
	black := color.RGBA{0, 0, 0, 255}
	blue := color.RGBA{44, 77, 232, 255}

	width := 500
	height := 500

	w, content := fc.Start("hello", width, height)
	content.AddObject(fc.Rect(width/2, height/2, width, height, black))
	content.AddObject(fc.Circle(width/2, height, width, blue))
	content.AddObject(fc.TextMid(width/2, height/2, "hello, world", width/10, white))
	content.AddObject(fc.Image(width/2, height/5, 200, 200, "earth.jpg"))

	interval := 50
	for x := interval; x < width; x += interval {
		content.AddObject(fc.Line(x, 0, x, height, 0.5, white))
		content.AddObject(fc.TextMid(x, height-15, fmt.Sprintf("%d", x), 10, white))
	}
	for y := interval; y < height; y += interval {
		content.AddObject(fc.Line(0, y, width, y, 0.5, white))
		content.AddObject(fc.TextMid(10, y, fmt.Sprintf("%d", y), 10, white))
	}
	fc.EndRun(w, content, width, height)
}
