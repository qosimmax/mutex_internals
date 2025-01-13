package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const unlocked = false
const locked = true

type BrokenMutex struct {
	state atomic.Bool
}

func (m *BrokenMutex) Lock() {
	for m.state.Load() {
	}

	m.state.Store(locked)
}

func (m *BrokenMutex) Unlock() {
	m.state.Store(unlocked)

}

const numGoroutines = 1000

func main() {
	var m BrokenMutex

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
