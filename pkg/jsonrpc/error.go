package jsonrpc

import (
	"encoding/json"
	"fmt"
	"strconv"
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
	Name  string     `json:"name"`
	Cause ErrorCause `json:"cause"`

	// Legacy - do not rely on them
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

func (err *Error) Error() string {
	return fmt.Sprintf("RPC error %s (%s)", err.Name, err.Cause.String())
}

type ErrorCause struct {
	Name string          `json:"name"`
	Info json.RawMessage `json:"info"`

	message *ErrorCauseMessage
}

type ErrorCauseMessage struct {
	ErrorMessage string `json:"error_message"`
}

func (cause *ErrorCause) UnmarshalJSON(b []byte) (err error) {
	var data struct {
		Name string          `json:"name"`
		Info json.RawMessage `json:"info"`
	}

	if err = json.Unmarshal(b, &data); err != nil {
		err = fmt.Errorf("unable to unmarshal error cause: %w", err)
		return
	}

	var info map[string]interface{}
	if err = json.Unmarshal(data.Info, &info); err != nil {
		err = fmt.Errorf("unable to unmarshal error cause info: %w", err)
		return
	}

	var message *ErrorCauseMessage
	if v, ok := info["error_message"]; ok {
		message = &ErrorCauseMessage{
			ErrorMessage: v.(string),
		}
	}

	*cause = ErrorCause{
		Name:    data.Name,
		Info:    data.Info,
		message: message,
	}

	return
}

func (cause ErrorCause) String() string {
	if cause.message != nil {
		return fmt.Sprintf("name=%s, message=%s", cause.Name, cause.message.ErrorMessage)
	}
	return fmt.Sprintf("name=%s, info=%s", cause.Name, strconv.Quote(string(cause.Info)))
}
