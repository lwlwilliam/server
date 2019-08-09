package templates

import (
	"github.com/lwlwilliam/server/response"
)

func NotFound(m *response.Message) {
	parse(m, response.NotFound)
}
