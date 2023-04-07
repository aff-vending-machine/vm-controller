package lcd2k

import (
	"context"
	"image/color"

	"github.com/aff-vending-machine/vm-controller/pkg/trace"
)

func (d *displayImpl) Draw(ctx context.Context) {
	_, span := trace.Start(ctx)
	defer span.End()

	d.m.Lock()
	defer d.m.Unlock()

	var rgba color.RGBA
	var dom color.RGBA
	for i := 0; i < d.screen.Width; i++ {
		for j := 0; j < d.screen.Height; j++ {
			rgba = d.canvas.RGBAAt(i, j)
			dom = d.dom.RGBAAt(i, j)
			if rgba.R != dom.R || rgba.G != dom.G || rgba.B != dom.B {
				d.fb.WritePixel(i, j, int(rgba.R), int(rgba.G), int(rgba.B), int(rgba.A))
				d.dom.SetRGBA(i, j, color.RGBA{R: rgba.R, G: rgba.G, B: rgba.B, A: rgba.A})
			}
		}
	}
}

func (d *displayImpl) drawScreen(i int, j int) {
	invi := d.screen.Width - i - 1
	invj := d.screen.Height - j - 1
	switch d.screen.Rotate {
	case 0:
		rgba := d.canvas.RGBAAt(i, j)
		d.fb.WritePixel(i, j, int(rgba.R), int(rgba.G), int(rgba.B), int(rgba.A))

	case 1:
		rgba := d.canvas.RGBAAt(i, invj)
		d.fb.WritePixel(j, i, int(rgba.R), int(rgba.G), int(rgba.B), int(rgba.A))

	case 2:
		rgba := d.canvas.RGBAAt(invi, invj)
		d.fb.WritePixel(i, j, int(rgba.R), int(rgba.G), int(rgba.B), int(rgba.A))

	case 3:
		rgba := d.canvas.RGBAAt(invi, j)
		d.fb.WritePixel(j, i, int(rgba.R), int(rgba.G), int(rgba.B), int(rgba.A))
	}
}
