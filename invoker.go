package goapp

import "sync"

func (a *App) AddInvoker(f any) {
	a.invokers = append(a.invokers, f)
}

func (a *App) AddInvokers(fs ...any) {
	a.invokers = append(a.providers, fs...)
}

func (a *App) runInvokers() {
	if a.concurrentInvokers {
		var wg sync.WaitGroup
		for _, invoker := range a.invokers {
			wg.Add(1)
			invoker := invoker
			go func() {
				defer wg.Done()
				if err := a.container.Invoke(invoker); err != nil {
					panic(err)
				}
			}()
		}
		wg.Wait()
	} else {
		for _, invoker := range a.invokers {
			if err := a.container.Invoke(invoker); err != nil {
				panic(err)
			}
		}
	}
}
