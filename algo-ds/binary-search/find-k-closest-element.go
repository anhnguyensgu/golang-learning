package main

func findClosestElements(arr []int, k int, x int) []int {
	l, r := 0, len(arr)-1

	for l <= r {
		mid := (l + r) / 2
		if arr[mid] == x {
			l = mid
			break
		}

		if arr[mid] > x {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	ans := []int{}

	i, j := l-1, l
	for k > 0 {
		if i == -1 {
			ans = append(ans, arr[j])
			j++
			k--
			continue
		}
		if j >= len(arr) {
			ans = append(ans, arr[i])
			i--
			k--
			continue
		}
		cur1 := arr[i] - x
		if cur1 < 0 {
			cur1 = -cur1
		}
		cur2 := arr[j] - x
		if cur2 < 0 {
			cur2 = -cur2
		}
		if cur1 <= cur2 {
			ans = append(ans, arr[i])
			i--
		} else {
			ans = append(ans, arr[j])
			j++
		}
		k--
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i] < ans[j]
	})
	return ans
}

func main() {

}
