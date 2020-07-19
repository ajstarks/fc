// fchart -- command line chart using fc chart packages
package main

import (
	"flag"
	"fmt"
	"image/color"
	"io"
	"os"

	"github.com/ajstarks/fc"
	"github.com/ajstarks/fc/chart"
)

func main() {

	// Command line options
	var width, height, xlabel int
	var barwidth, linewidth, linespacing, dotsize, textsize, ty, frameOp, top, bottom, left, right float64
	var chartitle, yrange, yaxfmt, dcolor string
	var zb, line, bar, hbar, scatter, lego, showtitle, showgrid bool
	flag.IntVar(&width, "w", 600, "canvas width")
	flag.IntVar(&height, "h", 600, "canvas height")
	flag.IntVar(&xlabel, "xlabel", 1, "x-xaxis label")
	flag.Float64Var(&barwidth, "barwidth", 0.5, "bar width")
	flag.Float64Var(&linewidth, "linewidth", 0.25, "bar width")
	flag.Float64Var(&linespacing, "ls", barwidth*4, "bar width")
	flag.Float64Var(&dotsize, "dotsize", 0.5, "bar width")
	flag.Float64Var(&textsize, "textsize", 1.5, "bar width")
	flag.Float64Var(&top, "top", 80, "bar width")
	flag.Float64Var(&bottom, "bottom", 40, "bar width")
	flag.Float64Var(&left, "left", 10, "bar width")
	flag.Float64Var(&right, "right", 90, "bar width")
	flag.Float64Var(&ty, "ty", 5, "title position relative to the top")
	flag.Float64Var(&frameOp, "frame", 0, "frame opacity")
	flag.StringVar(&yrange, "yrange", "", "y axis range (min,max,step")
	flag.StringVar(&chartitle, "chartitle", "", "chart title")
	flag.StringVar(&yaxfmt, "yfmt", "%v", "yaxis format")
	flag.StringVar(&dcolor, "color", "steelblue", "color")
	flag.BoolVar(&showtitle, "title", true, "show the title")
	flag.BoolVar(&showgrid, "grid", false, "show y axis grid")
	flag.BoolVar(&zb, "zero", true, "zero minumum")
	flag.BoolVar(&bar, "bar", false, "bar chart")
	flag.BoolVar(&line, "line", false, "line chart")
	flag.BoolVar(&hbar, "hbar", false, "horizontal bar")
	flag.BoolVar(&scatter, "scatter", false, "scatter chart")
	flag.BoolVar(&lego, "lego", false, "lego chart")
	flag.Parse()

	var input io.Reader
	var ferr error

	// Read from stdin or specified file
	if len(flag.Args()) == 0 {
		input = os.Stdin
	} else {
		input, ferr = os.Open(flag.Args()[0])
		if ferr != nil {
			fmt.Fprintf(os.Stderr, "%v\n", ferr)
			os.Exit(1)
		}
	}

	// Read in the data
	chart, err := chart.DataRead(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(2)
	}

	// Define the canvas
	canvas := fc.NewCanvas(fmt.Sprintf("Chart: %s", chart.Title), width, height)

	// Define the colors
	datacolor := fc.ColorLookup(dcolor)
	labelcolor := color.RGBA{100, 100, 100, 255}
	bgcolor := color.RGBA{255, 255, 255, 255}
	canvas.Rect(50, 50, 100, 100, bgcolor)

	// Set the chart attributes
	chart.Zerobased = zb
	chart.Top, chart.Bottom, chart.Left, chart.Right = top, bottom, left, right

	// Draw the data
	chart.Color = datacolor
	if frameOp > 0 {
		chart.Frame(canvas, frameOp)
	}
	if line {
		chart.Line(canvas, linewidth)
	}
	if bar {
		chart.Bar(canvas, barwidth)
	}
	if scatter {
		chart.Scatter(canvas, dotsize)
	}
	if hbar {
		chart.HBar(canvas, barwidth, linespacing, textsize)
	}
	if lego {
		chart.Lego(canvas, dotsize)
	}

	// Draw labels, axes if specified
	chart.Color = labelcolor
	if line || bar || scatter {
		chart.Label(canvas, textsize, xlabel)
		if len(yrange) > 0 {
			var yaxmin, yaxmax, yaxstep float64
			if n, err := fmt.Sscanf(yrange, "%v,%v,%v", &yaxmin, &yaxmax, &yaxstep); n == 3 && err == nil {
				chart.YAxis(canvas, textsize, yaxmin, yaxmax, yaxstep, yaxfmt, showgrid)
			}
		}
	}

	// Draw the chart titles
	if len(chartitle) > 0 {
		chart.Title = chartitle
	}
	if showtitle && len(chart.Title) > 0 {
		chart.CTitle(canvas, textsize*2, ty)
	}

	// Show the chart
	canvas.EndRun()
}
