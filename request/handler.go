package request

import (
	"io"
	"log"
	"net"
	"strings"

	"bytes"
	"github.com/lwlwilliam/server/response"
	"strconv"
	"time"
)

func Handler(conn net.Conn) {
	log.Printf("Handling the request from %s.", conn.RemoteAddr().String())
	defer conn.Close()

	b := make([]byte, 1)
	buff := bytes.NewBuffer(nil)
	conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	for {
		n, err := conn.Read(b)
		if err != nil && err != io.EOF {
			log.Printf("Read error: %s\n", err.Error())
			break
		}

		buff.Write(b[:n])

		// 请求头结束标记
		if b[0] == '\n' && buff.Len() > 4 && string(buff.String()[buff.Len()-4:buff.Len()]) == "\r\n\r\n" {
			headers := strings.Split(buff.String()[:buff.Len()-4], "\r\n")
			hasContentLen := false
			contentLen := 0
			for _, header := range headers {
				if strings.HasPrefix(strings.ToLower(header), "content-length") {
					hasContentLen = true
					contentLen, err = strconv.Atoi(strings.TrimSpace(strings.Split(header, ":")[1]))
					if err != nil {
						log.Println("strconv.Atoi error:", err)
					}
					break
				}
			}

			if hasContentLen {
				for contentLen > 0 {
					n, err := conn.Read(b)
					if err != nil && err != io.EOF {
						log.Printf("Read error: %s\n", err.Error())
						break
					}

					buff.Write(b[:n])
					contentLen--
				}
			}

			break
		}
	}

	reqString := buff.String()

	// 解析请求行
	reqLine := strings.TrimSpace(strings.Split(reqString, "\n")[0]) // 请求行
	var m response.Message
	parseReqLine(&m, reqLine)
	log.Println(m.Version, m.Code, m.Text, m.Headers, m.Body)
	m.Response(conn)

	// 构造响应报文
	log.Printf("Has Reponsed the %s.\n", conn.RemoteAddr())
}
