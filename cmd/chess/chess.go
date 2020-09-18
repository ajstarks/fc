package main

import (
	"image/color"

	"github.com/ajstarks/fc"
)

// row makes n squares spaced 2 times their width, at (x,y)
func row(canvas fc.Canvas, n int, x, y, size float64, fillcolor color.RGBA) {
	for i := 0; i < n; i++ {
		canvas.CornerRect(x, y, size, size, fillcolor)
		x += size * 2
	}
}

func main() {
	width, height := 500, 500
	black := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}
	gray := color.RGBA{128, 128, 128, 255}

	canvas := fc.NewCanvas("chess", width, height)
	canvas.Rect(50, 50, 100, 100, gray) // canvas background

	left := 10.0
	top := 90.0

	x := left
	y := top
	size := 10.0
	nrows := 8
	bgsize := size * float64(nrows)
	canvas.CornerRect(left, top, bgsize, bgsize, white) // board background
	for i := 0; i < nrows; i++ {                        // for every row make alternating squares
		if i%2 == 0 {
			x = left + size
		} else {
			x = left
		}
		row(canvas, nrows/2, x, y, size, black) // make the row
		y -= size                               // move down
	}
	canvas.EndRun()
}
