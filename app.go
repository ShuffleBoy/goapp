package app

import (
	"go.uber.org/dig"
)

type App struct {
	container *dig.Container
	providers []any
	invokers  []any
}

func NewApp() *App {
	app := new(App)
	app.container = dig.New()

	return app
}

func (a *App) Start() {
	a.initializeProviders()
	a.initializeInvokers()
}
