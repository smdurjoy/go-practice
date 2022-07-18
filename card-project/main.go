package main

func main() {
	cards := newDeck()
	
	hand, remainingCards := deal(cards, 5)

	// can call print function since those variables are deck type
	hand.print()
	remainingCards.print()
}