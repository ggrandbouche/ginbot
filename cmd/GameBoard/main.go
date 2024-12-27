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

// function to shuffle a deck of cards into a random order
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

func (board *GameBoard) DealHands() {
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			board.hand1 = append(board.hand1, board.DealCard())
		} else {
			board.hand2 = append(board.hand2, board.DealCard())
		}
	}
}

func printHand(hand []Card) {
	value := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	suit := []string{"♠", "♦", "♣", "♥"}
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

func (board *GameBoard) SortHands() {
	board.hand1 = sortHand(board.hand1)
	board.hand2 = sortHand(board.hand2)
}

func sortHand(hand []Card) []Card {
	for i := 0; i < len(hand)-1; i++ {
		min := i
		for j := i + 1; j < len(hand); j++ {
			if hand[j].value+hand[j].suit*13 < hand[min].value+hand[min].suit*13 {
				min = j
			}
		}
		//perform swap
		temp := hand[i]
		hand[i] = hand[min]
		hand[min] = temp
	}

	return hand
}

func groupify(hand []Card) int {
	//get the straights
	//get the pairs
	//compare point values
    return 0
}

func findStraights(hand []Card) [][]Card{
	var straight [][]Card
    for i := 0; i < len(hand) - 2; i++ {
        if(hand[i].suit == hand[i+2].suit && hand[i].value - hand[i+2].value == -2){
                group := []Card{hand[i], hand[i+1], hand[i+2]} 
                for j := i+3; j < len(hand); j++ {
                    if(hand[j].suit == hand[j-1].suit && hand[j].value - hand[j-1].value == 1){
                        group = append(group, hand[j])
                        i = j - 1
                    }else{
                        i = j - 1
                        break
                    }
                }
                straight = append(straight, group)
        }
    }
    return straight
}

func findMatches(hand []Card) [][]Card {
	//slice to return - has all card slices with 3-4 cards
	var matches [][]Card
	//map for all card value matches of 1-4 cards
	var allMatches = make(map[int][]Card)
	//loop through hand
	for i := 0; i < len(hand); i++ {
		//get the value and existence boolean 
		//from the map for the current card value
		val, keyExists := allMatches[hand[i].value]
		if (keyExists) {
			//add this card to the appropriate group
			allMatches[hand[i].value] = append(val, hand[i])
		} else {
			//make new group in map
			allMatches[hand[i].value] = []Card{hand[i]}
		}
	}
	//loop through groups to return the groups of 3+ cards
	for _, match := range allMatches {
		if len(match) >= 3 {
			matches = append(matches, match)
		}
	}
	return matches
}

func main() {
	var board = GameBoard{}
	board.IntitializeBoard()

	board.Shuffle()
	board.DealHands()
	board.SortHands()


    
    // test := []Card{{value: 1, suit: 0},{value: 2, suit: 0},{value: 2, suit: 1},
    //     {value: 4, suit: 0},{value: 7, suit: 0},{value: 2, suit: 0},{value: 9, suit: 0},
    //     {value: 2, suit: 0}}
	test := board.hand1
    printHand(test)
    var straight [][]Card = findMatches(test)
    fmt.Println(straight)
}

