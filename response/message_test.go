package response

import "testing"

func TestMessageBuild(t *testing.T) {
	tests := []struct {
		input    Message
		expected string
	}{
		{
			struct {
				Line    string
				Headers []string
				Body    string
			}{
				Line:    "HTTP/1.1 200 OK",
				Headers: []string{"Content-Type:text/html;charset=utf-8"},
				Body:    "Hello world!",
			},
			"HTTP/1.1 200 OK\r\nContent-Type:text/html;charset=utf-8\r\n\r\nHello world!",
		},
	}

	for _, tt := range tests {
		got := tt.input.Build()
		if got != tt.expected {
			t.Errorf("Build: expected=%q, got=%q", tt.expected, got)
		}
	}
}
