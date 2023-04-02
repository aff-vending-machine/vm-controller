package lcd_display

import (
	"context"
	"image"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/property"
	"github.com/aff-vending-machine/vm-controller/pkg/trace"
	"github.com/golang/freetype"
	"github.com/rs/zerolog/log"
)

func (d *displayImpl) AddText(ctx context.Context, props *property.Text) {
	_, span := trace.Start(ctx)
	defer span.End()

	// Truetype stuff
	props.Font.SetFontSize(props.Size)
	props.Font.SetClip(d.canvas.Bounds())
	props.Font.SetDst(d.canvas)
	props.Font.SetSrc(&image.Uniform{props.Color})
	//props.Font.SetHinting(font.HintingFull) // stack overflow

	pt := freetype.Pt(props.PosX, props.PosY)
	_, err := props.Font.DrawString(props.Label, pt)
	if err != nil {
		log.Error().Interface("property", props).Err(err).Msg("unable to draw")
	}
}
