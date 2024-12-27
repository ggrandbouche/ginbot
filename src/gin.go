package main

import (
	"fmt"
)

func gameLoop(board *GameBoard) (bool, int) {
	var gameOver = false

	for !gameOver {
		var input int

		board.SortHands()

		fmt.Println("Player 1's hand: ")
		printHand(board.hand1)
		fmt.Println("Draw new card(1) or take top of discard pile(2)")
		fmt.Print("This is the top of the discard pile: ")
		printCard(board.discard[len(board.discard)-1])
		fmt.Print("\n> ")
		fmt.Scan(&input)
		if input == 1 {
			var deltCard = board.DealCard()
			board.hand1 = append(board.hand1, deltCard)
			fmt.Print("You drew: ")
			printCard(deltCard)
		} else if input == 2 {
			board.hand1 = append(board.hand1, board.discard[len(board.discard)-1])
			board.discard = board.discard[:len(board.discard)-1]
		}
		fmt.Println("Please enter the index of the card you would like to discard, starting at index 0")
		sortHand(board.hand1)
		printHand(board.hand1)
		fmt.Println("> ")
		fmt.Scan(&input)
		board.discard = append(board.discard, board.hand1[input])
		board.hand1 = append(board.hand1[:input], board.hand1[input+1:]...)
		if groupify(board.hand1) == 0 {
			fmt.Println("Player one goes gin!")
			return true, groupify(board.hand2) + 20
		} else if groupify(board.hand1) <= 10 {
			fmt.Println("Would you like to knock? yes(1), no(2)\n> ")
			fmt.Scan(&input)
			if input == 1 {
				var points = groupify(board.hand2) - groupify(board.hand1)
				if points > 0 {
					fmt.Println("Player one wins with a knock")
					return true, points
				} else if points < 0 {
					fmt.Println("Player two wins on the knock")
					points = -points
					return false, points
				} else {
					fmt.Println("Players tied on the knock")
					return true, points
				}
			}
		}
		fmt.Println("\n\n\n")
		fmt.Println("Player 2's hand: ")
		printHand(board.hand2)
		fmt.Println("Draw new card(1) or take top of discard pile(2)")
		fmt.Print("This is the top of the discard pile: ")
		printCard(board.discard[len(board.discard)-1])
		fmt.Print("\n> ")
		fmt.Scan(&input)
		if input == 1 {
			var deltCard = board.DealCard()
			board.hand2 = append(board.hand1, deltCard)
			fmt.Print("You drew: ")
			printCard(deltCard)
		} else if input == 2 {
			board.hand2 = append(board.hand2, board.discard[len(board.discard)-1])
			board.discard = board.discard[:len(board.discard)-1]
		}
		fmt.Println("Please enter the index of the card you would like to discard, starting at index 0")
		sortHand(board.hand2)
		printHand(board.hand2)
		fmt.Println("> ")
		fmt.Scan(&input)
		board.discard = append(board.discard, board.hand2[input])
		board.hand2 = append(board.hand2[:input], board.hand2[input+1:]...)
		if groupify(board.hand2) == 0 {
			fmt.Println("Player one goes gin!")
			return true, groupify(board.hand1) + 20
		} else if groupify(board.hand2) <= 10 {
			fmt.Println("Would you like to knock? yes(1), no(2)\n> ")
			fmt.Scan(&input)
			if input == 1 {
				var points = groupify(board.hand1) - groupify(board.hand2)
				if points > 0 {
					fmt.Println("Player one wins with a knock")
					return true, points
				} else if points < 0 {
					fmt.Println("Player two wins on the knock")
					points = -points
					return false, points
				} else {
					fmt.Println("Players tied on the knock")
					return true, points
				}
			}
		}
		fmt.Println("\n\n\n")
	}
	return true, 0
}
