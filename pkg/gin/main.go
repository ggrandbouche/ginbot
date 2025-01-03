package gin

import (
	"fmt"
)

type Player struct {
	name string
	points int
}

func gin() {
	//header
	fmt.Println("\n\n\t\tWelcome to ginbot - a bot the plays gin\n\t")
	//initialize
	var board = GameBoard{}
	board.IntitializeBoard()
	board.Shuffle()
	board.DealHands()
	board.SortHands()

	var p1, p2 Player = Player{name:"", points:0}, Player{name:"", points:0}
	fmt.Println("Please enter a name for Player 1")
	fmt.Scan(&p1.name)
	fmt.Println("Please enter a name for Player 2")
	fmt.Scan(&p2.name)

	var playing bool = true
	//beginning game
	for playing {
		var player1Wins, pointsWon = gameLoop(&board, p1, p2)
		
		//keep track of points
		if pointsWon == 0 {
			//do nothing
		} else if player1Wins {
			p1.points += pointsWon
			fmt.Printf("%s gained %d points. \n", p1.name, pointsWon)
		} else {
			p2.points += pointsWon
			fmt.Printf("%s gained %d points. \n", p2.name, pointsWon)
		}
		fmt.Printf("%s has %d points. \n", p1.name, p1.points)
		fmt.Printf("%s has %d points. \n", p2.name, p2.points)
		
		var s string = ""
		for !(s=="p" || s=="P" || s=="q" || s=="Q"){
			fmt.Println("\n\n\tThanks for playing! Enter P to play again or Q to quit.")
			fmt.Scan(&s)
			if s == "q" || s == "Q" {
				return
			} else if s == "p" || s == "P" {
				continue
			}
		}

	}
    

}

