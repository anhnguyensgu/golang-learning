package main

import "fmt"

var listAns []string

func generateHelper(n int, i int, curSol []rune) {
	if n*2 == i {
		c := 0
		for _, r := range curSol {
			if r == '(' {
				c++
			} else if c > 0 {
				c--
			} else {
				return
			}
		}
		if c == 0 {
			listAns = append(listAns, string(curSol))
		}
	} else {
		curSol[i] = '('
		generateHelper(n, i+1, curSol)
		curSol[i] = ')'
		generateHelper(n, i+1, curSol)
	}

}

func generateParenthesis(n int) []string {
	listAns = []string{}
	curSol := make([]rune, n*2)
	generateHelper(n, 0, curSol)
	return listAns
}

func main() {
	fmt.Println(len(generateParenthesis(3)))
}
