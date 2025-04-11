package http

import (
	"article/pkg/response"
)

// ParseErrorCode ...
func ParseErrorCode(err string) response.Response {
	errResp := response.Error{}

	errResp.Msg = errResp.Msg + " | " + err

	return response.Response{
		Error: errResp,
	}
}
