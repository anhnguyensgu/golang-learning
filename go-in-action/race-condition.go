package main

import (
	"fmt"
	"sync"
)

var a = 10

func runningShit() {
	fmt.Println("sheet")
	m.Lock()
	a = 10
	m.Unlock()
}

var m sync.Mutex

func main() {
	go func() {
		for {
			fmt.Println("Thread B")
		}
	}()
	for {
		fmt.Println("Thread A")
	}
}
