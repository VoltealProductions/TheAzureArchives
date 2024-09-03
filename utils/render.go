package utils

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func getTemplate(templ string) (*template.Template, error) {
	fp := path.Join("./views", fmt.Sprintf("%s.html", templ))
	return template.ParseFiles(fp)
}

func RenderHtml(w http.ResponseWriter, r *http.Request, templ string, data any) {
	templat, err := getTemplate(templ)
	if err != nil {
		RenderHttpError(w, http.StatusInternalServerError, err)
		return
	}

	if err := templat.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RenderHttpError(w http.ResponseWriter, code int, err error) {
	http.Error(w, err.Error(), code)
}
