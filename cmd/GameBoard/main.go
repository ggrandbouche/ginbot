package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

/*
Game board structures
*/
type Card struct {
	value, suit int
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

    board.Shuffle()
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

func (board *GameBoard) DealCard() Card {
	var card Card = board.deck[len(board.deck)-1]
	board.deck = board.deck[:len(board.deck)-1]
	return card
}

