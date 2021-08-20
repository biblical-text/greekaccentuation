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
func SyllableAddAccent(s string, a Accent) string {
	// TODO
	/*
		o, n, c := onsetNucleusCoda(s)
		if isBreathing(o) {
			return AddDiacritic(AddDiacritic(n, o), a) + c
		} else {
			return o + string(AddDiacritic(n, a)) + c
		}

	*/
	return s
}

// AddAccentuation takes the syllables of a word and ....
func AddAccentuation(s []string, accentuation Accentuation) string {
	// TODO
	/*
		pos, accent := accentuation.Value()
		final := []string{""}
		if pos > 1 {
			final = s[1-pos:]
		}
		parts := append(s[:-pos], append(SyllableAddAccent(s[-pos], accent), final...)...)
		return strings.Join(parts, "")
	*/
	return ""
}

func DisplayAccentuation(accentuation Accentuation) string {
	return strings.ToLower(accentuation.Name())
}

func MakeOxytone(word string) string {
	return AddAccentuation(Syllabify(word), OXYTONE)
}

func MakeParoxytone(word string) string {
	return AddAccentuation(Syllabify(word), PAROXYTONE)
}

func MakeProparoxytone(word string) string {
	return AddAccentuation(Syllabify(word), PROPAROXYTONE)
}

func MakePerispomenon(word string) string {
	syllables := Syllabify(word)
	for _, i := range possibleAccentuations(syllables, true, false) {
		if i == PERISPOMENON {
			return AddAccentuation(syllables, PERISPOMENON)
		}
	}
	return AddAccentuation(syllables, OXYTONE)
}

func MakeProperispomenon(word string) string {
	syllables := Syllabify(word)
	if accentationInSet(PROPERISPOMENON, possibleAccentuations(syllables, true, false)) {
		return AddAccentuation(syllables, PROPERISPOMENON)
	}
	return AddAccentuation(syllables, PAROXYTONE)
}

func GetAccentuation(w string) Accentuation {
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
	//TODO
	/*

		ultima_length := SyllableLength(s[-1], treat_final_AI_OI_short)
		var penult_length string
		if len(s) >= 2 {
			penult_length = SyllableLength(s[-2], False)
		}
		if ultima_length == Length.UNKNOWN && default_short {
			ultima_length = Length.SHORT
		}
		if penult_length == Length.UNKNOWN && default_short {
			penult_length = Length.SHORT
		}

		yield = append(yield, OXYTONE)

		if !(ultima_length == Length.SHORT) {
			yield = append(yield, PERISPOMENON)
		}

		if len(s) >= 2 && !(penult_length == Length.LONG && ultima_length == Length.SHORT) {
			yield = append(yield, PAROXYTONE)
		}

		if len(s) >= 2 && !(penult_length == Length.SHORT || ultima_length == Length.LONG) {
			yield = append(yield, PROPERISPOMENON)
		}

		if len(s) >= 3 && !(ultima_length == Length.LONG) {
			yield = append(yield, PROPAROXYTONE)
		}

	*/
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
	return pre + AddAccentuation(s, ll[0])
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
		return pre + AddAccentuation(s, PROPERISPOMENON)
	}
	if accentationInSet(PAROXYTONE, accentuations) {
		return pre + AddAccentuation(s, PAROXYTONE)
	}
	if accentationInSet(OXYTONE, accentuations) { // fall back to an oxytone if necessary
		return pre + AddAccentuation(s, OXYTONE)
	}
	// TODO: It is unclear what happens in the python code here. Guess something.
	return w
}

//func Persistent(w string, lemma string, default_short=False) {
func Persistent(w string, lemma string, defaultShort bool) string {
	w = strings.ReplaceAll(w, "|", "")

	/*
	   accentuation := GetAccentuation(lemma)
	   if accentuation == NO_ACCENTUATION {
	       return ""
	   }

	   place, accent := accentuation.Value()
	   s := Syllabify(w)
	   possible := PossibleAccentuations(s, default_short)
	   place2 := len(s) - len(Syllabify(lemma)) + place
	   accent_pair := (place2, accent)
	   if accent_pair not in possible {
	       if accent == ACUTE and (place2, CIRCUMFLEX) in possible {
	           accent_pair = (place2, CIRCUMFLEX)
	       } else if accent == CIRCUMFLEX && (place2, ACUTE) in possible) {
	           accent_pair = (place2, Accent.ACUTE)
	       } else {
	           for i := 1; i <= 4; i++ {
	               if (place2 - i, ACUTE) in possible {
	                   accent_pair = (place2 - i, ACUTE)
	                   break
	               }
	           }
	       }
	   }
	   return AddAccentuation(s, Accentuation(accent_pair))
	*/
	return "" //todo
}

func accentationInSet(acentation Accentuation, set []Accentuation) bool {
	for _, i := range set {
		if acentation == i {
			return true
		}
	}
	return false
}
