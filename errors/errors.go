package errors

import (
	"os"
)

const (
	Error = 1
)

func Fatal(err error) {
	if err != nil {
		os.Exit(Error)
	}
}
