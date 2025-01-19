package gin

import(
    "strings"
    "strconv"
    //"fmt"
)

type Output struct {
    Output string
    Player int  // 0, 1, 2 -- p1, p2, both
}

func Gin(input <-chan string, output chan<- Output) {
    board := GameBoard{
        deck: []Card{}, 
        discard: []Card{}, 
        p1:  Player{name:"", hand: []Card{}, pts:0}, 
        p2: Player{name:"", hand: []Card{}, pts:0},
    }
    board.IntitializeBoard()
    board.Shuffle()
    board.DealHands()

    output <- Output{Output: "I:Welcome to the game of gin\nPlease enter your name\n>", Player: 0}
    board.p1.name = <-input 
    output <- Output{Output: "I:Welcome to the game of gin\nPlease enter your name\n>", Player: 1}
    board.p2.name = <-input

    for curPlayer := 0; curPlayer != 2; curPlayer ^= 1{
        board.SortHands()
        output <- Output{Output: "I:" + turn(board, curPlayer) + "\n>", Player: curPlayer}
        tempOutput, gameOver := parser(&board, curPlayer, <-input)
        if gameOver {
            output <- Output{Output: tempOutput, Player: 2}
            break
        } else {
            output <- Output{Output: "I:" + tempOutput, Player: curPlayer}
            tempOutput, gameOver = parser(&board, curPlayer, <-input)
            board.SortHands()
            output <- Output{Output: turn(board, curPlayer) + tempOutput, Player: curPlayer}
        }
    }
   
}

func turn(gb GameBoard, player int) string {
    if player == 0 {
        return "Hand: " + printHand(gb.p1.hand) + "Discard pile: " + printCard(gb.discard[len(gb.discard)-1])
    } else {
        return "Hand: " + printHand(gb.p2.hand) + "Discard pile: " + printCard(gb.discard[len(gb.discard)-1])
    }
}

func parser(gb *GameBoard, player int, input string) (string, bool) {
    inputArr := strings.Fields(input)

    if strings.ToLower(inputArr[0]) == "draw" {
        if len(inputArr) > 2 && strings.ToLower(inputArr[1]) == "from" && strings.ToLower(inputArr[2]) == "discard" {
            return drawFromDiscard(gb, player), false
        } else if len(inputArr) == 1 {
            return drawCard(gb, player), false
        }
    } else if index, err := strconv.Atoi(inputArr[0]); err == nil && index <= 10{
        discard(gb, player, index)
        return "\n>", false
    } else if strings.ToLower(inputArr[0]) == "knock" {
        return knock(gb, player)
    } else if strings.ToLower(inputArr[0]) == "go" && strings.ToLower(inputArr[1]) == "gin" || strings.ToLower(inputArr[0]) == "gin" {
        return goGin(gb, player) 
    }

    return "invalid", false
}
