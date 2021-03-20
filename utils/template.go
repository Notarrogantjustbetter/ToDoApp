package utils

import (
	"net/http"
	"html/template"
)

var tmpl *template.Template

func LoadTemplate() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func ExecuteTemplate(w http.ResponseWriter, templ string, data interface{}) {
	tmpl.ExecuteTemplate(w, templ, data)
}