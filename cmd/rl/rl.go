// mondrian makes Composition II with Red Blue and Yellow by Piet Mondrian
package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/ajstarks/fc"
)

func main() {

	width := 500
	height := 500
	rand.Seed(time.Now().Unix())
	w, content := fc.Start("Random Lines", width, height)
	for i := 0; i < width; i++ {
		r := uint8(rand.Intn(255))
		c := color.RGBA{r, r, r, 255}
		fc.Line(content, i, 0, rand.Intn(width), height, 5, c)
	}
	fc.EndRun(w, content, width, height)
}
