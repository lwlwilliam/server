package request

import (
	"bytes"
	"log"
	"net"
	"os"
)

func Handler(conn net.Conn) {
	defer conn.Close()
	tmp := make([]byte, 1)
	buf := bytes.NewBuffer(nil)

	for {
		// TODO: 怎么读，读多少字节，这是个问题
		_, err := conn.Read(tmp)
		if err != nil {
			log.Printf("conn.Read: %s", err.Error())
			break
		}
		buf.Write(tmp)
	}

	_, err := buf.WriteTo(os.Stdout)
	if err != nil {
		log.Println("buf.WriteTo:", err.Error())
	}
}