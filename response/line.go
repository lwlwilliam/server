package response

import (
	"fmt"
	"strconv"
)

// 构造响应行
func Line(code int, HTTPVersion string) string {
	return fmt.Sprintf("%s %s %s",
		HTTPVersion,
		strconv.Itoa(code),
		status[code])
}
