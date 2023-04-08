package lcd2k

import (
	"context"
	"image"
	"os"
	"sync"

	"github.com/aff-vending-machine/vm-controller/config"
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/property"
	"github.com/aff-vending-machine/vm-controller/pkg/boot"
)

type displayImpl struct {
	m      sync.Mutex
	fb     *os.File
	canvas *image.RGBA
	bg     image.Image
	screen *property.Screen
}

// Constants for framebuffer
const (
	framebufferWidth  = 1440
	framebufferHeight = 2560
	framebufferBPP    = 16
	framebufferStride = framebufferWidth * framebufferBPP / 8
)

func New(conf config.BoardConfig) *displayImpl {
	fb, err := os.OpenFile(conf.LCDDevice, os.O_RDWR, 0660)
	boot.TerminateWhenError(err)
	boot.AddTerminateFn(func(ctx context.Context) {
		fb.Close()
	})

	rect := image.Rect(0, 0, framebufferWidth, framebufferHeight)

	return &displayImpl{
		fb:     fb,
		canvas: image.NewRGBA(rect),
		bg:     image.NewUniform(image.Black),
		screen: property.NewScreen(0, framebufferWidth, framebufferHeight),
	}
}
