package main

import (
	"flag"
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
	var imw, imh int

	flag.IntVar(&imw, "w", 0, "image width")
	flag.IntVar(&imh, "h", 0, "image height")
	flag.Parse()

	if len(flag.Args()) != 1 {
		fmt.Fprintln(os.Stderr, "specify an image file")
		os.Exit(1)
	}
	imagefile = flag.Args()[0]
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
	if imw == 0 {
		imw = im.Width
	}
	if imh == 0 {
		imh = im.Height
	}

	canvas := fc.NewCanvas(imagefile, imw, imh)
	canvas.Image(50, 50, imw, imh, imagefile)
	f.Close()
	canvas.EndRun()
}
