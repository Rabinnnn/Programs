package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	numCards := 12
	numPlayers := 4
	cardsPerPlayer := numCards / numPlayers

	// Calculate whether the deck size is divisible by 4
	mod := 1
	for mod = 1; mod <= numPlayers; mod++ {
		if (mod*numCards)%numPlayers == 0 {
			break
		}
	}

	if mod != 1 {
		fmt.Printf("Invalid deck size. The number of cards must be divisible by 4.\n")
		return
	}

	for player := 0; player < numPlayers; player++ {
		fmt.Printf("Player %d: ", player+1)

		for card := 0; card < cardsPerPlayer; card++ {
			if card > 0 {
				fmt.Printf(", ")
			}
			fmt.Printf("%d", deck[player*numCards/numPlayers+card])
		}

		fmt.Printf("\n")
	}
}
