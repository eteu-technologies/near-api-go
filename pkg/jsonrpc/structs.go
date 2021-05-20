package jsonrpc

import (
	"encoding/json"
)

type JSONRPC struct {
	JSONRPC string `json:"jsonrpc"`
	ID      string `json:"id"`
	Method  string `json:"method"`
}

type Request struct {
	JSONRPC
	Params interface{} `json:"params,omitempty"`
}

type Response struct {
	JSONRPC
	Error  *Error          `json:"error"`
	Result json.RawMessage `json:"result"`
}
