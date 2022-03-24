package config

import (
	"log"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tuingking/flamingo/infra/httpserver"
	"github.com/tuingking/flamingo/infra/logger"
	"github.com/tuingking/flamingo/infra/mysql"
	"github.com/tuingking/flamingo/internal/product"
)

const (
	ProductionEnv = "production"
	LocalEnv      = "local"
)

type Config struct {
	Meta Metadata

	// Infra
	Logger     logger.Config
	HttpServer httpserver.Config
	MySQL      mysql.Config

	// Product
	Product product.Config
}

type Metadata struct {
	Namespace    string
	GoVersion    string
	BuildVersion string
	BuildTime    string
	CommitHash   string
}

func Init(opts ...Option) Config {
	var cfg Config

	// default option
	opt := &option{
		configPath: "./config/",
		configFile: "config",
		configType: "yaml",
	}

	// override option
	for _, fn := range opts {
		fn(opt)
	}

	v := viper.New()
	v.AddConfigPath(opt.configPath)
	v.SetConfigName(opt.configFile)
	v.SetConfigType(opt.configType)

	if err := v.ReadInConfig(); err != nil {
		log.Fatal(errors.Wrap(err, "read config"))
	}

	if err := v.Unmarshal(&cfg); err != nil {
		log.Fatal(errors.Wrap(err, "unmarshal config"))
	}

	logrus.Infof("%-7s %s", "Config", "✅")

	return cfg
}

func (c *Config) SetMetadata(m Metadata) {
	c.Meta = m
}

type option struct {
	configFile string
	configPath string
	configType string
}

type Option func(*option)

func WithConfigFile(configFile string) Option {
	return func(o *option) {
		o.configFile = configFile
	}
}

func WithConfigPath(configPath string) Option {
	return func(o *option) {
		o.configPath = configPath
	}
}

func WithConfigType(configType string) Option {
	return func(o *option) {
		o.configType = configType
	}
}
