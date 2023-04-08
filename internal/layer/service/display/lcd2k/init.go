package lcd2k

import (
	"context"
	"image"
	"sync"

	"github.com/aff-vending-machine/vm-controller/config"
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/property"
	"github.com/aff-vending-machine/vm-controller/pkg/boot"
	gonutz "github.com/gonutz/framebuffer"
	kaey "github.com/kaey/framebuffer"
	"github.com/rs/zerolog/log"
)

type displayImpl struct {
	m        sync.Mutex
	gonutzfb *gonutz.Device
	kaeyfb   *kaey.Framebuffer
	canvas   *image.RGBA
	dom      *image.RGBA
	screen   *property.Screen
	bg       image.Image
}

func New(conf config.BoardConfig) *displayImpl {
	var gonutzDevice *gonutz.Device
	var kaeyDevice *kaey.Framebuffer
	var err error

	gonutzDevice, err = gonutz.Open(conf.LCDDevice)
	if err != nil {
		log.Error().Err(err).Msg("gonutz is not used")
	}

	var width int
	var height int
	var w int
	var h int

	if gonutzDevice == nil {
		boot.AddTerminateFn(func(ctx context.Context) {
			gonutzDevice.Close()
		})
		size := gonutzDevice.Bounds().Size()
		w = size.X
		h = size.Y

	} else {
		kaeyDevice, err = kaey.Init(conf.LCDDevice)
		boot.TerminateWhenError(err)
		boot.AddTerminateFn(func(ctx context.Context) {
			kaeyDevice.Close()
		})

		w, h = kaeyDevice.Size()
	}

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
		gonutzfb: gonutzDevice,
		kaeyfb:   kaeyDevice,
		canvas:   image.NewRGBA(rect),
		dom:      image.NewRGBA(rect),
		screen:   property.NewScreen(rotate, width, height),
		bg:       image.NewUniform(image.Black),
	}
}
