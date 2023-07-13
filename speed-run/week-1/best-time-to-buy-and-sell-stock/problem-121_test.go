package besttimetobuyandsellstock

import "testing"

func TestFunc(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}

	ans := maxProfit(prices)
	println(ans)

	prices = []int{7, 6, 4, 3, 1}
	ans = maxProfit(prices)
	println(ans)
}
