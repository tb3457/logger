package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var log *logrus.Logger

type Level logrus.Level

func Register(level Level, file *os.File) {
	Out := io.MultiWriter(os.Stdout)
	if file != nil {
		Out = io.MultiWriter(os.Stdout, file)
	}
	log = &logrus.Logger{
		Out:       Out,
		Formatter: new(logrus.TextFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.Level(level),
	}
}

func Info(msg string) {
	if log != nil {
		log.Info(msg)
	}
}

func Infof(format string, data ...any) {
	if log != nil {
		log.Infof(format, data)
	}
}

func Panic(msg string) {
	if log != nil {
		log.Panic(msg)
	}
}

func Panicf(format string, data ...any) {
	if log != nil {
		log.Panicf(format, data)
	}
}

func Fatal(msg string) {
	if log != nil {
		log.Fatal(msg)
	}
}

func Fatalf(format string, data ...any) {
	if log != nil {
		log.Fatalf(format, data)
	}
}

func Error(msg string) {
	if log != nil {
		log.Error(msg)
	}
}

func Errorf(format string, data ...any) {
	if log != nil {
		log.Errorf(format, data)
	}
}

func Warn(msg string) {
	if log != nil {
		log.Warn(msg)
	}
}

func Warnf(format string, data ...any) {
	if log != nil {
		log.Warnf(format, data)
	}
}

func Debug(msg string) {
	if !debug() {
		return
	}
	if log != nil {
		log.Debug(msg)
	}
}

func Debugf(format string, data ...any) {
	if !debug() {
		return
	}
	if log != nil {
		log.Debugf(format, data)
	}
}

func Trace(msg string) {
	if !trace() {
		return
	}
	if log != nil {
		log.Trace(msg)
	}
}

func Tracef(format string, data ...any) {
	if !trace() {

	}
	if log != nil {
		log.Tracef(format, data)
	}
}

func debug() bool {
	return os.Getenv("DEBUG") != ""
}

func trace() bool {
	return os.Getenv("TRACE") != ""
}
