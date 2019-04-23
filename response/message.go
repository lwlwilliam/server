// 生成响应报文
package response

import (
	"net"
	"bytes"
	"log"
)

// 根据请求行构造响应报文
func Message(conn net.Conn, respLine ResponseLine, body []byte)  {
	buf := bytes.NewBuffer(nil)
	log.Println(respLine)

	// 响应行
	//buf.WriteString(string(respLine.Version))	// TODO: 不知道什么原因，获取到的会多了一个字符
	buf.WriteString("HTTP/1.1")
	buf.Write([]byte(" "))
	buf.WriteString(respLine.Code)
	buf.Write([]byte(" "))
	buf.Write(respLine.Status)
	buf.WriteString("\n")

	// 响应头
	buf.WriteString(TEXT_PLAIN)
	buf.WriteString("\n\n")

	// 响应体
	buf.Write(body)

	buf.WriteTo(conn)
}
