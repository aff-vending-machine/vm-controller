package display

import (
	"context"
	"image"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/property"
)

type LCD interface {
	ClearCanvas(context.Context)
	GetProperty(context.Context) *property.Screen
	SetBackground(context.Context, image.Image)
	SetGrayScale(context.Context)
	AddText(context.Context, *property.Text)
	AddImage(context.Context, *property.Image)
	Draw(context.Context)
}
