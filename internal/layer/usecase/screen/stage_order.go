package screen

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/property"
)

func (g *usecaseImpl) StageOrder(ctx context.Context, item hardware.Item, data *hardware.Data) {
	g.addCart(ctx, item, data.Cart)

	var step int
	total := data.TotalQuantity()
	if total == 0 {
		if len(item.SlotCode) == 0 {
			step = 0
		} else {
			step = 1
		}
	} else {
		if len(item.SlotCode) == 0 {
			step = 2
		} else {
			step = 3
		}
	}

	switch step {
	case 0:

	case 1:
		g.addButton(ctx, "red", property.NewLeftButton("* BACK"))

	case 2:
		g.addButton(ctx, "red", property.NewLeftButton("* BACK"))
		g.addButton(ctx, "green", property.NewRightButton("# SUM"))

	case 3:
		g.addButton(ctx, "red", property.NewLeftButton("* BACK"))
		g.addButton(ctx, "green", property.NewRightButton("# SUM"))
	}

	g.display.Draw(ctx)
}
