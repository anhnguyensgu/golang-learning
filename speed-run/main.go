package main

import "fmt"

type Provider interface {
	run()
}

type Concrete struct{}

func (c Concrete) run() {
	fmt.Println("struct 1 running")
}

func runWithStrategy(p Provider) {
	p.run()
}

func main() {
	concrete := Concrete{}
	runWithStrategy(concrete)
}
