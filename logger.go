package logger

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

type Fields = logrus.Fields

// Logger is the common interface batchcorp
type Logger interface {
	// Debug sends out a debug message with the given arguments to the logger.
	Debug(args ...interface{})
	// Debugf formats a debug message using the given arguments and sends it to the logger.
	Debugf(format string, args ...interface{})
	// Info sends out an informational message with the given arguments to the logger.
	Info(args ...interface{})
	// Infof formats an informational message using the given arguments and sends it to the logger.
	Infof(format string, args ...interface{})
	// Warn sends out a warning message with the given arguments to the logger.
	Warn(args ...interface{})
	// Warnf formats a warning message using the given arguments and sends it to the logger.
	Warnf(format string, args ...interface{})
	// Error sends out an error message with the given arguments to the logger.
	Error(args ...interface{})
	// Errorf formats an error message using the given arguments and sends it to the logger.
	Errorf(format string, args ...interface{})
	// Withfield creates a logrus.Entry for the logger.
	WithField(key string, args ...interface{}) *logrus.Entry
	// WithFields creats logrus.Entry and add multiple fields to the logger.
	WithFields(Fields) *logrus.Entry
}

type logger struct {
	entry *logrus.Entry
}

// Debug
func (l logger) Debug(args ...interface{}) {
	l.Debug(args...)
}

// Debugf
func (l logger) Debugf(format string, args ...interface{}) {
	l.Debugf(format, args...)
}

// Error
func (l logger) Error(args ...interface{}) {
	l.Error(args...)
}

// Errorf
func (l logger) Errorf(format string, args ...interface{}) {
	l.Errorf(format, args...)
}

// Info
func (l logger) Info(args ...interface{}) {
	l.Info(args...)
}

// Infof
func (l logger) Infof(format string, args ...interface{}) {
	l.Infof(format, args...)
}

// Warn
func (l logger) Warn(args ...interface{}) {
	l.Warn(args...)
}

// Warnf
func (l logger) Warnf(format string, args ...interface{}) {
	l.Warnf(format, args...)
}

// WithField
func (l logger) WithField(format string, args ...interface{}) *logrus.Entry {
	return l.WithField(format, args...)
}

// WithFields
func (l logger) WithFields(fields Fields) *logrus.Entry {
	return l.WithFields(fields)
}

// New returns a common logger
func New() Logger {
	log := logrus.New()
	return logger{entry: logrus.NewEntry(log)}
}

// NewNopLogger returns a logger that discards all log messages
func NewNoopLogger() Logger {
	log := logrus.New()
	log.Out = ioutil.Discard
	return logger{entry: logrus.NewEntry(log)}
}
