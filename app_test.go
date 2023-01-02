package goapp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewApp(t *testing.T) {
	a := NewApp()
	assert.NotEmpty(t, a)
	assert.Empty(t, a.providers)
	assert.Empty(t, a.invokers)
}

func TestApp_Start(t *testing.T) {
	a := NewApp()
	assert.NotPanics(t, a.Start)
}
