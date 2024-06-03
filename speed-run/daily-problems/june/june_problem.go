package june

func appendCharacters(s string, t string) int {
	idx, subseqIdx := 0, 0
	for idx < len(s) && subseqIdx < len(t) {
		if s[idx] == t[subseqIdx] {
			subseqIdx++
		}
		idx++
	}
	if subseqIdx == len(t) {
		return 0
	}

	return len(t) - subseqIdx
}
