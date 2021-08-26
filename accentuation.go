package greekaccentuation

import (
	"sort"
	"strings"
)

type Accentuation int

const (
	NO_ACCENTUATION Accentuation = 0
	OXYTONE         Accentuation = 1
	PERISPOMENON    Accentuation = 2
	PAROXYTONE      Accentuation = 3
	PROPERISPOMENON Accentuation = 4
	PROPAROXYTONE   Accentuation = 5
)

var Accentuations []Accentuation = []Accentuation{OXYTONE, PERISPOMENON, PAROXYTONE, PROPERISPOMENON, PROPAROXYTONE}

type ByAccentReverse []Accentuation

func (a ByAccentReverse) Len() int           { return len(a) }
func (a ByAccentReverse) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAccentReverse) Less(i, j int) bool { return int(a[i]) > int(a[j]) }

func findMatchingAccentuation(position int, accent Accent) Accentuation {
	for _, a := range Accentuations {
		pos, acc := a.Value()
		if pos != position {
			continue
		}
		if acc != accent {
			continue
		}
		return a
	}
	return NO_ACCENTUATION
}

func accentuationInSet(a Accentuation, set []Accentuation) bool {
	for _, s := range set {
		if a == s {
			return true
		}
	}
	return false
}

func (e Accentuation) Value() (int, Accent) {
	switch e {
	case OXYTONE:
		return 1, ACUTE
	case PERISPOMENON:
		return 1, CIRCUMFLEX
	case PAROXYTONE:
		return 2, ACUTE
	case PROPERISPOMENON:
		return 2, CIRCUMFLEX
	case PROPAROXYTONE:
		return 3, ACUTE
	}
	return 0, 0
}

func (e Accentuation) Name() string {
	switch e {
	case OXYTONE:
		return "OXYTONE"
	case PERISPOMENON:
		return "PERISPOMENON"
	case PAROXYTONE:
		return "PAROXYTONE"
	case PROPERISPOMENON:
		return "PROPERISPOMENON"
	case PROPAROXYTONE:
		return "PROPAROXYTONE"
	}
	return ""
}

func (e Accentuation) Position() int {
	a, _ := e.Value()
	return a
}

func (e Accentuation) Character() Accent {
	_, b := e.Value()
	return b
}

// SyllableAddAccent places an accent on the correct location
// within a syllable
func syllableAddAccent(s string, a Accent) string {
	o, n, c := onsetNucleusCoda(s)
	ro := []rune(o)
	rn := []rune(n)
	if len(ro) == 1 && breathing(ro[0]) != nil {
		return string(AddDiacritic(AddDiacritic(rn, ro[0]), a.Rune())) + c
	} else {
		return o + string(AddDiacritic(rn, a.Rune())) + c
	}

}

// AddAccentuation takes the syllables of a word and ....
func addAccentuation(s []string, accentuation Accentuation) string {
	if len(s) == 0 {
		return ""
	}
	pos, accent := accentuation.Value()
	pre := ""
	final := ""
	idx := len(s) - pos // Place the accent on this syllable
	if idx+1 < len(s) { // a b c   3
		final = strings.Join(s[idx+1:], "")
	}
	if idx > 0 {
		pre = strings.Join(s[0:idx], "")
	}
	return pre + syllableAddAccent(s[len(s)-pos], accent) + final
}

func DisplayAccentuation(accentuation Accentuation) string {
	return strings.ToLower(accentuation.Name())
}

func MakeOxytone(word string) string {
	return addAccentuation(Syllabify(word), OXYTONE)
}

func MakeParoxytone(word string) string {
	return addAccentuation(Syllabify(word), PAROXYTONE)
}

func MakeProparoxytone(word string) string {
	return addAccentuation(Syllabify(word), PROPAROXYTONE)
}

func MakePerispomenon(word string) string {
	syllables := Syllabify(word)
	for _, i := range possibleAccentuations(syllables, true, false) {
		if i == PERISPOMENON {
			return addAccentuation(syllables, PERISPOMENON)
		}
	}
	return addAccentuation(syllables, OXYTONE)
}

func MakeProperispomenon(word string) string {
	syllables := Syllabify(word)
	if accentationInSet(PROPERISPOMENON, possibleAccentuations(syllables, true, false)) {
		return addAccentuation(syllables, PROPERISPOMENON)
	}
	return addAccentuation(syllables, PAROXYTONE)
}

func getAccentuation(w string) Accentuation {
	u := syllableAccent(ultima(w))
	if u == ACUTE {
		return OXYTONE
	} else if u == CIRCUMFLEX {
		return PERISPOMENON
	}
	p := penult(w)
	if p != "" {
		pa := syllableAccent(p)
		if pa == ACUTE {
			return PAROXYTONE
		} else if pa == CIRCUMFLEX {
			return PROPERISPOMENON
		}
	}
	a := antepenult(w)
	if a != "" {
		aa := syllableAccent(a)
		if aa == ACUTE {
			return PROPAROXYTONE
		}
	}
	return 0
}

