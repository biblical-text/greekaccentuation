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
	return runesInList([]rune{a, b}, [][]rune{
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
	return runesHavePrefix(candidate, [][]rune{
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
	characters := []rune(norm.NFD.String(word))
	state := 0
	currentSyllable := []rune{}
	result := []string{}
	// Walk backwards from the end of the string and break
	// at the appropriate positions.
	for i := len(characters) - 1; i >= 0; i-- {
		ch := characters[i]
		//fmt.Println(i, string(ch), state, "isvowel", IsVowel(ch), "--", word, " <-->", string(currentSyllable), result)
		switch state {
		case 0:
			// Eat characters until we have eaten our first vowel, then change state
			currentSyllable = append([]rune{ch}, currentSyllable...)
			if IsVowel(ch) {
				state = 1
			}
		case 1:
			// We have eaten a vowel, now just take in legitimate vowel combinations
			// or the consonante that appears at the start of the syllable. ἴαμα
			if IsVowel(ch) || ch == ROUGH.Rune() || ch == ACUTE.Rune() || ch == SMOOTH.Rune() {
				if currentSyllable[0] == ACUTE.Rune() {
					currentSyllable = append([]rune{ch}, currentSyllable...)
				} else if currentSyllable[0] == ROUGH.Rune() {
					currentSyllable = append([]rune{ch}, currentSyllable...)
				} else if currentSyllable[0] == SMOOTH.Rune() {
					currentSyllable = append([]rune{ch}, currentSyllable...)
				} else if isDipthong(ch, currentSyllable[0]) {
					if len(currentSyllable) > 1 && (currentSyllable[1] == 'ι' || currentSyllable[1] == 'Ι') {
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
			// We have eaten a full syllable, but we might need to eat a
			// preceeding consonant.
			if IsVowel(ch) || ch == ACUTE.Rune() {
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
	for i, _ := range result {
		result[i] = norm.NFC.String(result[i])
	}
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

// onsetNucleusCoda splits the parts of a syllable to facilitate accentation.
// Returns composed (not decomposed) unicode format
func onsetNucleusCoda(s string) (string, string, string) {
	letters := []rune(norm.NFC.String(s))
	var onset []rune
	var nucleus []rune
	var coda []rune

	if s == "" {
		return "", "", ""
	}

	li := -1
	didBreak := false
	for i, ch := range letters {
		li = i
		if IsVowel(ch) {
			if i == 0 && breathing(ch) != nil {
				onset = []rune{breathing(letters[0]).Rune()}
				didBreak = true
				break
			} else if i == 0 && len(letters) > 1 && breathing(letters[1]) != nil {
				onset = []rune{breathing(letters[1]).Rune()}
				didBreak = true
				break
			} else {
				if i > 0 {
					onset = letters[:i]
				} else {
					onset = []rune{}
				}
				didBreak = true
				break
			}
		}
	}
	if !didBreak {
		return s, "", ""
	}

	for j, ch := range letters[li:] {
		if !IsVowel(ch) && !isBreathing(ch) {
			nucleus = letters[li : li+j]
			coda = letters[li+j:]
			break
		}
	}
	if len(nucleus) == 0 {
		nucleus = letters[li:]
	}
	if len(onset) == 1 && breathing(onset[0]) != nil {
		nucleus = stripBreathing(nucleus)
	}

	return string(onset), string(nucleus), string(coda)
}

func rime(s string) string {
	_, n, c := onsetNucleusCoda(s)
	return n + c
}

func body(s string) string {
	o, n, _ := onsetNucleusCoda(s)
	ro := []rune(o)
	if len(ro) == 1 && breathing(ro[0]) != nil {
		return addNecessaryBreathing(n, Breathing(breathing(ro[0]).Rune()))
	}
	return o + n
}

func syllableLength(s string, finalPosition ...bool) Length {
	n := []rune(nucleus(s)) // Middle part of syllable

	if len(n) == 0 {
		// TODO: I dont know if a hard fail is important here
		//raise ValueError("'{}' does not contain a nucleus".format(s))
		return UNKNOWN
	}

	r := rime(s) // Middle and last part of syllable

	if len(n) > 1 {
		var b []rune
		for _, ch := range []rune(r) {
			b = append(b, Base(ch))
		}

		if len(finalPosition) > 0 {
			if finalPosition[0] {
				if runesInList(b, [][]rune{[]rune("αι"), []rune("οι")}) {
					return SHORT
				} else {
					return LONG
				}
			} else {
				return LONG
			}
		} else {
			if runesInList(b, [][]rune{[]rune("αι"), []rune("οι")}) {
				return UNKNOWN
			} else {
				return LONG
			}
		}
	} else {
		rn := n[0]
		if iotaSubscript(rn) == IOTA {
			return LONG
		} else {
			b := Base(rn)
			if b == 'ε' || b == 'ο' || length(rn) == SHORT {
				return SHORT
			} else if b == 'η' || b == 'ω' || length(rn) == LONG {
				return LONG
			} else { // αιυ
				return UNKNOWN
			}
		}
	}
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

func Rebreath(word string) string {
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

//
func addNecessaryBreathing(w string, breathing Breathing) string {
	if w == "" {
		return w
	}
	s := Syllabify(w)
	if len(s) == 0 {
		return w
	}

	o, ns, c := onsetNucleusCoda(s[0])
	n := []rune(ns)
	if o == "" {
		var lastVowel int = -1
		var pre []rune
		var post []rune
		for i, ch := range n {
			if runeInArray(unicode.ToLower(Base(ch)), []rune("αεηιουω")) {
				lastVowel = i
			}
		}
		if lastVowel > 0 {
			pre = n[0:lastVowel]
		}
		if lastVowel+1 < len(n) {
			post = n[lastVowel+1:]
		}
		n = append(append(pre, AddBreathing(n[lastVowel], breathing)), post...)
		k := o + string(n) + c + strings.Join(s[1:], "")
		return norm.NFKC.String(k)
	} else {
		return w
	}
}
