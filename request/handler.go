package request

import (
	"github.com/lwlwilliam/server/conf"
	"github.com/lwlwilliam/server/errors"
	"github.com/lwlwilliam/server/parser"
	"io"
	"log"
	"net"
	"strings"

	"bytes"
	"github.com/lwlwilliam/server/errors/templates"
	"github.com/lwlwilliam/server/response"
	"strconv"
	"time"
)

func Handler(conn net.Conn) {
	log.Printf("Handling the request from %s.", conn.RemoteAddr().String())
	defer conn.Close()

	// 设置读缓冲以及超时
	b := make([]byte, 1)
	buff := bytes.NewBuffer(nil)
	err := conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	errors.Fatal(err)

	for {
		n, err := conn.Read(b)
		if err != nil && err != io.EOF {
			log.Printf("request.Handler Read error: %s\n", err.Error())
			break
		}

		buff.Write(b[:n])

		// 请求头结束标记
		if b[0] == '\n' &&
			buff.Len() > 4 && // 这是为了确保下面的 buff.Len() - 4 不会溢出
			string(buff.String()[buff.Len()-4:buff.Len()]) == (conf.LineFeed+conf.LineFeed) { // 读取已有缓冲的最后 4 个字节

			headers := strings.Split(buff.String()[:buff.Len()-4], conf.LineFeed)

			// 获取请求报文长度
			hasContentLen := false
			contentLen := 0
			for _, header := range headers {

				header = strings.ToLower(header)
				if strings.HasPrefix(header, "content-length:") {
					hasContentLen = true

					contentLenStr := strings.TrimSpace(strings.Split(header, ":")[1])
					contentLen, err = strconv.Atoi(contentLenStr)

					if err != nil {
						log.Println("request.Handler Atoi error:", err)
					}

					break
				}
			}

			// 如果报文标明了长度，则继续读取 contentLen 个字节
			if hasContentLen {
				for contentLen > 0 {
					n, err := conn.Read(b)
					if err != nil && err != io.EOF {
						log.Printf("request.Handler Read error: %s\n", err.Error())
						break
					}

					buff.Write(b[:n])
					contentLen--
				}
			}

			break
		}
	}

	// 完整的请求报文
	requestMessage := buff.String()
	var message response.Message

	// 解析请求行
	requestLine := strings.TrimSpace(strings.Split(requestMessage, conf.LineFeed)[0])
	lineStruct, err := parser.RequestLine(requestLine)
	if err != nil {
		log.Println("Request line:", requestLine)
		log.Println("Parse request line:", err)
		log.Printf("Request line struct: %s, %s, %s\n", lineStruct.Method, lineStruct.Path, lineStruct.HTTPVersion)
	}

	switch lineStruct.Method {
	case "GET":
		message = Get(lineStruct)
	case "POST":
		fallthrough
	case "PUT":
		fallthrough
	case "DELETE":
		fallthrough
	case "HEAD":
		fallthrough
	default:
		templates.BadRequest(&message)
	}

	messageStr := message.Build()
	//log.Println("Response message:", messageStr)
	conn.Write([]byte(messageStr))

	log.Printf("Has Reponsed the %s.\n", conn.RemoteAddr())
}
