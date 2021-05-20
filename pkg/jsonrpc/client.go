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

const (
	CodeParseError     = -32700
	CodeInvalidRequest = -32600
	CodeMethodNotFound = -32601
	CodeInvalidParams  = -32602
	CodeInternalError  = -32603

	CodeServerErrorRangeStart = -32099
	CodeServerErrorRangeEnd   = -32000
)

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
	Error  *JSONRPCError   `json:"error"`
	Result json.RawMessage `json:"result"`
}

type JSONRPCError struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

func (err JSONRPCError) Error() string {
	return fmt.Sprintf("JSON-RPC error '%s' (%d) %s", err.Message, err.Code, string(err.Data))
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

	return parseRPCBody(response.Body)
}

func parseRPCBody(body io.ReadCloser) (res JSONRPCResponse, err error) {
	defer func() { _ = body.Close() }()
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()

	if err = decoder.Decode(&res); err != nil {
		return
	}

	return
}
