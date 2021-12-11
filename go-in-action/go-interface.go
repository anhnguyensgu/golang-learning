package main

import (
	"fmt"
	"sync"
	"time"
)

var test bool = true
var num int = 2

type Vertex struct {
	X int
	Y int
	z int
}

func add(x, y int) (sum, a, b int) {
	a = x
	b = y
	sum = a + b
	return
}

type Abser interface {
	setZ(z int)
}

func changePointer(p *int, newValue int) {
	*p = newValue
}

func (v *Vertex) setZ(z int) {
	v.z = z
}
func say(s string) {
	for i := 0; i < 2; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int, wg *sync.WaitGroup) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
	wg.Done()
}

func channelSelect(channel chan int, quit chan int) {
	x := 0
	for {
		select {
		case x = <-channel:
			fmt.Printf("from channel %d\n", x)
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	myNum := 123
	p := &myNum

	var myArry [3]int
	myArry[0] = 1

	sliceLiteral := []int{1, 2, 3}

	mySlice := sliceLiteral[0:2]
	sliceLiteral = append(sliceLiteral, 0, 1, 2, 3, 4)

	fmt.Printf("my arr %+v\n", myArry)
	fmt.Printf("my unknow size arr %+v\n", sliceLiteral)
	fmt.Printf("my slice %+v\n", mySlice)

	fmt.Printf("pointer before %d\n", *p)

	changePointer(&myNum, 234)
	fmt.Printf("pointer after %d\n", *p)

	vertex := Vertex{X: 2, Y: 2}
	vertexPtr := &vertex

	fmt.Printf("my vertex %+v\n", vertex)
	vertex.X = 2
	fmt.Printf("my vertex after %+v\n", vertex)

	fmt.Printf("my vertex from pointer %+v\n", *vertexPtr)

	vertexPtr.X = 3
	vertexPtr.setZ(1)
	vertex.setZ(2)

	var abser Abser
	abser = vertexPtr
	abser.setZ(3)
	fmt.Printf("my vertex from pointer after %+v\n", *vertexPtr)
	go say("world")
	say("hello")

	s := []int{7, 2, 8, -9, 4, 0}

	var wg sync.WaitGroup
	wg.Add(2)
	channel := make(chan int, 2)
	go sum(s[:len(s)/2], channel, &wg)
	go sum(s[len(s)/2:], channel, &wg)
	wg.Wait()
	close(channel)

	result := 0
	for v := range channel {
		result += v
	}
	fmt.Printf("result %d\n", result)

	newChannel := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 2; i++ {
			fmt.Println("send to channel")
			newChannel <- i
		}
		quit <- 0
	}()
	channelSelect(newChannel, quit)
}
