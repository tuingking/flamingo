package rest

import (
	"net/http"
)

func (rs *RestHandler) Home(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"title": "Home",
	}

	rs.tpl.ExecuteTemplate(w, "index.html", data)
}
