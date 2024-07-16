package subpackages

import (
	"math"
	"strconv"
	"strings"
)

func AddLoanDetails(bankObject *Bank, loanDetails string) {
	words := strings.Fields(loanDetails)
	customerName := words[2]
	principal, err1 := strconv.Atoi(words[3])
	tenure, err2 := strconv.Atoi(words[4])
	rateOfInterest, err3 := strconv.Atoi(words[5])

	if err1 != nil || err2 != nil || err3 != nil {
		return
	}

	//fetch the loan object of customerName
	customerLoanObject := bankObject.customersMap[customerName]

	customerLoanObject.principal = principal
	customerLoanObject.RateOfInterest = rateOfInterest
	customerLoanObject.Tenure = tenure
	customerLoanObject.TotalEMIs = (tenure * NumberOfMonths)
	customerLoanObject.AmoutToBeRecovered = principal + CalculateInterest(&customerLoanObject)
	customerLoanObject.EMI = int(math.Ceil((float64(customerLoanObject.AmoutToBeRecovered) / float64((customerLoanObject.TotalEMIs)))))

	bankObject.customersMap[customerName] = customerLoanObject
}

func AlterLoanDetails(bankObject *Bank, loanDetails string) {
	words := strings.Fields(loanDetails)
	customerName := words[2]
	lumpSum, err1 := strconv.Atoi(words[3])
	emiNumber, err2 := strconv.Atoi(words[4])
	if err1 != nil || err2 != nil {
		return
	}
	//fetch the loan object of customerName
	customerLoanObject := bankObject.customersMap[customerName]
	lastestLumpSumAmountInMap := getlastestLumpSumAmount(customerLoanObject.LumpSum)
	customerLoanObject.LumpSum[emiNumber] = lastestLumpSumAmountInMap + lumpSum
	bankObject.customersMap[customerName] = customerLoanObject
}
