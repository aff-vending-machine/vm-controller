package main

import (
	"context"
	"fmt"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/config"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/smartedc"
	smartedc_serial "github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/serial/smartedc"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/log"
)

func init() {
	log.New()
}

func main() {

	initLog(config.BootConfig{App: config.AppConfig{ENV: "local"}})

	fmt.Printf("\n\nMAIN\n\n")
	ss := smartedc_serial.New(config.SmartEDCConfig{Port: "/dev/ttyACM0"})

	fmt.Printf("\n\nSALE\n\n")
	res, err := ss.Sale(
		context.TODO(),
		&smartedc.SaleRequest{
			TradeType:       "CARD",
			Amount:          1,
			TransactionType: "SALE",
			POSRefNo:        "TEST-0001",
		})

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", res)

	if res.ResponseMsg != "SUCCESS" {
		return
	}
	fmt.Printf("\n\nVOID\n\n")

	res2, err := ss.Void(
		context.TODO(),
		&smartedc.VoidRequest{
			TradeType:        "CARD",
			InvoiceNo:        "000018", // res.InvoiceNo,
			CardApprovalCode: "003254", // res.CardApprovalCode,
			Amount:           1,
			TransactionType:  "VOID",
			POSRefNo:         "TEST-0001",
		})

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		return
	}

	fmt.Printf("\n\nDONE %+v\n\n", res2)
}

func initLog(conf config.BootConfig) {
	if conf.App.ENV == "local" {
		log.SetOutput(log.ColorConsole())
	}
	log.SetLogLevel(conf.App.LogLevel)
}
