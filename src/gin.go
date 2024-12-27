package main

import (
	"fmt"
)

func gameLoop(board *GameBoard, p1 Player, p2 Player) (bool, int) {
	var gameOver = false

	for !gameOver {
		var input int

		board.SortHands()

		fmt.Print("\n\n\n-----------------------------------------\n")
		fmt.Printf("\n%s's hand: ", p1.name)
		printHand(board.hand1)
		fmt.Print("This is the top of the discard pile: ")
		printCard(board.discard[len(board.discard)-1])
		fmt.Println("\nDraw new card(1) or take top of discard pile(2)")
		fmt.Print("> ")
		fmt.Scan(&input)
		if input == 1 {
			var deltCard = board.DealCard()
			board.hand1 = append(board.hand1, deltCard)
			fmt.Print("You drew: ")
			printCard(deltCard)
			fmt.Println()
		} else if input == 2 {
			board.hand1 = append(board.hand1, board.discard[len(board.discard)-1])
			board.discard = board.discard[:len(board.discard)-1]
		} else {
			fmt.Println("Please enter 1 or 2!")
		}
		fmt.Println("Please enter the index of the card you would like to discard, starting at index 0")
		sortHand(board.hand1)
		printHand(board.hand1)
		fmt.Println("[0]  [1]  [2]  [3]  [4]  [5]  [6]  [7]  [8]  [9]  [10]")
		fmt.Print("> ")
		fmt.Scan(&input)
		board.discard = append(board.discard, board.hand1[input])
		board.hand1 = append(board.hand1[:input], board.hand1[input+1:]...)
		P1_pts := groupify(board.hand1)
		P2_pts := groupify(board.hand2)
		fmt.Printf("You have %d deadwood points in your hand\n", P1_pts)
		if P1_pts == 0 {
			fmt.Printf("%s goes gin!", p1.name)
			return true, P2_pts + 25
		} else if P1_pts <= 10 {
			fmt.Printf("Would you like to knock with %d points? yes(1), no(2)\n> ", P1_pts)
			fmt.Scan(&input)
			if input == 1 {
				fmt.Printf("%s has %d deadwood points.\n%s has %d deadwood points.\n", 
							p1.name, P1_pts, p2.name, P2_pts)
				var pointDiff = P2_pts - P1_pts
				if pointDiff > 0 {
					fmt.Printf("%s wins with a knock! \n", p1.name)
					return true, pointDiff
				} else if pointDiff < 0 {
					fmt.Printf("%s undercuts the knock! \n", p2.name)
					pointDiff = -pointDiff
					return false, pointDiff + 20
				} else {
					fmt.Println("Players tied on the knock! ")
					return true, pointDiff
				}
			}
		}

		//Player 2's turn
		fmt.Print("\n\n\n-----------------------------------------\n")
		fmt.Printf("\n%s's hand: ", p2.name)
		printHand(board.hand2)
		fmt.Print("This is the top of the discard pile: ")
		printCard(board.discard[len(board.discard)-1])
		fmt.Print("\nDraw new card(1) or take top of discard pile(2)\n> ")
		fmt.Scan(&input)
		if input == 1 {
			var deltCard = board.DealCard()
			board.hand2 = append(board.hand2, deltCard)
			fmt.Print("You drew: ")
			printCard(deltCard)
			fmt.Println()
		} else if input == 2 {
			board.hand2 = append(board.hand2, board.discard[len(board.discard)-1])
			board.discard = board.discard[:len(board.discard)-1]
		} else {
			fmt.Println("Please enter 1 or 2!")
		}
		fmt.Println("Please enter the index of the card you would like to discard, starting at index 0")
		sortHand(board.hand2)
		printHand(board.hand2)
		fmt.Println("[0]  [1]  [2]  [3]  [4]  [5]  [6]  [7]  [8]  [9]  [10]")
		fmt.Print("> ")
		fmt.Scan(&input)
		board.discard = append(board.discard, board.hand2[input])
		board.hand2 = append(board.hand2[:input], board.hand2[input+1:]...)
		P1_pts = groupify(board.hand1)
		P2_pts = groupify(board.hand2)
		fmt.Printf("You have %d deadwood points in your hand\n", P2_pts)
		if P2_pts == 0 {
			fmt.Printf("%s goes gin!\n", p2.name)
			return true, P1_pts + 25
		} else if P2_pts <= 10 {
			fmt.Printf("Would you like to knock with %d points? yes(1), no(2)\n> ", P2_pts)
			fmt.Scan(&input)
			if input == 1 {
				fmt.Printf("%s has %d deadwood points.\n%s has %d deadwood points.\n", 
							p1.name, P1_pts, p2.name, P2_pts)
				var pointDiff = P2_pts - P1_pts
				if pointDiff > 0 {
					fmt.Printf("%s wins with a knock! ", p2.name)
					return false, pointDiff
				} else if pointDiff < 0 {
					fmt.Printf("%s undercuts the knock! ", p1.name)
					pointDiff = -pointDiff
					return true, pointDiff
				} else {
					fmt.Println("Players tied on the knock! ")
					return true, pointDiff
				}
			}
		}
	}
	return true, 0
}
