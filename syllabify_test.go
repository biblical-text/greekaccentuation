package greekaccentuation

import "testing"

func TestIsVowel(t *testing.T) {
	if !IsVowel('ι') {
		t.Fatal("IsVowel() failed")
	}
	if !IsVowel('α') {
		t.Fatal("IsVowel() failed")
	}
	if !IsVowel('ἀ') {
		t.Fatal("IsVowel() failed")
	}
	if !IsVowel('ἁ') {
		t.Fatal("IsVowel() failed")
	}
	if !IsVowel('υ') {
		t.Fatal("IsVowel() failed")
	}
	if IsVowel('χ') {
		t.Fatal("IsVowel() failed")
	}
	if !IsVowel('υ') {
		t.Fatal("IsVowel() failed")
	}
	if IsVowel('σ') {
		t.Fatal("IsVowel() failed")
	}
	if !IsVowel('ὄ') {
		t.Fatal("IsVowel() failed")
	}
	if !IsVowel('ᾀ') {
		t.Fatal("IsVowel() failed")
	}
}

func TestIsDipthong(t *testing.T) {
	if !isDipthong('α', 'ι') {
		t.Fatal("IsDipthong() failed")
	}
	if !isDipthong('η', 'υ') {
		t.Fatal("IsDipthong() failed")
	}
	if isDipthong('χ', 'χ') {
		t.Fatal("IsDipthong() failed")
	}
	if isDipthong('σ', 'υ') {
		t.Fatal("IsDipthong() failed")
	}
}

func TestIsValidConsonantCluster(t *testing.T) {
	if !isValidConsonantCluster('π', []rune("ν")) {
		t.Fatal("isValidConsonantCluster() failed")
	}
	if !isValidConsonantCluster('π', []rune("νε")) {
		t.Fatal("isValidConsonantCluster() failed")
	}
	if isValidConsonantCluster('σ', []rune("ν")) {
		t.Fatal("isValidConsonantCluster() failed")
	}
	if !isValidConsonantCluster('σ', []rune("τρ")) {
		t.Fatal("isValidConsonantCluster() failed")
	}
	if !isValidConsonantCluster('σ', []rune("τρα")) {
		t.Fatal("isValidConsonantCluster() failed")
	}
}

func TestDisplayWord(t *testing.T) {
	if DisplayWord([]string{"α", "β", "σ"}) != "α.β.σ" {
		t.Fatal("DisplayWord() failed")
	}
	if DisplayWord([]string{"α", "β"}) != "α.β" {
		t.Fatal("DisplayWord() failed")
	}
	if DisplayWord([]string{}) != "" {
		t.Fatal("DisplayWord() failed")
	}
	if DisplayWord(Syllabify("γυναικός")) != "γυ.ναι.κός" {
		t.Fatalf("Syllabify() failed: %v", DisplayWord(Syllabify("γυναικός")))
	}
	if DisplayWord(Syllabify("καταλλάσσω")) != "κα.ταλ.λάσ.σω" {
		t.Fatal("Syllabify() failed")
	}
	if DisplayWord(Syllabify("γγγ")) != "γγγ" {
		t.Fatal("Syllabify() failed")
	}
}

func TestSyllabify(t *testing.T) {
	if !ArrayEqual(Syllabify("γυναικός"), []string{"γυ", "ναι", "κός"}) {
		t.Fatalf("Syllabify() failed: %v", Syllabify("γυναικός"))
	}
	if !ArrayEqual(Syllabify("καταλλάσσω"), []string{"κα", "ταλ", "λάσ", "σω"}) {
		t.Fatal("Syllabify() failed")
	}
	if !ArrayEqual(Syllabify("γγγ"), []string{"γγγ"}) {
		t.Fatal("Syllabify() failed")
	}
	// TODO: I am not yet sure of the form of ῡ́ and why it is relevant.
	//if !ArrayEqual(Syllabify("φῡ́ω"), []string{"φῡ́", "ω"}) {
	//	t.Fatalf("Syllabify() failed: %v", Syllabify("φῡ́ω"))
	//}
}

func ArrayEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for x, _ := range a {
		if a[x] != b[x] {
			return false
		}
	}
	return true
}
