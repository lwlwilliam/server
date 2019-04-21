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

	log.Printf("Listening %s:%s", *hostname, *port)
	ld, err := net.Listen("tcp", *hostname+":"+*port)
	error2.Fatal(err)

	for {
		conn, err := ld.Accept()
		log.Printf("Accept connection from %s.", conn.RemoteAddr().String())
		if err != nil {
			log.Println(err)
			err = conn.Close()
			log.Println(err)
			continue
		}

		go request.Handler(conn)
	}
}