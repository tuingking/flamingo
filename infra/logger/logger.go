package logger

import (
	"runtime"

	"github.com/sirupsen/logrus"
)

type Logger interface {
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

func New() Logger {
	return &logger{
		log: logrus.New(),
	}
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
