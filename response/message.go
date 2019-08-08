package response

import (
	"fmt"
	"github.com/lwlwilliam/server/conf"
	"strconv"
	"strings"
	"time"
)

type Message struct {
	Line    string
	Headers []string
	Body    string
}

// 构造报文
func (m *Message) Build() string {
	date := "Date: " + time.Now().Format(time.RFC1123)
	contentLen := "Content-Length: " + strconv.Itoa(len(m.Body))

	m.Headers = append(m.Headers,
		conf.Server,
		date,
		contentLen)

	return fmt.Sprintf("%s%s%s%s%s%s",
		m.Line, conf.LineFeed,
		strings.Join(m.Headers, conf.LineFeed), conf.LineFeed,
		conf.LineFeed,
		m.Body)
}
