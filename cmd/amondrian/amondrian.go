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

	var border float32 = 6.0
	w3 := width / 3
	w6 := w3 / 2
	w23 := w3 * 2
	w36 := w3 - w6
	ww6 := width - w6
	hw6 := height - w6
	hw3 := height - w3
	b2 := border * 2

	w, canvas := fc.AbsStart("Mondrian", width, height)
	fc.AbsCornerRect(canvas, 0, 0, w3, w3, white)       // upper left white square
	fc.AbsCornerRect(canvas, 0, w3, w3, w3, white)      // middle left white square
	fc.AbsCornerRect(canvas, 0, w23, w3, w3, blue)      // lower left blue square
	fc.AbsCornerRect(canvas, w3, 0, w23, w23, red)      // large red square
	fc.AbsCornerRect(canvas, w3, w23, w23, w3, white)   // lower-middle white rectangle
	fc.AbsCornerRect(canvas, ww6, hw3, w36, w6, white)  // lower right white square
	fc.AbsCornerRect(canvas, ww6, hw6, w36, w6, yellow) // lower right yellow square

	fc.AbsLine(canvas, 0, 0, 0, height, b2, black)          // left border
	fc.AbsLine(canvas, width, 0, width, height, b2, black)  // right border
	fc.AbsLine(canvas, 0, 0, width, 0, b2, black)           // top border
	fc.AbsLine(canvas, 0, height, width, height, b2, black) // botom border

	fc.AbsLine(canvas, 0, w3, w3, w3, border, black)         // bottom of upper left white square
	fc.AbsLine(canvas, w3, 0, w3, height, border, black)     // right border for left-hand squares
	fc.AbsLine(canvas, 0, w23, width, w23, border, black)    // two-thirds border
	fc.AbsLine(canvas, ww6, hw3, ww6, height, border, black) // left border for small squares
	fc.AbsLine(canvas, ww6, hw6, width, hw6, border, black)  // top/bottom of small squares

	fc.AbsEndRun(w, canvas, width, height)
}
