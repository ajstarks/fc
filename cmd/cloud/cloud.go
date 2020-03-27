package main

import (
	"image/color"
	"math/rand"

	"github.com/ajstarks/fc"
)

func rn(n int) float64 {
	return float64(rand.Intn(n))
}

func cloud(canvas fc.Canvas, x, y, r float64, color color.RGBA) {
	w := r
	h := w / 2
	hw := w / 2
	qw := w / 4
	qh := h / 4
	small := r * 0.5
	big := r * 0.75
	med := r * 0.60

	canvas.Circle(x-hw, y, small, color)
	canvas.Circle(x+hw, y, small, color)
	canvas.Circle(x+qw, y+qh, big, color)
	canvas.Circle(x-qw, y+qh, med, color)
	canvas.Rect(x, y, w, h, color)
}

func main() {
	width, height := 500, 500
	white := color.RGBA{255, 255, 255, 255}
	canvas := fc.NewCanvas("clouds", width, height)
	canvas.Rect(50, 50, 100, 100, white)
	for i := 0; i < 100; i++ {
		x, y, size := rn(100), rn(100), rn(20)
		r, g, b := uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255))
		cloud(canvas, x, y, size, color.RGBA{r, g, b, 255})
	}
	canvas.EndRun()
}
