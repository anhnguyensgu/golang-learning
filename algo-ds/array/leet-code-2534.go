package main

import "fmt"

// A \/ B = A + B - (A ^ B)
// => A \/ B + (A ^ B)= A + B

func NumOfSetBits(n int) int {
	count := 0
	for n != 0 {
		count += n % 2
		n >>= 1
	}
	return count
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

type Set = map[int]bool

func countExcellentPairs(nums []int, k int) int64 {
	var ans int64 = 0
	bitCountMap := map[int]Set{}
	for _, n := range nums {
		currentBitCount := NumOfSetBits(n)
		if _, ok := bitCountMap[currentBitCount]; !ok {
			bitCountMap[currentBitCount] = Set{}
		}
		bitCountMap[currentBitCount][n] = true
	}

	set := map[int]bool{}
	for _, n := range nums {
		if _, ok := set[n]; ok {
			continue
		}
		set[n] = true
		currentBitCount := NumOfSetBits(n)
		need := max(0, k-currentBitCount)
		for setBit, numList := range bitCountMap {
			if setBit >= need {
				ans = ans + int64(len(numList))
			}
		}
	}
	return ans
}

func main() {
	nums := []int{1, 2, 3, 1}

	k := 3
	ans := countExcellentPairs(nums, k)
	fmt.Println(ans)
}
