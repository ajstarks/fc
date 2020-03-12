// Package fc is fyne high-level canvas
package fc

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
)

// Text places text
func Text(x, y int, s string, size int, color color.RGBA) *canvas.Text {
	t := &canvas.Text{Text: s, Color: color, TextSize: size}
	adj := size / 5
	p := fyne.Position{X: x, Y: y - (size + adj)}
	t.Move(p)
	return t
}

// TextMid centers text
func TextMid(x, y int, s string, size int, color color.RGBA) *canvas.Text {
	t := Text(x, y, s, size, color)
	t.Alignment = fyne.TextAlignCenter
	return t
}

// TextEnd end-aligns text
func TextEnd(x, y int, s string, size int, color color.RGBA) *canvas.Text {
	t := Text(x, y, s, size, color)
	t.Alignment = fyne.TextAlignTrailing
	return t
}

// Line draws a line
func Line(x1, y1, x2, y2 int, size float32, color color.RGBA) *canvas.Line {
	p1 := fyne.Position{X: x1, Y: y1}
	p2 := fyne.Position{X: x2, Y: y2}
	l := &canvas.Line{StrokeColor: color, StrokeWidth: size, Position1: p1, Position2: p2}
	return l
}

// Circle draws a circle centered at (x,y)
func Circle(x, y, r int, color color.RGBA) *canvas.Circle {
	r2 := r / 2
	p1 := fyne.Position{X: x - r2, Y: y - r2}
	p2 := fyne.Position{X: x + r2, Y: y + r2}
	c := &canvas.Circle{FillColor: color, Position1: p1, Position2: p2}
	return c
}

// CornerRect makes a rectangle
func CornerRect(x, y, w, h int, color color.RGBA) *canvas.Rectangle {
	r := &canvas.Rectangle{FillColor: color}
	r.Move(fyne.Position{X: x, Y: y})
	r.Resize(fyne.Size{Width: w, Height: h})
	return r
}

// Rect makes a rectangle centered at x,y
func Rect(x, y, w, h int, color color.RGBA) *canvas.Rectangle {
	r := &canvas.Rectangle{FillColor: color}
	r.Move(fyne.Position{X: x - (w / 2), Y: y - (h / 2)})
	r.Resize(fyne.Size{Width: w, Height: h})
	return r
}

// Image places the image centered at x, y
func Image(x, y, w, h int, name string) *canvas.Image {
	i := canvas.NewImageFromFile(name)
	i.Move(fyne.Position{X: x - (w / 2), Y: y - (h / 2)})
	i.Resize(fyne.Size{Width: w, Height: h})
	return i
}

// CornerImage places the image centered at x, y
func CornerImage(x, y, w, h int, name string) *canvas.Image {
	i := canvas.NewImageFromFile(name)
	i.Move(fyne.Position{X: x, Y: y})
	i.Resize(fyne.Size{Width: w, Height: h})
	return i
}
