# fc - high-level canvas for the fyne toolkit

## methods

### Text -- return text at (x,y) with the specified size and color

	func Text(x, y int, s string, size int, color color.RGBA) *canvas.Text

### TextMid -- return Center text

	func TextMid(x, y int, s string, size int, color color.RGBA) *canvas.Text

### TectEnd -- return End-Aligned text

	func TextEnd(x, y int, s string, size int, color color.RGBA) *canvas.Text

### Line -- return  a colored line from (x1,y1) to (x2, y2)

	func Line(x1, y1, x2, y2 int, size float32, color color.RGBA) *canvas.Line

### Circle --- return a circle object centered at (x,y) with radius r

	func Circle(x, y, r int, color color.RGBA) *canvas.Circle

### CornerRect -- return a rectangle with (x,y) at the upper left, with dimension (w, h)

	func CornerRect(x, y, w, h int, color color.RGBA) *canvas.Rectangle

### Rect -- return a rectangle centered at (x,y), with dimension (w, h)

	func Rect(x, y, w, h int, color color.RGBA) *canvas.Rectangle

### Return an image named as (name) centered at (x, y) with dimensions (w, h)

	func Image(x, y, w, h int, name string) *canvas.Image

### Return an named image with upper left at (x, y) with dimensions (w, h)

	func CornerImage(x, y, w, h int, name string) *canvas.Image
