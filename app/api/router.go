package api

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/tuingking/flamingo/internal/home"
	"github.com/tuingking/flamingo/internal/product"
)

type Handler struct {
	Product product.API
	Home    home.Web
}

func NewRouter(h Handler) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Heartbeat("/ping"))
	mux.Use(middleware.RequestID)
	mux.Use(middleware.Logger)
	mux.Use(middleware.RealIP)

	// health
	mux.Route("/application", func(r chi.Router) {
		r.Get("/health", health)
	})

	// Web
	mux.Get("/", h.Home.Index)

	// API route
	mux.Route("/api", func(r chi.Router) {
		r.Route("/products", func(r chi.Router) {
			r.Get("/", h.Product.GetAllProducts)
		})
	})

	return mux
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world")
}

func health(w http.ResponseWriter, r *http.Request) {
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
