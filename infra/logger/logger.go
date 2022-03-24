package logger

import (
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
}

type logger struct {
	log *logrus.Logger
}

type Config struct {
	Format string
	Level  string
}

func New(cfg Config) Logger {
	log := logrus.New()
	log.SetOutput(os.Stdout)

	// formatter
	switch cfg.Format {
	case "json":
		log.SetFormatter(&logrus.JSONFormatter{})
	default:
		log.SetFormatter(&logrus.TextFormatter{})
	}

	// level
	switch cfg.Level {
	case "panic":
		log.SetLevel(logrus.PanicLevel)
	case "fatal":
		log.SetLevel(logrus.FatalLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	case "warn":
		log.SetLevel(logrus.WarnLevel)
	case "info":
		log.SetLevel(logrus.InfoLevel)
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	case "trace":
		log.SetLevel(logrus.TraceLevel)
	default:
		log.SetLevel(logrus.DebugLevel)
	}

	logrus.Infof("%-7s %s", "Logger", "✅")

	return &logger{
		log: log,
	}
}

func (l *logger) Debug(args ...interface{}) {
	l.log.WithField("goroutines", runtime.NumGoroutine()).Debug(args...)
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.log.WithField("goroutines", runtime.NumGoroutine()).Debugf(format, args...)
}

func (l *logger) Info(args ...interface{}) {
	l.log.WithField("goroutines", runtime.NumGoroutine()).Info(args...)
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.log.WithField("goroutines", runtime.NumGoroutine()).Infof(format, args...)
}

func (l *logger) Warn(args ...interface{}) {
	l.log.WithField("goroutines", runtime.NumGoroutine()).Warn(args...)
}

func (l *logger) Warnf(format string, args ...interface{}) {
	l.log.WithField("goroutines", runtime.NumGoroutine()).Warnf(format, args...)
}

func (l *logger) Error(args ...interface{}) {
	l.log.WithField("goroutines", runtime.NumGoroutine()).Error(args...)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.log.WithField("goroutines", runtime.NumGoroutine()).Errorf(format, args...)
}
