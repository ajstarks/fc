package main

import (
	"flag"
	"fmt"
	"image/color"
	"os"

	"github.com/ajstarks/fc"
	"github.com/ajstarks/fc/chart"
)

func main() {
	var width, height, xlabel int
	var barwidth float64
	var yax string
	var zb bool
	flag.IntVar(&width, "w", 500, "canvas width")
	flag.IntVar(&height, "h", 500, "canvas height")
	flag.IntVar(&xlabel, "xlabel", 1, "x-xaxis label")
	flag.Float64Var(&barwidth, "barwidth", 0.5, "bar width")
	flag.StringVar(&yax, "yrange", "", "y axis range (min,max,step")
	flag.BoolVar(&zb, "zero", true, "zero minumum")
	flag.Parse()

	canvas := fc.NewCanvas("Chart", width, height)
	canvas.Rect(50, 50, 100, 100, color.RGBA{255, 255, 255, 255})
	chart, err := chart.DataRead(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	chart.Top = 80
	chart.Bottom = 40
	chart.Zerobased = zb
	
	if len(yax) > 0 {
		var yaxmin, yaxmax, yaxstep float64
		if n, err := fmt.Sscanf(yax, "%v,%v,%v", &yaxmin, &yaxmax, &yaxstep); n == 3 && err == nil {
			chart.YAxis(canvas, 1.5, yaxmin, yaxmax, yaxstep, "%v", true)
		}
	}
	chart.CTitle(canvas, 3, 2)
	chart.Label(canvas, 1.5, xlabel)
	chart.Color = color.RGBA{176,196,222,255}
	chart.Bar(canvas, barwidth)
	canvas.EndRun()
}
