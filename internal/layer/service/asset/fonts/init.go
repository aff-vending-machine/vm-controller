package fonts

import (
	"github.com/aff-vending-machine/vm-controller/config"
	"github.com/golang/freetype"
)

type assetImpl struct {
	path  string
	cache map[string]*freetype.Context
}

func New(conf config.AssetConfig) *assetImpl {
	return &assetImpl{
		path:  conf.FontPath,
		cache: make(map[string]*freetype.Context),
	}
}
