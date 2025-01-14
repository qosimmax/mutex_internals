package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

type Mutex struct {
	ownerTicket    atomic.Int64
	nextFreeTicket atomic.Int64
}

func (m *Mutex) Lock() {
	ticket := m.nextFreeTicket.Add(1)
	for m.ownerTicket.Load() != ticket-1 {
		runtime.Gosched()
	}
}

func (m *Mutex) Unlock() {
	m.ownerTicket.Add(1)
}

const numGoroutines = 1000

func main() {
	var m Mutex

	wg := sync.WaitGroup{}
	wg.Add(numGoroutines)

	value := 0
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()

			m.Lock()
			value++
			m.Unlock()

		}()
	}

	wg.Wait()

	fmt.Println(value)
}
