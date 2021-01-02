package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/ajstarks/fc"
)

func main() {
	var imagefile string
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "specify an image file")
		os.Exit(1)
	}
	imagefile = os.Args[1]
	f, ferr := os.Open(imagefile)
	if ferr != nil {
		fmt.Fprintf(os.Stderr, "%v\n", ferr)
		os.Exit(2)
	}
	im, _, err := image.DecodeConfig(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", imagefile, err)
		os.Exit(3)
	}

	canvas := fc.NewCanvas(imagefile, im.Width, im.Height)
	canvas.Image(50, 50, im.Width, im.Height, imagefile)
	f.Close()
	canvas.EndRun()
}
