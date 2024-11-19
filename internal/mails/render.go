package mails

import (
	"bytes"
	"fmt"
	"html/template"
	"reflect"

	"github.com/studentkickoff/gobp/pkg/config"
	"github.com/studentkickoff/gobp/pkg/mjml"
)

var layoutTmpl *template.Template

func init() {
	var err error
	layoutTmpl, err = template.ParseFiles("views/mail/layouts/layout.mjml")
	if err != nil {
		panic(err)
	}
}

func renderHTMLMail(name string, data interface{}) (string, error) {
	tmpl, err := template.Must(layoutTmpl.Clone()).ParseFiles(fmt.Sprintf("views/mail/%s.mjml", name))
	if err != nil {
		return "", err
	}

	// add assetHo
	if data == nil {
		data = struct{ AssetHost string }{AssetHost: config.GetDefaultString("app.host", "http://localhost:3000")}
	} else {
		stype := reflect.ValueOf(data).Elem()
		field := stype.FieldByName("AssetHost")
		if field.IsValid() {
			field.SetString(config.GetDefaultString("app.host", "http://localhost:3000"))
		}
	}

	tmplWriter := new(bytes.Buffer)
	err = tmpl.Execute(tmplWriter, data)
	if err != nil {
		return "", err
	}
	mjmlStr := tmplWriter.String()

	return mjml.Convert(mjmlStr)
}
