package greekaccentuation

import (
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

// IsVowel returns true if a character is a vowel. Accents
// and iota subscripts are ignored
func IsVowel(ch rune) bool {
	switch Base(ch) {
	case 'α', 'ε', 'η', 'ι', 'ο', 'υ', 'ω':
		return true
	default:
		return false
	}
}

// IsDipthong tests if a rune string is a valid dipthong
func isDipthong(a, b rune) bool {
	a = unicode.ToLower(a)
	b = unicode.ToLower(b)
	return RunesInList([]rune{a, b}, [][]rune{
		[]rune("αι"),
		[]rune("ει"),
		[]rune("οι"),
		[]rune("υι"),
		[]rune("αυ"),
		[]rune("ευ"),
		[]rune("ου"),
		[]rune("ηυ"),
	})
}

// isValidConsonantCluster returns true if this consonant
// combination would be considered valid.
func isValidConsonantCluster(ch rune, syllable []rune) bool {
	candidate := append([]rune{ch}, syllable...)
	return RunesHavePrefix(candidate, [][]rune{
		[]rune("βδ"), []rune("βλ"), []rune("βρ"),
		[]rune("γλ"), []rune("γν"), []rune("γρ"),
		[]rune("δρ"),
		[]rune("θλ"), []rune("θν"), []rune("θρ"),
		[]rune("κλ"), []rune("κν"), []rune("κρ"), []rune("κτ"),
		[]rune("μν"),
		[]rune("πλ"), []rune("πν"), []rune("πρ"), []rune("πτ"),
		[]rune("σβ"), []rune("σθ"), []rune("σκ"), []rune("σμ"),
		[]rune("σπ"), []rune("στ"), []rune("σφ"), []rune("σχ"), []rune("στρ"),
		[]rune("φθ"), []rune("φλ"), []rune("φ"),
		[]rune("χλ"), []rune("χρ"),
	})
}

// DisplayWord is a helper function that displays a syllable
// array as a string.
func DisplayWord(parts []string) string {
	return strings.Join(parts, ".")
}

// Syllabify splits a word into a string array of syllables.
func Syllabify(word string) []string {
	characters := []rune(norm.NFC.String(word))
	state := 0
	currentSyllable := []rune{}
	result := []string{}
	for i := len(characters) - 1; i >= 0; i-- {
		ch := characters[i]
		switch state {
		case 0:
			currentSyllable = append([]rune{ch}, currentSyllable...)
			if IsVowel(ch) {
				state = 1
			}
		case 1:
			if IsVowel(ch) {
				if isDipthong(ch, currentSyllable[0]) {
					if len(currentSyllable) > 1 && currentSyllable[1] == 'ι' {
						result = append([]string{string(currentSyllable[1:])}, result...)
						currentSyllable = append([]rune{ch}, currentSyllable[0])
					} else {
						currentSyllable = append([]rune{ch}, currentSyllable...)
					}
				} else {
					result = append([]string{string(currentSyllable)}, result...)
					currentSyllable = []rune{ch}
				}
			} else {
				currentSyllable = append([]rune{ch}, currentSyllable...)
				state = 2
			}
		case 2:
			if IsVowel(ch) {
				result = append([]string{string(currentSyllable)}, result...)
				currentSyllable = []rune{ch}
				state = 1
			} else {
				if isValidConsonantCluster(ch, currentSyllable) {
					currentSyllable = append([]rune{ch}, currentSyllable...)
				} else {
					result = append([]string{string(currentSyllable)}, result...)
					currentSyllable = []rune{ch}
					state = 0
				}
			}
		}
	}
	result = append([]string{string(currentSyllable)}, result...)
	return result
}

// ultima returns the last syllable, or an empty string
func ultima(word string) string {
	s := Syllabify(word)
	if len(s) == 0 {
		return ""
	}
	return s[len(s)-1]
}

// penult returns the second last syllable, or an empty string
func penult(word string) string {
	s := Syllabify(word)
	if len(s) < 2 {
		return ""
	}
	return s[len(s)-2]
}

// antepenult returns the third last syllable, or an empty string
func antepenult(word string) string {
	s := Syllabify(word)
	if len(s) < 3 {
		return ""
	}
	return s[len(s)-3]
}

func onset(s string) string {
	o, _, _ := onsetNucleusCoda(s)
	return o
}

func nucleus(s string) string {
	_, n, _ := onsetNucleusCoda(s)
	return n
}

func coda(s string) string {
	_, _, c := onsetNucleusCoda(s)
	return c
}

// onsetNucleusCoda splits the parts of a syllable to facilitate accentation
func onsetNucleusCoda(s string) (string, string, string) {
	onset := ""
	nucleus := ""
	coda := ""
	// TODO
	/*
		    for i, ch := range []rune(s) {
		        if IsVowel(ch) {
		            if i == 0 && breathing(ch) {
		                onset = breathing(ch)
		                break
					} else if i == 0 && len(s) > 1 && breathing(s[1]) {
		                onset = breathing(s[1])
		                break
					} else {
						if i > 0 {
							onset = s[:i]
						}
		                break
					}
				}
			}
		    if onset == "" {
		        return s, "", ""
			}
		    for j, ch := range []rune(s[i:]) {
		        if !isVowel(ch) && !isBreathing(ch) {
		            nucleus = s[i:i + j]
		            coda = s[i + j:]
		            break
				}
			}
		    if nucleus == "" {
		        nucleus = s[i:]
		        coda = ""
			}
		    if isinstance(onset, Breathing) {
		        nucleus = strip_breathing(nucleus)
			}
	*/
	return onset, nucleus, coda
}

func rime(s string) string {
	_, n, c := onsetNucleusCoda(s)
	return n + c
}

func body(s string) string {
	o, n, _ := onsetNucleusCoda(s)
	// TODO
	/*
		if isBreathing(o, Breathing) {
			return addNecessaryBreathing(n, o)
		}
	*/
	return o + n
}

func syllableLength(s string, final string) Length {
	// TODO
	return UNKNOWN
}

func syllableAccent(s string) Accent {
	n := nucleus(s)
	if n != "" {
		for _, ch := range []rune(n) {
			a := accent(ch)
			if a != nil {
				return a.(Accent)
			}
		}
	}
	return 0
}

func rebreath(word string) string {
	if word == "" {
		return ""
	}
	if strings.HasPrefix(word, "h") {
		word = addNecessaryBreathing(word[1:], ROUGH)
	} else {
		word = addNecessaryBreathing(word, SMOOTH)
	}
	word = removeRedundantMacron(word)
	return word
}

func addNecessaryBreathing(w string, breathing Breathing) string {
	//TODO
	return w
}

// RemoveAccentFromRune strips all accents from a character. Returns true
// if the character is valid for a standard greek string.
func RemoveAccentFromRune(r rune) (rune, bool) {
	switch r {
	case 'α', 'ᾰ', 'ά', 'ἀ', 'ά', 'ὰ', 'ἄ', 'ἁ', 'ᾶ', 'ᾴ', 'ᾳ', 'ἆ', 'ἅ', 'ᾄ', 'ᾅ', 'ᾷ', 'ἃ', 'ἂ', 'ᾀ':
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
