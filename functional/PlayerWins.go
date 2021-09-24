package functional

// PlayerWins determines if the player has a winning blackjack hand.
func PlayerWins(playerHand string, dealerHand string) bool {

	playerScore := getHandSum(playerHand)
	dealerScore := getHandSum(dealerHand)

	return (playerScore > dealerScore && playerScore < 21 || dealerScore > 21)
}

func getHandSum(hand string) int {
	result := 0
	for _, value := range hand {
		if value > 47 && value < 58 {
			result += (int(value) - 48)
		} else if value == 65 {
			result += 11
		} else {
			result += 10
		}
	}
	return result
}
