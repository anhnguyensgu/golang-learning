package taskscheduler

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B', 'C', 'C', 'C', 'D', 'D', 'E'}
	n := 2

	ans := leastInterval(tasks, n)
	fmt.Println(ans)
}
