package main

import "fmt"

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	if arr[0] > arr[1] {
		return false
	}
	ans := true

	isPeak := false

	for i := 1; i < len(arr); i++ {
		if isPeak {
			if arr[i-1] <= arr[i] {
				return false
			}
		} else {
			if arr[i-1] >= arr[i] {
				if !isPeak && arr[i-1] != arr[i] {
					isPeak = true
				} else {

					return false
				}
			}
		}
	}
	return ans && isPeak
}

func main() {
	fmt.Println(validMountainArray([]int{1, 2, 3, 1, 0}))
	fmt.Println(validMountainArray([]int{1, 2, 3, 3, 0}))
	fmt.Println(validMountainArray([]int{1, 2, 3, -1, 0}))
}
