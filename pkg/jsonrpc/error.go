package jsonrpc

import (
	"encoding/json"
	"fmt"
)

const (
	CodeParseError     = -32700
	CodeInvalidRequest = -32600
	CodeMethodNotFound = -32601
	CodeInvalidParams  = -32602
	CodeInternalError  = -32603

	CodeServerErrorRangeStart = -32099
	CodeServerErrorRangeEnd   = -32000
)

type Error struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

func (err Error) Error() string {
	return fmt.Sprintf("JSON-RPC error '%s' (%d) %s", err.Message, err.Code, string(err.Data))
}
