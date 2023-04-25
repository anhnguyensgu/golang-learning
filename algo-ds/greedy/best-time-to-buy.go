package main

func maxProfit(prices []int) int {
	ans := 0
	from, to := 0, 0
	for i, price := range prices {
		if price < prices[from] {
			from, to = i, i
		} else if price > prices[to] {
			to = i
			if cur := prices[to] - prices[from]; ans < cur {
				ans = cur
			}
		}
	}

	return ans
}

func main() {

}
