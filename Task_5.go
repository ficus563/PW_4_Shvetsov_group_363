package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"sync"
	_ "os"
)

func readfiles(filename string, r chan<- string) {
	d, err := ioutil.ReadFile(filename)
	if err != nil {
		r <- fmt.Sprintf("Ошибка чтения файла %s: %v", filename, err)
		return
	}
	var h = md5.Sum(d)
	r <- fmt.Sprintf("%s: %x", filename, h)
}

func main() {
	var files = []string{"file1.txt", "file2.txt"}
	var result = make(chan string, len(files))
	var syn sync.WaitGroup
	for _, f := range files {
		syn.Add(1)
		
		go func(f string) {
			defer syn.Done()
			readfiles(f, result)
		}(f)
	}
	go func() {
		syn.Wait()
		close(result)
	}()
	for res := range result {
		fmt.Println(res)
	}
}

//work