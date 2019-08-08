package response

import (
	"github.com/lwlwilliam/server/conf"
	"testing"
)

func TestLine(t *testing.T) {
	tests := []struct {
		code        int
		httpVersion string
		expected    string
	}{
		{
			200, conf.DefaultHTTPVersion, "HTTP/1.1 200 OK",
		},
		{
			404, conf.DefaultHTTPVersion, "HTTP/1.1 404 Not Found",
		},
	}

	for _, tt := range tests {
		res := Line(tt.code, tt.httpVersion)
		if res != tt.expected {
			t.Errorf("Get: expected=%q, got=%q.", tt.expected, res)
		}
	}
}
