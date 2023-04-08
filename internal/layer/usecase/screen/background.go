package screen

import (
	"context"
	"fmt"
)

func (g *usecaseImpl) Background(ctx context.Context, stage string) {
	filename := fmt.Sprintf("stage-%s-bg.png", stage)
	bg, err := g.imgAsset.GetImage(ctx, filename)
	if err != nil {
		bg = g.imgAsset.Blank(ctx)
	}
	g.display.SetBackground(ctx, bg)
	//g.display.Draw(ctx)
}
