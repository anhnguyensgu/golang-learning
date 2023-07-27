package evaluate_reverse_polish_notation

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	//input := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	input := []string{"2", "1", "+", "3", "*"}
	//input := []string{"3", "11", "5", "+", "-"}
	//input := []string{"3", "11", "+", "5", "-"}
	ans := evalRPN(input)
	fmt.Println(ans)

	//input2 := []string{"4", "13", "5", "/", "+"}
	//ans2 := evalRPN(input2)
	//fmt.Println(ans2)
}
