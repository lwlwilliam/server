package templates

import "github.com/lwlwilliam/server/response"

func BadRequest() string {
	return parse(response.StatusText(response.BadRequest))
}