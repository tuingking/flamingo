package main

import (
	"html/template"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	"github.com/tuingking/flamingo/app/api"
	"github.com/tuingking/flamingo/config"
	"github.com/tuingking/flamingo/infra/logger"
	"github.com/tuingking/flamingo/infra/mysql"
	"github.com/tuingking/flamingo/internal/home"
	"github.com/tuingking/flamingo/internal/product"
)

func main() {
	// config
	cfg := config.Init("config/config.yaml")

	// Infra
	logger := logger.New()
	sql := mysql.Init(cfg.MySQL)

	// Domain - Product
	productRepo := product.NewRepository(sql)
	productSvc := product.NewService(productRepo)

	// Web
	tpl := template.Must(template.ParseGlob("web/templates/*"))

	// API handler
	apihandler := api.Handler{
		Product: product.NewAPI(logger, productSvc),
		Home:    home.NewWeb(tpl),
	}

	// mux
	mux := api.NewRouter(apihandler)

	// app
	app := api.NewApp(cfg, logger, mux)

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// serve app in another goroutine
	go app.ListenAndServe()
	logger.Info("AFTER LISTEN AND SERVE")

	// waiting for os signal
	<-quit

	logger.Info("server shutdown gracefully")
	if err := app.Shutdown(); err != nil {
		logger.Error(errors.Wrap(err, "failed to shutdown server"))
	}
}
