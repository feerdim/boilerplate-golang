package util

import (
	"bytes"
	"html/template"
)

func ParseTemplateHTML(pathFile string, data interface{}) (string, error) {
	template, err := template.ParseFiles(pathFile)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)

	if err = template.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
