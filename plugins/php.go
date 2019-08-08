package plugins

import (
	"log"
)

func PHP(file string) (string, error) {
	log.Println("PHP file:", file)
	return public("php", []string{"-f", file})
}
