package main

import (
	"bufio"
	"flag"
	"fmt"
	"image/color"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/ajstarks/fc"
)

// NameValue is a name,value pair
type NameValue struct {
	label string
	note  string
	value float64
}

// ChartBox holds the essential data for making a chart
type ChartBox struct {
	title                    string
	data                     []NameValue
	color                    color.RGBA
	top, bottom, left, right float64
	minvalue, maxvalue       float64
	zerobased                bool
}

const (
	largest  = 1.797693134862315708145274237317043567981e+308
	smallest = -largest
)

var labelcolor = color.RGBA{100, 100, 100, 255}

// DataRead reads tab separated values into a ChartBox
func DataRead(r io.Reader) (ChartBox, error) {
	var d NameValue
	var data []NameValue
	var err error
	maxval := smallest
	minval := largest
	title := ""
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		t := scanner.Text()
		if len(t) == 0 { // skip blank lines
			continue
		}
		if t[0] == '#' && len(t) > 2 { // process titles
			title = strings.TrimSpace(t[1:])
			continue
		}
		fields := strings.Split(t, "\t")
		if len(fields) < 2 {
			continue
		}
		if len(fields) == 3 {
			d.note = fields[2]
		} else {
			d.note = ""
		}
		d.label = fields[0]
		d.value, err = strconv.ParseFloat(fields[1], 64)
		if err != nil {
			d.value = 0
		}
		if d.value > maxval {
			maxval = d.value
		}
		if d.value < minval {
			minval = d.value
		}
		data = append(data, d)
	}
	err = scanner.Err()
	return ChartBox{
		title:     title,
		data:      data,
		minvalue:  minval,
		maxvalue:  maxval,
		zerobased: true,
		color:     color.RGBA{0, 0, 0, 255},
		left:      10,
		right:     90,
		top:       90,
		bottom:    50,
	}, err
}

func zerobase(usez bool, n float64) float64 {
	if usez {
		return 0
	}
	return n
}

// Bar makes a (column) bar chart
func (c *ChartBox) Bar(canvas fc.Canvas, size float64) {
	dlen := float64(len(c.data) - 1)
	ymin := zerobase(c.zerobased, c.minvalue)
	for i, d := range c.data {
		x := fc.MapRange(float64(i), 0, dlen, c.left, c.right)
		y := fc.MapRange(d.value, ymin, c.maxvalue, c.bottom, c.top)
		canvas.Line(x, c.bottom, x, y, size, c.color)
	}
}

// HBar makes a horizontal bar chart
func (c *ChartBox) HBar(canvas fc.Canvas, size, linespacing, textsize float64) {
	y := c.top
	xmin := zerobase(c.zerobased, c.minvalue)
	for _, d := range c.data {
		canvas.EText(c.left-2, y-size/2, textsize, d.label, labelcolor)
		x2 := fc.MapRange(d.value, xmin, c.maxvalue, c.left, c.right)
		canvas.Line(c.left, y, x2, y, size, c.color)
		y -= linespacing
	}
}

// Line makes a line chart
func (c *ChartBox) Line(canvas fc.Canvas, size float64) {
	n := len(c.data)
	fn := float64(n - 1)
	ymin := zerobase(c.zerobased, c.minvalue)
	for i := 0; i < n-1; i++ {
		v1 := c.data[i].value
		v2 := c.data[i+1].value
		x1 := fc.MapRange(float64(i), 0, fn, c.left, c.right)
		y1 := fc.MapRange(v1, ymin, c.maxvalue, c.bottom, c.top)
		x2 := fc.MapRange(float64(i+1), 0, fn, c.left, c.right)
		y2 := fc.MapRange(v2, ymin, c.maxvalue, c.bottom, c.top)
		canvas.Line(x1, y1, x2, y2, size, c.color)
	}
}

// Label draws the x axis labels
func (c *ChartBox) Label(canvas fc.Canvas, size float64, n int) {
	fn := float64(len(c.data) - 1)
	for i, d := range c.data {
		x := fc.MapRange(float64(i), 0, fn, c.left, c.right)
		if i%n == 0 {
			canvas.CText(x, c.bottom-(size*2), size, d.label, c.color)
		}
	}
}

// Scatter makes a scatter chart
func (c *ChartBox) Scatter(canvas fc.Canvas, size float64) {
	dlen := float64(len(c.data) - 1)
	ymin := zerobase(c.zerobased, c.minvalue)
	for i, d := range c.data {
		x := fc.MapRange(float64(i), 0, dlen, c.left, c.right)
		y := fc.MapRange(d.value, ymin, c.maxvalue, c.bottom, c.top)
		canvas.Circle(x, y, size, c.color)
	}
}

