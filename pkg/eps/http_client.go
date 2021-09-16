package eps

import (
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"
)

// NewHttpClient builds a http client
// if the env variables PROXY or HTTP_PROXY are set
// the http client uses these
func NewHttpClient() http.Client {
	c := http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   60 * time.Second,
				KeepAlive: 60 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   20 * time.Second,
			ResponseHeaderTimeout: 20 * time.Second,
			ExpectContinueTimeout: 3 * time.Second,
		},
	}

	proxyEvnUrl := ""
	if len(os.Getenv("PROXY")) > 0 {
		log.WithField("url", os.Getenv("PROXY")).Info("use proxy")
		proxyEvnUrl = os.Getenv("PROXY")
	}
	if len(os.Getenv("HTTP_PROXY")) > 0 {
		log.WithField("url", os.Getenv("HTTP_PROXY")).Info("use proxy")
		proxyEvnUrl = os.Getenv("HTTP_PROXY")
	}
	if len(proxyEvnUrl) > 0 {
		proxyUrl, err := url.Parse(proxyEvnUrl)
		if err != nil {
			log.Fatal(err)
		}
		c = http.Client{
			Transport: &http.Transport{
				Dial: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).Dial,
				TLSHandshakeTimeout:   10 * time.Second,
				ResponseHeaderTimeout: 10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
				Proxy:                 http.ProxyURL(proxyUrl),
			},
		}
		log.Info("proxy http client")
	} else {
		log.Info("default http client")
	}

	return c
}
