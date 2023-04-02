package display

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/property"
)

func (g *usecaseImpl) StageSummary(ctx context.Context, cart []hardware.Item) {
	g.addCart(ctx, hardware.Item{}, cart)
	g.addButton(ctx, "red", property.NewLeftButton("* Back"))
	g.addButton(ctx, "green", property.NewRightButton("# Pay"))
	g.display.Draw(ctx)
}
