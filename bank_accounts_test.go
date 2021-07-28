package fake

import (
	"testing"
)

func TestBankAccounts(t *testing.T) {
	for _, lang := range GetLangs() {
		SetLang(lang)

		v := BankAccountType()
		if v == "" {
			t.Errorf("BankAccountType failed with lang %s", lang)
		}

		v = BankAccountNum("")
		if v == "" {
			t.Errorf("BankAccountNum failed with lang %s", lang)
		}

		v = BankAccountNum("ad")
		if v == "" {
			t.Errorf("BankAccountNum failed with lang %s", lang)
		}
	}
}
