package config

import (
	"io/ioutil"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/tuingking/flamingo/infra/httpserver"
	"github.com/tuingking/flamingo/infra/logger"
	"github.com/tuingking/flamingo/infra/mysql"
	"gopkg.in/yaml.v2"
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
}

type Metadata struct {
	Namespace    string
	GoVersion    string
	BuildVersion string
	BuildTime    string
	CommitHash   string
}

func Init(file string) Config {
	var cfg Config

	f, err := ioutil.ReadFile(file)
	if err != nil {
		logrus.Fatal(errors.Wrap(err, "read config file"))
	}

	err = yaml.Unmarshal(f, &cfg)
	if err != nil {
		logrus.Fatal(errors.Wrap(err, "unmarshal yaml"))
	}

	return cfg
}

func (c *Config) SetMetadata(m Metadata) {
	c.Meta = m
}
