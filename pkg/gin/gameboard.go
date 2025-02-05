package gin

import (
	"math/rand/v2"
)

/*
Game board structures
*/
type Card struct {
	value, suit int
}
type Player struct {
    name string
    hand []Card 
    pts int
}

type GameBoard struct {
	deck    []Card
	discard []Card
    p1 Player
    p2 Player
}

func (board *GameBoard) IntitializeBoard() {
	//create every card, from val 0-12 and suit 0-3
	for i := 0; i < 52; i++ {
		board.deck = append(board.deck, Card{value: i % 13, suit: i / 13})
	}

	//initialize hand1, hand2, discard
	board.p1.hand = make([]Card, 0, 11)
	board.p2.hand = make([]Card, 0, 11)
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
			board.p1.hand = append(board.p1.hand, board.DealCard())
		} else {
			board.p2.hand = append(board.p2.hand, board.DealCard())
		}
	}
    board.discard = append(board.discard, board.DealCard())
}

func printHand(hand []Card) string {
	value := []string{" A", " 2", " 3", " 4", " 5", " 6", " 7", " 8", " 9", "10", " J", " Q", " K"}
	suit := []string{"♠", "♦", "♣", "♥"}
    var ret string
	for x, card := range hand {
		ret += value[card.value] + suit[card.suit]
		if (x != len(hand)-1) {ret += ", "}
	}
    ret += "\n"
    return ret
}

func printCard(card Card) string {
	value := []string{" A", " 2", " 3", " 4", " 5", " 6", " 7", " 8", " 9", "10", " J", " Q", " K"}
	suit := []string{"♠", "♦", "♣", "♥"}
    cardString :=  value[card.value ] + suit[card.suit] + "\n"
    return cardString
}

func (board *GameBoard) DealCard() Card {
	var card Card = board.deck[len(board.deck)-1]
	board.deck = board.deck[:len(board.deck)-1]
	return card
}

func (board *GameBoard) SortHands() {
	board.p1.hand = sortHand(board.p1.hand)
	board.p2.hand = sortHand(board.p2.hand)
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
	straights := findStraights(hand)
	//get the pairs
	matches := findMatches(hand)
	//find conflicts
	for a, straight := range straights {
		for b, straightCard := range straight {
			for c, match := range matches {
				for d, matchCard := range match {
					if straightCard == matchCard {
						//check if one group is length 4, keep it there
						if (len(straights[a]) > 3 && (b == 0 || b == len(straights[a])-1)) { //remove from this one
							straights[a] = append(straights[a][:b], straights[a][b+1:]...)
						} else if len(matches[c]) > 3 { //remove from this one
							matches[c] = append(matches[c][:d], matches[c][d+1:]...)
						} else { // both length 3
							straightValue := 0
							matchValue := 0
							for _, card := range straight {
								straightValue += card.value
							}
							for _, card := range match {
								matchValue += card.value
							}
							if (straightValue > matchValue) { //remove from matches
								matches = append(matches[:c], matches[c+1:]...)
							} else { //remove from straights
								straights = append(straights[:a], straights[a+1:]...)
							}
							
						}
					}
				}
			}
		}
	}
	//remove groups of len < 3
	for x,straight := range straights {
		if len(straight) < 3 {
			straights = append(straights[:x], straights[x+1:]...)
		}
	}
	for y,match := range matches {
		if len(match) < 3 {
			matches = append(matches[:y], matches[y+1:]...)
		}
	}
	//calculate hand total
	handTotal := calcTotal(hand) - calcTotalGroups(straights) - calcTotalGroups(matches)
	
    return handTotal
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

func calcTotal(hand []Card) int {
	total := 0
	for _, card := range hand {
		total += card.value
	}
	return total }
func calcTotalGroups(hand [][]Card) int {
	total := 0
	for _, slice := range hand {
		for _, card := range slice {
			total += card.value
		}
	}
	return total
}


func getPlayer(gb *GameBoard, n int) *Player {
    if n == 0 { 
        return &gb.p1 
    } else { 
        return &gb.p2 
    }
}
