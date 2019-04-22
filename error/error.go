// error function
package error

import (
	"log"
	"os"
)

func Fatal(extra string, err error) {
	errHandler(extra, err)
	os.Exit(int(FATAL))
}

func Warning(extra string, err error) {
	errHandler(extra, err)
}

func errHandler(extra string, err error) {
	if err != nil {
		log.Printf("%s: %s\n", extra, err.Error())
	}
}
