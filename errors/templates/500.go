package templates

import "github.com/lwlwilliam/server/response"

func InternalServerError() string {
	return parse(response.StatusText(response.InternalServerError))
}
