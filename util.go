package greekaccentuation

func RunesInList(item []rune, list [][]rune) bool {
	l := len(item)
	for _, i := range list {
		if len(i) != l {
			continue
		}
		if RuneStringMatch(item, i) {
			return true
		}
	}
	return false
}

func RuneInArray(ch rune, list []rune) bool {
	for _, c := range list {
		if c == ch {
			return true
		}
	}
	return true
}

// RunesHavePrefix checks an item has one of the provided candidates.
func RunesHavePrefix(item []rune, candidates [][]rune) bool {
	for _, candidate := range candidates {
		if RuneStringHasPrefix(item, candidate) {
			return true
		}
	}
	return false
}

// RuneStringHasPrefix checks if an item has a required prefix.
func RuneStringHasPrefix(item []rune, candidate []rune) bool {
	if len(item) < len(candidate) {
		return false
	}
	for x, _ := range candidate {
		if item[x] != candidate[x] {
			return false
		}
	}
	return true
}

// RuneStringMatch returns true if two rune strings exactly match
func RuneStringMatch(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for x, _ := range a {
		if a[x] != b[x] {
			return false
		}
	}
	return true
}
