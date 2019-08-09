package templates

import "github.com/lwlwilliam/server/response"

func InternalServerError(m *response.Message) {
	parse(m, response.InternalServerError)
}
