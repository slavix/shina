package utils

import (
	"bytes"
	"github.com/spf13/viper"
	"html/template"
	"net/http"
	"path/filepath"
)

type HTMLData struct {
	CSRFToken string
	Path      string
}

func RenderHTML(w http.ResponseWriter, r *http.Request, page string, data *HTMLData) {

	files := []string{
		filepath.Join(viper.GetString("site.HTMLDir"), "base.html"),
		filepath.Join(viper.GetString("site.HTMLDir"), page),
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
