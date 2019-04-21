// error function
package error

import (
	"log"
	"os"
)

func Fatal(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(FATAL)
	}
}

func Warning(err error) {
	if err != nil {
		log.Fatal(WARNING)
	}
}
