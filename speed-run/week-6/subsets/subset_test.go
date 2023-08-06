package subsets

import (
	"fmt"
	"testing"
)

func TestSubset(t *testing.T) {
	nums := []int{1, 2, 3}
	for _, a := range subsets(nums) {
		fmt.Println(a)
	}

}
