package main

import (
	"flag"
	"fmt"
	"image/color"
	"os"

	"github.com/ajstarks/fc"
	"github.com/ajstarks/fc/chart"
)

func composite(canvas fc.Canvas) error {
	sr, err := os.Open("sin.d")
	if err != nil {
		return err
	}
	cr, err := os.Open("cos.d")
	if err != nil {
		return err
	}
	sine, err := chart.DataRead(sr)
	if err != nil {
		return err
	}
	cosine, err := chart.DataRead(cr)
	if err != nil {
		return err
	}
	sr.Close()
	cr.Close()

	sine.Zerobased, cosine.Zerobased = false, false

	cosine.Frame(canvas, 5)
	cosine.Label(canvas, 1.5, 10)
	cosine.YAxis(canvas, 1.2, -1.0, 1.0, 1.0, "%0.2f", true)
	cosine.Color = color.RGBA{0, 128, 0, 255}
	sine.Color = color.RGBA{128, 0, 0, 255}
	cosine.Scatter(canvas, 1)
	sine.Scatter(canvas, 1)
	sine.Bar(canvas, 0.2)
	

	sine.Left = 10
	sine.Right = sine.Left + 40
	sine.Top, cosine.Top = 30, 30
	sine.Bottom, cosine.Bottom = 10, 10

	sine.CTitle(canvas, 2, 2)
	sine.Frame(canvas, 5)
	sine.Bar(canvas, 0.1)

	offset := 45.0
	cosine.Left = sine.Left + offset
	cosine.Right = sine.Right + offset

	cosine.CTitle(canvas, 2, 2)
	cosine.Frame(canvas, 5)
	cosine.Bar(canvas, 0.1)

	return nil
}

func main() {
	var width, height int
	flag.IntVar(&width, "w", 500, "canvas width")
	flag.IntVar(&height, "h", 500, "canvas height")
	flag.Parse()

	canvas := fc.NewCanvas("Chart", width, height)
	canvas.Rect(50, 50, 100, 100, color.RGBA{255, 255, 255, 255})
	if err := composite(canvas); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	canvas.EndRun()
}
