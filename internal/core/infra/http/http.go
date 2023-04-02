package http

import (
	"net/http"

	"github.com/aff-vending-machine/vm-controller/config"
)

type apiImpl struct {
	client *http.Client
}

func New(cfg config.HTTPConfig) *apiImpl {
	var client *http.Client
	if cfg.Cert {
		client = createClientWithCert(cfg)
	} else {
		client = createClient(cfg)
	}

	return &apiImpl{
		client: client,
	}
}
