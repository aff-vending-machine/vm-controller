package preload

import (
	"context"
	"strings"

	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/payment_channel"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/payment_channel/request"
	"github.com/aff-vending-machine/vm-controller/pkg/boot"
	"github.com/rs/zerolog/log"
)

func InitPromptPay(uc payment_channel.Usecase) {
	ctx := context.TODO()

	channel, err := uc.Get(ctx, []string{"channel:=:promptpay"})
	if err != nil && !strings.Contains(err.Error(), "not found") {
		boot.TerminateWhenError(err)
		return
	}
	if channel != nil {
		log.Info().Msg("promptpay channel is already exist")
		return
	}

	err = uc.Create(ctx, &request.Create{
		Name:         "Ksher - PromptPay",
		Channel:      "promptpay",
		Vendor:       "Ksher",
		Active:       true,
		Host:         "",
		MerchantID:   "",
		MerchantName: "",
		BillerCode:   "",
		BillerID:     "",
		StoreID:      "",
		TerminalID:   "",
	})
	log.Info().Msg("create promptpay (LugentPay) channel")
	boot.TerminateWhenError(err)
}

func InitCreditCard(uc payment_channel.Usecase) {
	ctx := context.TODO()

	channel, err := uc.Get(ctx, []string{"channel:=:creditcard"})
	if err != nil && !strings.Contains(err.Error(), "not found") {
		boot.TerminateWhenError(err)
		return
	}
	if channel != nil {
		log.Info().Msg("creditcard channel is already exist")
		return
	}

	err = uc.Create(ctx, &request.Create{
		Name:         "Link2500 - CreditCard",
		Channel:      "creditcard",
		Vendor:       "K",
		Active:       true,
		Host:         "vm-edc-link2500",
		MerchantID:   "000001",
		MerchantName: "",
		BillerCode:   "",
		BillerID:     "",
		StoreID:      "",
		TerminalID:   "",
	})
	log.Info().Msg("create creditcard (GHL) channel")
	boot.TerminateWhenError(err)
}
