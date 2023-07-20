package longestsubstringwithoutreaptingcharacters

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	s := "aaa"
	s2 := "abc"
	s3 := "aba"
	fmt.Println(lengthOfLongestSubstring(s))
	fmt.Println(lengthOfLongestSubstring(s2))
	fmt.Println(lengthOfLongestSubstring(s3))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
