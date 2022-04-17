package main

import "fmt"

func maxProfit(prices []int) int {
	ans := 0
	for i := 1; i < len(prices); i++ {
		if prices[i-1] < prices[i] {
			ans = ans + prices[i] - prices[i-1]
		}
	}
	return ans
}
func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(maxProfit([]int{5, 4, 3, 2, 1}))

}
