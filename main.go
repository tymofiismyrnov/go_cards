package main

func main() {
	cards := newDeck()
	cards.saveToFile("cards")
	cards = newDeckFromFile("cards")
	dealedCards, restDeck := deal(cards, 5)
	dealedCards.print()
	restDeck.shuffle()
	restDeck.print()
}
