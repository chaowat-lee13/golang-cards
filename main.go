package main

func main() {
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 14)

	// hand.print()
	// remainingCards.print()

	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
}
