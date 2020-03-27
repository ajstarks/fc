package main

import (
	"image/color"
	"math/rand"

	"fyne.io/fyne"
	"github.com/ajstarks/fc"
)

func cloud(canvas *fyne.Container, x, y, r int, color color.RGBA) {
	small := r / 2
	medium := (r * 6) / 10
	fc.AbsCircle(canvas, x, y, r, color)
	fc.AbsCircle(canvas, x+r, y+small, small, color)
	fc.AbsCircle(canvas, x-r-small, y+small, small, color)
	fc.AbsCircle(canvas, x-r, y, medium, color)
	fc.AbsCornerRect(canvas, x-r-small, y, r*2+small, r, color)
}

func main() {
	width, height := 500, 500
	white := color.RGBA{255, 255, 255, 255}
	w, canvas := fc.AbsStart("clouds", width, height)
	fc.AbsCornerRect(canvas, 0, 0, width, height, white)
	for i := 0; i < 100; i++ {
		x, y, size := rand.Intn(width), rand.Intn(height), rand.Intn(50)
		r, g, b := uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255))
		cloud(canvas, x, y, size, color.RGBA{r, g, b, 255})
	}
	fc.AbsEndRun(w, canvas, width, height)
}
