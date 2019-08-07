package response

import (
	"errors"
	"strconv"
)

var StatusText = map[string]string{
	strconv.Itoa(OK):                  "OK",
	strconv.Itoa(NotFound):            "Not Found",
	strconv.Itoa(BadRequest):          "Bad Request",
	strconv.Itoa(InternalServerError): "Internal Server Error",
}

func Text(code string) (text string, err error) {
	if _, ok := StatusText[code]; ok {
		text = StatusText[code]
		return text, nil
	}

	return "", errors.New("invalid status code")
}
