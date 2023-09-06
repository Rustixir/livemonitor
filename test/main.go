package main

import (
	"context"
	"livemonitor"
	"time"
)

func main() {
	go livemonitor.ListenAndServe(context.Background(), time.Second, ":6565")
	Dispatcher(100, time.Second)
}

func Dispatcher(spawnRate int, interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		<-ticker.C
		for i := 0; i < spawnRate; i++ {
			go Worker(30)
		}

	}
}

func Worker(maxProcess int) {
	var index = 0
	for {
		println("===> lorem ipsum")
		index++
		if index == maxProcess {
			return
		}
	}
}
