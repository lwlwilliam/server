package main

import (
	"flag"
	"log"
	"net"

	error2 "github.com/lwlwilliam/httpServer/error"
	"github.com/lwlwilliam/httpServer/request"
)

func main() {
	hostname := flag.String("h", "localhost", "hostname")
	port := flag.String("p", "8080", "port")

	ld, err := net.Listen("tcp", *hostname+":"+*port)
	error2.Fatal("Listening", err)

	for {
		conn, err := ld.Accept()
		if err != nil {
			log.Println(err)
			err = conn.Close()
			log.Println(err)
			continue
		}

		go request.Handler(conn)
	}
}