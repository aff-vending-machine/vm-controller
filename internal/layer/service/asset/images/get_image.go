package images

import (
	"context"
	"encoding/base64"
	"image"
	"image/png"
	"os"
	"strings"

	"github.com/aff-vending-machine/vm-controller/pkg/trace"
)

func (c *assetImpl) GetImage(ctx context.Context, name string) (image.Image, error) {
	_, span := trace.Start(ctx)
	defer span.End()

	if c.cache[name] != nil {
		return c.cache[name], nil
	}

	bytes, err := os.ReadFile(c.path + "/" + name)
	if err != nil {
		return nil, err
	}

	b64 := base64.StdEncoding.EncodeToString(bytes)
	input := base64.NewDecoder(base64.StdEncoding, strings.NewReader(b64))
	result, err := png.Decode(input)
	if err != nil {
		return nil, err
	}

	c.cache[name] = result
	return c.cache[name], nil
}
