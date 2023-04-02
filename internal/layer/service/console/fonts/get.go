package font_console

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/pkg/trace"
	"github.com/golang/freetype"
)

func (c *consoleImpl) Get(ctx context.Context) *freetype.Context {
	_, span := trace.Start(ctx)
	defer span.End()

	return freetype.NewContext()
}
