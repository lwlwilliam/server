// error code
package error

type errorCode int

const (
	FATAL   = errorCode(1)
	WARNING = errorCode(2)
)
