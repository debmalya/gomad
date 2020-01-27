package log

import (
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000",
		FullTimestamp:   true,
	})
}

// Debugf
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

// Infof
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// Printf
func Printf(format string, args ...interface{}) {
	logger.Printf(format, args...)
}

// Warnf
func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

// Warningf
func Warningf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

// Errorf
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

//Fatalf
func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

//Panicf
func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}

// Debug
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Info
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Print
func Print(args ...interface{}) {
	logger.Info(args...)
}

// Warn
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// Warning
func Warning(args ...interface{}) {
	logger.Warn(args...)
}

// Error
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Fatal
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Panic
func Panic(args ...interface{}) {
	logger.Panic(args...)
}
