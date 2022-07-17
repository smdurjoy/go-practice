package main

import "fmt"

func main() {
	cards := deck{"One card", "Two card", newCard()}
	cards = append(cards, "Hello card")
	
	
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "New Card"
}