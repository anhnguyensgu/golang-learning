package partitionequalsubsetsum

import (
	"fmt"
	"testing"
)

func TestPartition(t *testing.T) {
	/* nums := []int{2, 2, 1, 1}
	fmt.Println(canPartition(nums))
	nums = []int{1, 2, 5}
	fmt.Println(canPartition(nums)) */
	nums := []int{14, 9, 8, 4, 3, 2}
	fmt.Println(canPartition(nums))
}
