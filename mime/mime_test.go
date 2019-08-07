package mime

import "testing"

func TestGet(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"html", "Content-Type:text/html;charset=utf-8",
		},
		{
			"json", "Content-Type:application/json;charset=utf-8",
		},
		{
			"none", "Content-Type:text/plain;charset=utf-8",
		},
	}

	for _, tt := range tests {
		res := Get(tt.input)
		if res != tt.expected {
			t.Errorf("Get: expected=%q, got=%q", tt.expected, res)
		}
	}
}
