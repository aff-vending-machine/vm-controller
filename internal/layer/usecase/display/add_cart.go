package display

import (
	"context"
	"fmt"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/property"
)

func (g *usecaseImpl) addCart(ctx context.Context, item hardware.Item, cart []hardware.Item) {
	screen := g.display.GetProperty(ctx)
	text := property.InitText(g.fontAsset.Get(ctx), 200, 800, screen.Width, screen.Height)

	label := ""
	total := 0
	price := 0.0

	for _, item := range cart {
		label = fmt.Sprintf("%s: %s", item.SlotCode, ellipsis(item.Name, 16))
		g.display.AddText(ctx, text.Left(label))

		label = fmt.Sprintf("%d", item.Quantity)
		g.display.AddText(ctx, text.Right(label))

		total += item.Quantity
		price += float64(item.Quantity) * item.Price
		text = text.NextLine()
	}

	scLen := len(item.SlotCode)

	if scLen > 0 {
		label = fmt.Sprintf("%s%s: %s", item.SlotCode, "___"[:3-scLen], ellipsis(item.Name, 16))
		g.display.AddText(ctx, text.Left(label))

		if len(item.SlotCode) == 3 {
			g.display.AddText(ctx, text.Right("_"))
		}
	}

	text = text.Reset(200, 1600)
	g.display.AddText(ctx, text.Left("Total"))

	label = fmt.Sprintf("%d", total)
	g.display.AddText(ctx, text.Right(label))

	text = text.NextLine()
	g.display.AddText(ctx, text.Left("Payment"))

	label = fmt.Sprintf("%0.02f", price)
	g.display.AddText(ctx, text.Right(label))
}
