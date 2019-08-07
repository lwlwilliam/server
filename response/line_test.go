package response

import "testing"

func TestLine(t *testing.T) {
	tests := []struct {
		code        int
		httpVersion string
		expected    string
	}{
		{
			200, "HTTP/1.1", "HTTP/1.1 200 OK",
		},
		{
			404, "HTTP/1.1", "HTTP/1.1 404 Not Found",
		},
	}

	for _, tt := range tests {
		res := Line(tt.code, tt.httpVersion)
		if res != tt.expected {
			t.Errorf("Get: expected=%q, got=%q.", tt.expected, res)
		}
	}
}
