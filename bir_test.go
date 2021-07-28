package fake

import (
	"testing"
)

func TestBIR(t *testing.T) {
	for _, lang := range GetLangs() {
		SetLang(lang)

		v := BIRRDO()
		if v == "" {
			t.Errorf("BIRRDO failed with lang %s", lang)
		}

		v = BIRNumber()
		if v == "" {
			t.Errorf("BIRNumber failed with lang %s", lang)
		}
	}
}
