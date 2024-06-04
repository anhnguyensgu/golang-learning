package june

import (
	"fmt"
	"testing"
)

func TestAppendCharacter(tctx *testing.T) {
	var s, t string
	var expect, actual int

	s = "coaching"
	t = "coding"
	expect = 4
	actual = appendCharacters(s, t)
	if expect != actual {
		fmt.Printf("expect %v but actual %v\n", expect, actual)
		tctx.Fail()
	}

	s = "abcde"
	t = "a"
	expect = 0
	actual = appendCharacters(s, t)
	if expect != actual {
		fmt.Printf("expect %v but actual %v\n", expect, actual)
		tctx.Fail()
	}

	s = "z"
	t = "abcde"
	expect = 5
	actual = appendCharacters(s, t)
	if expect != actual {
		fmt.Printf("expect %v but actual %v\n", expect, actual)
		tctx.Fail()
	}
}

func TestLongestPalindrom(t *testing.T) {
	var s string
	var expect, actual int

	s = "abccccdd"
	expect = 7
	actual = longestPalindrome(s)
	if expect != actual {
		t.Fail()
	}

	s = "a"
	expect = 1
	actual = longestPalindrome(s)
	if expect != actual {
		t.Fail()
	}

	s = "aaa"
	expect = 3
	actual = longestPalindrome(s)
	if expect != actual {
		t.Fail()
	}
}
