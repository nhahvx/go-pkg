package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

// Interface -.
type Interface interface {
	Debug(message interface{}, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message interface{}, args ...interface{})
	Fatal(message interface{}, args ...interface{})
}

var AppLogger *Logger

// Logger -.
type Logger struct {
	logger *logrus.Logger
}

var _ Interface = (*Logger)(nil)

// New -.
func New() *Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
		PrettyPrint:     true,
	})

	log.SetLevel(logrus.DebugLevel)
	return &Logger{
		logger: log,
	}
}

// Debug -.
func (l *Logger) Debug(message interface{}, args ...interface{}) {
	l.logger.WithFields(logrus.Fields{
		"title": message,
	}).Debug(args...)
}

// Info -.
func (l *Logger) Info(message string, args ...interface{}) {
	l.logger.WithFields(logrus.Fields{
		"title": message,
	}).Info(args...)
}

// Warn -.
func (l *Logger) Warn(message string, args ...interface{}) {
	l.logger.WithFields(logrus.Fields{
		"title": message,
	}).Warn(args...)
}

// Error -.
func (l *Logger) Error(message interface{}, args ...interface{}) {
	l.logger.WithFields(logrus.Fields{
		"title": message,
	}).Error(args...)
}

// Fatal -.
func (l *Logger) Fatal(message interface{}, args ...interface{}) {
	l.logger.WithFields(logrus.Fields{
		"title": message,
	}).Fatal(args...)

	os.Exit(1)
}
