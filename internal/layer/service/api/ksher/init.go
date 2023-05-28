package ksher

import (
	"net/http"
)

type apiImpl struct {
	client *http.Client
}

const CSCANB_PATH = "/api/v1/cscanb/orders"

func New(client *http.Client) *apiImpl {
	return &apiImpl{
		client: client,
	}
}
