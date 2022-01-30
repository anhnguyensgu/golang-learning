package main

import "fmt"

func URLify(a string, size int) string {
	count := 0
	for i := size - 1; i >= 0; i-- {
		if a[i] == ' ' {
			count++
		}
	}

	sizeAfterTransformation := size + count*2
	var ans = make([]rune, size+(count*2))
	j := sizeAfterTransformation - 1
	for i := size - 1; i >= 0; i-- {
		if a[i] == ' ' {
			ans[j] = '0'
			ans[j-1] = '2'
			ans[j-2] = '%'
			j = j - 3
		} else {
			ans[j] = rune(a[i])
			j--
		}
	}
	return string(ans)
}

func main() {
	fmt.Println(URLify("abc a", 5))
	fmt.Println(URLify("Mr John Smith", 13))
}
