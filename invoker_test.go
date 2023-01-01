package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApp_AddInvoker(t *testing.T) {
	a := NewApp()
	dummyFunc := func() {}
	a.AddInvoker(dummyFunc)
	assert.NotEmpty(t, a.invokers)
	assert.Len(t, a.invokers, 1)
}

func TestApp_AddInvokers(t *testing.T) {
	a := NewApp()
	dummyFunc := func() {}
	dummyFunc2 := func() {}
	a.AddInvokers(dummyFunc, dummyFunc2)
	assert.NotEmpty(t, a.invokers)
	assert.Len(t, a.invokers, 2)
}

func TestApp_initializeInvokers(t *testing.T) {
	a := NewApp()
	res := make(chan int, 2)
	dummyFunc := func() {
		res <- 1
	}
	dummyFunc2 := func() {
		res <- 2
	}
	a.AddInvokers(dummyFunc, dummyFunc2)
	a.initializeInvokers()
	assert.Len(t, res, 2)
}