// Grid makes a grid
func Grid(canvas fc.Canvas, left, bottom, width, height, size float64, color color.RGBA) {
	for x := left; x <= left+width; x += size {
		canvas.Line(x, bottom, x, bottom+height, 0.1, color)
	}
	for y := bottom; y <= bottom+height; y += size {
		canvas.Line(left, y, left+width, y, 0.2, color)
	}
}

// YAxis makes the Y axis with optional grid lines
func (c *ChartBox) YAxis(canvas fc.Canvas, size, min, max, step float64, format string, gridlines bool) {
	w := c.right - c.left
	y := c.bottom
	nlabels := (max - min) / step
	n := (c.top - c.bottom) / nlabels
	for v := min; v <= max; v += step {
		if gridlines {
			canvas.Line(c.left, y, c.left+w, y, 0.05, color.RGBA{128, 128, 128, 255})
		}
		canvas.EText(c.left-2, y-(size/3), size, fmt.Sprintf(format, v), c.color)
		y += n
	}
}

// Title makes a title
func (c *ChartBox) Title(canvas fc.Canvas, size, offset float64) {
	midx := c.left + ((c.right - c.left) / 2)
	canvas.CText(midx, c.top+offset, size, c.title, c.color)
}

// Frame makes a filled frame with the specified opacity (0-100)
func (c *ChartBox) Frame(canvas fc.Canvas, op float64) {
	a := c.color.A // Save opacity
	w := c.right - c.left
	h := c.top - c.bottom
	fa := uint8((op / 100) * 255.0)
	c.color.A = fa
	canvas.Rect(c.left+w/2, c.bottom+h/2, w, h, c.color)
	c.color.A = a // Restore opacity
}

func composite(canvas fc.Canvas) error {
	sr, err := os.Open("sin.d")
	if err != nil {
		return err
	}
	cr, err := os.Open("cos.d")
	if err != nil {
		return err
	}
	sine, err := DataRead(sr)
	if err != nil {
		return err
	}
	cosine, err := DataRead(cr)
	if err != nil {
		return err
	}
	sr.Close()
	cr.Close()

	sine.zerobased, cosine.zerobased = false, false

	cosine.Frame(canvas, 5)
	cosine.Label(canvas, 1.5, 10)
	cosine.YAxis(canvas, 1.2, -1.0, 1.0, 1.0, "%0.2f", true)
	cosine.color = color.RGBA{0, 128, 0, 255}
	sine.color = color.RGBA{128, 0, 0, 255}
	cosine.Scatter(canvas, 0.75)
	sine.Scatter(canvas, 0.75)

	sine.left = 10
	sine.right = sine.left + 40
	sine.top, cosine.top = 30, 30
	sine.bottom, cosine.bottom = 10, 10

	sine.Title(canvas, 2, 2)
	sine.Frame(canvas, 5)
	sine.Scatter(canvas, 0.5)

	offset := 45.0
	cosine.left = sine.left + offset
	cosine.right = sine.right + offset

	cosine.Title(canvas, 2, 2)
	cosine.Frame(canvas, 5)
	cosine.Scatter(canvas, 0.5)

	return nil
}

func playground(canvas fc.Canvas) error {
	chart, err := DataRead(os.Stdin)
	if err != nil {
		return err
	}
	red := color.RGBA{127, 0, 0, 255}
	black := color.RGBA{0, 0, 0, 255}

	textsize := 1.2
	chart.top = 90
	chart.left = 15
	chart.bottom = 70
	chart.Title(canvas, 2, 2)
	chart.Frame(canvas, 5)
	chart.YAxis(canvas, textsize, 0, 800000, 100000, "%0.f", false)
	chart.Label(canvas, textsize, 5)
	chart.color = black
	chart.Scatter(canvas, 0.5)
	chart.color = red
	chart.Line(canvas, 0.1)
	chart.Bar(canvas, 0.25)

	chart.left = 60
	chart.right = 90
	chart.top = 65
	chart.bottom = 55
	chart.color = color.RGBA{0, 0, 127, 255}
	chart.Bar(canvas, 0.2)

	chart.top -= 15
	chart.bottom -= 15
	chart.Line(canvas, 0.05)

	chart.top -= 15
	chart.bottom -= 15
	chart.Scatter(canvas, 0.5)

	chart.top = 65
	chart.left = 15
	chart.right = 50
	chart.bottom = 4
	chart.color = red
	chart.HBar(canvas, 0.2, 1.1, 1)
	chart.Frame(canvas, 5)
	return nil
}
func main() {
	var width, height int
	flag.IntVar(&width, "w", 500, "canvas width")
	flag.IntVar(&height, "h", 500, "canvas height")
	flag.Parse()

	canvas := fc.NewCanvas("Chart", width, height)
	canvas.Rect(50, 50, 100, 100, color.RGBA{255, 255, 255, 255})
	err := composite(canvas)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	canvas.EndRun()
}
