// elections: show election results on a state grid
package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"

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
func readData(r io.Reader) ([]egrid, int, int, string, error) {
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
	return data, min, max, title, scanner.Err()
}

// process walks the data, making the visualization
func process(canvas fc.Canvas, startx, starty, rowsize, colsize float64, data []egrid, min, max int, title string) {
	println("processing...")
	amin := area(float64(min))
	amax := area(float64(max))
	beginPage(canvas, "black", "white")
	showtitle(canvas, title, starty+15)
	for _, d := range data {
		x := startx + (float64(d.row) * colsize)
		y := starty - (float64(d.col) * rowsize)
		r := maprange(area(float64(d.population)), amin, amax, 2, colsize)
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
func beginPage(canvas fc.Canvas, bgcolor, textcolor string) {
	canvas.CornerRect(0, 0, 1200, 900, fc.ColorLookup(bgcolor))
}

// endPage ends a page
func endPage(canvas fc.Canvas) {
	ctext(canvas, 50, 5, 1.5, "The area of a circle denotes state population: source U.S. Census", "gray")
	fmt.Println("eslide")
}

// beginDoc starts the document
func beginDoc() {
	fmt.Println("deck")
}

// endDoc ends the document
func endDoc() {
	fmt.Println("edeck")
}

func main() {
	width, height := 1200, 900
	canvas := fc.NewCanvas("elections", width, height)
	for _, f := range os.Args[1:] { // for every file, make a page
		r, err := os.Open(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		data, min, max, title, err := readData(r)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		process(canvas, 10, 75, 9, 7, data, min, max, title)
		canvas.EndRun()
		r.Close()
	}
	endDoc()
}
