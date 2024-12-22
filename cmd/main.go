package main

import (
	"fmt"
	"math/rand"
	// "math"
)

type Card struct {
	value, suit int
}

// Define the interface
type DeckOpps interface {
	BuildDeck();
	Shuffle();
}

type Deck struct {
	my_deck [52]Card
}

//function to initialize new deck of cards
func (d *Deck) BuildDeck() {
	//create every card, from val 0-12 and suit 0-3
	for i := 0; i < len(d.my_deck); i++ {
		d.my_deck[i] = Card{value: i%13, suit: i/13}
	}
}

//function to shuffle a deck of cards into a random order
func (d *Deck) Shuffle() {
	// go's math/rand std lib comes with a shuffle funciton.
	// it expects the number of items and a the function to perform a swap, 
	// which is super easy in go thanks this wierd syntax
	rand.Shuffle(
		len(d.my_deck), // num items
		func(i, j int) { //swap function
			d.my_deck[i], d.my_deck[j] = d.my_deck[j], d.my_deck[i]
		})
}

func main() {
	//init deck of cards
	var newDeck = Deck{}
	newDeck.BuildDeck()
	fmt.Println("newDeck:", newDeck.my_deck)

	//shuffle deck of cards
	newDeck.Shuffle()
	fmt.Println("\nnewDeck: ", newDeck)

}
