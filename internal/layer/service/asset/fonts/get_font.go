package fonts

import (
	"context"
	"os"

	"github.com/aff-vending-machine/vm-controller/pkg/trace"
	"github.com/golang/freetype"
)

func (c *assetImpl) GetFont(ctx context.Context, name string) (*freetype.Context, error) {
	_, span := trace.Start(ctx)
	defer span.End()

	if c.cache[name] != nil {
		return c.cache[name], nil
	}

	fontByte, err := os.ReadFile(c.path + "/" + name)
	if err != nil {
		return nil, err
	}

	truetypeFont, err := freetype.ParseFont(fontByte)
	if err != nil {
		return nil, err
	}

	fontCtx := freetype.NewContext()
	fontCtx.SetFont(truetypeFont)

	c.cache[name] = fontCtx
	return fontCtx, nil
}
