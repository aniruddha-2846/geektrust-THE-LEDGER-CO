package subpackages

var ListedBanks = make(map[string]Bank)
var NumberOfMonths = 12

type Loan struct {
	principal          int
	RateOfInterest     int //xx%
	Tenure             int //in years
	TotalEMIs          int
	EMI                int
	AmoutToBeRecovered int
	LumpSum            map[int]int
}
type Bank struct {
	BankName     string
	customersMap map[string]Loan
}
