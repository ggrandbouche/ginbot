package main

import (
	"fmt"
)

func main() {
	var board = GameBoard{}
	board.IntitializeBoard()

	board.Shuffle()
	board.DealHands()
	board.SortHands()


    
    test := []Card{	{value: 1, suit: 0},
					{value: 2, suit: 0},
					{value: 3, suit: 0},
        			{value: 4, suit: 0},
					{value: 3, suit: 1},
					{value: 3, suit: 2},
					{value: 7, suit: 3},
        			{value: 10, suit: 3}, 
					{value: 11, suit: 3}, 
					{value: 12, suit: 3}}
	// test := board.hand1
    printHand(test)
    fmt.Println(groupify(test))
}

