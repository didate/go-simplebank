package util

const (
	USD = "USD"
	EUR = "EUR"
	GNF = "GNF"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, GNF:
		return true
	}
	return false
}
