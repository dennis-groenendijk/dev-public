package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	value := 0

	switch {
	case card == "two":
		value = 2
	case card == "three":
		value = 3
	case card == "four":
		value = 4
	case card == "five":
		value = 5
	case card == "six":
		value = 6
	case card == "seven":
		value = 7
	case card == "eight":
		value = 8
	case card == "nine":
		value = 9
	case card == "ten", card == "jack", card == "queen", card == "king":
		value = 10
	case card == "ace":
		value = 11
	default:
		value = 0
	}
	return value
	panic("Please implement the ParseCard function")
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    playerScore := ParseCard(card1) + ParseCard(card2)
	dealerScore := ParseCard(dealerCard)

    switch {
        case card1 == "ace" && card2 == "ace":
    		return "P"
        case playerScore == 21:
    		if dealerScore >= 10 {
                return "S"
            } else {
            	return "W"
            }
        case playerScore >= 17 && playerScore <= 20:
    		return "S"
        case playerScore >= 12 && playerScore <= 16:
    		if dealerScore >= 7 {
                return "H"
            } else {
            	return "S"
            }
        case playerScore <= 11:
    		return "H"
    }
	panic("Please implement the FirstTurn function")
}
