package asset

import (
	"context"
	"image"

	"github.com/golang/freetype"
)

type Fonts interface {
	Get(ctx context.Context) *freetype.Context
	GetFont(ctx context.Context, name string) (*freetype.Context, error)
}

type Images interface {
	Blank(ctx context.Context) image.Image
	GetImage(ctx context.Context, name string) (image.Image, error)
}
