package main

import "fmt"
import "strings"

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
	deck    []Card
	hand1   []Card
	hand2   []Card
	discard []Card
}

func (board *GameBoard) IntitializeBoard() {
	//create every card, from val 0-12 and suit 0-3
	for i := 0; i < 52; i++ {
		board.deck = append(board.deck, Card{value: i % 13, suit: i / 13})
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
    board.DealHands()
    PrintHand(board.hand1)
}

func (board *GameBoard) DealHands() {
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			board.hand1 = append(board.hand1, board.DealCard())
		} else {
			board.hand2 = append(board.hand2, board.DealCard())
		}
	}
}

func PrintHand(hand []Card) {
	value := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	suit := []string{"♥", "♦", "♣", "♠"}
    var builder strings.Builder
	for _, card := range hand {
		builder.WriteString(fmt.Sprintf("%s%s, ", value[card.value], suit[card.suit]))
	}
    str := builder.String()
    str = str[:len(str)-2]
    fmt.Println(str)
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

func (board *GameBoard) DealCard() Card {
	var card Card = board.deck[len(board.deck)-1]
	board.deck = board.deck[:len(board.deck)-1]
	return card
}

