// handle GET request
package request

import (
	"github.com/lwlwilliam/server/conf"
	"github.com/lwlwilliam/server/mime"
	"github.com/lwlwilliam/server/parser"
	"github.com/lwlwilliam/server/plugins"
	"github.com/lwlwilliam/server/response"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Get(lineStruct parser.LineStruct) response.Message {
	log.Println(lineStruct.Path, strings.TrimPrefix(lineStruct.Path, "/"))
	file := strings.TrimSuffix(conf.DocumentRoot, "/") + conf.PathSeparator +
		strings.TrimPrefix(lineStruct.Path, "/")
	ext := strings.TrimPrefix(filepath.Ext(file), ".")

	var m response.Message
	switch ext {
	case "":
		file = file + "index.html"
		fallthrough
	case "htm":
		ext = "html"
	case "jpeg":
		ext = "jpg"
	}

	log.Println("GET:", file, "ext:", ext)
	buildMessage(file, ext, lineStruct, &m)

	return m
}

// 内部错误
func internalError(lineStruct parser.LineStruct, m *response.Message) {
	m.Line = response.Line(response.InternalServerError, lineStruct.HTTPVersion)
	m.Headers = append(m.Headers, mime.Get("plain"))
	m.Body = response.StatusText(response.InternalServerError)
}

// 获取请求对应的文件
func buildMessage(file, ext string, lineStruct parser.LineStruct, m *response.Message) {
	// TODO: 有没有可能根据 fn 对应的名字来自动调用包函数
	// 动态页面
	if fn, ok := conf.Plugins[ext]; ok {
		log.Println("request PHP")
		var body string
		var err error

		switch fn {
		case "PHP":
			body, err = plugins.PHP(file)
		}

		if err != nil {
			internalError(lineStruct, m)
		} else {
			m.Line = response.Line(response.OK, lineStruct.HTTPVersion)
			m.Headers = append(m.Headers, mime.Get("html"))
			m.Body = body
		}
		return
	}

	// 获取请求文件
	fd, err := os.Open(file)
	if err != nil {
		internalError(lineStruct, m)
		return
	}

	content, err := ioutil.ReadAll(fd)
	if err != nil {
		internalError(lineStruct, m)
		return
	} else {
		m.Line = response.Line(response.OK, lineStruct.HTTPVersion)
		m.Headers = append(m.Headers, mime.Get(ext))
		m.Body = string(content)
	}
}
