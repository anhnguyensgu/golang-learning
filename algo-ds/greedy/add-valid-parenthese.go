package main

import "fmt"

func makeValid(s string) int {
	ans := 0
	c := 0
	pos := [][2]int{}

	for i, r := range s {
		if r == '(' {
			c++
			fmt.Printf("start %d\n", i)
			fmt.Println(pos)
			pos = append(pos, [2]int{i,1})
		} else if c > 0 {
			c--
			fmt.Printf("close %d\n", i)
			pos = pos[:len(pos)-1]
		} else {
			ans++
			fmt.Printf("open %d\n", i)
			pos = append(pos, [2]int{i,0})
		}
	}

	if c > 0 {
		ans += c
	}
	fmt.Println(pos)

	return ans
}

func main() {
	fmt.Println(makeValid(")(()("))
}
