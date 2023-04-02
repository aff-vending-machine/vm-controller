package lcd_display

import (
	"context"
	"image"
	"image/draw"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/property"
	"github.com/aff-vending-machine/vm-controller/pkg/trace"
	"github.com/nfnt/resize"
)

func (d *displayImpl) AddImage(ctx context.Context, props *property.Image) {
	_, span := trace.Start(ctx)
	defer span.End()

	// Set the expected size
	dst := resize.Resize(
		uint(props.Width),
		uint(props.Height),
		props.Image,
		resize.Lanczos3,
	)

	// Set position
	rect := image.Rect(
		props.PosX,
		props.PosY,
		props.PosX+props.Width,
		props.PosY+props.Height,
	)

	draw.Draw(d.canvas, rect, dst, image.Point{}, draw.Over)
}
