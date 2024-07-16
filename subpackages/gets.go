package subpackages

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func CalculateInterest(customerLoanObject *Loan) int {
	return (customerLoanObject.principal * customerLoanObject.RateOfInterest * customerLoanObject.Tenure) / 100
}

func getlastestLumpSumAmount(lumpSumMap map[int]int) int {
	count := 0
	for keys := range lumpSumMap {
		count++
		if count == len(lumpSumMap) {
			return lumpSumMap[keys]
		}
	}
	return 0
}
func lowerBound(emiNumber int, lumSumpEMInumbers []int) int {
	lo, hi := 0, len(lumSumpEMInumbers)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if lumSumpEMInumbers[mid] == emiNumber {
			return lumSumpEMInumbers[mid]
		} else if lumSumpEMInumbers[mid] > emiNumber {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return lumSumpEMInumbers[hi]
}
func CalculateRecoveredAmount(emiNumber int, customerLoanObject *Loan) int {
	var keys []int
	for k := range customerLoanObject.LumpSum {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	earliestEMInumber := lowerBound(emiNumber, keys)
	return (emiNumber * customerLoanObject.EMI) + (customerLoanObject.LumpSum[earliestEMInumber])
}

func CalculateNoOfEMIsLeft(customerLoanObject *Loan, recoveredAmount int) int {
	return int(math.Ceil(float64((customerLoanObject.AmoutToBeRecovered - recoveredAmount)) / float64(customerLoanObject.EMI)))
}

func FetchBankObject(bankName string) (Bank, bool) {
	bankObject, found := ListedBanks[bankName]
	return bankObject, found
}

func GenerateOutput(bankName string, details string, bankObject *Bank) {
	words := strings.Fields(details)
	customerName := words[2]
	emiNumber, err := strconv.Atoi(words[3])
	if err != nil {
		return
	}
	customerLoanObject := bankObject.customersMap[customerName]
	//! this function is the root cause of the errors
	recoveredAmount := CalculateRecoveredAmount(emiNumber, &customerLoanObject)
	NoOfEMIsLeft := CalculateNoOfEMIsLeft(&customerLoanObject, recoveredAmount)
	fmt.Printf("%s %s %d %d\n", bankName, customerName, recoveredAmount, NoOfEMIsLeft)
}
