package main

import (
	"context"
	"html/template"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/pkg/errors"
	"github.com/tuingking/flamingo/config"
	"github.com/tuingking/flamingo/handler/rest"
	"github.com/tuingking/flamingo/infra/httpserver"
	"github.com/tuingking/flamingo/infra/logger"
	"github.com/tuingking/flamingo/infra/mysql"
	"github.com/tuingking/flamingo/internal/product"
)

var (
	Namespace    string
	BuildVersion string
	BuildTime    string
	CommitHash   string
)

func main() {
	meta := config.Metadata{
		Namespace:    Namespace,
		GoVersion:    runtime.Version(),
		BuildVersion: BuildVersion,
		BuildTime:    BuildTime,
		CommitHash:   CommitHash,
	}

	// config
	cfg := config.Init(
		config.WithConfigFile("config"),
		config.WithConfigType("yaml"),
	)
	cfg.SetMetadata(meta)

	// Infra
	logger := logger.New(cfg.Logger)
	sql := mysql.New(cfg.MySQL)

	logger.Infof("Meta: %+v", cfg.Meta)
	logger.Infof("MySQL: %+v", cfg.MySQL)

	// Domain - Product
	productRepo := product.NewRepository(sql)
	productSvc := product.NewService(productRepo)

	// web template
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
