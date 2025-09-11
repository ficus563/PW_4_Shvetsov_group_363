package main

import (
	"fmt"
	"sync"
	"time"
)

func number_interval(number *sync.WaitGroup) {
	defer number.Done()
	for s := 1; s <= 5; s++ {
		fmt.Println(s)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var n sync.WaitGroup
	n.Add(1)
	go number_interval(&n)
	n.Wait()
}

