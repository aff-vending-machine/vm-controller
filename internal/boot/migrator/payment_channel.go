package migrator

import (
	"context"
	"strings"

	"vm-controller/internal/core/interface/payment_channel"
	"vm-controller/internal/layer/usecase/payment_channel/request"
	"vm-controller/pkg/boot"

	"github.com/rs/zerolog/log"
)

func InitTestPay(uc payment_channel.Usecase) {
	ctx := context.TODO()

	channel, err := uc.Get(ctx, &request.Get{Channel: "testpay"})
	if err != nil && !strings.Contains(err.Error(), "not found") {
		boot.TerminateWhenError(err)
		return
	}
	if channel != nil {
		log.Info().Msg("testpay channel is already exist")
		return
	}

	err = uc.Create(ctx, &request.Create{
		Name:    "Test",
		Channel: "testpay",
		Vendor:  "at44",
		Active:  true,
	})
	boot.TerminateWhenError(err)
	log.Info().Msg("create test channel")
}

func InitPromptPay(uc payment_channel.Usecase) {
	ctx := context.TODO()

	channel, err := uc.Get(ctx, &request.Get{Channel: "promptpay"})
	if err != nil && !strings.Contains(err.Error(), "not found") {
		boot.TerminateWhenError(err)
		return
	}
	if channel != nil {
		log.Info().Msg("promptpay channel is already exist")
		return
	}

	err = uc.Create(ctx, &request.Create{
		Name:    "Ksher - PromptPay",
		Channel: "promptpay",
		Vendor:  "Ksher",
		Active:  true,
	})
	boot.TerminateWhenError(err)
	log.Info().Msg("create promptpay (Ksher) channel")
}

func InitCreditCard(uc payment_channel.Usecase) {
	ctx := context.TODO()

	channel, err := uc.Get(ctx, &request.Get{Channel: "creditcard"})
	if err != nil && !strings.Contains(err.Error(), "not found") {
		boot.TerminateWhenError(err)
		return
	}
	if channel != nil {
		log.Info().Msg("creditcard channel is already exist")
		return
	}

	err = uc.Create(ctx, &request.Create{
		Name:       "Link2500 - CreditCard",
		Channel:    "creditcard",
		Vendor:     "Kasikorn Bank",
		Active:     true,
		Host:       "vm-link2500",
		MerchantID: "000001",
	})
	boot.TerminateWhenError(err)
	log.Info().Msg("create creditcard (Link2500) channel")
}
