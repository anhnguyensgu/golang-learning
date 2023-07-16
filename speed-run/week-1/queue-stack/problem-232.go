package queuestack

type stack []int

func (s *stack) add(x int) {
	*s = append(*s, x)
}

func (s *stack) pop() int {
	a := *s
	r := a[len(a)-1]
	*s = a[0 : len(a)-1]
	return r
}

func (s *stack) peek() int {
	return (*s)[len(*s)-1]
}

func (s *stack) empty() bool {
	return len(*s) == 0
}

// type MyQueue struct {
// 	s1 *stack
// 	s2 *stack
// }

// func Constructor() MyQueue {
// 	return MyQueue{
// 		s1: &stack{},
// 		s2: &stack{},
// 	}
// }

// func (this *MyQueue) Push(x int) {
// 	this.s1.add(x)
// 	for !this.s1.empty() {
// 		this.s2.add(this.s1.pop())
// 	}
// 	for !this.s2.empty() {
// 		this.s1.add(this.s2.pop())
// 	}
// }

// func (this *MyQueue) Pop() int {
// 	return this.s1.pop()
// }

// func (this *MyQueue) Peek() int {
// 	return this.s1.peek()
// }

// func (this *MyQueue) Empty() bool {
// 	return this.s1.empty()
// }

// Solution 2
type MyQueue struct {
	s1 *stack
	s2 *stack
}

func Constructor() MyQueue {
	return MyQueue{
		s1: &stack{},
		s2: &stack{},
	}
}

func (this *MyQueue) Push(x int) {
	this.s1.add(x)
}

func (this *MyQueue) Pop() int {
	if this.s2.empty() {
		for !this.s1.empty() {
			this.s2.add(this.s1.pop())
		}
	}
	return this.s2.pop()
}

func (this *MyQueue) Peek() int {
	if this.s2.empty() {
		return (*this.s1)[0]
	}
	return this.s2.peek()
}

func (this *MyQueue) Empty() bool {
	return this.s1.empty() && this.s2.empty()
}
