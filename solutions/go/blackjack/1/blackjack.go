package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
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
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	playerTotal := ParseCard(card1) + ParseCard(card2)
	dealerValue := ParseCard(dealerCard)

	switch {
	// Split aces
	case playerTotal == 22:
		return "P"

	// Blackjack rules
	case playerTotal == 21 && dealerValue < 10:
		return "W"

	case playerTotal == 21:
		return "S"

	// Stand on strong hands
	case playerTotal >= 17 && playerTotal <= 20:
		return "S"

	// Stand on 12–16 if dealer is weak
	case playerTotal >= 12 && playerTotal <= 16 && dealerValue >= 2 && dealerValue <= 6:
		return "S"

	// Otherwise hit
	default:
		return "H"
	}
}


