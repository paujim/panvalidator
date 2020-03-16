package panvalidator

// CardType ...
type CardType string

// Values for CardType
const (
	Amex       CardType = "AMEX"
	Discover   CardType = "Discover"
	MasterCard CardType = "MasterCard"
	Visa       CardType = "VISA"
	Unknown    CardType = "Unknown"
)

func isAmex(number int) bool {
	number = number / 1e13
	return 34 == number || 37 == number
}

func isDiscover(number int) bool {
	return 6011 == number/1e12
}

func isMasterCard(number int) bool {
	number = number / 1e14
	return number > 50 && number < 56
}

func isVisa(number int) bool {
	return 4 == number/1e12 || 4 == number/1e15
}

func getCardType(number int) CardType {
	cardType := Unknown

	if isAmex(number) {
		cardType = Amex
	} else if isDiscover(number) {
		cardType = Discover
	} else if isMasterCard(number) {
		cardType = MasterCard
	} else if isVisa(number) {
		cardType = Visa
	}
	return cardType
}

// IsValid ...
func IsValid(number int) (bool, CardType) {
	var sum int
	cardType := getCardType(number)

	if cardType == Unknown {
		return false, cardType
	}

	for i := 0; number > 0; i++ {
		lastDigit := number % 10

		if i%2 != 0 {
			lastDigit = lastDigit * 2
			if lastDigit > 9 {
				lastDigit = lastDigit%10 + 1
			}
		}
		sum += lastDigit
		number = number / 10
	}
	return sum%10 == 0, cardType
}
