package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func furthestBuilding(heights []int, bricks int, ladders int) int {
	curSum := 0
	last := heights[0]
	topMax := &IntHeap{}
	curTopSum := 0
	heap.Init(topMax)
	for index, height := range heights {
		// fmt.Printf("current %d %d \n", height, last)
		if height > last {
			curDif := height - last
			topMax.Push(curDif)
			if topMax.Len() > ladders {
				curTopSum += curDif - topMax.Pop()
			}
			fmt.Printf("top max %o %d %d \n", topMax, curTopSum, curSum)
			if curSum-curTopSum > bricks {
				// fmt.Printf("index %d \n", index)
				return index - 1
			}
		}
		last = height
	}
	return len(heights) - 1
}

func main() {

	heights := []int{4, 2, 7, 6, 9, 14, 12}
	bricks := 5
	ladders := 1
	fmt.Println(furthestBuilding(heights, bricks, ladders))
}
