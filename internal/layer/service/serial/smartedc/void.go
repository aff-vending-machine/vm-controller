package smartedc_serial

import (
	"context"
	"encoding/xml"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/smartedc"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/trace"
	"github.com/rs/zerolog/log"
	"github.com/tarm/serial"
)

func (e *smartedcImpl) Void(ctx context.Context, req *smartedc.VoidRequest) (*smartedc.VoidResult, error) {
	_, span := trace.Start(ctx)
	defer span.End()

	stream, err := serial.OpenPort(e.config)
	if err != nil {
		return nil, err
	}
	defer stream.Close()

	if err := stream.Flush(); err != nil {
		return nil, err
	}

	// amount * 100
	req.Amount *= 100
	payload, _ := xml.Marshal(req)
	log.Info().Bytes("payload", payload).Msg("EDC send")

	// POS send payload to EDC
	_, err = stream.Write(payload)
	if err != nil {
		return nil, err
	}

	// EDC response to POS
	res := make([]byte, 1024)
	n, err := stream.Read(res)
	if err != nil {
		return nil, err
	}
	log.Info().Bytes("result", res[:n]).Msg("EDC received")

	var result smartedc.VoidResult
	err = xml.Unmarshal(res[:n], &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
