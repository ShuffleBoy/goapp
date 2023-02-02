package goapp

type Option func(app *App)

func WithConcurrentInvokers() Option {
	return func(app *App) {
		app.concurrentInvokers = true
	}
}

func (a *App) applyOptions(options ...Option) {
	for _, option := range options {
		option(a)
	}
}
