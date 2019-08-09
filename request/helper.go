package request

import (
	"os"
)

func FileNotExist(file string) bool {
	_, err := os.Stat(file)
	return os.IsNotExist(err)
}
