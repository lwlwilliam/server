package main

import (
	"flag"
	"log"
	"net"
	"os"

	error2 "github.com/lwlwilliam/httpServer/error"
	"github.com/lwlwilliam/httpServer/request"
)

const (
	HOST = "localhost"
	PORT = "8000"
)

func main() {
	hostname := flag.String("h", HOST, "hostname")
	port := flag.String("p", PORT, "port")

	ld, err := net.Listen("tcp", *hostname+":"+*port)
	if err != nil {
		log.Printf("Listening: %s\n", err.Error())
		os.Exit(error2.FATAL)
	}

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
