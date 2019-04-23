// response package defines the format of response message.
package response

import (
	"bytes"
	"net"
	"strings"
)

// 响应报文类型
type Message struct {
	Version string
	Code    string
	Text    string
	Headers []string
	Body    string
}

// 响应
func (m *Message) Response(c net.Conn) (err error) {
	buf := bytes.NewBuffer(nil)

	// 响应行
	// TODO: 不知道什么原因，解析的报文请求头版本会多了一个字符
	m.Version = HTTPVersion
	respline := strings.Join([]string{m.Version, m.Code, m.Text, Linefeed}, " ")

	// 响应头
	headers := strings.Join(m.Headers, "")

	message := strings.Join([]string{respline, headers, Linefeed, m.Body}, "")
	buf.WriteString(message)
	_, err = buf.WriteTo(c)

	return err
}
