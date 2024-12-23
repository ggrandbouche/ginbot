package main

import (
	"fmt"
	"math/rand"
	// "math"
)

type Card struct {
	value, suit int
}

// Define the Deck Operations interface
type DeckOpps interface {
	BuildDeck();
	Shuffle();
	DealCard() Card;
}

//define the Deck struct
type Deck struct {
	my_deck [52]Card //all 52 cards
	topCardIndex int //decrement to keep track of the top card
}

//function to initialize new deck of cards
func (d *Deck) BuildDeck() {
	//create every card, from val 0-12 and suit 0-3
	for i := 0; i < len(d.my_deck); i++ {
		d.my_deck[i] = Card{value: i%13, suit: i/13}
	}

	//initialize topCardIndex
	d.topCardIndex = len(d.my_deck)-1
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

func (d *Deck) DealCard() Card {
	var topCard Card = d.my_deck[d.topCardIndex]
	//take top card off of deck
	d.my_deck[d.topCardIndex].suit = -1
	d.my_deck[d.topCardIndex].value = -1

	//update topCardIndex
	d.topCardIndex--

	return topCard
}

func main() {
	//init deck of cards
	var newDeck = Deck{}
	newDeck.BuildDeck()

	//shuffle deck of cards
	newDeck.Shuffle()

	fmt.Println("\nnewDeck: ", newDeck)

	c := newDeck.DealCard()
	fmt.Print("\ntop card: ", c)
	fmt.Print("\nnewDeck after dealcard; ", newDeck)
	fmt.Print("\nnewDeck.topCardIndex; ", newDeck.topCardIndex)
	fmt.Print("\n")
}
