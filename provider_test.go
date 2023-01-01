package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testDep struct{}
type testDep2 struct{}

func depInitializer() *testDep {
	return new(testDep)
}

func depInitializer2() *testDep2 {
	return new(testDep2)
}

func TestApp_AddProvider(t *testing.T) {
	a := NewApp()
	assert.Empty(t, a.providers)
	a.AddProvider(depInitializer)
	assert.NotEmpty(t, a.providers)
	assert.Len(t, a.providers, 1)
}

func TestApp_AddProviders(t *testing.T) {
	a := NewApp()
	assert.Empty(t, a.providers)
	a.AddProviders(depInitializer, depInitializer2)
	assert.NotEmpty(t, a.providers)
	assert.Len(t, a.providers, 2)
}

func TestApp_initializeProviders(t *testing.T) {
	a := NewApp()

	type testDep3 struct{}
	provider := func(_ *testDep, _ *testDep2) *testDep3 {
		return new(testDep3)
	}
	a.AddProviders(depInitializer, depInitializer2, provider)
	assert.NotPanics(t, a.initializeProviders)
}
