package templates

import "github.com/lwlwilliam/server/response"

func BadRequest(m *response.Message) {
	parse(m, response.BadRequest)
}
