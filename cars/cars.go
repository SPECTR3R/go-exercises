package cars

// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
func SuccessRate(speed int) float64 {
	switch {
	case speed == 0:
		return 0.0
	case speed >= 1 && speed <= 4:
		return 1
	case speed >= 5 && speed <= 8:
		return 0.9
	case speed >= 9 && speed <= 10:
		return 0.77

	}
	return 0.0
}

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	baseRate := speed * 221
	return SuccessRate(speed) * float64(baseRate)
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	ratePerHour := CalculateProductionRatePerHour(speed)
	return int(ratePerHour / 60)
}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
	if ratePerHour := CalculateProductionRatePerHour(speed); ratePerHour > limit {
		return limit
	} else {
		return ratePerHour
	}
}

func ParseCard(card string) int {
	switch card {
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "queen", "king", "jack":
		return 10
	case "ace":
		return 11
	default:
		return 0
	}
}

func LargeHand(isBlackjack bool, dealerScore int) string {
	switch {
	case isBlackjack && dealerScore < 10:
		return "W"
	case isBlackjack && dealerScore >= 10:
		return "S"
	}
	return "P"
}

func SmallHand(handScore, dealerScore int) string {
	isInRange := handScore >= 12 && handScore <= 16
	switch {
	case handScore >= 17, isInRange && dealerScore <= 6:
		return "S"
	case isInRange && dealerScore >= 7:
		return "H"
	default:
		return "H"
	}
}
