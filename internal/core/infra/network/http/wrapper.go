package http

import (
	"net/http"

	"github.com/aff-vending-machine/vm-controller/configs"
)

type Wrapper struct {
	*http.Client
}

func New(cfg configs.HTTPConfig) *Wrapper {
	var client *http.Client
	if cfg.Cert {
		client = createClientWithCert(cfg)
	} else {
		client = createClient(cfg)
	}

	return &Wrapper{
		client,
	}
}
