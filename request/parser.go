package request

import (
	"github.com/lwlwilliam/server/response"
	"github.com/lwlwilliam/server/server"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
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
		case "php":
			m.Headers = append(m.Headers, response.ContentType["html"])
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

	/***************************************************/
	// TODO: 暂时先用这种简陋的方法处理 php（在接收到 php 文件请求时，调用 php 解释器，接收解释器运行结果，再返回给浏览器），先理解大概运行机制，以后慢慢完善吧
	segments := strings.Split(path, "/")
	for _, seg := range segments {
		if strings.HasSuffix(seg, ".php") {
			log.Println("PHP file:", wd+string(os.PathSeparator)+seg)
			cmd := exec.Command("php", "-f", wd+string(os.PathSeparator)+seg)
			output, err := cmd.CombinedOutput()
			if err != nil {
				log.Println("PHP error:", err)
				code = response.NotFound
				text, _ = response.Text(code)
				body = text
				return
			}

			log.Printf("PHP output: %s\n", output)

			code = response.OK
			text, _ = response.Text(code)
			body = string(output)
			return
		}
	}
	/***************************************************/

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
