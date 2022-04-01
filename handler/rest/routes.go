package rest

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/tuingking/flamingo/docs"
)

func NewHttpHandler(h RestHandler) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// health
	r.Route("/application", func(r chi.Router) {
		r.Get("/health", h.Health)
	})

	// swagger
	r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("./doc.json")))

	// Web
	r.Route("/", func(r chi.Router) {
		r.Get("/", h.Home)
	})

	// public API
	r.Post("/api/auth/token", h.IssueAccessToken)
	r.Post("/api/account", h.RegisterNewAccount)

	// protected API
	r.Mount("/api", privateRoute(h))

	return r
}

func privateRoute(h RestHandler) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	// r.Use(panics.HTTPRecoveryMiddleware)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
	}))
	// r.Use(flaware.Satpam(h.auth))

	// API route
	r.Route("/products", func(r chi.Router) {
		r.Get("/", h.GetAllProducts)
		r.Post("/seed/{n}", h.SeedProduct)
	})

	return r
}

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
