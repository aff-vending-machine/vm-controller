package lcd2k

import (
	"context"
	"image"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/property"
	"github.com/rs/zerolog/log"
)

func (*consoleImpl) ClearCanvas(context.Context) {
}

func (*consoleImpl) GetProperty(context.Context) *property.Screen {
	return &property.Screen{
		Rotate: 0,
		Width:  0,
		Height: 0,
	}
}

func (*consoleImpl) SetBackground(context.Context, image.Image) {
	log.Debug().Msg("SCREEN: set background")
}

func (*consoleImpl) SetGrayScale(context.Context) {
}

func (*consoleImpl) AddText(context.Context, *property.Text) {
}

func (*consoleImpl) AddImage(context.Context, *property.Image) {
}

func (*consoleImpl) Draw(context.Context) {
}
