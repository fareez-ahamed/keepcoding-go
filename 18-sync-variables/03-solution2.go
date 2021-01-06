package main

import (
	"fmt"
	"sync"
)

const turns = 1000000

// Counter is a type
type Counter struct {
	count int
	sync.Mutex
}

func main() {
	done := make(chan struct{})
	counter := Counter{count: 0}

	go func() {
		for i := 0; i < turns; i++ {
			counter.Lock()
			counter.count += i * 10
			counter.Unlock()
		}
		done <- struct{}{}
	}()

	go func() {
		for i := 0; i < turns; i++ {
			counter.Lock()
			counter.count -= i * 10
			counter.Unlock()
		}
		done <- struct{}{}
	}()

	<-done
	<-done
	fmt.Println(counter.count)
}
