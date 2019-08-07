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

// 请求行解释
func Line(line string) (lineS LineStruct, err error) {
	segments := strings.Split(line, " ")
	if len(segments) != 3 {
		lineS = LineStruct{"", "", "", ""}
		return lineS, errors.New("bad request line")
	}

	method := segments[0]
	HTTPVersion := segments[2]

	pathinfo := strings.Split(segments[1], "?")
	path := pathinfo[0]
	query := ""
	if len(pathinfo) == 2 {
		query = pathinfo[1]
	}

	lineS = LineStruct{method, path, query, HTTPVersion}
	err = nil

	return
}
