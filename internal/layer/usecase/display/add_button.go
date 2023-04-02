package display

import (
	"context"
	"fmt"
	"image"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/property"
)

func (g *usecaseImpl) addButton(ctx context.Context, name string, props *property.Button) {
	filename := fmt.Sprintf("obj-%s-btn.png", name)
	btn, err := g.imgAsset.GetImage(ctx, filename)
	if err != nil {
		btn = image.NewRGBA(image.Rect(0, 0, int(props.Width), int(props.Height)))
	}

	g.display.AddImage(ctx, &property.Image{
		Image:  btn,
		PosX:   props.PosX,
		PosY:   props.PosY,
		Width:  props.Width,
		Height: props.Height,
	})

	g.display.AddText(ctx, property.InitText(
		g.fontAsset.Get(ctx),
		props.PosX,
		props.PosY+props.Height/2-20,
		props.Width,
		props.Height).
		Center(props.Label))
}
