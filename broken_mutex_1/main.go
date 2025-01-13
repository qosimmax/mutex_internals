package main

import (
	"fmt"
	"sync"
)

const unlocked = false
const locked = true

type BrokenMutex struct {
	state bool
}

func (m *BrokenMutex) Lock() {
	for m.state {
	}

	m.state = locked
}

func (m *BrokenMutex) Unlock() {
	m.state = unlocked

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
