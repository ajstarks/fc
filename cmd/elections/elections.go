// elections: show election results on a state grid
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/ajstarks/fc"
)

// Data file structure
type egrid struct {
	name       string
	party      string
	row        int
	col        int
	population int
}

// One election "frame"
type election struct {
	title    string
	min, max int
	data     []egrid
}

type options struct {
	width, height               int
	top, left, rowsize, colsize float64
}

var partymap = map[string]string{"r": "red", "d": "blue", "i": "gray"}

// maprange maps one range into another
func maprange(value, low1, high1, low2, high2 float64) float64 {
	return low2 + (high2-low2)*(value-low1)/(high1-low1)
}

// area computes the area of a circle
func area(v float64) float64 {
	return math.Sqrt((v / math.Pi)) * 2
}

// atoi converts a string to an integer
func atoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return v
}

// readData reads election data into the data structure
func readData(r io.Reader) (election, error) {
	var d egrid
	var data []egrid
	title := ""
	scanner := bufio.NewScanner(r)
	min, max := math.MaxInt32, -math.MaxInt32
	for scanner.Scan() {
		t := scanner.Text()
		if len(t) == 0 { // skip blank lines
			continue
		}
		if t[0] == '#' && len(t) > 2 { // get the title
			title = t[2:]
			continue
		}
		fields := strings.Split(t, "\t")
		if len(fields) < 5 { // skip incomplete records
			continue
		}
		// name,col,row,party,population
		d.name = fields[0]
		d.col = atoi(fields[1])
		d.row = atoi(fields[2])
		d.party = fields[3]
		d.population = atoi(fields[4])
		data = append(data, d)
		// compute min, max
		if d.population > max {
			max = d.population
		}
		if d.population < min {
			min = d.population
		}
	}
	var e election
	e.title = title
	e.min = min
	e.max = max
	e.data = data
	return e, scanner.Err()
}

// process walks the data, making the visualization
func process(canvas fc.Canvas, opts options, e election) {
	amin := area(float64(e.min))
	amax := area(float64(e.max))
	beginPage(canvas, "black")
	showtitle(canvas, e.title, opts.top+15)
	for _, d := range e.data {
		x := opts.left + (float64(d.row) * opts.colsize)
		y := opts.top - (float64(d.col) * opts.rowsize)
		r := maprange(area(float64(d.population)), amin, amax, 2, opts.colsize)
		circle(canvas, x, y, r, partymap[d.party])
		ctext(canvas, x, y-0.5, 1.2, d.name, "white")
	}
	endPage(canvas)
}

// showtitle shows the title and subhead
func showtitle(canvas fc.Canvas, s string, top float64) {
	fields := strings.Fields(s) // year, democratic, republican, third-party (optional)
	if len(fields) < 3 {
		return
	}
	suby := top - 7
	ctext(canvas, 50, top, 3.6, fields[0]+" US Presidential Election", "white")
	legend(canvas, 20, suby, 2.0, fields[1], partymap["d"])
	legend(canvas, 80, suby, 2.0, fields[2], partymap["r"])
	if len(fields) > 3 {
		legend(canvas, 50, suby, 2.0, fields[3], partymap["i"])
	}
}

// circle makes a circle
func circle(canvas fc.Canvas, x, y, r float64, color string) {
	canvas.Circle(x, y, r, fc.ColorLookup(color))
}

// ctext makes centered text
func ctext(canvas fc.Canvas, x, y, ts float64, s string, color string) {
	canvas.CText(x, y, ts, s, fc.ColorLookup(color))
}

// ltext makes left-aligned text
func ltext(canvas fc.Canvas, x, y, ts float64, s string, color string) {
	canvas.Text(x, y, ts, s, fc.ColorLookup(color))
}

// legend makes the subtitle
func legend(canvas fc.Canvas, x, y, ts float64, s string, color string) {
	ltext(canvas, x, y, ts, s, "white")
	circle(canvas, x-ts, y+ts/3, ts/2, color)
}

// beginPage starts a page
func beginPage(canvas fc.Canvas, bgcolor string) {
	canvas.Rect(50, 50, 100, 100, fc.ColorLookup(bgcolor))
}

// endPage ends a page
func endPage(canvas fc.Canvas) {
	ctext(canvas, 50, 5, 1.5, "The area of a circle denotes state population: source U.S. Census", "gray")
}

// back shows the previous frame
func back(c fc.Canvas, opts options, e []election, n *int, limit int) {
	*n--
	if *n < 0 {
		*n = limit - 1
	}
	process(c, opts, e[*n])
}

// forward shows the next frame
func forward(c fc.Canvas, opts options, e []election, n *int, limit int) {
	*n++
	if *n > limit-1 {
		*n = 0
	}
	process(c, opts, e[*n])
}

func main() {

	var opts options
	flag.Float64Var(&opts.top, "top", 75, "top value")
	flag.Float64Var(&opts.left, "left", 7, "top value")
	flag.Float64Var(&opts.rowsize, "rowsize", 9, "top value")
	flag.Float64Var(&opts.colsize, "colsize", 7, "top value")
	flag.IntVar(&opts.width, "width", 1200, "canvas width")
	flag.IntVar(&opts.height, "height", 900, "canvas height")
	flag.Parse()
	// Read in the data
	var elections []election
	for _, f := range flag.Args() {
		r, err := os.Open(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		e, err := readData(r)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		elections = append(elections, e)
		r.Close()
	}
	frames := len(elections)
	if frames < 1 {
		fmt.Fprintln(os.Stderr, "no data read")
		os.Exit(1)
	}

	// initialize
	c := fc.NewCanvas("elections", opts.width, opts.height)
	w := c.Window
	n := 0
	process(c, opts, elections[n]) // show the first frame

	// make the toolbars
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.NavigateBackIcon(), func() { back(c, opts, elections, &n, frames) }),    // previous frame
		widget.NewToolbarAction(theme.NavigateNextIcon(), func() { forward(c, opts, elections, &n, frames) }), // next frame
	)
	// add the content
	w.SetContent(fyne.NewContainerWithLayout(layout.NewBorderLayout(toolbar, nil, nil, nil), toolbar, c.Container))
	w.Resize(fyne.NewSize(opts.width, opts.height+toolbar.Size().Height))

	// run it!
	w.ShowAndRun()

}
