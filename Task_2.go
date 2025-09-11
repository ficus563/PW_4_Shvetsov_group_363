package main

import (
	"fmt"
	"sync"
)


func worker(a int, jobs <-chan int, results chan<- int) {
	for j:= range jobs {
		results <- j * j
	}
}


func main() {
	var jobs = make(chan int, 10)
	var results = make(chan int, 10)
	var k sync.WaitGroup

	for i := 1; i <= 3; i++ {
		k.Add(1)
		go func(w int) {
			defer k.Done()
			worker(w, jobs, results)
		}(i)
	}

	for b := 1; b <= 10; b++ {
		jobs <- b
	}
	close(jobs)

	go func() {
		k.Wait()
		close(results)
	}()

	for c := range results {
		fmt.Println(c, "\n")
	}
}