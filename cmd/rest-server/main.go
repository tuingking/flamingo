package main

import (
	"context"
	"html/template"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	"github.com/tuingking/flamingo/config"
	"github.com/tuingking/flamingo/handler/rest"
	"github.com/tuingking/flamingo/infra/httpserver"
	"github.com/tuingking/flamingo/infra/logger"
	"github.com/tuingking/flamingo/infra/mysql"
	"github.com/tuingking/flamingo/internal/product"
)

func main() {
	// config
	cfg := config.Init("config/config.yaml")

	// Infra
	logger := logger.New(cfg.Logger)
	sql := mysql.New(cfg.MySQL)

	// Domain - Product
	productRepo := product.NewRepository(sql)
	productSvc := product.NewService(productRepo)

	// Web
	tpl := template.Must(template.ParseGlob("web/templates/*"))

	// rest handler
	restHandler := rest.NewRestHandler(
		logger,
		tpl,
		productSvc,
	)

	// http handler
	mux := rest.NewHttpHandler(restHandler)

	// server
	server := httpserver.New(cfg.HttpServer, logger, mux)

	// graceful shutdown
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

		// waiting for os signal
		<-quit

		logger.Info("server shutdown gracefully")
		if err := server.Shutdown(context.Background()); err != nil {
			logger.Error(errors.Wrap(err, "failed to shutdown server"))
		}
	}()

	// serve
	server.ListenAndServe()

	logger.Info("bye bye 👋👋")
}
