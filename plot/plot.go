package main

import (
	"bytes"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
	"image/color"
	"os"
)

func main() {
	p := plot.New()

	p.Title.Text = "My very very very\nlong Title"
	p.X.Min = 0
	p.X.Max = 20
	p.Y.Min = 0
	p.Y.Max = 10

	p.X.Label.Text = "X-axis"
	p.Y.Label.Text = "Y-axis"

	f1 := plotter.NewFunction(func(x float64) float64 { return 5 })
	f1.LineStyle.Color = color.RGBA{R: 255, G: 255}

	f2 := plotter.NewFunction(func(x float64) float64 { return 6 })
	f2.LineStyle.Color = color.RGBA{R: 255, B: 255}

	labels, err := plotter.NewLabels(plotter.XYLabels{
		XYs: []plotter.XY{
			{X: 2.5, Y: 2.5},
			{X: 7.5, Y: 2.5},
			{X: 7.5, Y: 7.5},
			{X: 2.5, Y: 7.5},
		},
		Labels: []string{"Agg", "Bgg", "Cgg", "Dgg"},
	})
	p.Add(f1, f2, labels)
	p.Add(plotter.NewGrid())

	//p.Legend.Add("fg1", f1)
	//p.Legend.Add("fg2", f2)
	//p.Legend.Top = true

	c := vgimg.PngCanvas{
		Canvas: vgimg.New(20*vg.Centimeter, 15*vg.Centimeter),
	}

	d := draw.New(c)
	p.Draw(d)
	p.DrawGlyphBoxes(d)

	buf := new(bytes.Buffer)
	_, err = c.WriteTo(buf)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("glyphbox.png", buf.Bytes(), 0644)
}
