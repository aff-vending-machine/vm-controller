package display

import (
	"context"
	"image"
	"image/color"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/property"
	"github.com/aff-vending-machine/vm-controller/pkg/module/flow"
)

func (g *usecaseImpl) Error(ctx context.Context, e error) {
	screen := g.display.GetProperty(ctx)

	filename := "obj-error-msg.png"
	border, err := g.imgAsset.GetImage(ctx, filename)
	if err != nil {
		border = image.NewRGBA(image.Rect(0, 0, 320, 1200))
	}

	g.display.SetGrayScale(ctx)
	g.display.AddImage(ctx, &property.Image{
		Image:  border,
		PosX:   120,
		PosY:   int(screen.Height/2) - 160,
		Width:  1200,
		Height: 320,
	})

	text := property.InitText(g.fontAsset.Get(ctx), 0, 0, screen.Width, screen.Height)
	text.Color = color.White
	text.Size = 72

	switch e {
	case flow.ErrInvalidKey:
		g.display.AddText(ctx, text.Center("PLEASE CORRECT THE KEY"))

	case flow.ErrInvalidSlot:
		g.display.AddText(ctx, text.Center("NO ITEM"))

	case flow.ErrEmptyItem:
		g.display.AddText(ctx, text.Center("OUT OF STOCK"))

	case flow.ErrItemIsNotEnough:
		g.display.AddText(ctx, text.Center("ITEM IS NOT ENOUGH"))

	case flow.ErrMachineE0:
		g.display.AddText(ctx, text.Center("PLEASE CONTACT AGENT"))

	case flow.ErrMachineE1:
		g.display.AddText(ctx, text.Center("PRESS # TO OPEN GATE"))

	case flow.ErrMachineE2:
		g.display.AddText(ctx, text.Center("PLEASE CONTACT AGENT"))

	default:
		g.display.AddText(ctx, text.Center(e.Error()))
	}

	g.display.Draw(ctx)
}
