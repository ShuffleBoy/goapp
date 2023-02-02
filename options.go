package goapp

type Option func(app *App)

func WithConcurrentInvokers() Option {
	return func(app *App) {
		app.concurrentInvokers = true
	}
}

func (a *App) applyOptions() {
	for _, option := range a.options {
		option(a)
	}
}
