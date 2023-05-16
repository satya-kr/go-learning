package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func displayMessage(ch chan string) {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Second)
			ch <- fmt.Sprintf("This is value %d in channel", getRand())
		}()
	}
	wg.Wait()
	close(ch)
}

func getRand() int {
	return rand.Intn(100)
}

func main() {
	ch := make(chan string, 3)
	go displayMessage(ch)
	for c := range ch {
		fmt.Println(c)
	}
}
