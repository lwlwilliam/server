package request

import (
	"io"
	"log"
	"net"
	"strings"

	"github.com/lwlwilliam/httpServer/response"
)

func Handler(conn net.Conn) {
	log.Printf("Handling the request from %s.", conn.RemoteAddr().String())
	defer conn.Close()

	buf := make([]byte, 1024)
	//for {
		// TODO: 怎么读，读多少字节，这是个问题，暂时用一个大的 slice 确保读完所有内容吧
		n, err := conn.Read(buf)
		if err != nil && err != io.EOF {
			log.Printf("Read: %s\n", err.Error())
			//	break
		}
	//}

	reqString := string(buf[:n])

	// 解析请求行
	reqStringLine := strings.Split(reqString, "\n")[0] // 请求行
	respLine, body := parseReqLine(reqStringLine)

	// 构造响应报文
	response.Message(conn, respLine, body)
	log.Printf("Has Reponsed the %s correctly.\n", conn.RemoteAddr())
}
