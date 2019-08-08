package templates

import (
	"github.com/lwlwilliam/server/response"
)

func NotFound() string {
	return parse(response.StatusText(response.NotFound))
}
