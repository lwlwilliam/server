package response

import (
	"fmt"
	"github.com/lwlwilliam/server/conf"
	"strings"
)

type Message struct {
	Line    string
	Headers []string
	Body    string
}

// 构造报文
func (m *Message) Build() string {
	return fmt.Sprintf("%s%s%s%s%s%s",
		m.Line, conf.LineFeed,
		strings.Join(m.Headers, conf.LineFeed), conf.LineFeed,
		conf.LineFeed,
		m.Body)
}
