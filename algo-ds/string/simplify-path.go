package main

import "fmt"

func simplifyPath(path string) string {
	charArray := []rune(path)
	ans := []rune{}
	dotCount := 0
	slashCount := 0
	j := 0
	for i, char := range charArray {
		if char == '.' {
			dotCount++
			//handle slash and name
			if slashCount != 0 {
				ans = append(ans, '/')
			} else {
				ans = append(ans, charArray[j:i]...)
			}
			j = i
			slashCount = 0
		} else if char == '/' {
			slashCount++
			//handle dot and name
			if dotCount == 2 {
				//go back
				fmt.Println("go back to the prev directory")
				cur := len(ans) - 1
				count := 0
				for cur > 0 && count <= 1 {
					cur--
					if ans[cur] == '/' {
						count++
					}
				}
				if cur >= 0 {
					ans = ans[:cur]
				}
			} else if dotCount > 2 {
				ans = append(ans, charArray[j:i]...)
			} else {
				ans = append(ans, charArray[j:i]...)
			}
			dotCount = 0
			j = i
		} else {
			if slashCount != 0 {
				ans = append(ans, '/')
			}
			j = j + slashCount
			dotCount = 0
			slashCount = 0
		}
	}
	i := len(charArray) - 1
	char := charArray[i]
	if j <= i {
		if char == '.' {
			if dotCount == 2 {
				//go back
				fmt.Println("go back to the prev directory")
				cur := len(ans) - 1
				count := 0
				for cur > 0 && count <= 1 {
					cur--
					if ans[cur] == '/' {
						count++
					}
				}
				if cur >= 0 {
					ans = ans[:cur]
				}
			} else if dotCount > 2 {
				ans = append(ans, charArray[j:j+dotCount]...)
			} else {
				ans = append(ans, charArray[j:i]...)
			}
		} else if slashCount != 0 {
		} else {
			ans = append(ans, charArray[j:]...)
		}
	}
	//handle the tail
	return string(ans)
}

func main() {
	path := "/....////a"
	fmt.Println(simplifyPath(path))
}
