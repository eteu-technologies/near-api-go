package jsonrpc

import (
	"net/http"
	"net/url"
)

type ClientOpts struct {
	NetworkAddr string
	HTTPClient  *http.Client
}

type ClientOptFn func(*ClientOpts) error

func WithNetworkAddr(networkAddr string) ClientOptFn {
	return func(co *ClientOpts) (err error) {
		if _, err = url.Parse(networkAddr); err != nil {
			return
		}

		co.NetworkAddr = networkAddr
		return
	}
}

func WithHTTPClient(client *http.Client) ClientOptFn {
	return func(co *ClientOpts) (err error) {
		co.HTTPClient = client
		return
	}
}
