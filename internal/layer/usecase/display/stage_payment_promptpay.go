package display

import (
	"context"
	"encoding/base64"
	"fmt"
	"image/color"
	"image/png"
	"strings"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/property"
	"github.com/rs/zerolog/log"
)

func (g *usecaseImpl) StagePaymentPromptPay(ctx context.Context, qrcode string, price float64) {
	screen := g.display.GetProperty(ctx)

	logo, err := g.imgAsset.GetImage(ctx, "promptpay.png")
	if err != nil {
		g.display.AddText(ctx, property.InitText(g.fontAsset.Get(ctx), 0, -480, screen.Width, screen.Height).Center("PromptPay"))
	} else {
		g.display.AddImage(ctx, &property.Image{
			Image:  logo,
			PosX:   int(screen.Width/2) - 300,
			PosY:   720,
			Width:  600,
			Height: 200,
		})
	}

	b64 := strings.ReplaceAll(qrcode, "data:image/png;base64,", "")
	input := base64.NewDecoder(base64.StdEncoding, strings.NewReader(b64))
	src, err := png.Decode(input)
	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to handle queue polling")
		return
	}

	g.display.AddImage(ctx, &property.Image{
		Image:  src,
		PosX:   int(screen.Width/2) - 400,
		PosY:   int(screen.Height/2) - 320,
		Width:  800,
		Height: 800,
	})

	text := property.InitText(g.fontAsset.Get(ctx), 0, 800, screen.Width, screen.Height)
	text.Color = color.White
	label := fmt.Sprintf("Price:  %.02f  Baht", price)
	g.display.AddText(ctx, text.Center(label))

	g.addButton(ctx, "red", property.NewLeftButton("* BACK"))

	g.display.Draw(ctx)
}
