// confetti -- random shapes
package main

import (
	"flag"
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"fyne.io/fyne"
	"github.com/ajstarks/fc"
)

func rn(n int) int {
	return rand.Intn(n)
}

func rn8(n int) uint8 {
	return uint8(rn(n))
}

func main() {
	var nshapes, width, height, maxsize int
	var timing bool
	flag.IntVar(&nshapes, "n", 500, "number of shapes")
	flag.IntVar(&width, "w", 500, "canvas width")
	flag.IntVar(&height, "h", 500, "canvas height")
	flag.IntVar(&maxsize, "size", width/10, "max size")
	flag.BoolVar(&timing, "t", false, "timing")

	flag.Parse()

	var bt time.Time
	if timing {
		bt = time.Now()
	}

	w, content := fc.Start("Confetti", width, height)

	for i := 0; i < nshapes; i++ {
		x := rn(width)
		y := rn(height)
		color := color.RGBA{rn8(255), rn8(255), rn8(255), rn8(255)}
		if i%2 == 0 {
			w := rn(maxsize)
			h := rn(maxsize)
			content.AddObject(fc.Rect(x, y, w, h, color))
		} else {
			content.AddObject(fc.Circle(x, y, rn(maxsize), color))
		}
	}

	w.Resize(fyne.NewSize(width, height))
	w.SetFixedSize(true)
	w.SetPadded(false)
	w.SetContent(content)
	if timing {
		fmt.Printf("rendering time=%v\n", time.Since(bt))
	}
	w.ShowAndRun()
}
