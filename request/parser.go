package request

import (
	"github.com/lwlwilliam/server/response"
	"github.com/lwlwilliam/server/server"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// 解析请求行
func parseReqLine(m *response.Message, line string) (err error) {
	linePart := strings.Split(line, " ")
	m.Headers = []string{
		server.Version,
		//response.ContentType["plain"],
	}

	if len(linePart) != 3 {
		//http.StatusBadRequest
		m.Version = response.HTTPVersion
		m.Code = response.BadRequest
		m.Text, _ = response.Text(m.Code)
		m.Headers = append(m.Headers, response.ContentType["plain"])
		m.Body = m.Text
		return nil
	}

	httpVerb, path := linePart[0], linePart[1]
	m.Version = linePart[2]

	switch httpVerb {
	case "GET":
		m.Code, m.Text, m.Body = get(path)

		if (path == "/" || strings.HasSuffix(path, ".html")) && m.Code == response.OK {
			m.Headers = append(m.Headers, response.ContentType["html"])
		} else {
			m.Headers = append(m.Headers, response.ContentType["plain"])
		}

	case "POST":
		fallthrough
	case "HEAD":
		fallthrough
	case "DELETE":
		fallthrough
	default:
		//http.StatusBadRequest
		m.Code = response.BadRequest
		m.Text, _ = response.Text(m.Code)
		m.Headers = append(m.Headers, response.ContentType["plain"])
		m.Body = m.Text
	}

	return
}

// GET 方法处理
func get(path string) (code string, text string, body string) {
	var wd string
	if server.Htdocs == "" {
		wd, _ = os.Getwd()
	} else {
		wd = server.Htdocs
	}

	if path == "/" {
		path = "/index.html"
	}

	path = wd + path
	log.Printf("GET %s\n", path)

	file, err := os.Open(path)
	if err != nil {
		//http.StatusNotFound
		code = response.NotFound
		text, _ = response.Text(code)
		body = text
		return
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		//http.StatusInternalServerError
		code = response.InternalServerError
		text, _ = response.Text(code)
		body = text
	} else {
		code = response.OK
		text, _ = response.Text(code)
		body = string(content)
	}
	return code, text, body
}
