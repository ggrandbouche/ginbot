package main

import (
	"fmt"
	"math/rand/v2"
)

/*
	Game board structures
*/
type Card struct {
	value, suit int
}

type GameBoard struct {
	deck []Card
	hand1 []Card
	hand2 []Card
	discard []Card
}

func (board *GameBoard) IntitializeBoard() {
	//create every card, from val 0-12 and suit 0-3
	for i := 0; i < 52; i++ {
		board.deck = append(board.deck, Card{value: i%13, suit: i/13})
	}

	//initialize hand1, hand2, discard
	board.hand1 = make([]Card, 0, 11)
	board.hand2 = make([]Card, 0, 11)
	board.discard = make([]Card, 0, 52)

}

//function to shuffle a deck of cards into a random order
func (board *GameBoard) Shuffle() {
	// go's math/rand std lib comes with a shuffle funciton.
	// it expects the number of items and a the function to perform a swap, 
	// which is super easy in go thanks this wierd syntax
	rand.Shuffle(
		len(board.deck), // num items
		func(i, j int) { //swap function
			board.deck[i], board.deck[j] = board.deck[j], board.deck[i]
		})
}


func main() {
	var board = GameBoard{}
	board.IntitializeBoard()
	fmt.Println(board)
	board.Shuffle()
	fmt.Println(board)
}

