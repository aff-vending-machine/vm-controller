package fonts

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/pkg/trace"
	"github.com/golang/freetype"
	"github.com/rs/zerolog/log"
)

func (a *assetImpl) Get(ctx context.Context) *freetype.Context {
	_, span := trace.Start(ctx)
	defer span.End()

	if a.cache["current"] == nil {
		res, err := a.GetFont(ctx, "NotoSansMono-Regular.ttf")
		if err != nil {
			log.Error().Err(err).Msg("unable to read default font")
			return freetype.NewContext()
		}

		a.cache["current"] = res
		return res
	}

	return a.cache["current"]
}
