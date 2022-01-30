package main

import (
	"fmt"
)

func numberOfBeams(bank []string) int {
	count := 0
	multi := 1
	i := 0
	for ; i < len(bank); i++ {
		cur := 0
		for _, c := range bank[i] {
			if c == '1' {
				cur = cur + 1
			}
		}
		if cur != 0 {
			multi = cur
			break
		}
	}

	for index, r := range bank {
		if index <= i {
			continue
		}

		cur := 0
		for _, c := range r {
			if c == '1' {
				cur = cur + 1
			}
		}
		if cur != 0 {
			count = multi*cur + count
			multi = cur
		}
	}
	return count
}

func main() {
	fmt.Println(numberOfBeams([]string{"000", "111", "000", "110"}))
}
