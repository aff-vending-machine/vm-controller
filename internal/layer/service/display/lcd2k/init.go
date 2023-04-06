package lcd2k

import (
	"context"
	"image"
	"sync"

	"github.com/aff-vending-machine/vm-controller/config"
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/property"
	"github.com/aff-vending-machine/vm-controller/pkg/boot"
	"github.com/kaey/framebuffer"
)

type displayImpl struct {
	m      sync.Mutex
	fb     *framebuffer.Framebuffer
	canvas *image.RGBA
	screen *property.Screen
	bg     image.Image
}

func New(conf config.BoardConfig) *displayImpl {
	device, err := framebuffer.Init(conf.LCDDevice)
	boot.TerminateWhenError(err)
	boot.AddTerminateFn(func(ctx context.Context) {
		device.Close()
	})

	var width int
	var height int

	w, h := device.Size()
	rotate := conf.LCDRotate % 4

	if rotate%2 == 0 {
		width = w
		height = h
	} else {
		width = h
		height = w
	}
	rect := image.Rect(0, 0, width, height)

	return &displayImpl{
		fb:     device,
		canvas: image.NewRGBA(rect),
		screen: property.NewScreen(rotate, width, height),
		bg:     image.NewUniform(image.Black),
	}
}
