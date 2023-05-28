package payment_channel

import (
	"context"

	"vm-controller/internal/layer/usecase/payment_channel/request"
	"vm-controller/pkg/helpers/db"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) Set(ctx context.Context, req *request.Set) error {
	if v := validate.Struct(req); !v.Validate() {
		return errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	for _, ch := range req.Data {
		channel, err := uc.paymentChannelRepo.FindOne(ctx, db.NewQuery().AddWhere("channel = ?", ch.Channel))
		if err != nil {
			return err
		}

		updated := map[string]interface{}{
			"name":          ch.Name,
			"vendor":        ch.Vendor,
			"is_enable":     ch.IsEnable,
			"host":          ch.Host,
			"merchant_id":   ch.MerchantID,
			"merchant_name": ch.MerchantName,
			"biller_code":   ch.BillerCode,
			"biller_id":     ch.BillerID,
			"token":         ch.Token,
			"store_id":      ch.StoreID,
			"terminal_id":   ch.TerminalID,
		}

		_, err = uc.paymentChannelRepo.Update(ctx, db.NewQuery().AddWhere("id = ?", channel.ID), updated)
		if err != nil {
			log.Error().Err(err).Interface("req", req).Msg("unable to update payment channel")
		}

	}

	return nil
}
