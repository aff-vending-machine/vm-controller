package lcd_display

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/property"
	"github.com/aff-vending-machine/vm-controller/pkg/trace"
)

func (d *displayImpl) GetProperty(ctx context.Context) *property.Screen {
	_, span := trace.Start(ctx)
	defer span.End()

	return d.screen
}
