package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	afterBuy := -prices[0]
	afterSell := 0
	afterHolding := 0
	afterCooldown := 0

	for i := 1; i < len(prices); i++ {
		temp := afterBuy
		afterBuy = maxValue(afterBuy, maxValue(afterHolding-prices[i], afterCooldown-prices[i]))
		afterHolding = maxValue(afterHolding, afterCooldown)
		afterCooldown = afterSell
		afterSell = temp + prices[i]
	}

	return maxValue(afterSell, maxValue(afterHolding, afterCooldown))
}

func minValue(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func maxValue(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	prices := []int{1, 2, 3, 0, 2}
	result := maxProfit(prices)
	fmt.Println(result)

}
