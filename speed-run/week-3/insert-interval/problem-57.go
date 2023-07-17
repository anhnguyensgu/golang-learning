package insertinterval

type Interval [][]int

func (a Interval) Len() int {
	return len(a)
}

func (a Interval) Less(i, j int) bool {
	return a[i][0] < a[j][0]
}

func (a Interval) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	ans := [][]int{}

	for i := 0; i < len(intervals); i++ {
		if newInterval[0] > intervals[i][0] {
			if newInterval[0] <= intervals[i][1] {
				newInterval[0] = intervals[i][0]
				if newInterval[1] < intervals[i][1] {
					newInterval[1] = intervals[i][1]
				}
			} else {
				ans = append(ans, intervals[i])
			}
		} else {
			if newInterval[1] >= intervals[i][0] {
				if newInterval[1] <= intervals[i][1] {
					newInterval[1] = intervals[i][1]
				}
			} else {
				ans = append(ans, newInterval)
				newInterval = intervals[i]
			}
		}
	}

	ans = append(ans, newInterval)

	return ans
}
