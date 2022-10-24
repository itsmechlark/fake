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

type bankAccount struct {
	vendor     string
	min_length int
	max_length int
}

var bankAccountTypes = []string{
	"Checking",
	"Savings",
	"Money Market",
	"Certificate of Deposit",
}

var bankAccounts = map[string]bankAccount{
	"aub":                 {"Asia United Bank", 12, 12},
	"bpi":                 {"BPI", 10, 12},
	"citibank":            {"Citibank, N.A.", 10, 10},
	"ewbanker":            {"East West Bank", 11, 12},
	"maybank":             {"Maybank Phils. Inc.", 11, 12},
	"ucpb":                {"United Coconut Planters Bank", 12, 12},
	"rcbcsavings":         {"RCBC Savings Bank", 10, 10},
	"businessbank":        {"Philippine Business Bank", 12, 12},
	"metrobank":           {"Metrobank", 12, 13},
	"chinabank":           {"China Bank", 10, 12},
	"alliedbank":          {"Allied Bank", 10, 16},
	"bpifamily":           {"BPI Family Savings Bank", 10, 12},
	"pbcom":               {"PBCOM", 12, 12},
	"businessbanksavings": {"Philippine Business Bank, Inc., A Savings Bank", 12, 12},
	"robinsonsbank":       {"Robinsons Bank", 12, 15},
	"sterlingbank":        {"Sterling Bank", 12, 12},
	"dbp":                 {"Development Bank of the Philippines", 10, 10},
	"landbank":            {"Land Bank", 10, 12},
	"rcbc":                {"RCBC", 10, 12},
	"unionbank":           {"Union Bank of the Philippines", 10, 12},
	"bankcom":             {"Bank of Commerce", 11, 12},
}

func BankName() string {
	n := len(bankAccounts)
	var vendors []string
	for _, cc := range bankAccounts {
		vendors = append(vendors, cc.vendor)
	}

	return vendors[r.Intn(n)]
}

func BankAccountType() string {
	return bankAccountTypes[r.Intn(len(bankAccountTypes))]
}

func BankAccountNumDefault() string {
	return BankAccountNum("")
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
	for i := 0; i < bank.max_length; i++ {
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
