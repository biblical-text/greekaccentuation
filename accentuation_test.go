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
}

func TestAddAccentuation(t *testing.T) {
	if addAccentuation(Syllabify("θεος"), OXYTONE) != "θεός" {
		t.Fatalf("AddAccentuation() failed. %v != %v", "θεός",
			addAccentuation(Syllabify("θεος"), OXYTONE))
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
