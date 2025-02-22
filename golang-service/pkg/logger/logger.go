package logger

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

type Interface interface {
	Debug(message interface{}, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message interface{}, args ...interface{})
	Fatal(message interface{}, args ...interface{})
}

type Logger struct {
	logger *slog.Logger
}

var _ Interface = (*Logger)(nil)

func New(level string) *Logger {
	var logLevel slog.Level

	switch strings.ToLower(level) {
	case "error":
		logLevel = slog.LevelError
	case "warn":
		logLevel = slog.LevelWarn
	case "info":
		logLevel = slog.LevelInfo
	case "debug":
		logLevel = slog.LevelDebug
	default:
		logLevel = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		Level: logLevel,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))

	return &Logger{
		logger: logger,
	}
}

func (l *Logger) Debug(message interface{}, args ...interface{}) {
	l.msg(slog.LevelDebug, message, args...)
}

func (l *Logger) Info(message string, args ...interface{}) {
	l.log(slog.LevelInfo, message, args...)
}

func (l *Logger) Warn(message string, args ...interface{}) {
	l.log(slog.LevelWarn, message, args...)
}

func (l *Logger) Error(message interface{}, args ...interface{}) {
	l.msg(slog.LevelError, message, args...)
}

func (l *Logger) Fatal(message interface{}, args ...interface{}) {
	l.msg(slog.LevelError, message, args...)
	os.Exit(1)
}

func (l *Logger) log(level slog.Level, message string, args ...interface{}) {
	if len(args) == 0 {
		l.logger.Log(nil, level, message)
	} else {
		l.logger.Log(nil, level, fmt.Sprintf(message, args...))
	}
}

func (l *Logger) msg(level slog.Level, message interface{}, args ...interface{}) {
	switch msg := message.(type) {
	case error:
		l.log(level, msg.Error(), args...)
	case string:
		l.log(level, msg, args...)
	default:
		l.log(level, fmt.Sprintf("%s message %v has unknown type %v", level, message, msg), args...)
	}
}
