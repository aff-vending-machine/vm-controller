package screen

import (
	"context"
)

func (g *usecaseImpl) StageIdle(ctx context.Context) {
	g.display.Draw(ctx)
}
