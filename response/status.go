package response

const (
	OK                  = 200
	BadRequest          = 400
	NotFound            = 404
	InternalServerError = 500
)

var status = map[int]string{
	OK:                  "OK",
	BadRequest:          "Bad Request",
	NotFound:            "Not Found",
	InternalServerError: "Internal Server Error",
}

// 获取状态短语
func StatusText(code int) string {
	if text, ok := status[code]; ok {
		return text
	}

	return status[BadRequest]
}
