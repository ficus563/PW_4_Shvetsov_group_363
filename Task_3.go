package main

import (
	"fmt"
	"sync"
	"time"
		
)

func main() {
	var rec_channel = make(chan struct{})
	var tick = time.Tick(200 * time.Millisecond)
	var ad sync.WaitGroup
	
	for s := 0; s < 15; s++ {
		ad.Add(1)
		go func(n int) {
			defer ad.Done()
			rec_channel <- struct{}{}
			<-tick
			var times = time.Now().Format("15:04:05")
			fmt.Println("Запрос", n, "обработан в", times) 
		}(s)
	}

	go func () {	
		ad.Wait()
		close(rec_channel)
	}()

	for range rec_channel {

	}
}
