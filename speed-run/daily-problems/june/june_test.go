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

func TestCommonChars(t *testing.T) {
	var words []string
	var ans []string

	words = []string{"bella", "label", "roller"}
	ans = commonChars(words)
	fmt.Println(ans)

}

func TestIsNStraightHand(t *testing.T) {
	var hand []int
	var groupSize int
	var actual, expect bool

	expect = true
	// hand = []int{1, 2, 3, 6, 2, 3, 4, 7, 8}
	hand = []int{1, 2, 3, 1, 2, 3}
	groupSize = 3
	if actual = isNStraightHand(hand, groupSize); actual != expect {
		t.Fail()
	}
}
