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

func TestSyllableAccent(t *testing.T) {
	if syllableAccent("") != NO_ACCENT {
		t.Fatal("syllableAccent() failed")
	}
	if syllableAccent("κός") != ACUTE {
		t.Fatalf("syllableAccent() failed, returned %v", syllableAccent("κός"))
	}
	if syllableAccent("ναι") != NO_ACCENT {
		t.Fatal("syllableAccent() failed")
	}
	if syllableAccent("φῶς") != CIRCUMFLEX {
		t.Fatal("syllableAccent() failed")
	}
}

func TestIotaSubscript(t *testing.T) {
	if iotaSubscript('ᾳ') == nil {
		t.Fatalf("iotaSubscript() failed, should not be nil")
	}
	if iotaSubscript('ᾳ') != IOTA {
		t.Fatalf("iotaSubscript() failed, returned %c", iotaSubscript('ᾳ').Rune())
	}
	if iotaSubscript(' ') != nil {
		t.Fatalf("iotaSubscript() failed, returned %c", iotaSubscript(' ').Rune())
	}
	if iotaSubscript('α') != nil {
		t.Fatalf("iotaSubscript() failed")
	}
}

func TestSyllableLength(t *testing.T) {
	if syllableLength("") != UNKNOWN {
		t.Fatal("syllableLength() failed")
	}
	if syllableLength("κός") != SHORT {
		t.Fatalf("syllableLength() failed, returned %v", syllableLength("κός").Name())
	}
	if syllableLength("σω") != LONG {
		t.Fatal("syllableLength() failed")
	}
	if syllableLength("τοῦ") != LONG {
		t.Fatal("syllableLength() failed")
	}
	if syllableLength("ᾳ") != LONG {
		t.Fatal("syllableLength() failed")
	}
	if syllableLength("ναι") != UNKNOWN {
		t.Fatal("syllableLength() failed")
	}
	if syllableLength("οἰ", false) != LONG {
		t.Fatal("syllableLength() failed")
	}
	if syllableLength("ναι", true) != SHORT {
		t.Fatal("syllableLength() failed")
	}
	if syllableLength("ναι", false) != LONG {
		t.Fatal("syllableLength() failed")
	}
	if syllableLength("οἰ", false) != LONG {
		t.Fatal("syllableLength() failed")
	}
}
func TestOnsetNucleusCoda(t *testing.T) {

	{
		o := onset("κός")
		if o != "κ" {
			t.Fatal("Onset() failed")
		}
	}

	{
		n := nucleus("κός")
		if n != "ό" {
			t.Fatalf("Nucleus() failed, returned %s", n)
		}
	}

	{
		c := coda("κός")
		if c != "ς" {
			t.Fatal("Coda() failed")
		}
	}

	{
		o, n, c := onsetNucleusCoda("κός")
		if o != "κ" {
			t.Fatal("OnsetNucleusCoda() failed")
		}
		if n != "ό" {
			t.Fatalf("OnsetNucleusCoda() failed, returned %s - %s - %s", o, n, c)
		}
		if c != "ς" {
			t.Fatal("OnsetNucleusCoda() failed")
		}
	}

	{
		o, n, c := onsetNucleusCoda("ναι")
		if o != "ν" {
			t.Fatal("OnsetNucleusCoda() failed")
		}
		if n != "αι" {
			t.Fatal("OnsetNucleusCoda() failed")
		}
		if c != "" {
			t.Fatal("OnsetNucleusCoda() failed")
		}
	}

	{
		o, n, c := onsetNucleusCoda("βββ")
		if o != "βββ" {
			t.Fatal("OnsetNucleusCoda() failed")
		}
		if n != "" {
			t.Fatal("OnsetNucleusCoda() failed")
		}
		if c != "" {
			t.Fatal("OnsetNucleusCoda() failed")
		}
	}

}

func TestBody(t *testing.T) {

	{
		o := body("κός")
		if o != "κό" {
			t.Fatal("body() failed")
		}
	}

	{
		o := body("ό")
		if o != "ό" {
			t.Fatal("body() failed")
		}
	}

	{
		o := body("ὅ")
		if o != "ὅ" {
			t.Fatal("body() failed")
		}
	}

	{
		o := body("οἴδ")
		if o != "οἴδ" {
			t.Fatal("body() failed")
		}
	}

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
