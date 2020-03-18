// eclipse illustrates the eclipse
package main

import (
	"image/color"
	"github.com/ajstarks/fc"
)

func main() {
	width, height := 500, 500
	h2 := height/2
	r := width / 10
	black := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}

	w, content := fc.Start("Eclipse", width, height)
	for x, y := 50, h2; x < 450; x += 75 {
		content.AddObject(fc.Circle(x, h2, r+2, white))
		content.AddObject(fc.Circle(x, y, r, black))
		y += 10
	}
	fc.EndRun(w, content, width, height)
}
