package goapp

import (
	"go.uber.org/dig"
)

type App struct {
	container *dig.Container
	providers []any
	invokers  []any

	// Options
	concurrentInvokers bool
}

func NewApp(options ...Option) *App {
	app := new(App)
	app.container = dig.New()
	app.applyOptions(options...)
	return app
}

func (a *App) Start() {
	a.initializeProviders()
	a.runInvokers()
}
