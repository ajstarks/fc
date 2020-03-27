# fc - high-level canvas for the fyne toolkit

## Introduction

There are methods for Text (begin, centered, and end aligned), Circles, Lines, Rectangles, and Images.


## Absolute methods: uses absolute coordinate system and fyne structures directly

### AbsStart -- begin the application

	func AbsStart(name string, w, h int) (fyne.Window, *fyne.Container)


### AbsEndRun -- run the app

	func AbsEnd(window fyne.Window, content *fyne.Container, w, h int) {

### AbsText -- Add text at (x,y) with the specified size and color

	func AbsText(c *fyne.Container, x, y int, s string, size int, color color.RGBA)

### AbsTextMid -- Add Center text

	func AbsTextMid(c *fyne.Container, x, y int, s string, size int, color color.RGBA)

### AbsTectEnd -- Add End-Aligned text

	func AbsTextEnd(c *fyne.Container, x, y int, s string, size int, color color.RGBA)

### AbsLine -- Add  a colored line from (x1,y1) to (x2, y2)

	func AbsLine(c *fyne.Container, x1, y1, x2, y2 int, size float32, color color.RGBA)

### AbsCircle --- Add a circle object centered at (x,y) with radius r

	func AbsCircle(c *fyne.Container, x, y, r int, color color.RGBA)

### AbsCornerRect -- Add a rectangle with (c *fyne.Container, x,y) at the upper left, with dimension (w, h)

	func AbsCornerRect(c *fyne.Container, x, y, w, h int, color color.RGBA)

### AbsRect -- Add a rectangle centered at (c *fyne.Container, x,y), with dimension (w, h)

	func AbsRect(c *fyne.Container, x, y, w, h int, color color.RGBA)

### AbsAdd an image named as name centered at (x, y) with dimensions ( w, h)

	func AbsImage(c *fyne.Container, x, y, w, h int, name string)

### AbsAdd an named image with upper left at (x, y) with dimensions (w, h)

	func AbsCornerImage(c *fyne.Container, x, y, w, h int, name string)