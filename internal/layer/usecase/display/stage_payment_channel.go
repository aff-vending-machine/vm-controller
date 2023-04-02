package display

import (
	"context"
	"fmt"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/property"
)

func (g *usecaseImpl) StagePaymentChannel(ctx context.Context, channels []entity.PaymentChannel) {
	g.display.ClearCanvas(ctx)

	g.addButton(ctx, "red", property.NewLeftButton("* BACK"))

	screen := g.display.GetProperty(ctx)
	text := property.InitText(g.fontAsset.Get(ctx), 200, 800, screen.Width, screen.Height)
	text.Size = 80
	g.display.AddText(ctx, text.Left("Payment Channel:"))
	text.PosY = text.PosY + 20
	for i, channel := range channels {
		text = text.NextLine()
		text.PosY = text.PosY + 20
		msg := fmt.Sprintf("%d. %s", i+1, channel.Name)
		g.display.AddText(ctx, text.Left(msg))
	}

	g.display.Draw(ctx)
}
