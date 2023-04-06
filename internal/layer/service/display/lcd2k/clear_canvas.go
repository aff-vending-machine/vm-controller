package lcd2k

import (
	"context"
	"image"
	"image/draw"

	"github.com/aff-vending-machine/vm-controller/pkg/trace"
)

func (d *displayImpl) ClearCanvas(ctx context.Context) {
	_, span := trace.Start(ctx)
	defer span.End()

	draw.Draw(d.canvas, d.canvas.Rect, d.bg, image.Point{}, draw.Over)
}
