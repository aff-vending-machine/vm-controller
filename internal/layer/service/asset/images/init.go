package images

import (
	"image"

	"github.com/aff-vending-machine/vm-controller/config"
)

type assetImpl struct {
	path  string
	cache map[string]image.Image
}

func New(cfg config.AssetConfig) *assetImpl {
	return &assetImpl{path: cfg.ImagePath, cache: make(map[string]image.Image)}
}
