package response

import "errors"

type ResponseLine struct {
	Version []byte
	Code string
	Status []byte
}

const (
	OK                    = "200"
	NOT_FOUND             = "404"
	BAD_REQUEST           = "400"
	INTERNAL_SERVER_ERROR = "500"
)

func Status(code string) (text []byte, err error) {
	var status = map[string]string{
		OK:                    "OK",
		NOT_FOUND:             "Not Found",
		BAD_REQUEST:           "Bad Request",
		INTERNAL_SERVER_ERROR: "Internal Server Error",
	}

	if _, ok := status[code]; ok {
		text = []byte(status[code])
		return text, nil
	}

	return nil, errors.New("Invalid status code.")
}
