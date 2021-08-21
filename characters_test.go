package greekaccentuation

import "testing"

func TestBase(t *testing.T) {
	if Base('ᾳ') != 'α' {
		t.Fatal("Base() failed")
	}
	if Base('Ι') != 'Ι' {
		t.Fatalf("Base() failed, returns %v", Base('Ι'))
	}
	if Base('ι') != 'ι' {
		t.Fatal("Base() failed")
	}
	if Base('α') != 'α' {
		t.Fatal("Base() failed")
	}
	if Base('ᾷ') != 'α' {
		t.Fatal("Base() failed")
	}
	if Base('ἄ') != 'α' {
		t.Fatal("Base() failed")
	}

}

func TestAddDiacritic(t *testing.T) {
	if string(AddDiacritic([]rune("α"), IOTA.Rune())) != "ᾳ" {
		t.Fatalf("AddDiacritic() failed. Returned: %v",
			string(AddDiacritic([]rune{'α'}, IOTA.Rune())))
	}
	if string(AddDiacritic(AddDiacritic([]rune{'ο'}, ROUGH.Rune()), ACUTE.Rune())) != "ὅ" {
		t.Fatalf("AddDiacritic() failed. Returned: %v",
			string(AddDiacritic(AddDiacritic([]rune{'ο'}, ROUGH.Rune()), ACUTE.Rune())))
	}
	if string(AddDiacritic([]rune("ου"), CIRCUMFLEX.Rune())) != "οῦ" {
		t.Fatalf("AddDiacritic() failed. Returned: %v",
			string(AddDiacritic([]rune("ου"), CIRCUMFLEX.Rune())))
	}
}

func TestAddBreathing(t *testing.T) {
	if AddBreathing('α', ROUGH) != 'ἁ' {
		t.Fatalf("AddBreathing() failed. Returned: %v",
			string(AddBreathing('α', ROUGH)))
	}
	if AddBreathing('α', SMOOTH) != 'ἀ' {
		t.Fatalf("AddBreathing() failed. Returned: %v",
			string(AddBreathing('α', SMOOTH)))
	}
}

func TestStripLength(t *testing.T) {
	if string(stripLength([]rune(Recessive("δεικνῡς", true, false)))) != "δείκνυς" {
		t.Fatalf("StripLength() failed. Returned: %v",
			string(stripLength([]rune(Recessive("δεικνῡς", true, false)))))
	}
}
func TestRecessive(t *testing.T) {
	if Recessive("εἰσηλθον", true, false) != "εἴσηλθον" {
		t.Fatalf("Recessive() failed. Returned: %v",
			Recessive("εἰσηλθον", true, false))
	}
}
