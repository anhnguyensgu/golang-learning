package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- 200
	}()
	go func() {
		time.Sleep(1 * time.Second)
		c2 <- 100
	}()

	for i := 1; i <= 3; i++ {
		select {
		case v1 := <-c1:
			fmt.Printf("from chan 1 %d\n", v1)
		case v2 := <-c2:
			fmt.Printf("from chan 2 %d\n", v2)
			// default:
			// 	fmt.Println("no activity")
			// 	break
		}
	}
	fmt.Println("done")
}
