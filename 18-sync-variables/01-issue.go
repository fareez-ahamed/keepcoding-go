package main

import (
	"fmt"
)

const turns = 100000

func main() {
	num := 0
	done := make(chan struct{})
	go func() {
		for i := 0; i < turns; i++ {
			num++
		}
		done <- struct{}{}
	}()
	go func() {
		for i := 0; i < turns; i++ {
			num--
		}
		done <- struct{}{}
	}()

	<-done
	<-done
	fmt.Println(num)
}
