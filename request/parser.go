package request

import (
	"github.com/lwlwilliam/httpServer/response"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"
)

// 解析请求行
func parseReqLine(line string) (r response.ResponseLine, content []byte) {
	reqSlice := strings.Split(line, " ")
	if len(reqSlice) != 3 {
		//http.StatusBadRequest
		r.Code = response.BAD_REQUEST
		r.Status, _ = response.Status(r.Code)
        content = r.Status
        return
	}

	httpVerb, path := reqSlice[0], reqSlice[1]
	r.Version = []byte(reqSlice[2])

	switch httpVerb {
	case "GET":
		r.Code, r.Status, content = get(path)
	case "POST":
		//http.StatusBadRequest
		r.Code = response.BAD_REQUEST
		r.Status, _ = response.Status(r.Code)
		content = r.Status
	case "HEAD":
		//http.StatusBadRequest
		r.Code = response.BAD_REQUEST
		r.Status, _ = response.Status(r.Code)
		content = r.Status
	case "DELETE":
		//http.StatusBadRequest
		r.Code = response.BAD_REQUEST
		r.Status, _ = response.Status(r.Code)
		content = r.Status
	default:
		//http.StatusBadRequest
		r.Code = response.BAD_REQUEST
		r.Status, _ = response.Status(r.Code)
		content = r.Status
	}

	return
}

// GET 方法处理
func get(path string) (code string, status []byte, body []byte) {
	wd, _ := os.Getwd()
	//log.Printf("working directory: %s\n", wd)
	//log.Printf("request path: %s\n", path)
	log.Println(runtime.GOOS)
	if path == "/" {
		path = "index.html"
	}

	path = wd + path
	file, err := os.Open(path)
	if err != nil {
		//http.StatusNotFound
		code = response.NOT_FOUND
		status, _ = response.Status(code)
		body = status
	}

	body, err = ioutil.ReadAll(file)
	if err != nil {
		//http.StatusInternalServerError
		code = response.INTERNAL_SERVER_ERROR
		status, _ = response.Status(code)
		body = status
	} else {
		code = response.OK
		status, _ = response.Status(code)
	}
	return code, status, body
}
