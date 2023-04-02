package display

import (
	"context"
	"fmt"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/property"
)

func (g *usecaseImpl) addGrabItem(ctx context.Context, cart []hardware.Item) {
	screen := g.display.GetProperty(ctx)
	text := property.InitText(g.fontAsset.Get(ctx), 200, 800, screen.Width, screen.Height)

	label := ""
	grab := 0
	total := 0
	price := 0.0

	for _, item := range cart {
		label = fmt.Sprintf("%s: %s", item.SlotCode, ellipsis(item.Name, 16))
		g.display.AddText(ctx, text.Left(label))

		label = fmt.Sprintf("%d / %d", item.Received, item.Quantity)
		g.display.AddText(ctx, text.Right(label))

		grab += item.Received
		total += item.Quantity
		price += float64(item.Quantity) * item.Price
		text = text.NextLine()
	}

	text = text.Reset(0, 400)
	g.display.AddText(ctx, text.Center("TOTAL"))

	text = text.Reset(200, 1600)
	text = text.NextLine()
	g.display.AddText(ctx, text.Left("Received Item"))

	label = fmt.Sprintf("%d / %d", grab, total)
	g.display.AddText(ctx, text.Right(label))
}
