package greekaccentuation

import (
	"golang.org/x/text/unicode/norm"
)

type RuneInterface interface {
	Rune() rune
}

type Breathing rune

const (
	SMOOTH Breathing = '\u0313'
	ROUGH  Breathing = '\u0314'
	PSILI  Breathing = SMOOTH
	DASIA  Breathing = ROUGH
)

var Breathings []RuneInterface = []RuneInterface{SMOOTH, ROUGH, PSILI, DASIA}

func (e Breathing) Rune() rune {
	return rune(e)
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

var Diacritics []RuneInterface = []RuneInterface{DIAERESIS}

func (e Diacritic) Rune() rune {
	return rune(e)
}

type Subscript rune

const (
	IOTA          Subscript = '\u0345'
	YPOGEGRAMMENI Subscript = IOTA
)

var Subscripts []RuneInterface = []RuneInterface{IOTA, YPOGEGRAMMENI}

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
func ExtractDiacritic(diacritics []RuneInterface, unknownValue RuneInterface) ExtractDiacriticFunction {
	return func(ch rune) RuneInterface {
		decomposedForm := []rune(norm.NFD.String(string([]rune{ch})))
		for _, diacritic := range diacritics {
			if RuneInArray(diacritic.Rune(), decomposedForm) {
				return diacritic
			}
		}
		return unknownValue
	}
}

// AddDiacritic attaches the specified diacritic to the specified character
func AddDiacritic(base rune, diacritic rune) rune {
	c := []rune(norm.NFC.String(string([]rune{base, diacritic})))
	return c[0]
}

// AddBreathing attaches the specified breathing to a character
func AddBreathing(ch rune, breathing Breathing) rune {
	decomposed := []rune(norm.NFD.String(string([]rune{ch})))
	if len(decomposed) > 1 && decomposed[1] == LONG.Rune() {
		d := append(append(decomposed[0:2], breathing.Rune()), decomposed[2:]...)
		return []rune(norm.NFC.String(string(d)))[0]
	} else {
		d := append(append(decomposed[0:1], breathing.Rune()), decomposed[1:]...)
		return []rune(norm.NFC.String(string(d)))[0]
	}
}

type RemoveDiacriticFunction func(text []rune) []rune

// Given an Enum of Unicode diacritics, return a function that takes a
// string and returns the string without those diacritics.
func RemoveDiacritic(diacritics []RuneInterface) RemoveDiacriticFunction {
	return func(text []rune) []rune {
		return text
		/*
		   return unicodedata.normalize("NFC", "".join(
		       ch
		       for ch in unicodedata.normalize("NFD", text)
		       if ch not in diacritics)
		   )
		*/
	}
}

var breathing = ExtractDiacritic(Breathings, nil)
var stripBreathing = RemoveDiacritic(Breathings)

var accent = ExtractDiacritic(Accents, nil)
var stripAccents = RemoveDiacritic(Accents)

var diaeresis = ExtractDiacritic(Diacritics, nil)

var iotaSubscript = ExtractDiacritic(Subscripts, nil)
var ypogegrammeni = iotaSubscript

var length = ExtractDiacritic(Lengths, UNKNOWN)
var stripLength = RemoveDiacritic(Lengths)

// If a circumflex, no need for macron indicating length
func removeRedundantMacron(word string) string {
	decomposed := []rune(norm.NFD.String(string(word)))

	if RuneStringHasInfix(decomposed, []rune{'\u0304', '\u0342'}) {
		return string(stripLength(decomposed))
	} else {
		return word
	}
}
