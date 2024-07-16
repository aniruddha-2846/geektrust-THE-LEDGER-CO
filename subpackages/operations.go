package subpackages

import (
	"strings"
)

func ProcessLoan(line string) {
	words := strings.Fields(line)
	bankName := words[1]
	customerName := words[2]
	bankObject, found := FetchBankObject(bankName)
	if !found {
		//create a new object of the bankName
		bankObject := new(Bank)
		loanObject := new(Loan)
		bankObject.customersMap = make(map[string]Loan)
		loanObject.LumpSum = make(map[int]int)
		loanObject.LumpSum[0] = 0
		bankObject.customersMap[customerName] = *loanObject
		ListedBanks[bankName] = *bankObject
		AddLoanDetails(bankObject, line)
		return
	}
	AddLoanDetails(&bankObject, line)
}

func ProcessPayment(line string) {
	words := strings.Fields(line)
	bankName := words[1]
	bankObject, found := FetchBankObject(bankName)
	if !found {
		return
	}
	AlterLoanDetails(&bankObject, line)
}

func ProcessBalance(line string) {
	words := strings.Fields(line)
	bankName := words[1]
	bankObject, found := FetchBankObject(bankName)
	if !found {
		return
	}
	GenerateOutput(bankName, line, &bankObject)
}
