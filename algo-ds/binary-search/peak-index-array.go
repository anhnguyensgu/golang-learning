package main

import "fmt"

func peakIndexInMountainArray(arr []int) int {
	l := 0
	r := len(arr) - 1

	for l <= r {
		mid := l + (r-l)/2

		if mid == 0 || (arr[mid] > arr[mid-1] && arr[mid] < arr[mid+1]) {
			fmt.Println("on the left")
			l = mid + 1
		} else if mid == len(arr)-1 || (arr[mid] < arr[mid-1] && arr[mid] > arr[mid+1]) {

			fmt.Println("on the right")
			r = mid - 1
		} else {
			return mid
		}
		fmt.Printf("l %d r %d \n", l, r)
	}
	return -1
}
func main() {
	fmt.Println(peakIndexInMountainArray([]int{3, 5, 3, 2, 0}))
	// fmt.Println(peakIndexInMountainArray([]int{18, 29, 38, 59, 98, 100, 99, 98, 90}))
}
