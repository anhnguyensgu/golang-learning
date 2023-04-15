package main

import (
	"container/heap"
	"fmt"
)

func eatenApples(apples []int, days []int) int {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	ans := 0

	for i, apple := range apples {
		heap.Push(&pq, [2]int{i + days[i], apple})

		for len(pq) > 0 && (pq[0][0] <= i || pq[0][1] <= 0) {
			heap.Pop(&pq)
		}

		if len(pq) > 0 {
			ans++
			pq[0][1]--
		}
	}

	i := len(apples)
	for len(pq) > 0 {
		if pq[0][0] <= i || pq[0][1] <= 0 {
			heap.Pop(&pq)
			continue
		}
		ans++
		pq[0][1]--
		i++
	}

	return ans
}

type PriorityQueue [][2]int

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i][0] < pq[j][0]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	newValue := x.([2]int)
	*pq = append(*pq, newValue)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func main() {
	fmt.Println(eatenApples([]int{2, 1, 1, 4, 5}, []int{10, 10, 6, 4, 2}))
}
