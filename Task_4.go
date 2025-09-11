package main

import (
	"fmt"
	"sync"
	"net/http"
)

func main() {
	var URLS = []string{
		"https://github.com/",
	}

	var pool_Size = 3

	var channel_jobs = make(chan string, len(URLS)) 
	var results = make(chan string, len(URLS))

	var WG sync.WaitGroup
	for p:= 0; p < pool_Size; p++ {
		WG.Add(1)
		go worker(p, channel_jobs, results, &WG)
	}

	for _, url := range URLS {
		channel_jobs <- url
	}
	close(channel_jobs)

	go func () {
		WG.Wait()
		close(results)
	}()

	fmt.Println("result")
	for result := range results {
		fmt.Println(result)
	}
}

func worker(id int, channel_jobs <-chan string, results chan <- string, WG *sync.WaitGroup) {
	defer WG.Done()

	for url :=range channel_jobs {
		status, err := check_url(url)
		if err != nil {
			results <- fmt.Sprintf("Ошибка",id, url, err)
		} else {
			results <- fmt.Sprintf("Cтатус", id, url, status)
		}
	}
}

func check_url(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return resp.Status, nil
}

//work