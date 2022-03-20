package home

import (
	"html/template"
	"net/http"
)

type Web struct {
	tpl *template.Template
}

func NewWeb(tpl *template.Template) Web {
	return Web{
		tpl: tpl,
	}
}

func (web *Web) Index(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"title": "Home",
	}

	web.tpl.ExecuteTemplate(w, "index.html", data)
}
