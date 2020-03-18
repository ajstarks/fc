// eclipse illustrates the eclipse
package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"github.com/ajstarks/fc"
)

func main() {
	w := app.New().NewWindow("Eclipse")

	width, height := 500, 500
	w2, h2 := width/2, height/2
	r := width / 10
	black := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}
	content := fyne.NewContainer(fc.Rect(w2, h2, width, height, black))

	for x, y := 50, h2; x < 450; x += 75 {
		content.AddObject(fc.Circle(x, h2, r+2, white))
		content.AddObject(fc.Circle(x, y, r, black))
		y += 10
	}
	w.Resize(fyne.NewSize(width, height))
	w.SetFixedSize(true)
	w.SetPadded(false)
	w.SetContent(content)
	w.ShowAndRun()
}
