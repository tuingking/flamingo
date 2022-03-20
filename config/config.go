package config

import (
	"io/ioutil"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/tuingking/flamingo/infra/mysql"
	"gopkg.in/yaml.v2"
)

type Config struct {
	HttpServer struct {
		Port         string
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
	}

	MySQL mysql.Config
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
