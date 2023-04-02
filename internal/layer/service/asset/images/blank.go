package images

import (
	"context"
	"image"

	"github.com/aff-vending-machine/vm-controller/pkg/trace"
)

func (c *assetImpl) Blank(ctx context.Context) image.Image {
	_, span := trace.Start(ctx)
	defer span.End()

	return image.NewRGBA(image.Rectangle{
		image.Point{0, 0},
		image.Point{1440, 2560},
	})
}
