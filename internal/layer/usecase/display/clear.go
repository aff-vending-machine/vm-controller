package display

import (
	"context"
)

func (g *usecaseImpl) Clear(ctx context.Context) {
	g.display.ClearCanvas(ctx)
	g.display.Draw(ctx)
}
