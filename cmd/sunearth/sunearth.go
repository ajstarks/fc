package main

import (
	"image/color"
	"math/rand"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"github.com/ajstarks/fc"
)

func main() {

	w := app.New().NewWindow("sun+earth")

	white := color.RGBA{255, 255, 255, 255}
	black := color.RGBA{0, 0, 0, 255}
	blue := color.RGBA{44, 77, 232, 255}
	yellow := color.RGBA{255, 248, 231, 255}

	width := 500
	height := 500

	earthsize := 6
	sunsize := earthsize * 109

	rect := fc.Rect(width/2, height/2, width, height, black)
	earth := fc.Circle(150, 50, earthsize, blue)
	sun := fc.Circle(width, height, sunsize, yellow)
	content := fyne.NewContainer(rect, sun, earth)
	for i := 0; i < width; i++ {
		x, y := rand.Intn(width), rand.Intn(height)
		content.AddObject(fc.Line(x, y, x, y+1, 0.4, white))
	}
	w.Resize(fyne.NewSize(width, height))
	w.SetFixedSize(true)
	w.SetPadded(false)
	w.SetContent(content)
	w.ShowAndRun()
}
