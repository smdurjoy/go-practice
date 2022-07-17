package main

func main() {
	cards := deck{"One card", "Two card", newCard()}
	cards = append(cards, "Hello card")
	
	cards.print();
}

func newCard() string {
	return "New Card"
}