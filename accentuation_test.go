package greekaccentuation

import (
	"fmt"
	"testing"

	"golang.org/x/text/unicode/norm"
)

func TestSyllableAddAccent(t *testing.T) {
	if syllableAddAccent("κος", ACUTE) != "κός" {
		t.Fatalf("syllableAddAccent() failed.")
	}
	if syllableAddAccent("ος", ACUTE) != "ός" {
		t.Fatalf("syllableAddAccent() failed.")
	}
	if syllableAddAccent("φως", CIRCUMFLEX) != "φῶς" {
		t.Fatalf("syllableAddAccent() failed.")
	}
	if syllableAddAccent("ου", CIRCUMFLEX) != "οῦ" {
		t.Fatalf("syllableAddAccent() failed. Returned: %v",
			syllableAddAccent("ου", CIRCUMFLEX))
	}
	if syllableAddAccent("ἀν", ACUTE) != "ἄν" {
		t.Fatalf("syllableAddAccent() failed. Returned: %v",
			syllableAddAccent("ἀν", ACUTE))
	}
}

func TestAddAccentuation(t *testing.T) {
	if addAccentuation(Syllabify("θεος"), OXYTONE) != "θεός" {
		t.Fatalf("AddAccentuation() failed. %v != %v", "θεός",
			addAccentuation(Syllabify("θεος"), OXYTONE))
	}
	if addAccentuation(Syllabify("ἀνθρωπος"), PROPAROXYTONE) != "ἄνθρωπος" {
		t.Fatalf("AddAccentuation() failed. ἄνθρωπος != %v",
			addAccentuation(Syllabify("ἄνθρωπος"), PROPAROXYTONE))
	}
	if addAccentuation(Syllabify("Ἰαρεδ"), OXYTONE) != "Ἰαρέδ" {
		t.Fatalf("AddAccentuation() failed. Ἰαρέδ != %v",
			addAccentuation(Syllabify("Ἰαρεδ"), OXYTONE))
	}
}

func TestPossibleAccentuations(t *testing.T) {
	{
		s := Syllabify("εγινωσκου")
		accentuations := possibleAccentuations(s, true, false)
		possible := []string{}
		for _, a := range accentuations {
			possible = append(possible, addAccentuation(s, a))
		}
		expected := []string{"εγινωσκού", "εγινωσκοῦ", "εγινώσκου"}
		if !stringArrayMatch(possible, expected) {
			t.Fatalf("possibleAccentuations() failed. %v != %v", possible, expected)
		}
	}
}

func TestPersistent(t *testing.T) {
	if Persistent("ἀνθρωπος", "ἄνθρωπος", false) != "ἄνθρωπος" {
		t.Fatalf("Persistent() failed. Returned %s", Persistent("ἀνθρωπος", "ἄνθρωπος", false))
	}
	if Persistent("ἀνθρωπου", "ἄνθρωπος", false) != "ἀνθρώπου" {
		t.Fatalf("Persistent() failed.")
	}
	if Persistent("καταβαινον", "καταβαίνων", false) != "καταβαῖνον" {
		t.Fatalf("Persistent() failed.")
	}
	// If unaccented, return a blank string. (Why?)
	if Persistent("Ααρων", "Ααρων", false) != "" {
		t.Fatalf("Persistent() failed. Returned %s", Persistent("Ααρων", "Ααρων", false))
	}
	if Persistent("Ἀαρων", "Ἀαρών", false) != "Ἀαρών" {
		t.Fatalf("Persistent() failed. Returned %s", Persistent("Ἀαρων", "Ἀαρών", false))
	}
	// Missing breathing causes panic
	//if Persistent("ααρων", "ἀαρών", false) != "ἀαρών" {
	//	t.Fatalf("Persistent() failed. Returned %s", Persistent("ααρων", "Ἀαρών", false))
	//}

	if Persistent("Ἰαννης", "Ἰάννης", false) != "Ἰάννης" {
		t.Fatalf("Persistent() failed. Returned %s", Persistent("Ἰαννης", "Ἰάννης", false))
	}

	if Persistent("Ἰάκωβος", "Ἰάκωβος", false) != "Ἰάκωβος" {
		t.Fatalf("Persistent() failed. Returned: %s. Syllables: %s, Bytes: %v",
			Persistent("Ἰάκωβος", "Ἰάκωβος", false),
			DisplayWord(Syllabify("Ἰάκωβος")),
			decomposedBytes(Persistent("Ἰάκωβος", "Ἰάκωβος", false)))
	}
	if Persistent("Ἰαρεδ", "Ἰαρέδ", false) != "Ἰαρέδ" {
		t.Fatalf("Persistent() failed. Returned %s", Persistent("Ἰαρεδ", "Ἰαρέδ", false))
	}
	if Persistent("περιπατει", "περιπατῶ", false) != "περιπατεῖ" {
		t.Fatalf("Persistent() failed. Returned %s", Persistent("περιπατει", "περιπατέω", false))
	}
	if Persistent("περιπατεις", "περιπατῶ", false) != "περιπατεῖς" {
		t.Fatalf("Persistent() failed. Returned %s", Persistent("περιπατεις", "περιπατεῖς", false))
	}
	//if Persistent("περιπατει", "περιπατέω", false) != "περιπατεῖ" {
	//	t.Fatalf("Persistent() failed. Returned %s", Persistent("περιπατει", "περιπατέω", false))
	//}
	//if Persistent("Ιαρεδ", "Ἰαρέδ", false) != "Ἰαρέδ" {
	//	t.Fatalf("Persistent() failed. Returned %s", Persistent("Ιαρεδ", "Ἰαρέδ", false))
	//}

}

func TestFixBrokenUnicodeRunes(t *testing.T) {
	//fmt.Println("####", "Ἰάκωβος", decomposedBytes(string(fixBrokenUnicode("Ἰάκωβος"))))
	if !RuneArrayEqual(fixBrokenUnicodeRunes(
		[]rune{921, 787, 945, 769, 769, 954, 969, 946, 959, 962}),
		[]rune{921, 787, 945, 769, 954, 969, 946, 959, 962}) {
		t.Fatalf("fixBrokenUnicodeRunes() failed")
	}
}

func RuneArrayEqual(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for x, _ := range a {
		if a[x] != b[x] {
			fmt.Println("array item ", x, " match failed", a[x], "!=", b[x])
			return false
		}
	}
	return true
}

func decomposedBytes(s string) []rune {
	return []rune(norm.NFD.String(s))
}

func stringArrayMatch(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
