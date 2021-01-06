package main

import (
	"fmt"
	"sync"
)

const turns = 100000

func main() {
	done := make(chan struct{})
	mutex := sync.Mutex{}
	num := 0

	go func() {
		for i := 0; i < turns; i++ {
			mutex.Lock()
			num += i * 10
			mutex.Unlock()
		}
		done <- struct{}{}
	}()

	go func() {
		for i := 0; i < turns; i++ {
			mutex.Lock()
			num -= i * 10
			mutex.Unlock()
		}
		done <- struct{}{}
	}()

	<-done
	<-done
	fmt.Println(num)
}
