package main

import (
	"fmt"
	"math/rand"
	"time"
)

type database struct {
	Source string
	Data   string
}

func sourcess(name string, data string, done chan<- database) {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Intn(500)+100) * time.Millisecond)
	done <- database{Source: name, Data: data}
}

func main() {
	done := make(chan database)
	sources := map[string]string{
		"source 1": "дистрибутивы Linux",
		"source 2": "Hackintosh",
	}
	for name, data := range sources {
		go sourcess(name, data, done)
	}
	flc := <-done
	fmt.Printf("Найденный результат:\nИсточник: %s\nИнформация: %s\n", flc.Source, flc.Data)
}
//work