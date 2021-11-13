package greekaccentuation

func runesInList(item []rune, list [][]rune) bool {
	l := len(item)
	for _, i := range list {
		if len(i) != l {
			continue
		}
		if runeStringMatch(item, i) {
			return true
		}
	}
	return false
}

func runeInArray(ch rune, list []rune) bool {
	for _, c := range list {
		if c == ch {
			return true
		}
	}
	return false
}

// RunesHavePrefix checks an item has one of the provided candidates.
func runesHavePrefix(item []rune, candidates [][]rune) bool {
	for _, candidate := range candidates {
		if runeStringHasPrefix(item, candidate) {
			return true
		}
	}
	return false
}

// RuneStringHasPrefix checks if an item has a required prefix.
func runeStringHasPrefix(item []rune, candidate []rune) bool {
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

// RuneStringHasInfix checks if a rune array contains a string of runes
func runeStringHasInfix(item []rune, candidate []rune) bool {
	itemLength := len(item)
	if itemLength < len(candidate) {
		return false
	}
	if len(candidate) == 0 {
		return true
	}
	for x, _ := range item {
		match := true
		for y, _ := range candidate {
			if x+y >= itemLength {
				match = false
				break
			}
			if item[x+y] != candidate[y] {
				match = false
			}
		}
		if match {
			return true
		}
	}
	return false
}

// RuneStringMatch returns true if two rune strings exactly match
func runeStringMatch(a, b []rune) bool {
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

func RemoveAccentsFromString(s string) string {
	o := ""

	for _, c := range []rune(s) {
		r, _ := RemoveAccentFromRune(c)
		o = o + string(r)
	}

	return o
}

// RemoveAccentFromRune strips all accents from a character. Returns true
// if the character is valid for a standard greek string.
func RemoveAccentFromRune(r rune) (rune, bool) {
	switch r {
	case 'α', 'ᾰ', 'ά', 'ἀ', 'ά', 'ὰ', 'ἄ', 'ἁ', 'ᾶ', 'ᾴ', 'ᾳ', 'ἆ', 'ἅ', 'ᾄ', 'ᾅ', 'ᾷ', 'ἃ', 'ἂ', 'ᾀ', 'ᾱ':
		return 'α', true
	case 'Ἀ', 'A', 'Ἄ', 'Ἃ', 'Ἅ', 'ᾍ', 'Ἆ', 'Ἁ':
		return 'Α', true
	case 'Β':
		return 'Β', true
	case 'ῥ':
		return 'ρ', true
	case 'Δ':
		return 'δ', true
	case 'ε', 'έ', 'έ', 'ὲ', 'ἔ', 'ἑ', 'ἐ', 'ἕ', 'ἓ':
		return 'ε', true
	case 'Ε', 'Ἐ', 'Ἕ', 'Ἔ', 'Ἑ', 'Ἓ':
		return 'Ε', true
	case 'Ζ':
		return 'Ζ', true
	case 'η', 'ή', 'ῆ', 'ή', 'ὴ', 'ἡ', 'ἤ', 'ἦ', 'ἠ', 'ῄ', 'ῇ', 'ἧ', 'ἥ', 'ῃ', 'ἢ', 'ᾖ', 'ἣ', 'ᾗ', 'ᾐ', 'ᾔ', 'ᾑ':
		return 'η', true
	case 'Η', 'Ἡ', 'Ἤ', 'Ἢ', 'Ἥ', 'Ἠ', 'Ἦ', 'Ἣ', 'ᾜ':
		return 'Η', true
	case 'ι', 'ί', 'ί', 'ῖ', 'ἰ', 'ἴ', 'ἶ', 'ἱ', 'ϊ', 'ῒ', 'ἵ', 'ἷ', 'ΐ', 'ὶ', 'ἳ':
		return 'ι', true
	case 'Ι', 'Ἰ', 'Ἵ', 'Ἱ', 'Ἴ':
		return 'Ι', true
	case 'Κ':
		return 'κ', true
	case 'Μ':
		return 'μ', true
	case 'Ν':
		return 'ν', true
	case 'Ξ':
		return 'ξ', true
	case 'ο', 'ό', 'ό', 'ὁ', 'ὸ', 'ὄ', 'ὃ', 'ὀ', 'ὅ':
		return 'ο', true
	case 'Ο', 'Ὁ', 'Ὅ', 'Ὃ', 'Ὄ', 'Ὀ':
		return 'Ο', true
	case 'Ρ', 'Ῥ':
		return 'Ρ', true
	case 'T':
		return 'Τ', true
	case 'υ', 'ῡ', 'ὖ', 'ύ', 'ύ', 'ὺ', 'ὐ', 'ὑ', 'ῦ', 'ϋ', 'ὔ', 'ὕ', 'ὗ', 'ΰ', 'ὓ', 'ῢ', 'ὒ':
		return 'υ', true
	case 'Υ', 'Ὑ', 'Ὕ':
		return 'Υ', true
	case 'ω', 'ᾦ', 'ὦ', 'ὤ', 'ὠ', 'ὢ', 'ῶ', 'ώ', 'ὥ', 'ώ', 'ὼ', 'ᾠ', 'ὡ', 'ῴ', 'ῳ', 'ῷ', 'ὧ', 'ᾧ':
		return 'ω', true
	case 'Ὥ', 'Ὡ', 'ᾯ', 'Ὧ', 'Ὦ', 'Ὤ', 'Ὠ':
		return 'Ω', true
	case '῾', 0x0351, '᾽', '῞', '῝', '῎', '῍', 0x0300, 0x0301, 0x0302, 0x342:
		return 0, true
	case ' ', ',', '.', '-', '[', ']', '(', ')', ':', ';', '\'', '"':
		return r, true
	}
	return r, false
}
