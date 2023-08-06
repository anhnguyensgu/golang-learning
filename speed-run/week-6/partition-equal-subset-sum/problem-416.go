package partitionequalsubsetsum

import "fmt"

func canPartition(nums []int) bool {
	half := 0
	for _, n := range nums {
		half += n
	}
	if half%2 == 1 {
		return false
	}

	half = half / 2
	dp := make([][]bool, half+1)

	for i := 0; i <= len(nums); i++ {
		dp[i] = make([]bool, half+1)
		dp[i][0] = true
	}

	for j, coin := range nums {
		for i := 1; i <= half; i++ {
			if i-coin >= 0 {
				dp[j+1][i] = dp[j][i-coin] || dp[j][i]
			} else {
				dp[j+1][i] = dp[j][i]
			}
		}
		if dp[j+1][half] {
			return true
		}

	}

	for _, a := range dp {
		fmt.Println(a)
	}
	return false
}
