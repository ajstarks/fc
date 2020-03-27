// Package fc is fyne high-level canvas
package fc

import (
	"image/color"
	"math"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
)

// Canvas is where objects are drawn into
type Canvas struct {
	Window    fyne.Window
	Container *fyne.Container
	Width     float64
	Height    float64
}

// NewCanvas makes a new canvas
func NewCanvas(name string, w, h int) Canvas {
	c := Canvas{
		Window:    app.New().NewWindow(name),
		Container: fyne.NewContainer(IRect(w/2, h/2, w, h, color.RGBA{0, 0, 0, 255})),
		Width:     float64(w),
		Height:    float64(h),
	}
	return c
}

// MapRange -- given a value between low1 and high1, return the corresponding value between low2 and high2
func MapRange(value, low1, high1, low2, high2 float64) float64 {
	return low2 + (high2-low2)*(value-low1)/(high1-low1)
}

const pi = 3.14159265358979323846264338327950288419716939937510582097494459 // https://oeis.org/A000796

// Radians converts degrees to radians
func Radians(deg float64) float64 {
	return (deg * pi) / 180.0
}

// Polar returns the euclidian corrdinates from polar coordinates
func Polar(x, y, r, angle float64) (float64, float64) {
	px := (r * math.Cos(Radians(angle))) + x
	py := (r * math.Sin(Radians(angle))) + y
	return px, py
}

func pct(p float64, m float64) float64 {
	return ((p / 100.0) * m)
}

// dimen returns canvas dimensions from percentages (converting from x increasing left-right, y increasing top-bottom)
func dimen(xp, yp, w, h float64) (float64, float64) {
	return pct(xp, w), pct(100-yp, h)
}

// AbsStart initiates the canvas
func AbsStart(name string, w, h int) (fyne.Window, *fyne.Container) {
	return app.New().NewWindow(name), fyne.NewContainer(IRect(w/2, h/2, w, h, color.RGBA{0, 0, 0, 255}))
}

// EndRun shows the content and runs the app
func (canvas *Canvas) EndRun() {
	window := canvas.Window
	window.Resize(fyne.NewSize(int(canvas.Width), int(canvas.Height)))
	window.SetFixedSize(true)
	window.SetPadded(false)
	window.SetContent(canvas.Container)
	window.ShowAndRun()
}

