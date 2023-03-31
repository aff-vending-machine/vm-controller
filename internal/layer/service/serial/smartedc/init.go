package smartedc_serial

import (
	"time"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/config"
	"github.com/tarm/serial"
)

type smartedcImpl struct {
	config *serial.Config
}

func New(conf config.SmartEDCConfig) *smartedcImpl {
	config := &serial.Config{
		Name:        conf.Port,
		Baud:        9600,
		ReadTimeout: time.Duration(conf.TimeoutInSec) * time.Second,
		Size:        8,
		Parity:      serial.ParityOdd,
		StopBits:    serial.Stop1,
	}

	return &smartedcImpl{
		config: config,
	}
}
