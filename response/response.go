// response package defines the format of response message.
package response

import (
	"bytes"
	"net"
	"strings"
	"github.com/lwlwilliam/server/conf"
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
	respline := strings.Join([]string{m.Version, m.Code, m.Text, conf.LineFeed}, " ")

	// 响应头
	headers := strings.Join(m.Headers, "")

	// 响应体
	message := strings.Join([]string{respline, headers, conf.LineFeed, m.Body}, "")
	buf.WriteString(message)

	_, err = buf.WriteTo(c)

	return err
}
