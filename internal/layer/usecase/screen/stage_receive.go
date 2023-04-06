package screen

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
)

func (g *usecaseImpl) StageReceive(ctx context.Context, cart []hardware.Item) {
	g.addGrabItem(ctx, cart)
	g.display.Draw(ctx)
}
