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
