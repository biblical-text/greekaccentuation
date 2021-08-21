package greekaccentuation

import "testing"

func TestRunesInList(t *testing.T) {
	if !runesInList([]rune("αχ"), [][]rune{[]rune("αχα"), []rune("αχ")}) {
		t.Fatal("RunesInListIsDipthong() failed")
	}
	if runesInList([]rune("αχ"), [][]rune{[]rune("αχα"), []rune("αχι")}) {
		t.Fatal("RunesInListIsDipthong() failed")
	}
}

func TestRuneStringMatch(t *testing.T) {
	if !runeStringMatch([]rune("αχ"), []rune("αχ")) {
		t.Fatal("RunesStringMatch() failed")
	}
	if runeStringMatch([]rune("αχ"), []rune("αχα")) {
		t.Fatal("RunesStringMatch() failed")
	}
	if runeStringMatch([]rune("αχ"), []rune("αα")) {
		t.Fatal("RunesStringMatch() failed")
	}
}

func TestRuneStringHasPrefix(t *testing.T) {
	if !runeStringHasPrefix([]rune("αχαρφ"), []rune("αχ")) {
		t.Fatal("RuneStringHasPrefix() failed")
	}
	if !runeStringHasPrefix([]rune("αχ"), []rune("α")) {
		t.Fatal("RuneStringHasPrefix() failed")
	}
	if !runeStringHasPrefix([]rune("αχ"), []rune("")) {
		t.Fatal("RuneStringHasPrefix() failed")
	}
	if runeStringHasPrefix([]rune("αχ"), []rune("ααχα")) {
		t.Fatal("RuneStringHasPrefix() failed")
	}
	if runeStringHasPrefix([]rune("αχ"), []rune("αεφφ")) {
		t.Fatal("RuneStringHasPrefix() failed")
	}
	if runeStringHasPrefix([]rune("βχ"), []rune("α")) {
		t.Fatal("RuneStringHasPrefix() failed")
	}
}

func TestRuneStringHasInfix(t *testing.T) {
	if !runeStringHasInfix([]rune("αχαρφ"), []rune("")) {
		t.Fatal("RuneStringHasInfix() failed")
	}
	if !runeStringHasInfix([]rune("αχαρφ"), []rune("αχ")) {
		t.Fatal("RuneStringHasInfix() failed")
	}
	if !runeStringHasInfix([]rune("αχαρφ"), []rune("χ")) {
		t.Fatal("RuneStringHasInfix() failed")
	}
	if !runeStringHasInfix([]rune("αχαρφ"), []rune("ρφ")) {
		t.Fatal("RuneStringHasInfix() failed")
	}
	if runeStringHasInfix([]rune("αχαρφ"), []rune("ρεαιφ")) {
		t.Fatal("RuneStringHasInfix() failed")
	}
	if runeStringHasInfix([]rune("αχαρφ"), []rune("ι")) {
		t.Fatal("RuneStringHasInfix() failed")
	}
	if !runeStringHasInfix([]rune(""), []rune("")) {
		t.Fatal("RuneStringHasInfix() failed")
	}
}

func TestRunesHavePrefix(t *testing.T) {
	if !runesHavePrefix([]rune("αχαρφ"), [][]rune{[]rune("αχ")}) {
		t.Fatal("RuneStringHasPrefix() failed")
	}
	if runesHavePrefix([]rune("αχαρφ"), [][]rune{[]rune("χα")}) {
		t.Fatal("RuneStringHasPrefix() failed")
	}
	if !runesHavePrefix([]rune("αχαρφ"), [][]rune{[]rune("χα"), []rune("αχ")}) {
		t.Fatal("RuneStringHasPrefix() failed")
	}
}
