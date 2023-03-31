package repository

import (
	"context"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/trace"
)

func (r *Template[T]) FindOne(ctx context.Context, filter []string) (*T, error) {
	_, span := trace.Start(ctx)
	defer span.End()

	var data T
	tx := MakeQuery(r.DB, filter)

	result := tx.First(&data)
	if err := result.Error; err != nil {
		return nil, err
	}

	return &data, nil
}
