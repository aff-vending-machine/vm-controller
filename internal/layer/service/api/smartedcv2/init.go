package smartedcv2

import (
	"net/http"
	"time"
)

type apiImpl struct {
	client *http.Client
}

func New() *apiImpl {
	timeout := 30 * time.Second

	client := http.Client{
		Timeout: timeout,
	}

	return &apiImpl{
		client: &client,
	}
}
