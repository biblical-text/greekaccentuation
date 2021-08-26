package greekaccentuation

import (
	"golang.org/x/text/unicode/norm"
)

type RuneInterface interface {
	Rune() rune
}

type Breathing rune

const (
	NO_BREATHING Breathing = 0
	SMOOTH       Breathing = '\u0313'
	ROUGH        Breathing = '\u0314'
	PSILI        Breathing = SMOOTH //- GO doesnt handle this well
	DASIA        Breathing = ROUGH  //- GO doesnt handle this well
)

var Breathings []RuneInterface = []RuneInterface{SMOOTH, ROUGH}

func (e Breathing) Rune() rune {
	return rune(e)
}

func (e Breathing) Name() string {
	switch e {
	case NO_BREATHING:
		return "NO_BREATHING"
	case SMOOTH:
		return "SMOOTH"
	case ROUGH:
		return "ROUGH"
	default:
		return ""
	}
}

func isBreathing(c rune) bool {
	for b := range Breathings {
		if c == rune(b) {
			return true
		}

	}
	return false
}

type Accent rune

const (
	NO_ACCENT   Accent = 0
	ACUTE       Accent = '\u0301'
	GRAVE       Accent = '\u0300'
	CIRCUMFLEX  Accent = '\u0342'
	OXIA        Accent = ACUTE
	VARIA       Accent = GRAVE
	PERISPOMENI Accent = CIRCUMFLEX
)

var Accents []RuneInterface = []RuneInterface{ACUTE, GRAVE, CIRCUMFLEX, OXIA, VARIA, PERISPOMENI}

func (e Accent) Name() string {
	switch e {
	case NO_ACCENT:
		return "NO_ACCENT"
	case ACUTE:
		return "ACUTE"
	case GRAVE:
		return "GRAVE"
	case CIRCUMFLEX:
		return "CIRCUMFLEX"
	default:
		return ""
	}
}

func isAccent(c rune) bool {
	for a := range Accents {
		if c == rune(a) {
			return true
		}

	}
	return false
}

func (e Accent) Rune() rune {
	return rune(e)
}

type Diacritic rune

const (
	DIAERESIS Diacritic = '\u0308'
)

func (e Diacritic) Name() string {
	switch e {
	case DIAERESIS:
		return "DIAERESIS"
	default:
		return ""
	}
}

var Diacritics []RuneInterface = []RuneInterface{DIAERESIS}

func (e Diacritic) Rune() rune {
	return rune(e)
}

type Subscript rune

const (
	IOTA          Subscript = '\u0345'
	YPOGEGRAMMENI Subscript = IOTA
)

func (e Subscript) Name() string {
	switch e {
	case IOTA:
		return "IOTA"
	default:
		return ""
	}
}

var Subscripts []RuneInterface = []RuneInterface{IOTA}

func (e Subscript) Rune() rune {
	return rune(e)
}

type Length rune

const (
	SHORT   Length = '\u0306'
	LONG    Length = '\u0304'
	UNKNOWN Length = -1
)

var Lengths []RuneInterface = []RuneInterface{SHORT, LONG, UNKNOWN}

func (e Length) Name() string {
	switch e {
	case SHORT:
		return "SHORT"
	case LONG:
		return "LONG"
	default:
		return ""
	}
}

func (e Length) Rune() rune {
	return rune(e)
}

func Base(ch rune) rune {
	characters := []rune(norm.NFD.String(string([]rune{ch})))
	if len(characters) > 0 {
		return characters[0]
	}
	return 0
}

type ExtractDiacriticFunction func(ch rune) RuneInterface

// ExtractDiacritic accepts an Enum of Unicode diacritics, return a function that takes a
// Unicode character and returns the member of the collection the character
// has (or None).
//func ExtractDiacritic(Enum, unknownValue=None) ExtractDiacriticFunction {
func extractDiacritic(diacritics []RuneInterface, unknownValue RuneInterface) ExtractDiacriticFunction {
	return func(ch rune) RuneInterface {
		decomposedForm := []rune(norm.NFD.String(string([]rune{ch})))
		for _, diacritic := range diacritics {
			if runeInArray(diacritic.Rune(), decomposedForm) {
				return diacritic
			}
		}
		return unknownValue
	}
}

// AddDiacritic attaches the specified diacritic to the specified character
func AddDiacritic(base []rune, diacritic rune) []rune {
	return []rune(norm.NFC.String(string(append(base, diacritic))))
}

// AddBreathing attaches the specified breathing to a character
func AddBreathing(ch rune, breathing Breathing) rune {
	decomposed := []rune(norm.NFD.String(string([]rune{ch})))
	if len(decomposed) > 1 && decomposed[1] == LONG.Rune() {
		d := append([]rune{decomposed[0], decomposed[1], breathing.Rune()}, decomposed[2:]...)
		return []rune(norm.NFC.String(string(d)))[0]
	} else {
		d := append([]rune{decomposed[0], breathing.Rune()}, decomposed[1:]...)
		return []rune(norm.NFC.String(string(d)))[0]
	}
}

type RemoveDiacriticFunction func(text []rune) []rune

// Given an Enum of Unicode diacritics, return a function that takes a
// string and returns the string without those diacritics.
func removeDiacritic(diacritics []RuneInterface) RemoveDiacriticFunction {
	return func(text []rune) []rune {
		before := []rune(norm.NFD.String(string(text)))
		after := []rune{}
		for _, ch := range before {
			skip := false
			for _, d := range diacritics {
				if d.Rune() == ch {
					skip = true
				}
			}
			if !skip {
				after = append(after, ch)
			}
		}

		return []rune(norm.NFC.String(string(after)))
	}
}

var breathing = extractDiacritic(Breathings, nil)
var stripBreathing = removeDiacritic(Breathings)

var accent = extractDiacritic(Accents, nil)
var StripAccents = removeDiacritic(Accents)

var diaeresis = extractDiacritic(Diacritics, nil)

var iotaSubscript = extractDiacritic(Subscripts, nil)
var ypogegrammeni = iotaSubscript

var length = extractDiacritic(Lengths, UNKNOWN)
var stripLength = removeDiacritic(Lengths)

// If a circumflex, no need for macron indicating length
func removeRedundantMacron(word string) string {
	decomposed := []rune(norm.NFD.String(string(word)))

	if runeStringHasInfix(decomposed, []rune{'\u0304', '\u0342'}) {
		return string(stripLength(decomposed))
	} else {
		return word
	}
}
