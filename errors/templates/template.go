package templates

import (
	"bytes"
	"github.com/lwlwilliam/server/conf"
	"github.com/lwlwilliam/server/mime"
	"github.com/lwlwilliam/server/response"
	"html/template"
	"log"
)

var tmpl = `<html>
<head>
	<title>{{.Name}}</title>
</head>
<body>
	<h1>{{.Name}}</h1>
</body>
</html>`

type Title struct {
	Name string
}

func parse(m *response.Message, code int) error {
	var title Title
	t := template.New("template")
	t, err := t.Parse(tmpl)
	if err != nil {
		log.Println("template parse:", err)
		buildInternalServerErrorMessage(m)
		return err
	}

	buff := bytes.NewBuffer(nil)
	title.Name = response.StatusText(code)
	err = t.Execute(buff, title)
	if err != nil {
		log.Println("template execute:", err)
		buildInternalServerErrorMessage(m)
		return err
	}

	*m = response.Message{
		Line:    response.Line(code, conf.DefaultHTTPVersion),
		Headers: []string{mime.Get("html")},
		Body:    buff.String(),
	}

	return nil
}

func buildInternalServerErrorMessage(m *response.Message) {
	*m = response.Message{
		Line:    response.Line(response.InternalServerError, conf.DefaultHTTPVersion),
		Headers: []string{mime.Get("html")},
		Body:    response.StatusText(response.InternalServerError),
	}
}
