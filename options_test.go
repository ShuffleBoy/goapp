package goapp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApp_applyOptions(t *testing.T) {
	app := &App{}
	app.applyOptions(WithConcurrentInvokers())
	assert.Equal(t, true, app.concurrentInvokers)
}

func TestWithConcurrentInvokers(t *testing.T) {
	app := &App{}
	WithConcurrentInvokers()(app)
	assert.Equal(t, true, app.concurrentInvokers)
}
