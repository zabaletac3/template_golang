package lifecycle

import "sync"

type Workers struct {
	wg sync.WaitGroup
}

func NewWorkers() *Workers {
	return &Workers{}
}

func (w *Workers) Add(delta int) {
	w.wg.Add(delta)
}

func (w *Workers) Done() {
	w.wg.Done()
}

func (w *Workers) Wait() {
	w.wg.Wait()
}
