package main

import (
	"fmt"
	"sort"
)

type Asteroids []int

func (r Asteroids) Len() int           { return len(r) }
func (r Asteroids) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r Asteroids) Less(i, j int) bool { return r[i] < r[j] }

func asteroidsDestroyed(mass int, asteroids []int) bool {
	var sortableA Asteroids = asteroids
	sort.Sort(sortableA)

	for _,v := range sortableA {
		if v <= mass {
			mass = mass + v
		} else {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(asteroidsDestroyed(5 , []int{4,9,23,4}))
}

