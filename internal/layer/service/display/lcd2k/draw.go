package lcd2k

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/pkg/trace"
	"github.com/rs/zerolog/log"
)

func (d *displayImpl) Draw(ctx context.Context) {
	_, span := trace.Start(ctx)
	defer span.End()

	d.m.Lock()
	defer d.m.Unlock()

	// Convert the image to framebuffer format
	var fbImg [framebufferStride * framebufferHeight]byte
	imgSize := d.canvas.Bounds().Size()
	for y := 0; y < imgSize.Y; y++ {
		copy(fbImg[y*framebufferStride:(y+1)*framebufferStride], d.canvas.Pix[y*d.canvas.Stride:(y+1)*d.canvas.Stride])
	}

	// Double buffering: update the off-screen buffer
	_, err := d.fb.Write(fbImg[:])
	if err != nil {
		log.Error().Err(err).Msg("unable to write to framebuffer:")
		return
	}
}
