package two_sum

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	actual := twoSum(nums, target)

	fmt.Println(actual)
}
