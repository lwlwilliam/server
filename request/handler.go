package request

import (
	"io"
	"log"
	"net"
	"strings"

	"github.com/lwlwilliam/httpServer/response"
)

func Handler(conn net.Conn) {
	defer conn.Close()

	tmp := make([]byte, 1024)
	//for {
		// TODO: 怎么读，读多少字节，这是个问题
		n, err := conn.Read(tmp)
		if err != nil && err != io.EOF {
			log.Printf("Read: %s\n", err.Error())
			//	break
		}
	//}

	tmpstr := string(tmp[:n])
	reqLine := strings.Split(tmpstr, "\n")[0]
	headerLine, body := requestLine(reqLine)
	respStatus := string(headerLine.Version) + " " +
		headerLine.Code + " " +
		string(headerLine.Status) + "\r\n"

	respHeader := response.TEXT_PLAIN + "\r\n"
	emptyLine := "\r\n"
	respBody := string(body)

	_, err = conn.Write([]byte(respStatus + respHeader + emptyLine + respBody))
	if err != nil {
		log.Printf("Write: %s\n", err)
	}
}
