package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLogger(t *testing.T) {
	logger := NewLogger()
	assert.NotEmpty(t, logger)
	assert.Equal(t, &logrus.JSONFormatter{}, logger.Formatter)
}
