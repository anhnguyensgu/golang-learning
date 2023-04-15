package main

import "fmt"
func findKthPositive(arr []int, k int) int {   
  var a [10000]int
  for _, i := range arr {
    a[i]++
  }

  for i:= 1; i < len(a); i++ {
    if a[i] == 0 {
      k--
      if k == 0 {
        return i
      }
    }
  }
  return -1
}

func main() {
	fmt.Println(findKthPositive([]int{2,3,4,7,11}, 5))

}
