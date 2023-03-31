package mail_api

import "github.com/aff-vending-machine/vmc-rpi-ctrl/config"

type apiImpl struct {
	Host     string
	Port     int
	Username string
	Password string
}

func New(cfg config.MailConfig) *apiImpl {
	return &apiImpl{
		Host:     cfg.Host,
		Port:     cfg.Port,
		Username: cfg.Username,
		Password: cfg.Password,
	}
}
