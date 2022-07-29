package main

import (
	"log"
	"os"

	"github.com/tuingking/flamingo/config"
	"github.com/tuingking/flamingo/handler/cli"
	"github.com/tuingking/flamingo/infra/logger"
	"github.com/tuingking/flamingo/infra/mysql"
	"github.com/tuingking/flamingo/infra/sqlgorm"
	"github.com/tuingking/flamingo/internal/product"

	_runtime "runtime/pprof"
)

var (
	cpuProfile = "cpu.prof"
)

func main() {
	// Create file to stored the profiling
	cpuProf, err := os.Create(cpuProfile)
	if err != nil {
		log.Fatal(err)
	}
	defer cpuProf.Close()
	_runtime.StartCPUProfile(cpuProf)
	defer _runtime.StopCPUProfile()

	// Config
	cfg := config.Init(
		config.WithConfigFile("config"),
		config.WithConfigType("yaml"),
	)
	// Infra
	log := logger.New(cfg.Logger)
	sql := mysql.New(cfg.MySQL)
	grm := sqlgorm.New(cfg.SQLGorm)

	// Service
	productRepo := product.NewRepository(sql, grm)
	productSvc := product.NewService(cfg.Product.Service, log, productRepo)

	app := cli.CLI{
		Cfg:        &cfg,
		Log:        &log,
		ProductSvc: productSvc,
	}

	app.Execute()
}
