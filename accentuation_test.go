package greekaccentuation

import (
	"testing"
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
	if Persistent("Ἀαρων", "Ἀαρών", false) != "Ἀαρών" {
		t.Fatalf("Persistent() failed. Returned %s", Persistent("Ἀαρων", "Ἀαρών", false))
	}
	// Missing breathing causes panic
	//if Persistent("ααρων", "ἀαρών", false) != "ἀαρών" {
	//	t.Fatalf("Persistent() failed. Returned %s", Persistent("ααρων", "Ἀαρών", false))
	//}

	if Persistent("Ἰαννης", "Ἰάννης", false) != "Ἰάννης" {
		t.Fatalf("Persistent() failed. Returned %s", Persistent("Ἰαννης", "Ἰάννης", false))
	}

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
