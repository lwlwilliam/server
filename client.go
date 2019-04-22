package main

import (
	"fmt"
	"net"
	"os"
	"time"
	"log"
)

func main()  {
	// 阻塞 Dial
	//conn, err := net.Dial("tcp", "localhost:8888")

	// 带上超时机制的 Dial
	conn, err := net.DialTimeout("tcp", "localhost:8000", 30 * time.Second)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// read or write on conn
	conn.Write([]byte("Hello world!"))

	res := make([]byte, 1024)
	conn.Read(res)
	conn.Close()
	fmt.Println(string(res))
}
