package fonts

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/pkg/trace"
	"github.com/golang/freetype"
)

func (c *consoleImpl) GetFont(ctx context.Context, _ string) (*freetype.Context, error) {
	_, span := trace.Start(ctx)
	defer span.End()

	return freetype.NewContext(), nil
}
