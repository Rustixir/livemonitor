
![LiveMonitor](https://github.com/Rustixir/livemonitor/blob/main/logo.png)

# LiveMonitor
LiveMonitor can track your golang runtime in Real-Time.
Powered by Skyview(https://github.com/Rustixir/skyview)


![LiveMonitor](https://github.com/Rustixir/livemonitor/blob/main/screenshoot.png)

## What is it ??

Think you want to run a golang app and need to monitor memory statistics, garbage collection, 
number of created goroutines, just run LiveMonitor embedded in your program !! 
now you can track Golang runtime with browser in RealTime


Example Program:

``` 
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


```


you just need to add `go livemonitor.ListenAndServe(context.Background(), time.Second, ":6565")`
this to your program and then enjoy from it 
