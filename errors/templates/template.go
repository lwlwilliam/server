package templates

import (
	"bytes"
	"github.com/lwlwilliam/server/response"
	"html/template"
	"log"
)

var tmpl = `<html>
<head>
	<title>{{.title}}</title>
</head>
<body>
	<h3>{{.title}}</h3>
</body>
</html>`

func parse(title string) string {
	t := template.New("template")
	t, err := t.Parse(tmpl)
	if err != nil {
		log.Println("template parse:", err)
		return response.StatusText(response.InternalServerError)
	}

	buff := bytes.NewBuffer(nil)
	err = t.Execute(buff, title)
	if err != nil {
		log.Println("template parse:", err)
		return response.StatusText(response.InternalServerError)
	}

	return buff.String()
}
