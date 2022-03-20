package rest

import (
	"encoding/json"
	"net/http"
	"os"
)

func (rs *RestHandler) Health(w http.ResponseWriter, r *http.Request) {
	resp := map[string]interface{}{
		"name": os.Args[0],
		"status": map[string]string{
			"application": "OK",
		},
	}

	data, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
