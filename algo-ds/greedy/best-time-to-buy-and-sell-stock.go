package main

import "fmt"

func maxProfit(prices []int) int {
	from, to := 0, 0
	ans := -1
	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[from] {
			from = i
			to = i
		} else if prices[i] > prices[to] {
			to = i
			cur := prices[to] - prices[from]
			if ans < cur {
				ans = cur
			}
		}
	}
	if ans < 0 {
		ans = 0
	}
	return ans
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(maxProfit([]int{5, 4, 3, 2, 1}))
}
