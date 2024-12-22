package main

import (
	"fmt"
	// "math"
)

type Card struct {
	value, suit int
}

// Define the interface
type DeckOpps interface {
	BuildDeck();
}

type Deck struct {
	my_deck [52]Card
}

func (d *Deck) BuildDeck() {
	for i := 0; i < len(d.my_deck); i++ {
		d.my_deck[i] = Card{value: i%13, suit: i/13}
	}
}

func main() {
	var testingDeck DeckOpps 
	deck := Deck{}
	testingDeck = &deck
	testingDeck.BuildDeck()
	fmt.Println(deck.my_deck)

}
