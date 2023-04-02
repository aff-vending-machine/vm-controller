package display

import (
	"context"
	"fmt"
	"image/color"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/property"
)

func (g *usecaseImpl) StagePaymentCreditCard(ctx context.Context, price float64) {
	screen := g.display.GetProperty(ctx)

	text := property.InitText(g.fontAsset.Get(ctx), 0, -480, screen.Width, screen.Height)
	g.display.AddText(ctx, text.Center("Credit Card"))

	src, err := g.imgAsset.GetImage(ctx, "creditcard.png")
	if err != nil {
		src = g.imgAsset.Blank(ctx)
	}
	g.display.AddImage(ctx, &property.Image{
		Image:  src,
		PosX:   int(screen.Width/2) - 320,
		PosY:   int(screen.Height/2) - 320,
		Width:  640,
		Height: 640,
	})

	text = text.Reset(0, 800)
	text.Color = color.White
	label := fmt.Sprintf("Price:  %.02f  Baht", price)
	g.display.AddText(ctx, text.Center(label))

	g.addButton(ctx, "red", property.NewLeftButton("* BACK"))

	g.display.Draw(ctx)
}
