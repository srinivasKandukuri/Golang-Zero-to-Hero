package main

import (
	"sync"
)

// =============================WORK======================================//
type Worker interface {
	Task()
}

type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

func New(MaxGoroutines int) *Pool {
	pool := &Pool{
		work: make(chan Worker),
	}
	pool.wg.Add(MaxGoroutines)
	for i := 0; i < MaxGoroutines; i++ {
		go func() {
			for w := range pool.work {
				w.Task()
			}
			pool.wg.Done()
		}()
	}
	return pool
}

func (p *Pool) Run(w Worker) {
	p.work <- w
}

func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}

//=======================================END==================================================//

var names = []string{
	"steve",
	"bob",
	"sk",
	"pk",
}

type namePrinter struct {
	name string
}

func (m *namePrinter) Task() {
	println(m.name)
	//time.Sleep(1 * time.Second)
}

func main() {
	p := New(2)

	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		for _, name := range names {
			np := namePrinter{
				name: name,
			}
			go func() {
				p.Run(&np)
				wg.Done()
			}()
		}
	}
	wg.Wait()
	p.Shutdown()
}
