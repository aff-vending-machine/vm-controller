package link2500

import (
	"net/http"
)

type apiImpl struct {
	client *http.Client
}

const LINK2500_PATH = "/api/v1/link2500"

func New(client *http.Client) *apiImpl {
	return &apiImpl{
		client: client,
	}
}
