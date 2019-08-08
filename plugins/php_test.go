package plugins

import (
	"github.com/lwlwilliam/server/conf"
	"strings"
	"testing"
)

func TestPHP(t *testing.T) {
	dir := strings.TrimSuffix(conf.DocumentRoot, conf.PathSeparator) + conf.PathSeparator
	tests := []struct {
		file     string
		expected string
		err      error
	}{
		{
			dir + "php_test.php",
			"Hello world!",
			nil,
		},
	}

	for _, tt := range tests {
		got, err := PHP(tt.file)
		if err != tt.err {
			t.Errorf("PHP: expected=%q, got=%q", tt.err, err)
		}

		if got != tt.expected {
			t.Errorf("PHP: expected=%q, got=%q", tt.expected, got)
		}
	}
}
