package httpserver

import (
	"context"
	"net/http"
	"time"

	"github.com/tuingking/flamingo/infra/logger"
)

type HTTPServer interface {
	Close() error
	ListenAndServe() error
	// ListenAndServeTLS(certFile, keyFile string) error
	// RegisterOnShutdown(f func())
	// Serve(l net.Listener) error
	// ServeTLS(l net.Listener, certFile, keyFile string) error
	// SetKeepAlivesEnabled(v bool)
	Shutdown(ctx context.Context) error
}

type httpServer struct {
	cfg    Config
	logger logger.Logger
	server *http.Server
}

type Config struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func New(cfg Config, logger logger.Logger, h http.Handler) HTTPServer {
	return &httpServer{
		cfg:    cfg,
		logger: logger,
		server: &http.Server{
			Addr:         cfg.Port,
			ReadTimeout:  cfg.ReadTimeout,
			WriteTimeout: cfg.WriteTimeout,
			Handler:      h,
		},
	}
}

func (s *httpServer) Close() error {
	return s.server.Close()
}

func (s *httpServer) ListenAndServe() error {
	s.logger.Info("Server running on port ", s.cfg.Port)
	s.logger.Info("Read Timeout ", s.cfg.ReadTimeout)
	s.logger.Info("Write Timeout ", s.cfg.WriteTimeout)
	return s.server.ListenAndServe()
}

func (s *httpServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
