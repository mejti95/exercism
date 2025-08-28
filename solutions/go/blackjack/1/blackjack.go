package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
        case "ace":
        	return 11
        case "queen":
        	return 10
        case "king":
        	return 10
        case "ten":
			return 10
        case "jack":
        	return 10
        case "nine":
        	return 9
        case "eight":
        	return 8
        case "seven":
        	return 7
        case "six":
        	return 6
        case "five":
        	return 5
        case "four":
        	return 4
        case "three":
        	return 3
        case "two":
        	return 2
        default:
        	return 0
    }	
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    myCards := ParseCard(card1) + ParseCard(card2)
	if myCards == 22{
        return "P"
    }
    if myCards == 21{
        if ParseCard(dealerCard) == 10 || ParseCard(dealerCard) == 11{
        	return "S"
    	} 
        return "W"
    }
    if myCards >= 17 && myCards <= 20{
        return "S"
    }
    if myCards >= 12 && myCards <= 16{
        if ParseCard(dealerCard) >= 7 {
            return "H"
        }
        return "S"
    }
    if myCards <= 11{
        return "H"
    }
    return "dupa"
}
