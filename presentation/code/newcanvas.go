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
		Container: fyne.NewContainer(iRect(w/2, h/2, w, h, color.RGBA{0, 0, 0, 255})),
		Width:     float64(w),
		Height:    float64(h),
	}
	return c
}
