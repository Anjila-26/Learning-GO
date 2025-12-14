package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	// Your existing ParseCard function is correct
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
	// Parse the card values
	playerCard1 := ParseCard(card1)
	playerCard2 := ParseCard(card2)
	dealerValue := ParseCard(dealerCard)
    
	playerSum := playerCard1 + playerCard2

	// --- Alex's Strategy Logic ---

	// 1. If you have a pair of aces you must always split them.
	// (Since ParseCard returns 11 for "ace")
	if playerCard1 == 11 && playerCard2 == 11 {
		return "P" // Split
	}

	// 2. If you have a Blackjack (sum up to 21)
	if playerSum == 21 {
		// Dealer has an ace (11), face card, or ten (10)
		if dealerValue == 11 || dealerValue == 10 {
			return "S" // Stand
		}
		// Dealer does not have an ace, face card, or ten
		return "W" // Automatically Win
	}

	// 3. If your cards sum up to a value within the range [17, 20] you should always stand.
	if playerSum >= 17 && playerSum <= 20 {
		return "S" // Stand
	}

	// 4. If your cards sum up to a value within the range [12, 16]
	if playerSum >= 12 && playerSum <= 16 {
		// Always stand unless the dealer has a 7 or higher, in which case you should always hit.
		if dealerValue >= 7 {
			return "H" // Hit
		}
		return "S" // Stand
	}

	// 5. If your cards sum up to 11 or lower you should always hit.
	if playerSum <= 11 {
		return "H" // Hit
	}

	return "S" // Default fallback (e.g., if playerSum is > 21, though not possible with 2 cards)
}