//func possibleAccentuations(s []string, treat_final_AI_OI_short=True, default_short=False) {
func possibleAccentuations(s []string, treat_final_AI_OI_short bool, defaultShort bool) []Accentuation {
	var yield []Accentuation

	ultimaLength := syllableLength(s[len(s)-1], treat_final_AI_OI_short)
	var penultLength Length
	if len(s) >= 2 {
		penultLength = syllableLength(s[len(s)-2], false)
	}
	if ultimaLength == UNKNOWN && defaultShort {
		ultimaLength = SHORT
	}
	if penultLength == UNKNOWN && defaultShort {
		penultLength = SHORT
	}

	yield = append(yield, OXYTONE)

	if !(ultimaLength == SHORT) {
		yield = append(yield, PERISPOMENON)
	}

	if len(s) >= 2 && !(penultLength == LONG && ultimaLength == SHORT) {
		yield = append(yield, PAROXYTONE)
	}

	if len(s) >= 2 && !(penultLength == SHORT || ultimaLength == LONG) {
		yield = append(yield, PROPERISPOMENON)
	}

	if len(s) >= 3 && !(ultimaLength == LONG) {
		yield = append(yield, PROPAROXYTONE)
	}

	return yield
}

// Recessive does something that I am sure is interesting, but I don't know what.
//func Recessive(w string, treat_final_AI_OI_short=True, default_short=False)
func Recessive(w string, treat_final_AI_OI_short bool, default_short bool) string {
	pre := ""
	parts := strings.SplitN(w, "|", 2)
	if len(parts) > 1 {
		pre = parts[0]
		w = parts[1]
	}
	s := Syllabify(w)
	ll := possibleAccentuations(s, treat_final_AI_OI_short, default_short)
	sort.Sort(ByAccentReverse(ll))
	if len(ll) == 0 {
		return w
	}
	return pre + addAccentuation(s, ll[0])
}

//func OnPenult(w, default_short=False) {
func OnPenult(w string, default_short bool) string {
	pre := ""
	parts := strings.SplitN(w, "|", 2)
	if len(parts) > 1 {
		pre = parts[0]
		w = parts[1]
	}
	s := Syllabify(w)
	accentuations := possibleAccentuations(s, default_short, false)
	if accentationInSet(PROPERISPOMENON, accentuations) {
		return pre + addAccentuation(s, PROPERISPOMENON)
	}
	if accentationInSet(PAROXYTONE, accentuations) {
		return pre + addAccentuation(s, PAROXYTONE)
	}
	if accentationInSet(OXYTONE, accentuations) { // fall back to an oxytone if necessary
		return pre + addAccentuation(s, OXYTONE)
	}
	// TODO: It is unclear what happens in the python code here. Guess something.
	return w
}

// Persistent returns the accented form of a word. Returns an empty string
// if the dictionary entry contains no accent.
//func Persistent(w string, lemma string, default_short=False) {
func Persistent(word string, lemma string, defaultShort bool) string {
	w := strings.ReplaceAll(word, "|", "")

	// Get accentuation of the lemma
	accentuation := getAccentuation(lemma)
	if accentuation == NO_ACCENTUATION {
		// Why was this behaviour chosen? In this case I would prefer to
		// return an unaccented string for all alternate forms.
		return ""
	}
	place, accent := accentuation.Value()

	s := Syllabify(w)
	possible := possibleAccentuations(s, false, defaultShort)
	place2 := len(s) - len(Syllabify(lemma)) + place
	accentPair := findMatchingAccentuation(place2, accent)

	if !accentuationInSet(accentPair, possible) {
		opt1 := findMatchingAccentuation(place2, CIRCUMFLEX)
		opt2 := findMatchingAccentuation(place2, ACUTE)
		if accent == ACUTE && accentuationInSet(opt1, possible) {
			accentPair = opt1
		} else if accent == CIRCUMFLEX && accentuationInSet(opt2, possible) {
			accentPair = opt2
		} else {
			for i := 1; i <= 4; i++ {
				opt := findMatchingAccentuation(place2-i, ACUTE)
				if accentuationInSet(opt, possible) {
					accentPair = opt
					break
				}
			}
		}
	}

	return addAccentuation(s, Accentuation(accentPair))
}

func accentationInSet(acentation Accentuation, set []Accentuation) bool {
	for _, i := range set {
		if acentation == i {
			return true
		}
	}
	return false
}
