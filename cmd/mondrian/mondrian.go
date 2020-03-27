// mondrian makes Composition II with Red Blue and Yellow by Piet Mondrian
package main

import (
	"image/color"

	"github.com/ajstarks/fc"
)

func main() {

	width := 500
	height := 500

	black := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}
	blue := color.RGBA{0, 0, 255, 255}
	red := color.RGBA{255, 0, 0, 255}
	yellow := color.RGBA{255, 255, 0, 255}

	third := 100.0 / 3

	halft := third / 2
	qt := third / 4
	t2 := third * 2
	tq := 100.0 - qt
	t2h := t2 + halft

	border := 1.0
	b2 := border * 2

	canvas := fc.NewCanvas("Mondrian", width, height)
	canvas.CornerRect(0, 100, 100, 100, white)               // white background
	canvas.Rect(halft, halft, third, third, blue)            // lower left blue square
	canvas.Rect(t2, t2, t2, t2, red)                         // big red
	canvas.Rect(tq, qt, halft, halft, yellow)                // small yellow lower right
	canvas.Line(0, 0, 0, 100, b2, black)                     // left border
	canvas.Line(100, 0, 100, 100, b2, black)                 // right border
	canvas.Line(0, 0, 100, 0, b2, black)                     // top border
	canvas.Line(0, 100, 100, 100, b2, black)                 // bottom border
	canvas.Line(t2h, halft, t2h+halft, halft, border, black) // top of yellow square
	canvas.Line(third, 100, third, 0, border, black)         //  first column border
	canvas.Line(t2h, 0, t2h, third, border, black)           // left of small right squares
	canvas.Line(0, third, 100, third, border, black)         // top of bottom squares
	canvas.Line(0, t2, third, t2, border, black)             // border between left white squares

	canvas.EndRun()
}
