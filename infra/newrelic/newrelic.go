package newrelic

import (
	"os"
	"time"

	"github.com/newrelic/go-agent/v3/integrations/nrlogrus"
	nr "github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
)

type NewRelic interface {
	StartTransaction(name string) *nr.Transaction
	RecordCustomEvent(eventType string, params map[string]interface{})
	RecordCustomMetric(name string, value float64)
	WaitForConnection(timeout time.Duration) error
	Shutdown(timeout time.Duration)
}

type Config struct {
	AppName    string
	LicenseKey string
	Output     *os.File
}

func Init(cfg Config) NewRelic {
	app, err := nr.NewApplication(
		nr.ConfigAppName(cfg.AppName),
		nr.ConfigLicense(cfg.LicenseKey),
		nr.ConfigDebugLogger(cfg.Output),
		func(config *nr.Config) {
			logrus.SetLevel(logrus.DebugLevel)
			config.Logger = nrlogrus.StandardLogger()
		},
	)
	if err != nil {
		logrus.Panicf("unable to create New Relic Application, err= %s", err)
	}

	return app
}
