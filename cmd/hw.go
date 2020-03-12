package main

import (
	"image/color"
	"math/rand"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"github.com/ajstarks/fc"
)

func rn(n int) int {
	return rand.Intn(n)
}

func rn8(n int) uint8 {
	return uint8(rn(n))
}

func main() {

	white := color.RGBA{255, 255, 255, 255}
	black := color.RGBA{0, 0, 0, 255}
	blue := color.RGBA{44, 77, 232, 255}

	width := 792
	height := 612

	w := app.New().NewWindow("hello")
	rect := fc.Rect(width/2, height/2, width, height, black)
	circle := fc.Circle(width/2, height, width, blue)
	text := fc.TextMid(width/2, height/2, "hello, world", width/10, white)
	image := fc.Image(width/2, height/6, 200, 200, "earth.jpg")
	content := fyne.NewContainer(rect, circle, image, text)

	for i := 0; i < 500; i++ {
		x := rn(width)
		y := rn(height)
		color := color.RGBA{rn8(255), rn8(255), rn8(255), rn8(255)}
		content.AddObject(fc.Circle(x, y, rn(40), color))
	}

	// interval := 50
	// for x := interval; x < width; x += interval {
	// 	content.AddObject(fc.Line(x, 0, x, height, 0.5, white))
	// 	content.AddObject(fc.TextMid(x, height-15, fmt.Sprintf("%d", x), 10, white))
	// }
	// for y := interval; y < height; y += interval {
	// 	content.AddObject(fc.Line(0, y, width, y, 0.5, white))
	// 	content.AddObject(fc.TextMid(10, y, fmt.Sprintf("%d", y), 10, white))
	// }

	w.Resize(fyne.NewSize(width, height))
	w.SetFixedSize(true)
	w.SetPadded(false)
	w.SetContent(content)
	w.ShowAndRun()
}
