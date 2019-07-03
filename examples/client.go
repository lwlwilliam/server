package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	hostname := flag.String("h", "localhost", "hostname")
	port := flag.String("p", "8000", "port")
	flag.Parse()

	// 带上超时机制的 Dial
	conn, err := net.DialTimeout("tcp", *hostname+":"+*port,
		30*time.Second)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	conn.Write([]byte("Hello world!"))

	res := make([]byte, 1024)
	conn.Read(res)
	conn.Close()
	fmt.Println(string(res))
}
