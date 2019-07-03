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

	httpVerb, url := linePart[0], linePart[1]
	m.Version = linePart[2]

	switch httpVerb {
	case "GET":
		m.Code, m.Text, m.Body = get(url)

		path := strings.Split(url, "?")[0]
		seg := strings.Split(path, ".")
		var ext string
		if len(seg) == 2 {
			ext = seg[1]
		}

		switch ext {
		case "":
		case "html":
			fallthrough
		case "htm":
			m.Headers = append(m.Headers, response.ContentType["html"])
		case "png":
			m.Headers = append(m.Headers, response.ContentType["png"])
		case "jpg":
			fallthrough
		case "jpeg":
			m.Headers = append(m.Headers, response.ContentType["jpeg"])
		case "gif":
			m.Headers = append(m.Headers, response.ContentType["gif"])
		case "json":
			m.Headers = append(m.Headers, response.ContentType["json"])
		default:
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
func get(url string) (code string, text string, body string) {
	var wd string
	path := strings.Split(url, "?")[0]

	if server.DocumentRoot == "" {
		wd, _ = os.Getwd()
	} else {
		wd = server.DocumentRoot
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
