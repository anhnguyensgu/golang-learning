package segmentree

type SegmentTree interface {
	Update(id, l, r, i, v int)
	Get(id, l, r, i, v int) int
}

type segmentTree struct {
	values []int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (s segmentTree) Update(id, l, r, i, v int) {
	if i < l || i > r {
		return
	}

	if l == r {
		s.values[id] = v
		return
	}

	mid := (l + r) / 2
	s.Update(id*2, l, mid, i, v)
	s.Update(id*2+1, mid+1, r, i, v)

	s.values[id] = max(s.values[id*2], s.values[id*2]+1)
}

func (s segmentTree) Get(id, l, r, i, v int) int {
	//TODO implement me
	panic("implement me")
}

func Init() SegmentTree {
	return &segmentTree{
		values: make([]int, 10000),
	}
}
