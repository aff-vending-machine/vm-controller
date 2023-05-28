package http

import (
	"net/http"
	"time"

	"vm-controller/configs"
)

func createClient(cfg configs.HTTPConfig) *http.Client {
	timeout := time.Duration(cfg.TimeoutInSec) * time.Second

	return &http.Client{
		Timeout: timeout,
	}
}
