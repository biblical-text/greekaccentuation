package greekaccentuation

import "testing"

func TestRunesInList(t *testing.T) {
	if !RunesInList([]rune("αχ"), [][]rune{[]rune("αχα"), []rune("αχ")}) {
		t.Fatal("RunesInListIsDipthong() failed")
	}
	if RunesInList([]rune("αχ"), [][]rune{[]rune("αχα"), []rune("αχι")}) {
		t.Fatal("RunesInListIsDipthong() failed")
	}
}

func TestRuneStringMatch(t *testing.T) {
	if !RuneStringMatch([]rune("αχ"), []rune("αχ")) {
		t.Fatal("RunesStringMatch() failed")
	}
	if RuneStringMatch([]rune("αχ"), []rune("αχα")) {
		t.Fatal("RunesStringMatch() failed")
	}
	if RuneStringMatch([]rune("αχ"), []rune("αα")) {
		t.Fatal("RunesStringMatch() failed")
	}
}

func TestRuneStringHasPrefix(t *testing.T) {
	if !RuneStringHasPrefix([]rune("αχαρφ"), []rune("αχ")) {
		t.Fatal("RuneStringHasPrefix() failed")
	}
	if !RuneStringHasPrefix([]rune("αχ"), []rune("α")) {
		t.Fatal("RuneStringHasPrefix() failed")
	}
	if !RuneStringHasPrefix([]rune("αχ"), []rune("")) {
		t.Fatal("RuneStringHasPrefix() failed")
	}
	if RuneStringHasPrefix([]rune("αχ"), []rune("ααχα")) {
		t.Fatal("RuneStringHasPrefix() failed")
	}
	if RuneStringHasPrefix([]rune("αχ"), []rune("αεφφ")) {
		t.Fatal("RuneStringHasPrefix() failed")
	}
	if RuneStringHasPrefix([]rune("βχ"), []rune("α")) {
		t.Fatal("RuneStringHasPrefix() failed")
	}
}

func TestRuneStringHasInfix(t *testing.T) {
	if !RuneStringHasInfix([]rune("αχαρφ"), []rune("")) {
		t.Fatal("RuneStringHasInfix() failed")
	}
	if !RuneStringHasInfix([]rune("αχαρφ"), []rune("αχ")) {
		t.Fatal("RuneStringHasInfix() failed")
	}
	if !RuneStringHasInfix([]rune("αχαρφ"), []rune("χ")) {
		t.Fatal("RuneStringHasInfix() failed")
	}
	if !RuneStringHasInfix([]rune("αχαρφ"), []rune("ρφ")) {
		t.Fatal("RuneStringHasInfix() failed")
	}
	if RuneStringHasInfix([]rune("αχαρφ"), []rune("ρεαιφ")) {
		t.Fatal("RuneStringHasInfix() failed")
	}
	if RuneStringHasInfix([]rune("αχαρφ"), []rune("ι")) {
		t.Fatal("RuneStringHasInfix() failed")
	}
	if !RuneStringHasInfix([]rune(""), []rune("")) {
		t.Fatal("RuneStringHasInfix() failed")
	}
}

func TestRunesHavePrefix(t *testing.T) {
	if !RunesHavePrefix([]rune("αχαρφ"), [][]rune{[]rune("αχ")}) {
		t.Fatal("RuneStringHasPrefix() failed")
	}
	if RunesHavePrefix([]rune("αχαρφ"), [][]rune{[]rune("χα")}) {
		t.Fatal("RuneStringHasPrefix() failed")
	}
	if !RunesHavePrefix([]rune("αχαρφ"), [][]rune{[]rune("χα"), []rune("αχ")}) {
		t.Fatal("RuneStringHasPrefix() failed")
	}
}
