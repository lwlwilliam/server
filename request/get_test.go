package request

import (
	"github.com/lwlwilliam/server/parser"
	"github.com/lwlwilliam/server/response"
	"testing"
)

func TestGet(t *testing.T) {
	tests := []struct {
		input    parser.LineStruct
		expected response.Message
	}{
		{
			parser.LineStruct{
				Method:      "GET",
				Path:        "/InternalServerError",
				Query:       "",
				HTTPVersion: "HTTP/1.1",
			},
			response.Message{
				Line:    "HTTP/1.1 500 Internal Server Error",
				Headers: []string{"Content-Type:text/plain;charset=utf-8"},
				Body:    "Internal Server Error",
			},
		},
	}

	for _, tt := range tests {
		got := Get(tt.input)
		if got.Line != tt.expected.Line {
			t.Errorf("Get: expected=%q, got=%q", tt.expected.Line, got.Line)
		}

		if !testHeaders(got.Headers, tt.expected.Headers) {
			t.Errorf("Get: expected=%q, got=%q", tt.expected.Headers, got.Headers)
		}

		if got.Body != tt.expected.Body {
			t.Errorf("Get: expected=%q, got=%q", tt.expected.Body, got.Body)
		}
	}
}

func testHeaders(got []string, expected []string) bool {
	equal := true

	for idx, gotHeader := range got {
		if gotHeader != expected[idx] {
			equal = false
			break
		}
	}

	return equal
}
