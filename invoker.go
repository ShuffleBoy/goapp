package app

func (a *App) AddInvoker(f any) {
	a.invokers = append(a.invokers, f)
}

func (a *App) AddInvokers(fs ...any) {
	a.invokers = append(a.providers, fs...)
}

func (a *App) initializeInvokers() {
	for _, invoker := range a.invokers {
		if err := a.container.Invoke(invoker); err != nil {
			panic(err)
		}
	}
}
