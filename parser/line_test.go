package parser

import "testing"

func TestLine(t *testing.T) {
	tests := []struct {
		input       string
		method      string
		path        string
		query       string
		HTTPVersion string
		err         error
	}{
		{
			"GET / HTTP/1.1",
			"GET",
			"/",
			"",
			"HTTP/1.1",
			nil,
		},
		{
			"GET /hello?name=william HTTP/1.1",
			"GET",
			"/hello",
			"name=william",
			"HTTP/1.1",
			nil,
		},
	}

	for _, tt := range tests {
		lineS, err := Line(tt.input)
		if err != nil {
			t.Errorf("Line error: %s", err.Error())
		}

		if lineS.Method != tt.method || lineS.Path != tt.path ||
			lineS.Query != tt.query || lineS.HTTPVersion != tt.HTTPVersion {
			t.Errorf("Line: expected=%q, %q, %q, %q; got=%q, %q, %q, %q",
				tt.method, tt.path, tt.query, tt.HTTPVersion,
				lineS.Method, lineS.Path, lineS.Query, lineS.HTTPVersion)
		}
	}
}
