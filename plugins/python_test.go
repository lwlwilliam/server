package plugins

import (
	"github.com/lwlwilliam/server/conf"
	"strings"
	"testing"
)

func TestPython(t *testing.T) {
	dir := strings.TrimSuffix(conf.DocumentRoot, conf.PathSeparator) + conf.PathSeparator
	tests := []struct {
		file     string
		expected string
		err      error
	}{
		{
			dir + "test.py",
			"<h1>Hello Python!</h1>",
			nil,
		},
	}

	for _, tt := range tests {
		got, err := Python(tt.file)
		if err != tt.err {
			t.Errorf("Python: expected=%q, got=%q", tt.err, err)
		}

		if got != tt.expected {
			t.Errorf("Python: expected=%q, got=%q", tt.expected, got)
		}
	}
}
