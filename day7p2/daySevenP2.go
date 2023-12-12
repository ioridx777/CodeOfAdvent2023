package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	// "regexp"
	// "strconv"
	"strings"
)

var strengthTable = []int{}

type Player struct {
	hand     []int
	bid      int
	strength int
	ranking  int
	suit     int
}

func cardToInt(cardStr string) int {
	current := 0
	switch cardStr {
	case "A":
		current = 14
	case "K":
		current = 13
	case "Q":
		current = 12
	case "J":
		current = 1
	case "T":
		current = 10
	default:
		current, _ = strconv.Atoi(cardStr)
	}
	return current
}

func getStrength(hand []int) (int, int) {
	numberMap := map[int]int{}
	result := 0
	for _, val := range hand {
		numberMap[val]++
		result = result*100 + val
	}
	maxType := 0
	currentType := 0
	jCount := numberMap[1]

	fmt.Println(`number `, numberMap)
	for v, count := range numberMap {
		wildCount := jCount
		if v == 1 && count == 5 {
			maxType = 6
			continue
		}
		if v == 1 {
			wildCount = 0
			continue
		}
		fmt.Println(`hand `, hand, `: `, v)
		fmt.Println(`wild `, wildCount)
		fmt.Println(`count `, count)
		switch {
		case count+wildCount == 5: //five of a kind
			currentType = 6
		case count+wildCount == 4: //four of a kind
			currentType = 5
		case count == 3 && maxType == 1: //full house
			fallthrough
		case count == 2 && maxType == 3: //full house
			currentType = 4
		case count+wildCount == 3: //three of a kind
			currentType = 3
		case count == 2 && maxType == 1: //two pair
			currentType = 2
		case count+wildCount == 2:
			currentType = 1
		}
		if currentType > maxType {
			fmt.Println(`updating to `, currentType)
			maxType = currentType
		}

	}
	fmt.Println(`hand `, hand)
	fmt.Println(`max type `, maxType)
	return maxType*10000000000 + result, maxType
}

type PlayerList []Player

func (p PlayerList) Len() int {
	return len(p)
}

// func (p PlayerList) Less(i, j int) bool {
// 	if p[i].strength == p[j].strength {
// 		panic("Here it goes ")
// 	}
// 	return p[i].strength < p[j].strength
// }

func (p PlayerList) Less(i, j int) bool {
	if p[i].strength == p[j].strength {
		panic("Here it goes ")
	}
	if p[i].suit == p[j].suit {
		for k := 0; k < 5; k++ {
			if p[i].hand[k] == p[j].hand[k] {
				continue
			}
			return p[i].hand[k] < p[j].hand[k]
		}
	}
	return p[i].suit < p[j].suit
}
func (p PlayerList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	inputBytes, error := os.ReadFile("inputDay7.txt")

	if error != nil {
		fmt.Println("fail")
		return
	}
	inputText := string(inputBytes)

	linesOfText := strings.Split(inputText, "\n")
	playerList := PlayerList{}
	for _, line := range linesOfText {
		// fmt.Println(line)
		if len(line) == 0 {
			continue
		}
		hand := [5]int{}
		for i := 0; i < 5; i++ {
			hand[i] = cardToInt(line[i : i+1])
		}
		bid, _ := strconv.Atoi(line[6:])

		// fmt.Println(`bid`, line[6:])
		// fmt.Println(`bid`, bid)
		strength, suit := getStrength(hand[:])
		newPlayer := Player{
			hand:     hand[:],
			bid:      bid,
			strength: strength,
			suit:     suit,
		}
		playerList = append(playerList, newPlayer)
		// fmt.Println(getStrength(hand[:]))
	}
	// for playerIndex, player := range playerList {
	// 	fmt.Println(playerIndex, "n: ", player)
	// }
	sort.Sort(playerList)
	sum := 0
	for playerIndex, player := range playerList {
		fmt.Println(playerIndex, ": ", player)
		// sum += player.bid
		// fmt.Println(sum, " + ", player.bid, " * ", (playerIndex + 1))
		sum += (playerIndex + 1) * player.bid
	}

	fmt.Println(`hellowlrd`, sum)
}
