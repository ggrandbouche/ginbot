package main

import "fmt"

/*
	Game board structures
*/
type GameBoardOpps interface {

}

type GameBoard struct {
	deck []Card
	hand1 []Card
	hand2 []Card
	discard []Card
}



func main() {
	//Test Deck
	//init deck of cards
	var newDeck = Deck{}
	newDeck.BuildDeck()

	//shuffle deck of cards
	newDeck.Shuffle()

	fmt.Println("\nnewDeck: ", newDeck)

	c := newDeck.DealCard()
	fmt.Print("\ntop card: ", c)
	fmt.Print("\nnewDeck after dealcard; ", newDeck, )
	fmt.Println("\nnewDeck.topCardIndex; ", newDeck.topCardIndex)

	//reset deck
	newDeck.BuildDeck()
	newDeck.Shuffle()

	
}
