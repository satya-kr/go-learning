package main

import (
	"fmt"
	"sync"
)

var count = 0

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 100000; i++ {
		wg.Add(1)

		go increment(&wg, &m)
	}
	wg.Wait()
	fmt.Println("Count:", count)
}

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	count += 1
	m.Unlock()
}
