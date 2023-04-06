package lcd2k

import (
	"context"
	"image"
	"image/draw"

	"github.com/aff-vending-machine/vm-controller/pkg/trace"
)

func (d *displayImpl) SetBackground(ctx context.Context, img image.Image) {
	_, span := trace.Start(ctx)
	defer span.End()

	// Set position
	rect := image.Rect(0, 0, d.screen.Width, d.screen.Height)
	d.bg = img

	draw.Draw(d.canvas, rect, img, image.Point{}, draw.Src)
}
