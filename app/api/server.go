package api

import (
	"context"
	"net/http"

	"github.com/tuingking/flamingo/config"
	"github.com/tuingking/flamingo/infra/logger"
)

type App struct {
	cfg    config.Config
	logger logger.Logger
	server *http.Server
	mux    http.Handler
}

func NewApp(cfg config.Config, logger logger.Logger, mux http.Handler) App {
	return App{
		cfg:    cfg,
		logger: logger,
		mux:    mux,
	}
}

func (app *App) ListenAndServe() error {
	app.server = &http.Server{
		Addr:         app.cfg.HttpServer.Port,
		ReadTimeout:  app.cfg.HttpServer.ReadTimeout,
		WriteTimeout: app.cfg.HttpServer.WriteTimeout,
		Handler:      app.mux,
	}

	app.logger.Infof("Http Server run on port %s", app.cfg.HttpServer.Port)
	return app.server.ListenAndServe()
}

func (a *App) Shutdown() error {
	return a.server.Shutdown(context.Background())
}
