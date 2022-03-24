package rest

import (
	"net/http"
	"runtime"
)

func (h *RestHandler) Home(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"title": "Home",
		"stat": map[string]interface{}{
			"cpu":       runtime.NumCPU(),
			"goroutine": runtime.NumGoroutine(),
		},
	}

	h.tpl.ExecuteTemplate(w, "index.html", data)
}
