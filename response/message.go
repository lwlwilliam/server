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
func Build(line string, headers []string, body string) string {
	return fmt.Sprintf("%s%s%s%s%s%s",
		line, conf.LineFeed,
		strings.Join(headers, conf.LineFeed), conf.LineFeed,
		conf.LineFeed,
		body)
}
