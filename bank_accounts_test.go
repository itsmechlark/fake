package fake

import (
	"testing"
)

func TestBankAccounts(t *testing.T) {
	for _, lang := range GetLangs() {
		SetLang(lang)

		v := BankName()
		if v == "" {
			t.Errorf("BankName failed with lang %s", lang)
		}

		v = BankAccountType()
		if v == "" {
			t.Errorf("BankAccountType failed with lang %s", lang)
		}

		v = BankAccountNumDefault()
		if v == "" {
			t.Errorf("BankAccountNumDefault failed with lang %s", lang)
		}

		v = BankAccountNum("aub")
		if v == "" {
			t.Errorf("BankAccountNum failed with lang %s", lang)
		}
	}
}
