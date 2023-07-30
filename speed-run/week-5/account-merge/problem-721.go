package accountmerge

type Set map[string]bool

func (s Set) contains(value string) bool {
	_, ok := s[value]
	return ok
}

func (s Set) add(value string) {
	s[value] = true
}

func accountsMerge(accounts [][]string) [][]string {
	ans := [][]string{}
	mapSet := map[string][]Set{}
	for _, acc := range accounts {
		name := acc[0]
		if emailSets, ok := mapSet[name]; ok {
			for _, emailSet := range emailSets {
				found := false
				for _, a := range acc {
					if emailSet.contains(a) {
						found = true
						break
					}
				}
				if found {
					for _, a := range acc {
						emailSet.add(a)
					}
				}
			}
		} else {
			for i := 1; i < len(acc); i++ {
			}
		}
	}
	return ans
}
