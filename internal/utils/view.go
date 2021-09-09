package utils

import (
	"bytes"
	"html/template"
	"net/http"
	"path/filepath"
	"shina/internal/site"
	"shina/pkg/config"
)

func RenderHTML(w http.ResponseWriter, r *http.Request, page string, data *site.HTMLData) {
	files := []string{
		filepath.Join(config.GetString("site.HTMLDir"), "base.html"),
		filepath.Join(config.GetString("site.HTMLDir"), page+".page.html"),
	}

	ts, err := template.New("").ParseFiles(files...)
	if err != nil {
		ServerError(w, err)
		return
	}

	buf := new(bytes.Buffer)

	err = ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		ServerError(w, err)
		return
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		ServerError(w, err)
		return
	}
}