// AbsEndRun shows the content and runs the app using bare windows and containers
func AbsEndRun(window fyne.Window, c *fyne.Container, w, h int) {
	window.Resize(fyne.NewSize(w, h))
	window.SetFixedSize(true)
	window.SetPadded(false)
	window.SetContent(c)
	window.ShowAndRun()
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

// container methods, Absoulte coordinates

// AbsText places text within a container
func AbsText(c *fyne.Container, x, y int, s string, size int, color color.RGBA) {
	t := &canvas.Text{Text: s, Color: color, TextSize: size}
	adj := size / 5
	p := fyne.Position{X: x, Y: y - (size + adj)}
	t.Move(p)
	c.AddObject(t)
}

// AbsTextMid centers text within a container
func AbsTextMid(c *fyne.Container, x, y int, s string, size int, color color.RGBA) {
	t := IText(x, y, s, size, color)
	t.Alignment = fyne.TextAlignCenter
	c.AddObject(t)
}

// AbsTextEnd end-aligns text within a container
func AbsTextEnd(c *fyne.Container, x, y int, s string, size int, color color.RGBA) {
	t := IText(x, y, s, size, color)
	t.Alignment = fyne.TextAlignTrailing
	c.AddObject(t)
}

// AbsLine draws a line within a container
func AbsLine(c *fyne.Container, x1, y1, x2, y2 int, size float32, color color.RGBA) {
	p1 := fyne.Position{X: x1, Y: y1}
	p2 := fyne.Position{X: x2, Y: y2}
	c.AddObject(&canvas.Line{StrokeColor: color, StrokeWidth: size, Position1: p1, Position2: p2})
}

// AbsCircle is a containerized circle within a container
func AbsCircle(c *fyne.Container, x, y, r int, color color.RGBA) {
	p1 := fyne.Position{X: x - r, Y: y - r}
	p2 := fyne.Position{X: x + r, Y: y + r}
	c.AddObject(&canvas.Circle{FillColor: color, Position1: p1, Position2: p2})
}

// AbsCornerRect makes a rectangle within a container
func AbsCornerRect(c *fyne.Container, x, y, w, h int, color color.RGBA) {
	r := &canvas.Rectangle{FillColor: color}
	r.Move(fyne.Position{X: x, Y: y})
	r.Resize(fyne.Size{Width: w, Height: h})
	c.AddObject(r)
}

// AbsRect makes a rectangle centered at x,y within a container
func AbsRect(c *fyne.Container, x, y, w, h int, color color.RGBA) {
	r := &canvas.Rectangle{FillColor: color}
	r.Move(fyne.Position{X: x - (w / 2), Y: y - (h / 2)})
	r.Resize(fyne.Size{Width: w, Height: h})
	c.AddObject(r)
}

// AbsImage places the image centered at x, y within a container
func AbsImage(c *fyne.Container, x, y, w, h int, name string) {
	i := canvas.NewImageFromFile(name)
	i.Move(fyne.Position{X: x - (w / 2), Y: y - (h / 2)})
	i.Resize(fyne.Size{Width: w, Height: h})
	c.AddObject(i)
}

// AbsCornerImage places the image centered at x, y within a container
func AbsCornerImage(c *fyne.Container, x, y, w, h int, name string) {
	i := canvas.NewImageFromFile(name)
	i.Move(fyne.Position{X: x, Y: y})
	i.Resize(fyne.Size{Width: w, Height: h})
	c.AddObject(i)
}

//
// container methods, Percent coordinates
//

// Text places text within a container, using percent coordinates
func (canvas *Canvas) Text(x, y float64, size float64, s string, color color.RGBA) {
	x, y = dimen(x, y, canvas.Width, canvas.Height)
	size = pct(size, canvas.Width)
	AbsText(canvas.Container, int(x), int(y), s, int(size), color)
}

// CText places centered text using percent coordinates
func (canvas *Canvas) CText(x, y float64, size float64, s string, color color.RGBA) {
	x, y = dimen(x, y, canvas.Width, canvas.Height)
	size = pct(size, canvas.Width)
	AbsTextMid(canvas.Container, int(x), int(y), s, int(size), color)
}

// EText places end-aligned text within a container, using percent coordinates
func (canvas *Canvas) EText(x, y float64, size float64, s string, color color.RGBA) {
	x, y = dimen(x, y, canvas.Width, canvas.Height)
	size = pct(size, canvas.Width)
	AbsTextEnd(canvas.Container, int(x), int(y), s, int(size), color)
}

// Circle places a circle within a container, using percent coordinates
func (canvas *Canvas) Circle(x, y, r float64, color color.RGBA) {
	x, y = dimen(x, y, canvas.Width, canvas.Height)
	r = pct(r, canvas.Width)
	AbsCircle(canvas.Container, int(x), int(y), int(r/2), color)
}

// Line places a line within a container, using percent coordinates
func (canvas *Canvas) Line(x1, y1, x2, y2, size float64, color color.RGBA) {
	x1, y1 = dimen(x1, y1, canvas.Width, canvas.Height)
	x2, y2 = dimen(x2, y2, canvas.Width, canvas.Height)
	size = pct(size, canvas.Width)
	AbsLine(canvas.Container, int(x1), int(y1), int(x2), int(y2), float32(size), color)
}

// Rect places a rectangle centered on (x,y) within a container, using percent coordinates
func (canvas *Canvas) Rect(x, y, w, h float64, color color.RGBA) {
	x, y = dimen(x, y, canvas.Width, canvas.Height)
	w, h = dimen(w, h, canvas.Width, canvas.Height)
	AbsRect(canvas.Container, int(x), int(y), int(w), int(h), color)
}

// CornerRect places a rectangle with upper left corner  on (x,y) within a container, using percent coordinates
func (canvas *Canvas) CornerRect(x, y, w, h float64, color color.RGBA) {
	x, y = dimen(x, y, canvas.Width, canvas.Height)
	w, h = dimen(w, h, canvas.Width, canvas.Height)
	AbsCornerRect(canvas.Container, int(x), int(y), int(w), int(h), color)
}

// Image places an image centered at (x, y) within a container, using percent coordinates
func (canvas *Canvas) Image(x, y float64, w, h int, name string) {
	x, y = dimen(x, y, canvas.Width, canvas.Height)
	AbsImage(canvas.Container, int(x), int(y), w, h, name)
}
