package main

import (
	"log"
	"sync"
	"time"
)

var config map[string]string
var configOnce sync.Once = sync.Once{}

func loadConfig() {
	time.Sleep(100 * time.Millisecond)
	log.Println("Loading configuration")
	config = map[string]string{
		"hostname": "localhost",
	}
}

func getConfig(key string) string {
	if config == nil {
		configOnce.Do(loadConfig)
	}
	return config[key]
}

func doSomething(done chan struct{}) {
	getConfig("hostname")
	done <- struct{}{}
}

func main() {
	done := make(chan struct{})
	for i := 0; i < 5; i++ {
		go doSomething(done)
	}
	for i := 0; i < 5; i++ {
		<-done
	}
}
