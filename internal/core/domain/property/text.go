package property

import (
	"image/color"

	"github.com/golang/freetype"
)

type Text struct {
	Label string
	Font  *freetype.Context
	Size  float64
	Color color.Color
	PosX  int
	PosY  int
	Panel struct {
		Width  int
		Height int
	}
}

func InitText(font *freetype.Context, baseX int, baseY int, w int, h int) *Text {
	return &Text{
		Font:  font,
		Size:  80,
		Color: color.Black,
		PosX:  baseX,
		PosY:  baseY,
		Panel: struct {
			Width  int
			Height int
		}{
			Width:  w,
			Height: h,
		},
	}
}

func (p *Text) Left(label string) *Text {
	return &Text{
		Label: label,
		Font:  p.Font,
		Size:  p.Size,
		Color: p.Color,
		PosX:  p.PosX,
		PosY:  p.PosY,
	}
}

func (p *Text) Center(label string) *Text {
	height := p.Font.PointToFixed(p.Size) >> 6
	width := int(float64(len(label)) * p.Size * 0.625)

	return &Text{
		Label: label,
		Font:  p.Font,
		Size:  p.Size,
		Color: p.Color,
		PosX:  p.PosX + p.Panel.Width/2 - width/2,
		PosY:  p.PosY + p.Panel.Height/2 - int(height)/2,
	}
}

func (p *Text) Right(label string) *Text {
	width := int(float64(len(label)) * p.Size * 0.625)

	return &Text{
		Label: label,
		Font:  p.Font,
		Size:  p.Size,
		Color: p.Color,
		PosX:  p.Panel.Width - p.PosX - width,
		PosY:  p.PosY,
	}
}

func (p *Text) Reset(x, y int) *Text {
	return &Text{
		Font:  p.Font,
		Size:  p.Size,
		Color: p.Color,
		PosX:  x,
		PosY:  y,
		Panel: p.Panel,
	}
}

func (p *Text) NextLine() *Text {
	height := p.Font.PointToFixed(p.Size) >> 6
	newPosY := p.PosY + int(float64(height)*1.5)

	return &Text{
		Font:  p.Font,
		Size:  p.Size,
		Color: p.Color,
		PosX:  p.PosX,
		PosY:  newPosY,
		Panel: p.Panel,
	}
}

// func (p *Text) GetHeight() int {
// 	return int(p.Font.PointToFixed(p.Size) >> 6)
// }

// func (p *Text) GetWidth() int {
// 	return int(p.Size * 0.625)
// }

// func (p *Text) SetCenter(length int, cx int, cy int) {
// 	p.X = cx - (length * p.GetWidth() / 2)
// 	p.Y = cy - (p.GetHeight() / 2)
// }

// func (p *Text) NextLine(spacing float64) {
// 	p.Y += int(float64(p.GetHeight()) * spacing)
// }
