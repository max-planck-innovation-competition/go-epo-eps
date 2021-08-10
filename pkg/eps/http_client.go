package eps

import (
	"net/http"
	"time"
)

func NewHttpClient() http.Client {
	return http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			TLSHandshakeTimeout: 20 * time.Second,
		},
	}
}
