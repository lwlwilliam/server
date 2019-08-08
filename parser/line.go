package parser

import (
	"errors"
	"strings"
)

type LineStruct struct {
	Method      string
	Path        string
	Query       string
	HTTPVersion string
}

// 解释请求行
func RequestLine(line string) (lineS LineStruct, err error) {
	line = strings.TrimSpace(line)
	segments := strings.Split(line, " ")
	if len(segments) != 3 {
		lineS = LineStruct{"", "", "", ""}
		return lineS, errors.New("bad request line")
	}

	pathinfo := strings.Split(segments[1], "?")
	query := ""
	if len(pathinfo) == 2 {
		query = pathinfo[1]
	}

	lineS = LineStruct{
		Method:      segments[0],
		Path:        pathinfo[0],
		Query:       query,
		HTTPVersion: segments[2],
	}
	err = nil
	return
}
