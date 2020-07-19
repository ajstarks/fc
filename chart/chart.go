// Package chart makes charts using the fync canvas package
package chart

import (
	"bufio"
	"fmt"
	"image/color"
	"io"
	"math"
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
	Title                    string
	Data                     []NameValue
	Color                    color.RGBA
	Top, Bottom, Left, Right float64
	Minvalue, Maxvalue       float64
	Zerobased                bool
}

const (
	largest  = 1.797693134862315708145274237317043567981e+308
	smallest = -largest
)

var labelcolor = color.RGBA{100, 100, 100, 255}

// DataRead reads tab separated values into a ChartBox
// default values for the top, bottom, left, right (90,50,10,90) are filled in
// as is the default color, black
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
		Title:     title,
		Data:      data,
		Minvalue:  minval,
		Maxvalue:  maxval,
		Color:     color.RGBA{0, 0, 0, 255},
		Left:      10,
		Right:     90,
		Top:       90,
		Bottom:    50,
		Zerobased: true,
	}, err
}

// zerobase uses the correct base for scaling
func zerobase(usez bool, n float64) float64 {
	if usez {
		return 0
	}
	return n
}

// Bar makes a (column) bar chart
func (c *ChartBox) Bar(canvas fc.Canvas, size float64) {
	dlen := float64(len(c.Data) - 1)
	ymin := zerobase(c.Zerobased, c.Minvalue)
	for i, d := range c.Data {
		x := fc.MapRange(float64(i), 0, dlen, c.Left, c.Right)
		y := fc.MapRange(d.value, ymin, c.Maxvalue, c.Bottom, c.Top)
		canvas.Line(x, c.Bottom, x, y, size, c.Color)
	}
}

// HBar makes a horizontal bar chart
func (c *ChartBox) HBar(canvas fc.Canvas, size, linespacing, textsize float64) {
	y := c.Top
	xmin := zerobase(c.Zerobased, c.Minvalue)
	for _, d := range c.Data {
		canvas.EText(c.Left-2, y-size/2, textsize, d.label, labelcolor)
		x2 := fc.MapRange(d.value, xmin, c.Maxvalue, c.Left, c.Right)
		canvas.Line(c.Left, y, x2, y, size, c.Color)
		y -= linespacing
	}
}

// Line makes a line chart
func (c *ChartBox) Line(canvas fc.Canvas, size float64) {
	n := len(c.Data)
	fn := float64(n - 1)
	ymin := zerobase(c.Zerobased, c.Minvalue)
	for i := 0; i < n-1; i++ {
		v1 := c.Data[i].value
		v2 := c.Data[i+1].value
		x1 := fc.MapRange(float64(i), 0, fn, c.Left, c.Right)
		y1 := fc.MapRange(v1, ymin, c.Maxvalue, c.Bottom, c.Top)
		x2 := fc.MapRange(float64(i+1), 0, fn, c.Left, c.Right)
		y2 := fc.MapRange(v2, ymin, c.Maxvalue, c.Bottom, c.Top)
		canvas.Line(x1, y1, x2, y2, size, c.Color)
	}
}

// Label draws the x axis labels
func (c *ChartBox) Label(canvas fc.Canvas, size float64, n int) {
	fn := float64(len(c.Data) - 1)
	for i, d := range c.Data {
		x := fc.MapRange(float64(i), 0, fn, c.Left, c.Right)
		if i%n == 0 {
			canvas.CText(x, c.Bottom-(size*2), size, d.label, c.Color)
		}
	}
}

// Scatter makes a scatter chart
func (c *ChartBox) Scatter(canvas fc.Canvas, size float64) {
	dlen := float64(len(c.Data) - 1)
	ymin := zerobase(c.Zerobased, c.Minvalue)
	for i, d := range c.Data {
		x := fc.MapRange(float64(i), 0, dlen, c.Left, c.Right)
		y := fc.MapRange(d.value, ymin, c.Maxvalue, c.Bottom, c.Top)
		canvas.Circle(x, y, size, c.Color)
	}
}

// datasum returns the sum of the data
func datasum(data []NameValue) float64 {
	sum := 0.0
	for _, d := range data {
		sum += d.value
	}
	return sum
}

// dotgrid makes a grid 10x10 grid of dots colored by value
func dotgrid(canvas fc.Canvas, x, y, left, step float64, n int, fillcolor color.RGBA) (float64, float64) {
	edge := (((step * 0.3) + step) * 7) + left
	for i := 0; i < n; i++ {
		if x > edge {
			x = left
			y -= step
		}
		op := fillcolor.A
		canvas.Circle(x, y, 2*step*0.3, fillcolor)
		fillcolor.A = op - 70
		canvas.Rect(x, y, step*0.9, step*0.9, fillcolor)
		fillcolor.A = op
		x += step
	}
	return x, y
}

// Lego makes a lego/waffle chart
func (c *ChartBox) Lego(canvas fc.Canvas, step float64) {
	left := c.Left
	x := left
	y := c.Top

	sum := datasum(c.Data)
	for _, d := range c.Data {
		pct := (d.value / sum) * 100
		v := int(math.Round(pct))
		px, py := dotgrid(canvas, x, y, left, step, v, fc.ColorLookup(d.note))
		x = px
		y = py
	}
	y -= step * 2
	for _, d := range c.Data {
		pct := (d.value / sum) * 100
		v := int(math.Round(pct))
		canvas.Circle(left, y, 2*step*0.3, fc.ColorLookup(d.note))
		canvas.Text(left+step, y-step*0.2, step*0.5, fmt.Sprintf("%s (%.d%%)", d.label, v), fc.ColorLookup("rgb(120,120,120"))
		y -= step
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
	w := c.Right - c.Left
	ymin := zerobase(c.Zerobased, c.Minvalue)
	for v := min; v <= max; v += step {
		y := fc.MapRange(v, ymin, c.Maxvalue, c.Bottom, c.Top)
		if gridlines {
			canvas.Line(c.Left, y, c.Left+w, y, 0.05, color.RGBA{128, 128, 128, 255})
		}
		canvas.EText(c.Left-2, y-(size/3), size, fmt.Sprintf(format, v), c.Color)
	}
}

// CTitle makes a centered title
func (c *ChartBox) CTitle(canvas fc.Canvas, size, offset float64) {
	midx := c.Left + ((c.Right - c.Left) / 2)
	canvas.CText(midx, c.Top+offset, size, c.Title, c.Color)
}

// Frame makes a filled frame with the specified opacity (0-100)
func (c *ChartBox) Frame(canvas fc.Canvas, op float64) {
	a := c.Color.A // Save opacity
	w := c.Right - c.Left
	h := c.Top - c.Bottom
	fa := uint8((op / 100) * 255.0)
	c.Color.A = fa
	canvas.Rect(c.Left+w/2, c.Bottom+h/2, w, h, c.Color)
	c.Color.A = a // Restore opacity
}
