package courseschedule

func canFinish_dfs(numCourses int, prerequisites [][]int) bool {
	adj := [][]int{}
	for i := 0; i < numCourses; i++ {
		adj = append(adj, []int{})
	}

	for _, p := range prerequisites {
		adj[p[0]] = append(adj[p[0]], p[1])
	}

	visit := make([]bool, numCourses)
	inStack := make([]bool, numCourses)
	for i := 0; i < numCourses; i++ {
		if dfs(visit, inStack, adj, i) {
			return false
		}
	}

	return true
}

func dfs(visit, inStack []bool, adj [][]int, course int) bool {
	if inStack[course] {
		return true
	}
	if visit[course] {
		return false
	}

	visit[course] = true
	inStack[course] = true

	for _, neighbor := range adj[course] {
		if dfs(visit, inStack, adj, neighbor) {
			return true
		}
	}
	inStack[course] = false
	return false
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	inDegree := make([]int, numCourses)
	adj := make([][]int, numCourses)
	for _, p := range prerequisites {
		adj[p[0]] = append(adj[p[0]], p[1])
		inDegree[p[1]]++
	}

	q := []int{}

	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			q = append(q, i)
		}
	}

	visitedNode := 0

	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		visitedNode++
		for _, neighbor := range adj[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				q = append(q, neighbor)
			}
		}

	}

	return visitedNode == numCourses
}
