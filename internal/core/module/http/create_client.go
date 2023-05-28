package http

import (
	"net/http"
	"time"

	"github.com/aff-vending-machine/vm-controller/config"
)

func createClient(cfg config.HTTPConfig) *http.Client {
	timeout := time.Duration(cfg.TimeoutInSec) * time.Second

	return &http.Client{
		Timeout: timeout,
	}
}
