package screen

import (
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/asset"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/display"
)

type usecaseImpl struct {
	imgAsset  asset.Images
	fontAsset asset.Fonts
	display   display.LCD
}

func New(i asset.Images, f asset.Fonts, d display.LCD) *usecaseImpl {
	return &usecaseImpl{i, f, d}
}
