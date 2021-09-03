package utils

import (
	"bytes"
	"github.com/spf13/viper"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"shina/internal/site"
)

func RenderHTML(w http.ResponseWriter, r *http.Request, page string, data *site.HTMLData) {

	log.Println(viper.GetString("site.HTMLDir"))

	files := []string{
		filepath.Join(viper.GetString("site.HTMLDir"), "base.html"),
		filepath.Join(viper.GetString("site.HTMLDir"), page+".page.html"),
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
