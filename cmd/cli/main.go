package main

import (
	"github.com/tuingking/flamingo/config"
	"github.com/tuingking/flamingo/infra/logger"
	"github.com/tuingking/flamingo/infra/mysql"
	"github.com/tuingking/flamingo/internal/account"
)

var (
	accountSvc account.Service
)

func main() {
	// Config
	cfg := config.Init(
		config.WithConfigFile("config"),
		config.WithConfigType("yaml"),
	)

	// Infra
	logger := logger.New(cfg.Logger)
	sql := mysql.New(cfg.MySQL)

	accountRepo := account.NewRepository(sql)
	accountSvc = account.NewService(account.Config{}, logger, accountRepo)
}
