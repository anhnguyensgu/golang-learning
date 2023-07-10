package main

import "fmt"

func numsSameConsecDiff(n int, k int) []int {
	if n == 1 {
		return []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	}
	ans := []int{}
	//
	for i := 1; i < 10; i++ {
		helperDFS(n-1, k, i, &ans)
	}

	return ans
}

func helperDFS(n int, k int, num int, ans *[]int) {
	if n == 0 {
		a := *ans
		a = append(a, num)
		*ans = a
		return
	}

	cur := num % 10

	if cur+k >= 0 && cur+k < 10 {
		helperDFS(n-1, k, num*10+cur+k, ans)
	}

	if k != 0 && cur-k >= 0 && cur-k < 10 {
		helperDFS(n-1, k, num*10+cur-k, ans)
	}

}

func solveWithQueue(n int, k int) []int {
	queue := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for level := 1; level < n; level++ {
		nextQueue := []int{}
		for _, num := range queue {
			curDigit := num % 10
			nextDigits := []int{curDigit + k}
			if k != 0 {
				nextDigits = append(nextDigits, curDigit-k)
			}

			for _, next := range nextDigits {
				if next >= 0 && next < 10 {
					nextQueue = append(nextQueue, num*10+next)
				}
			}
		}
		queue = nextQueue
	}
	return queue
}

func main() {
	n, k := 3, 7
	fmt.Println(solveWithQueue(n, k))
	fmt.Println(numsSameConsecDiff(n, k))
}
