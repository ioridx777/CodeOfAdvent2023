package main

import (
	"fmt"
	"os"

	// "regexp"
	// "strconv"
	"strings"
)

var locMap = map[string][2]string{}
var instruct = []int{}

type Path struct {
	hitStep       []int
	startLoopStep int
	endLoopStep   int
}

func main() {
	inputBytes, error := os.ReadFile("inputDay8.txt")

	if error != nil {
		fmt.Println("fail")
		return
	}
	inputText := string(inputBytes)

	linesOfText := strings.Split(inputText, "\n")
	for i := 0; i < len(linesOfText[0]); i++ {
		fmt.Println(`destructing `, linesOfText[0][i:i+1])
		nextInstruct := (linesOfText[0][i : i+1])
		nextInstructNum := 0
		if nextInstruct == "L" {
			nextInstructNum = 0
		} else if nextInstruct == "R" {
			nextInstructNum = 1
		} else {
			continue
		}
		instruct = append(instruct, nextInstructNum)
	}
	fmt.Println(`final instruct `, instruct)
	for lineIndex, line := range linesOfText {
		if lineIndex < 2 {
			continue
		}
		if len(line) == 0 {
			continue
		}
		// fmt.Println(line)
		if len(line) == 0 {
			continue
		}
		locMap[line[0:3]] = [2]string{line[7:10], line[12:15]}
	}
	startingPoss := []string{}
	for pos := range locMap {
		if pos[len(pos)-1:] == "A" {
			startingPoss = append(startingPoss, pos)
		}
	}
	startingPoss = startingPoss[0:1]
	fmt.Println(startingPoss)

	// isReach := false
	// step := 0
	// // fastWalks := make(FastWalk, len(startingPoss))
	// repeatCheck := map[string]bool{}
	// for !isReach {
	// 	// if checkAllPos(startingPoss) {
	// 	// 	isReach = true
	// 	// 	break
	// 	// }
	// 	for _, nextInstruct := range instruct {
	// 		// startingPos = locMap[startingPos][nextInstruct]
	// 		// fmt.Println(step, `walk `, nextInstruct, ": ", startingPos)
	// 		nextPoss := []string{}
	// 		for _, pos := range startingPoss {
	// 			nextPos := locMap[pos][nextInstruct]
	// 			nextPoss = append(nextPoss, nextPos)
	// 		}
	// 		startingPoss = nextPoss
	// 		step++
	// 		if checkAllPos(startingPoss) {
	// 			isReach = true
	// 			break
	// 		}
	// 	}
	// 	fmt.Println(`this walk end in `, startingPoss)
	// 	if repeatCheck[startingPoss[0]] {
	// 		fmt.Println(`Is repeat `, startingPoss[0])
	// 		break
	// 	} else {
	// 		repeatCheck[startingPoss[0]] = true
	// 	}
	// }
	// fmt.Println(`hellowlrd`, step)
}
