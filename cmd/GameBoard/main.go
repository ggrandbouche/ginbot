package main

import "fmt"

/*
	Game board structures
*/
type Card struct {
	value, suit int
}

type GameBoardOpps interface {
	InitializeBoard()
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


func main() {
	//Test Deck
	// //init deck of cards
	// var newDeck = Deck{}
	// newDeck.BuildDeck()

	// //shuffle deck of cards
	// newDeck.Shuffle()

	// fmt.Println("\nnewDeck: ", newDeck)

	// c := newDeck.DealCard()
	// fmt.Print("\ntop card: ", c)
	// fmt.Print("\nnewDeck after dealcard; ", newDeck, )
	// fmt.Println("\nnewDeck.topCardIndex; ", newDeck.topCardIndex)

	// //reset deck
	// newDeck.BuildDeck()
	// newDeck.Shuffle()

	var board = GameBoard{}
	board.IntitializeBoard()
	fmt.Println(board)

	
}



// //function to shuffle a deck of cards into a random order
// func (d *Deck) Shuffle() {
// 	// go's math/rand std lib comes with a shuffle funciton.
// 	// it expects the number of items and a the function to perform a swap, 
// 	// which is super easy in go thanks this wierd syntax
// 	rand.Shuffle(
// 		len(d.my_deck), // num items
// 		func(i, j int) { //swap function
// 			d.my_deck[i], d.my_deck[j] = d.my_deck[j], d.my_deck[i]
// 		})
// }

// func (d *Deck) DealCard() Card {
// 	var topCard Card = d.my_deck[d.topCardIndex]
// 	//take top card off of deck
// 	d.my_deck[d.topCardIndex].suit = -1
// 	d.my_deck[d.topCardIndex].value = -1

// 	//update topCardIndex
// 	d.topCardIndex--

// 	return topCard
// }