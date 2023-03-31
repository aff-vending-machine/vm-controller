package preload

import (
	"context"
	"strings"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/usecase/payment_channel"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/usecase/payment_channel/request"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/boot"
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
		Name:         "LugentPay - PromptPay",
		Channel:      "promptpay",
		Vendor:       "LugentPay",
		Active:       true,
		Host:         "paymentgw.lugentpay.com:44083",
		MerchantID:   "TTROPQR",
		MerchantName: "TTROPQR",
		BillerCode:   "LGN",
		BillerID:     "310080016625002",
		StoreID:      "4033000",
		TerminalID:   "1001",
	})
	log.Info().Msg("create promptpay (LugentPay) channel")
	boot.TerminateWhenError(err)
}

func InitWechatPay(uc payment_channel.Usecase) {
	ctx := context.TODO()

	channel, err := uc.Get(ctx, []string{"channel:=:wechatpay"})
	if err != nil && !strings.Contains(err.Error(), "not found") {
		boot.TerminateWhenError(err)
		return
	}
	if channel != nil {
		log.Info().Msg("wechatpay channel is already exist")
		return
	}

	err = uc.Create(ctx, &request.Create{
		Name:         "LugentPay - WechatPay",
		Channel:      "wechatpay",
		Vendor:       "LugentPay",
		Active:       true,
		Host:         "paymentgw.lugentpay.com:44083",
		MerchantID:   "TTROPQR",
		MerchantName: "TTROPQR",
		BillerCode:   "LGN",
		BillerID:     "310080016625002",
		StoreID:      "4033000",
		TerminalID:   "1001",
	})

	log.Info().Msg("create wechatpay (LugentPay) channel")
	boot.TerminateWhenError(err)
}

func InitAliPay(uc payment_channel.Usecase) {
	ctx := context.TODO()

	channel, err := uc.Get(ctx, []string{"channel:=:alipay"})
	if err != nil && !strings.Contains(err.Error(), "not found") {
		boot.TerminateWhenError(err)
		return
	}
	if channel != nil {
		log.Info().Msg("alipay channel is already exist")
		return
	}

	err = uc.Create(ctx, &request.Create{
		Name:         "LugentPay - AliPay",
		Channel:      "alipay",
		Vendor:       "LugentPay",
		Active:       true,
		Host:         "paymentgw.lugentpay.com:44083",
		MerchantID:   "TTROPQR",
		MerchantName: "TTROPQR",
		BillerCode:   "LGN",
		BillerID:     "310080016625002",
		StoreID:      "4033000",
		TerminalID:   "1001",
	})
	log.Info().Msg("create alipay (LugentPay) channel")
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
		Name:         "GHL - CreditCard",
		Channel:      "creditcard",
		Vendor:       "GHL",
		Active:       true,
		Host:         "",
		MerchantID:   "",
		MerchantName: "",
		BillerCode:   "",
		BillerID:     "",
		StoreID:      "",
		TerminalID:   "",
	})
	log.Info().Msg("create creditcard (GHL) channel")
	boot.TerminateWhenError(err)
}
