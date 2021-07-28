package fake

import (
	"strings"
)

func IBAN() string {
	return lookup(lang, "ibans", true)
}

func BICS() string {
	return lookup(lang, "bics", true)
}

func BankName() string {
	return lookup(lang, "bank_names", true)
}

type bankAccount struct {
	vendor   string
	length   int
}

var bankAccounts = map[string]bankAccount{
	"ad":       {"Andorra", 24},
	"ae": {"United Arab Emirates", 23},
	"al":       {"Albania", 28},
	"at":   {"Austria", 20},
}

func BankAccountType() string {
	n := len(bankAccounts)
	var vendors []string
	for _, cc := range bankAccounts {
		vendors = append(vendors, cc.vendor)
	}

	return vendors[r.Intn(n)]
}

// BankAccount generated bank account number according to the bank number rules
func BankAccountNum(vendor string) string {
	if vendor != "" {
		vendor = strings.ToLower(vendor)
	} else {
		var vendors []string
		for v := range bankAccounts {
			vendors = append(vendors, v)
		}
		vendor = vendors[r.Intn(len(vendors))]
	}
	bank := bankAccounts[vendor]
	num := []rune("")
	for i := 0; i < bank.length; i++ {
		num = append(num, genBankAccountDigit(num))
	}
	return string(num)
}

func genBankAccountDigit(num []rune) rune {
	sum := 0
	for i := len(num) - 1; i >= 0; i-- {
		n := int(num[i])
		if i%2 != 0 {
			sum += n
		} else {
			if n*2 > 9 {
				sum += n*2 - 9
			} else {
				sum += n * 2
			}
		}
	}
	return rune(((sum/10+1)*10 - sum) % 10)
}
