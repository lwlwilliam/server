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
