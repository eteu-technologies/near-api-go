package jsonrpc

import (
	"encoding/json"
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
