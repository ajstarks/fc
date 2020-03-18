// mondrian makes Composition II with Red Blue and Yellow by Piet Mondrian
package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"github.com/ajstarks/fc"
)

func main() {

	width := 500
	height := 500
	w := app.New().NewWindow("Mondrian")

	black := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}
	blue := color.RGBA{0, 0, 255, 255}
	red := color.RGBA{255, 0, 0, 255}
	yellow := color.RGBA{255, 255, 0, 255}

	var border float32 = 4.0
	w3 := width / 3
	w6 := w3 / 2
	w23 := w3 * 2
	w36 := w3 - w6
	ww6 := width - w6
	hw6 := height - w6
	hw3 := height - w3
	b2 := border * 2
	r1 := fc.CornerRect(0, 0, w3, w3, white)       // upper left white square
	r2 := fc.CornerRect(0, w3, w3, w3, white)      // middle left white square
	r3 := fc.CornerRect(0, w23, w3, w3, blue)      // lower left blue square
	r4 := fc.CornerRect(w3, 0, w23, w23, red)      // large red square
	r5 := fc.CornerRect(w3, w23, w23, w3, white)   // lower-middle white rectangle
	r6 := fc.CornerRect(ww6, hw3, w36, w6, white)  // lower right white square
	r7 := fc.CornerRect(ww6, hw6, w36, w6, yellow) // lower right yellow square

	ll := fc.Line(0, 0, 0, height, b2, black)          // left border
	lr := fc.Line(width, 0, width, height, b2, black)  // right border
	lt := fc.Line(0, 0, width, 0, b2, black)           // top border
	lb := fc.Line(0, height, width, height, b2, black) // botom border

	l1 := fc.Line(0, w3, w3, w3, border, black)         // bottom of upper left white square
	l2 := fc.Line(w3, 0, w3, height, border, black)     // right border for left-hand squares
	l3 := fc.Line(0, w23, width, w23, border, black)    // two-thirds border
	l4 := fc.Line(ww6, hw3, ww6, height, border, black) // left border for small squares
	l5 := fc.Line(ww6, hw6, width, hw6, border, black)  // top/bottom of small squares
	content := fyne.NewContainer(r1, r2, r3, r4, r5, r6, r7, l1, l2, l3, l4, l5, ll, lr, lt, lb)

	w.Resize(fyne.NewSize(width, height))
	w.SetFixedSize(true)
	w.SetPadded(false)
	w.SetContent(content)
	w.ShowAndRun()
}
