package http

import (
	"net/http"

	"github.com/aff-vending-machine/vm-controller/config"
)

type Wrapper struct {
	*http.Client
}

func New(cfg config.HTTPConfig) *Wrapper {
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
