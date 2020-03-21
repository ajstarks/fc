# fc - high-level canvas for the fyne toolkit

## methods

## Start -- begin the application

	func Start(name string, w, h int) (c *fyne.Container, fyne.Window, *fyne.Container) {

### EndRun -- run the app

	func End(window fyne.Window, content *fyne.Container, w, h int) {

### Text -- Add text at (x,y) with the specified size and color

	func Text(c *fyne.Container, x, y int, s string, size int, color color.RGBA)

### TextMid -- Add Center text

	func TextMid(c *fyne.Container, x, y int, s string, size int, color color.RGBA)

### TectEnd -- Add End-Aligned text

	func TextEnd(c *fyne.Container, x, y int, s string, size int, color color.RGBA)

### Line -- Add  a colored line from (c *fyne.Container, x1,y1) to (c *fyne.Container, x2, y2)

	func Line(c *fyne.Container, x1, y1, x2, y2 int, size float32, color color.RGBA)

### Circle --- Add a circle object centered at (c *fyne.Container, x,y) with radius r

	func Circle(c *fyne.Container, x, y, r int, color color.RGBA)

### CornerRect -- Add a rectangle with (c *fyne.Container, x,y) at the upper left, with dimension (c *fyne.Container, w, h)

	func CornerRect(c *fyne.Container, x, y, w, h int, color color.RGBA)

### Rect -- Add a rectangle centered at (c *fyne.Container, x,y), with dimension (c *fyne.Container, w, h)

	func Rect(c *fyne.Container, x, y, w, h int, color color.RGBA)

### Add an image named as name centered at (x, y) with dimensions ( w, h)

	func Image(c *fyne.Container, x, y, w, h int, name string)

### Add an named image with upper left at (x, y) with dimensions (w, h)

	func CornerImage(c *fyne.Container, x, y, w, h int, name string)
