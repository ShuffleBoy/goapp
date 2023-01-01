package logger

import "github.com/sirupsen/logrus"

type Logger struct {
	*logrus.Logger
}

func NewLogger() *Logger {
	logger := logrus.New()
	logger.SetFormatter(new(logrus.JSONFormatter))
	return &Logger{logger}
}
