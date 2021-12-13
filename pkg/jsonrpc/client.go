package jsonrpc

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sync/atomic"
)

const JSONRPCVersion = "2.0"

type Client struct {
	URL string

	client    *http.Client
	nextReqId uint64
}

func NewClient(networkAddr string) (client Client, err error) {
	_, err = url.Parse(networkAddr)
	if err != nil {
		return
	}

	client.client = new(http.Client)
	client.URL = networkAddr
	atomic.StoreUint64(&client.nextReqId, 0)

	return
}

func (c *Client) nextId() uint64 {
	return atomic.AddUint64(&c.nextReqId, 1)
}

func (c *Client) CallRPC(ctx context.Context, method string, params interface{}) (res Response, err error) {
	reqId := fmt.Sprintf("%d", c.nextId())
	body, err := json.Marshal(Request{
		JSONRPC{JSONRPCVersion, reqId, method},
		params,
	})
	if err != nil {
		return
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, c.URL, bytes.NewBuffer(body))
	if err != nil {
		return
	}

	request.Header.Add("Content-Type", "application/json")

	response, err := c.client.Do(request)
	if err != nil {
		return
	}

	return parseRPCBody(response)
}

func parseRPCBody(r *http.Response) (res Response, err error) {
	//fmt.Printf("%#v\n", r)

	body := r.Body
	if body == nil {
		err = errors.New("nil body")
		return
	}
	defer func() { _ = body.Close() }()

	// TODO: check for Content-Type header
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()

	if err = decoder.Decode(&res); err != nil {
		return
	}

	return
}
