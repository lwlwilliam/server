package request

import (
	"bufio"
	"log"
	"net"
)

func Handler(conn net.Conn) {
	defer conn.Close()
	tmp := make([]byte, 1024)
	writer := bufio.NewWriter(conn)
	for {
		// TODO: 怎么读，读多少字节，这是个问题
		n, err := conn.Read(tmp)
		if err != nil {
			log.Printf("conn.Read: %s", err.Error())
			break
		}
		writer.Write(tmp[:n])
	}

	writer.Flush()
}
