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

	// Command line options
	var width, height, xlabel int
	var barwidth, linewidth, linespacing, dotsize, textsize, frameOp, top, bottom, left, right float64
	var chartitle, yrange string
	var zb, line, bar, hbar, scatter, showtitle, showgrid bool
	flag.IntVar(&width, "w", 800, "canvas width")
	flag.IntVar(&height, "h", 800, "canvas height")
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
	flag.Float64Var(&frameOp, "frame", 0, "frame opacity")
	flag.StringVar(&yrange, "yrange", "", "y axis range (min,max,step")
	flag.StringVar(&chartitle, "chartitle", "", "chart title")
	flag.BoolVar(&showtitle, "title", true, "show the title")
	flag.BoolVar(&showgrid, "grid", false, "show y axis grid")
	flag.BoolVar(&zb, "zero", true, "zero minumum")
	flag.BoolVar(&bar, "bar", false, "bar chart")
	flag.BoolVar(&line, "line", false, "line chart")
	flag.BoolVar(&hbar, "hbar", false, "horizontal bar")
	flag.BoolVar(&scatter, "scatter", false, "scatter chart")
	flag.Parse()

	// Read in the data
	canvas := fc.NewCanvas("Chart", width, height)
	chart, err := chart.DataRead(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	// Define the colors
	datacolor := color.RGBA{176, 196, 222, 255}
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

	// Draw labels, axes if specified
	chart.Color = labelcolor
	if line || bar || scatter {
		chart.Label(canvas, textsize, xlabel)
		if len(yrange) > 0 {
			var yaxmin, yaxmax, yaxstep float64
			if n, err := fmt.Sscanf(yrange, "%v,%v,%v", &yaxmin, &yaxmax, &yaxstep); n == 3 && err == nil {
				chart.YAxis(canvas, textsize, yaxmin, yaxmax, yaxstep, "%v", showgrid)
			}
		}
	}

	// Draw the chart titles
	if len(chartitle) > 0 {
		chart.Title = chartitle
	}
	if showtitle && len(chart.Title) > 0 {
		chart.CTitle(canvas, textsize*2, 2)
	}

	// Show the chart
	canvas.EndRun()
}
