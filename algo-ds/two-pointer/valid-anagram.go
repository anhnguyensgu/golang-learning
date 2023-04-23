package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapChar := map[byte]int{}
	for i := range s {
		mapChar[s[i]]++
		mapChar[t[i]]--
	}

	for _, count := range mapChar {
		if count != 0 {
			return false
		}
	}
	return true
}

func main() {

}
