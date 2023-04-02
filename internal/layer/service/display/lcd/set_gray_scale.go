package lcd_display

import (
	"context"
	"image"
	"image/draw"

	"github.com/aff-vending-machine/vm-controller/pkg/trace"
	"github.com/disintegration/imaging"
)

func (d *displayImpl) SetGrayScale(ctx context.Context) {
	_, span := trace.Start(ctx)
	defer span.End()

	gray := rgbaToGray(d.canvas)
	blurredImage := imaging.Blur(gray, 3.5)
	rect := image.Rect(0, 0, d.screen.Width, d.screen.Height)

	draw.Draw(d.canvas, rect, blurredImage, image.Point{}, draw.Src)
}

func rgbaToGray(img image.Image) *image.Gray {
	var (
		bounds = img.Bounds()
		gray   = image.NewGray(bounds)
	)
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			var rgba = img.At(x, y)
			gray.Set(x, y, rgba)
		}
	}

	return gray
}
