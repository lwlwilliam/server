package plugins

import (
	"log"
)

func Python(file string) (string, error) {
	log.Println("Python file:", file)
	return public("python", []string{file})
}
