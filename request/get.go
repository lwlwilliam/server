// handle GET request
package request

import (
	"github.com/lwlwilliam/server/conf"
	"github.com/lwlwilliam/server/errors/templates"
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
	file := strings.TrimSuffix(conf.DocumentRoot, "/") + conf.PathSeparator +
		strings.TrimPrefix(lineStruct.Path, "/")
	ext := strings.TrimPrefix(filepath.Ext(file), ".")

	var m response.Message
	switch ext {
	case "":
		if file == strings.TrimSuffix(conf.DocumentRoot, "/")+conf.PathSeparator {
			file = file + "index.html"
		}

		fallthrough
	case "htm":
		ext = "html"
	case "jpeg":
		ext = "jpg"
	}

	log.Println("GET:", file, "; ext:", ext)
	buildMessage(file, ext, lineStruct, &m)

	return m
}

// 获取请求对应的文件
func buildMessage(file, ext string, lineStruct parser.LineStruct, m *response.Message) {
	if FileNotExist(file) {
		templates.NotFound(m)
		return
	}

	if lineStruct.Method == "" || lineStruct.Path == "" || lineStruct.HTTPVersion == "" {
		templates.BadRequest(m)
		return
	}

	// TODO: 有没有可能根据 fn 对应的名字来自动调用包函数
	// 动态页面
	if fn, ok := conf.Plugins[ext]; ok {
		log.Println("Request Plugins")
		var body string
		var err error

		switch fn {
		case "PHP":
			body, err = plugins.PHP(file)
		case "Python":
			body, err = plugins.Python(file)
		}

		if err != nil {
			templates.InternalServerError(m)
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
		templates.InternalServerError(m)
		return
	}

	content, err := ioutil.ReadAll(fd)
	if err != nil {
		templates.InternalServerError(m)
		return
	} else {
		m.Line = response.Line(response.OK, lineStruct.HTTPVersion)
		m.Headers = append(m.Headers, mime.Get(ext))
		m.Body = string(content)
	}
}
