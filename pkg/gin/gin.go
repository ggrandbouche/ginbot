package gin

import(
    "strings"
    "strconv"
    "fmt"
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
    //gameOver := false

    output <- Output{Output: "I:Welcome to the game of gin\n Please enter your name\n>", Player: 2}
    board.p1.name = <-input 
    board.p2.name = <-input

    output <- Output{Output: "I:" + turn(board, 0) + "\n>", Player: 0}
    output <- Output{Output: "I:" + turn(board, 0) + parser(&board, 0, <-input), Player: 0}

    output <- Output{Output: "I:" + turn(board, 1) + "\n>", Player: 1}
    output <- Output{Output: "I:" + turn(board, 1) + parser(&board, 2, <-input), Player: 1}
    /*
    for !gameOver {
        fmt.Println("inside loop")
        output <- Output{Output: turn(board, 0) + parser(&board, 0, <-input), Player: 0}
        output <- Output{Output: turn(board, 0) + parser(&board, 0, <-input), Player: 0}

        output <- Output{Output: turn(board, 1) + parser(&board, 2, <-input), Player: 1}
        output <- Output{Output: turn(board, 1) + parser(&board, 2, <-input), Player: 1}
    }
    */
    
}

func turn(gb GameBoard, player int) string {
    if player == 0 {
        return "Hand: " + printHand(gb.p1.hand) + "Discard pile: " + printCard(gb.discard[len(gb.discard)-1])
    } else {
        return "Hand: " + printHand(gb.p2.hand) + "Discard pile: " + printCard(gb.discard[len(gb.discard)-1])
    }
}

func parser(gb *GameBoard, player int, input string) string {

    fmt.Println("parser called")
    inputArr := strings.Split(input, " ")

    if strings.ToLower(inputArr[0]) == "draw" {
        if strings.ToLower(inputArr[1]) == "from" && strings.ToLower(inputArr[2]) == "discard" {
            return drawFromDiscard(gb, player)
        } else {
            return drawCard(gb, player)
        }
    } else if strings.ToLower(inputArr[0]) == "discard" {
        index, err := strconv.Atoi(inputArr[1])
        if err != nil {
            fmt.Println("Error reading index of card to discard")
            return ""
        }
        return discard(gb, player, index)
    } else if strings.ToLower(inputArr[0]) == "knock" {
        return knock(gb, player)
    } else if strings.ToLower(inputArr[0]) == "go" && strings.ToLower(inputArr[1]) == "gin" {
        return goGin(gb, player) 
    }

    return "invalid" 
}
