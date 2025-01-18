package gin

import(
    "strconv"
)

func drawCard(gb *GameBoard, player int) string {
    var output string
    hand :=  getPlayer(gb, player).hand
    hand = append(hand, gb.deck[len(gb.deck)-1])
    gb.deck = gb.deck[:len(gb.deck)-1]

    output += printHand(hand) + "\n>" 

    return output
}

func discard(gb *GameBoard, player int, n int) string {
    var output string
    hand := getPlayer(gb, player).hand

    gb.discard = append(gb.discard, hand[n])

    output += "\n" + printHand(hand) + "\n"

    return output
}

func drawFromDiscard (gb *GameBoard, player int) string {
    var output string
    hand := getPlayer(gb, player).hand

    hand = append(hand, gb.discard[len(gb.discard)-1])
    gb.discard = gb.discard[:len(gb.discard)-1]

    output += "\n" + printHand(hand) + "\n>"

    return output
}

func knock(gb *GameBoard, player int) string {
    var output string

    current := getPlayer(gb, player)
    currentPts := groupify(current.hand)
    player ^= player
    other := getPlayer(gb, player)
    otherPts := groupify(other.hand)

    output += "\n" + current.name + " has " + strconv.Itoa(currentPts) + " deadwood points\n"
    output += other.name + " has " + strconv.Itoa(otherPts) + " deadwood points\n" 

    ptsDiff := currentPts - otherPts
    if(ptsDiff > 0) {
        output += current.name + " wins with a knock and gets " + strconv.Itoa(ptsDiff) + " points\n"
        current.pts += currentPts
    } else if(ptsDiff < 0) {
        ptsDiff = -ptsDiff
        output += other.name + " undercuts the knock and gets " + strconv.Itoa(ptsDiff+20) + " points\n"
        other.pts += ptsDiff + 20
    } else {
        output += "Players tied on the knock\n>"
    }
    return output
}

// need to check if any cards can be player on the other players hand later
func goGin(gb *GameBoard, player int) string {
    var output string

    current := getPlayer(gb, player)
    player ^= player
    other := getPlayer(gb, player)
    pts := 50+groupify(other.hand)

    output += current.name + " Went gin and gets " + strconv.Itoa(pts) + " points\n>"
    current.pts += pts
    
    return output
}


