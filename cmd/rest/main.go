package main

import (
	"context"
	"html/template"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/pkg/errors"
	"github.com/tuingking/flamingo/config"
	"github.com/tuingking/flamingo/handler/rest"
	"github.com/tuingking/flamingo/infra/httpserver"
	"github.com/tuingking/flamingo/infra/logger"
	"github.com/tuingking/flamingo/infra/mysql"
	"github.com/tuingking/flamingo/infra/sqlgorm"
	"github.com/tuingking/flamingo/internal/account"
	"github.com/tuingking/flamingo/internal/auth"
	"github.com/tuingking/flamingo/internal/product"
)

var (
	Namespace    string
	BuildVersion string
	BuildTime    string
	CommitHash   string
)

// @title Swagger Flamingo API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api
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
	// config.WithConfigFile("config"),
	// config.WithConfigType("yaml"),
	)
	cfg.SetMetadata(meta)

	// Infra
	logger := logger.New(cfg.Logger)
	sql := mysql.New(cfg.MySQL)
	grm := sqlgorm.New(cfg.SQLGorm)

	logger.Infof("Meta: %+v", cfg.Meta)
	logger.Infof("MySQL: %+v", cfg.MySQL)

	authSvc := auth.NewService(cfg.Auth, logger)
	accountRepo := account.NewRepository(sql)
	accountSvc := account.NewService(cfg.Account, logger, accountRepo)
	productRepo := product.NewRepository(sql, grm)
	productSvc := product.NewService(cfg.Product.Service, logger, productRepo)

	// web template
	tpl := template.Must(template.ParseGlob("web/templates/*"))

	// rest handler
	restHandler := rest.NewRestHandler(
		logger,
		tpl,
		authSvc,
		productSvc,
		accountSvc,
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

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			logger.Error(errors.Wrap(err, "failed to shutdown server"))
		}
		logger.Info("server shutdown gracefully 😁")
	}()

	// serve
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal(errors.Wrap(err, "ListenAndServe"))
	}

	logger.Info("bye bye 👋👋")
}
