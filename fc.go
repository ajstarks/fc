// Package fc is fyne high-level canvas
package fc

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
)

// Context is where objects are drawn into
type Context struct {
	container fyne.Container
}

// IText places text
func IText(x, y int, s string, size int, color color.RGBA) *canvas.Text {
	t := &canvas.Text{Text: s, Color: color, TextSize: size}
	adj := size / 5
	p := fyne.Position{X: x, Y: y - (size + adj)}
	t.Move(p)
	return t
}

// ITextMid centers text
func ITextMid(x, y int, s string, size int, color color.RGBA) *canvas.Text {
	t := IText(x, y, s, size, color)
	t.Alignment = fyne.TextAlignCenter
	return t
}

// ITextEnd end-aligns text
func ITextEnd(x, y int, s string, size int, color color.RGBA) *canvas.Text {
	t := IText(x, y, s, size, color)
	t.Alignment = fyne.TextAlignTrailing
	return t
}

// ILine draws a line
func ILine(x1, y1, x2, y2 int, size float32, color color.RGBA) *canvas.Line {
	p1 := fyne.Position{X: x1, Y: y1}
	p2 := fyne.Position{X: x2, Y: y2}
	l := &canvas.Line{StrokeColor: color, StrokeWidth: size, Position1: p1, Position2: p2}
	return l
}

// ICircle draws a circle centered at (x,y)
func ICircle(x, y, r int, color color.RGBA) *canvas.Circle {
	p1 := fyne.Position{X: x - r, Y: y - r}
	p2 := fyne.Position{X: x + r, Y: y + r}
	c := &canvas.Circle{FillColor: color, Position1: p1, Position2: p2}
	return c
}

// ICornerRect makes a rectangle
func ICornerRect(x, y, w, h int, color color.RGBA) *canvas.Rectangle {
	r := &canvas.Rectangle{FillColor: color}
	r.Move(fyne.Position{X: x, Y: y})
	r.Resize(fyne.Size{Width: w, Height: h})
	return r
}

// IRect makes a rectangle centered at x,y
func IRect(x, y, w, h int, color color.RGBA) *canvas.Rectangle {
	r := &canvas.Rectangle{FillColor: color}
	r.Move(fyne.Position{X: x - (w / 2), Y: y - (h / 2)})
	r.Resize(fyne.Size{Width: w, Height: h})
	return r
}

// IImage places the image centered at x, y
func IImage(x, y, w, h int, name string) *canvas.Image {
	i := canvas.NewImageFromFile(name)
	i.Move(fyne.Position{X: x - (w / 2), Y: y - (h / 2)})
	i.Resize(fyne.Size{Width: w, Height: h})
	return i
}

// ICornerImage places the image centered at x, y
func ICornerImage(x, y, w, h int, name string) *canvas.Image {
	i := canvas.NewImageFromFile(name)
	i.Move(fyne.Position{X: x, Y: y})
	i.Resize(fyne.Size{Width: w, Height: h})
	return i
}

// container methods

// Text places text within a container
func Text(c *fyne.Container, x, y int, s string, size int, color color.RGBA) {
	t := &canvas.Text{Text: s, Color: color, TextSize: size}
	adj := size / 5
	p := fyne.Position{X: x, Y: y - (size + adj)}
	t.Move(p)
	c.AddObject(t)
}

// TextMid centers text within a container
func TextMid(c *fyne.Container, x, y int, s string, size int, color color.RGBA) {
	t := IText(x, y, s, size, color)
	t.Alignment = fyne.TextAlignCenter
	c.AddObject(t)
}

// TextEnd end-aligns text within a container
func TextEnd(c *fyne.Container, x, y int, s string, size int, color color.RGBA) {
	t := IText(x, y, s, size, color)
	t.Alignment = fyne.TextAlignTrailing
	c.AddObject(t)
}

// Line draws a line within a container
func Line(c *fyne.Container, x1, y1, x2, y2 int, size float32, color color.RGBA) {
	p1 := fyne.Position{X: x1, Y: y1}
	p2 := fyne.Position{X: x2, Y: y2}
	c.AddObject(&canvas.Line{StrokeColor: color, StrokeWidth: size, Position1: p1, Position2: p2})
}

// Circle is a containerized circle within a container
func Circle(c *fyne.Container, x, y, r int, color color.RGBA) {
	p1 := fyne.Position{X: x - r, Y: y - r}
	p2 := fyne.Position{X: x + r, Y: y + r}
	c.AddObject(&canvas.Circle{FillColor: color, Position1: p1, Position2: p2})
}

// CornerRect makes a rectangle within a container
func CornerRect(c *fyne.Container, x, y, w, h int, color color.RGBA) {
	r := &canvas.Rectangle{FillColor: color}
	r.Move(fyne.Position{X: x, Y: y})
	r.Resize(fyne.Size{Width: w, Height: h})
	c.AddObject(r)
}

// Rect makes a rectangle centered at x,y within a container
func Rect(c *fyne.Container, x, y, w, h int, color color.RGBA) {
	r := &canvas.Rectangle{FillColor: color}
	r.Move(fyne.Position{X: x - (w / 2), Y: y - (h / 2)})
	r.Resize(fyne.Size{Width: w, Height: h})
	c.AddObject(r)
}

// Image places the image centered at x, y within a container
func Image(c *fyne.Container, x, y, w, h int, name string) {
	i := canvas.NewImageFromFile(name)
	i.Move(fyne.Position{X: x - (w / 2), Y: y - (h / 2)})
	i.Resize(fyne.Size{Width: w, Height: h})
	c.AddObject(i)
}

// CornerImage places the image centered at x, y within a container
func CornerImage(c *fyne.Container, x, y, w, h int, name string) {
	i := canvas.NewImageFromFile(name)
	i.Move(fyne.Position{X: x, Y: y})
	i.Resize(fyne.Size{Width: w, Height: h})
	c.AddObject(i)
}

// Start initiates the canvas
func Start(name string, w, h int) (fyne.Window, *fyne.Container) {
	return app.New().NewWindow(name), fyne.NewContainer(IRect(w/2, h/2, w, h, color.RGBA{0, 0, 0, 255}))
}

// EndRun shows the content and runs the app
func EndRun(window fyne.Window, content *fyne.Container, w, h int) {
	window.Resize(fyne.NewSize(w, h))
	window.SetFixedSize(true)
	window.SetPadded(false)
	window.SetContent(content)
	window.ShowAndRun()
}
