package rest

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/tuingking/flamingo/docs"
)

func NewHttpHandler(h RestHandler) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Heartbeat("/ping"))
	mux.Use(middleware.RequestID)
	mux.Use(middleware.Logger)
	mux.Use(middleware.RealIP)

	// health
	mux.Route("/application", func(r chi.Router) {
		r.Get("/health", h.Health)
	})

	// swagger
	mux.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("./doc.json")))

	// Web
	mux.Route("/", func(r chi.Router) {
		r.Get("/", h.Home)
	})

	// API route
	mux.Route("/api", func(r chi.Router) {
		r.Route("/products", func(r chi.Router) {
			r.Get("/", h.GetAllProducts)
			r.Post("/bulk", h.BulkCreate)
		})
	})

	return mux
}
