package main

import "fmt"

type node struct {
	Leafs []node
	value int
}

func dfs_traversal(root node) {
	fmt.Printf("node %d\n", root.value)

	if root.Leafs == nil {
		return
	}

	for _, n := range root.Leafs {
		dfs_traversal(n)
	}
}

func test(list *[][]int) {
	*list = nil
	// list[0] = append(list[0], 1)
}

func test2(a *int) {
	*a = 1
}

func main() {
	// root := node{[]node{node{nil, 2}}, 1}
	// dfs_traversal(root)

	adjecencyList := [][]int{{2}, {4}, {1, 4}, {2}}

	list := [][]int{{1, 2}}
	test(&list)
	fmt.Println(list)

	a := 2

	test2(&a)
	fmt.Print(a)
}
