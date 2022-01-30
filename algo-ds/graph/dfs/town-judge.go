package main

var degrees [10005]int

func findJudge(n int, trust [][]int) int {
	t1 := make([]int, n) // trusted by someone
	t2 := make([]int, n) // someone you trusted

	for _, r := range trust {
		t1[r[0] - 1]++
		t2[r[1] - 1]++
	}

	for i:= 0; i < n; i++ {
		if t1[i] == 0 && t2[i] == n - 1{
			return i + 1
		}
	}
	return -1
}

func main() {

}