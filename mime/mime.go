// retrieve mime type
package mime

import (
	"fmt"
	"github.com/lwlwilliam/server/conf"
)

var contentType = map[string]string{
	"plain": "text/plain",
	"html":  "text/html",
	"png":   "image/png",
	"jpg":   "image/jpeg",
	"gif":   "image/gif",
	"json":  "application/json",
}

// Content-Type 响应头
func Get(ext string) string {
	if tp, ok := contentType[ext]; ok {
		return fmt.Sprintf("Content-Type:%s;charset=%s",
			tp, conf.DefaultCharset)
	}

	return fmt.Sprintf("Content-Type:%s;charset=%s",
		contentType["plain"], conf.DefaultCharset)
}
