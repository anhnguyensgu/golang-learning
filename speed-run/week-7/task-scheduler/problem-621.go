package taskscheduler

import (
	"fmt"
	"sort"
)

func leastInterval(tasks []byte, n int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	maxFrequency := 0
	mapCount := map[int]int{}
	for _, t := range tasks {
		cur := int(t - 'A')
		if _, ok := mapCount[cur]; ok {
			mapCount[cur]++
		} else {
			mapCount[cur] = 1
		}

		maxFrequency = max(maxFrequency, mapCount[cur])

	}

	maxFrequencyCount := 0
	for _, c := range mapCount {
		if c == maxFrequency {
			maxFrequencyCount++
		}
	}
	fmt.Printf("%d %d\n", maxFrequency, maxFrequencyCount)
	ans := (n+1)*(maxFrequency-1) + maxFrequencyCount

	return max(ans, len(tasks))

}

func generateSolution(tasks []byte, n int) int {
	mapCount := [26]int{}
	for _, t := range tasks {
		cur := int(t - 'A')
		mapCount[cur]++
	}

	keys := []int{}
	for idx, val := range mapCount {
		if val > 0 {
			keys = append(keys, idx)
		}
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return mapCount[i] > mapCount[j]
	})

	c := len(keys)
	fmt.Println(keys)
	ans := 0
	for c > 0 {
		count := 0
		for i := 0; i <= n && i < len(keys); i++ {
			fmt.Printf("%d ", keys[i])
			mapCount[keys[i]]--
			count++
		}

		newQ := []int{}
		for _, val := range keys {
			if mapCount[val] > 0 {
				newQ = append(newQ, val)
			}
		}
		for count <= n && len(newQ) > 0 {
			count++
			fmt.Print("idle ")
		}
		ans += count
		fmt.Println()
		keys = newQ
		c = len(keys)
	}

	return ans
}
