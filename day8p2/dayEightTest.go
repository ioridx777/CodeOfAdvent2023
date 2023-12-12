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

func Amain() {
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
	isReach := false
	startingPos := "AAA"
	step := 0
	for !isReach {
		// if startingPos == "ZZZ" {
		// 	isReach = true
		// 	break
		// }
		fmt.Println(`Step `, step)
		if step == 31742 {
			fmt.Println(startingPos)
			break
		}
		for _, nextInstruct := range instruct {
			startingPos = locMap[startingPos][nextInstruct]
			// fmt.Println(step, `walk `, nextInstruct, ": ", startingPos)
			step++
			// if startingPos == "ZZZ" {
			// 	isReach = true
			// 	break
			// }
		}
		// fmt.Println(`this walk end in `, startingPos)
	}
	fmt.Println(`hellowlrd`, step)
}
