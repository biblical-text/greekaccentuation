package greekaccentuation

import "golang.org/x/text/unicode/norm"

type Breathing rune

const (
    SMOOTH Breathing  = '\u0313'
    ROUGH Breathing = '\u0314'
    PSILI Breathing = SMOOTH
    DASIA Breathing = ROUGH
)

func (e Breathing) Rune() rune {
    return rune(e)
}

func isBreathing(c rune) bool {
    return c == rune(SMOOTH) || c == rune(ROUGH)
}

type Accent rune 

const (
    ACUTE Accent = '\u0301'
    GRAVE Accent = '\u0300'
    CIRCUMFLEX Accent = '\u0342'
    OXIA Accent = ACUTE
    VARIA Accent = GRAVE
    PERISPOMENI Accent = CIRCUMFLEX
)
func (e Accent) Rune() rune {
    return rune(e)
}


type Diacritic rune

const (
    DIAERESIS Diacritic = '\u0308'
)

func (e Diacritic) Rune() rune {
    return rune(e)
}

type Subscript rune

const (
    IOTA Subscript = '\u0345'
    YPOGEGRAMMENI Subscript = IOTA
)
func (e Subscript) Rune() rune {
    return rune(e)
}

type Length rune

const( 
    SHORT Length= '\u0306'
    LONG Length = '\u0304'
    UNKNOWN Length = -1
)
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