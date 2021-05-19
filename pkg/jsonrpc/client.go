package jsonrpc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync/atomic"
)

const JSONRPCVersion = "2.0"

type JSONRPC struct {
	JSONRPC string `json:"jsonrpc"`
	ID      string `json:"id"`
	Method  string `json:"method"`
}

type JSONRPCRequest struct {
	JSONRPC
	Params interface{} `json:"params,omitempty"`
}

type JSONRPCResponse struct {
	JSONRPC
	Error  *json.RawMessage `json:"error"`
	Result json.RawMessage  `json:"result"`
}

type JSONRPCClient struct {
	URL string

	client    *http.Client
	nextReqId uint64
}

func NewClient(networkAddr string) (client JSONRPCClient, err error) {
	_, err = url.Parse(networkAddr)
	if err != nil {
		return
	}

	client.client = new(http.Client)
	client.URL = networkAddr
	atomic.StoreUint64(&client.nextReqId, 0)

	return
}

func (c *JSONRPCClient) nextId() uint64 {
	return atomic.AddUint64(&c.nextReqId, 1)
}

func (c *JSONRPCClient) CallRPC(ctx context.Context, method string, params interface{}) (res JSONRPCResponse, err error) {
	reqId := fmt.Sprintf("%d", c.nextId())
	body, err := json.Marshal(JSONRPCRequest{
		JSONRPC{JSONRPCVersion, reqId, method},
		params,
	})
	if err != nil {
		return res, err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, c.URL, bytes.NewBuffer(body))
	if err != nil {
		return res, err
	}

	request.Header.Add("Content-Type", "application/json")

	response, err := c.client.Do(request)
	if err != nil {
		return res, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		// TODO: better error handling
		return res, fmt.Errorf("status code %d", response.StatusCode)
	}

	return parseRPCBody(response.Body)
}

func parseRPCBody(body io.Reader) (res JSONRPCResponse, err error) {
	var decodedResponse JSONRPCResponse
	if err = json.NewDecoder(body).Decode(&res); err != nil {
		return
	}

	if rerr := decodedResponse.Error; rerr != nil {
		// TODO: custom error struct
		return res, fmt.Errorf("jsonrpc error %s", *rerr)
	}

	return
}
