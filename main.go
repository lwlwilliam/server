package main

import (
	"flag"
	"log"
	"net"
	"os"
	"runtime"

	error2 "github.com/lwlwilliam/server/error"
	"github.com/lwlwilliam/server/request"
)

const (
	HOST = "localhost"
	PORT = "8000"
)

func main() {
	wd, _ := os.Getwd()
	log.Printf("The program run at %s on %s.\n", wd, runtime.GOOS)

	hostname := flag.String("h", HOST, "hostname")
	port := flag.String("p", PORT, "port")
	flag.Parse()

	log.Printf("Listening %s:%s...\n", *hostname, *port)
	ld, err := net.Listen("tcp", *hostname+":"+*port)
	if err != nil {
		log.Printf("Listening: %s\n", err.Error())
		os.Exit(error2.FATAL)
	}

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