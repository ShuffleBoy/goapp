package goapp

func (a *App) AddProvider(constructor any) {
	a.providers = append(a.providers, constructor)
}

func (a *App) AddProviders(constructors ...any) {
	a.providers = append(a.providers, constructors...)
}

func (a *App) initializeProviders() {
	for _, provider := range a.providers {
		if err := a.container.Provide(provider); err != nil {
			panic(err)
		}
	}
}
