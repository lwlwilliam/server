package main

import (
	"flag"
	"log"
	"net"
	"os"
	"runtime"

	"github.com/lwlwilliam/server/errors"
	"github.com/lwlwilliam/server/request"
)

const (
	HOST = "localhost"
	PORT = "8000"
)

func main() {
	// working directory
	wd, _ := os.Getwd()
	log.Printf("The program run at %s on %s.\n", wd, runtime.GOOS)

	// flags
	hostname := flag.String("h", HOST, "hostname")
	port := flag.String("p", PORT, "port")
	flag.Parse()

	// listen
	log.Printf("Listening %s:%s...\n", *hostname, *port)
	ld, err := net.Listen("tcp", *hostname+":"+*port)
	if err != nil {
		log.Printf("Listening: %s\n", err.Error())
		os.Exit(errors.Error)
	}

	// accept
	log.Println("Accepting connections...")
	for {
		conn, err := ld.Accept()
		if err != nil {
			log.Printf("Accept: %s\n", err.Error())
			err = conn.Close()
			continue
		}

		go request.Handler(conn)
	}
}
