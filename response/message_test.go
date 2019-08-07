package response

import "testing"

func TestBuild(t *testing.T) {
	tests := []struct {
		line     string
		headers  []string
		body     string
		expected string
	}{
		{
			"HTTP/1.1 200 OK",
			[]string{"Content-Type:text/html;charset=utf-8"},
			"Hello world!",
			"HTTP/1.1 200 OK\r\nContent-Type:text/html;charset=utf-8\r\n\r\nHello world!",
		},
	}

	for _, tt := range tests {
		got := Build(tt.line, tt.headers, tt.body)
		if got != tt.expected {
			t.Errorf("Build: expected=%q, got=%q", tt.expected, got)
		}
	}
}
