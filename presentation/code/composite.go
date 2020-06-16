cosine.Frame(canvas, 5)
cosine.Label(canvas, 1.5, 10)
cosine.YAxis(canvas, 1.2, -1.0, 1.0, 1.0, "%0.2f", true)
cosine.Color = color.RGBA{0, 128, 0, 255}
sine.Color = color.RGBA{128, 0, 0, 255}
cosine.Scatter(canvas, 0.75)
sine.Scatter(canvas, 0.75)

sine.Left = 10
sine.Right = sine.Left + 40
sine.Top, cosine.Top = 30, 30
sine.Bottom, cosine.Bottom = 10, 10

sine.CTitle(canvas, 2, 2)
sine.Frame(canvas, 5)
sine.Scatter(canvas, 0.5)

offset := 45.0
cosine.Left = sine.Left + offset
cosine.Right = sine.Right + offset

cosine.CTitle(canvas, 2, 2)
cosine.Frame(canvas, 5)
cosine.Scatter(canvas, 0.5)
