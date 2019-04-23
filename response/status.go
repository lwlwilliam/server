package response

import "errors"

const (
	OK                  = "200"
	NotFound            = "404"
	BadRequest          = "400"
	InternalServerError = "500"
)

var StatusText = map[string]string{
	OK:                  "OK",
	NotFound:            "Not Found",
	BadRequest:          "Bad Request",
	InternalServerError: "Internal Server Error",
}

func Text(code string) (text string, err error) {
	if _, ok := StatusText[code]; ok {
		text = StatusText[code]
		return text, nil
	}

	return "", errors.New("invalid status code")
}
