package property

import img "image"

type Image struct {
	Image  img.Image
	PosX   int
	PosY   int
	Width  int
	Height int
}

func NewBackground(i img.Image) *Image {
	return &Image{
		Image:  i,
		PosX:   0,
		PosY:   0,
		Width:  i.Bounds().Dx(),
		Height: i.Bounds().Dy(),
	}
}
