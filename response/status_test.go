package response

import "testing"

func TestStatusText(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{
			1000, "Bad Request",
		},
		{
			200, "OK",
		},
		{
			400, "Bad Request",
		},
	}

	for _, tt := range tests {
		res := StatusText(tt.input)
		if res != tt.expected {
			t.Errorf("StatusText: expected=%q, got=%q", tt.expected, res)
		}
	}
}
