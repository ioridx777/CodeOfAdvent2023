package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var ballColor = []string{"blue", "green", "red"}
var ballCount = []int{14, 13, 12}

func processDay2Line(line string, ballType int) int {
	colorWlen := len(ballColor[ballType]) + 1

	r, _ := regexp.Compile("[0-9]+ " + ballColor[ballType])
	matchedIndex := r.FindStringIndex(line)
	// fmt.Println(line[matchedIndex[0]:matchedIndex[-len(" blue")]])
	if len(matchedIndex) == 0 {
		return 0
	}
	// fmt.Println(line[matchedIndex[0] : matchedIndex[1]-colorWlen])
	value, _ := strconv.Atoi(line[matchedIndex[0] : matchedIndex[1]-colorWlen])
	prevMatchIndex := matchedIndex
	newMatchIndex := r.FindStringIndex(line[prevMatchIndex[1]:])
	for len(newMatchIndex) != 0 {
		newValue, _ := strconv.Atoi(line[newMatchIndex[0]+prevMatchIndex[1] : +prevMatchIndex[1]+newMatchIndex[1]-colorWlen])
		// fmt.Println(`new value `, line[newMatchIndex[0]+prevMatchIndex[1]:+prevMatchIndex[1]+newMatchIndex[1]])
		if newValue > value {
			value = newValue
		}
		// fmt.Println(`prev`, prevMatchIndex)
		// fmt.Println(`new`, newMatchIndex)
		// fmt.Println(`total `, len(line))
		// fmt.Println(`A `, newMatchIndex[0]+1+prevMatchIndex[1])
		// fmt.Println(`B `, newMatchIndex[1]+1+prevMatchIndex[1])
		prevMatchIndex = []int{newMatchIndex[0] + prevMatchIndex[1], newMatchIndex[1] + prevMatchIndex[1]}
		// fmt.Println(`Remaining `, line[prevMatchIndex[1]:])
		// fmt.Println(prevMatchIndex[1], " vs ", len(line))
		if prevMatchIndex[1] >= len(line)-1 {
			break
		}
		newMatchIndex = r.FindStringIndex(line[prevMatchIndex[1]:])
	}
	fmt.Print(ballColor[ballType], `: `, value, `/`, ballCount[ballType])
	fmt.Println(` Success: `, value <= ballCount[ballType])
	return value
}

func main2p2() {
	// inputBytes, error := os.ReadFile("sampleDay2.txt")
	inputBytes, error := os.ReadFile("inputDay2.txt")

	if error != nil {
		fmt.Println("fail")
		return
	}
	inputText := string(inputBytes)

	linesOfText := strings.Split(inputText, "\n")
	sum := 0
	for _, line := range linesOfText {
		fmt.Println(line)
		if len(line) == 0 {
			continue
		}
		multi := processDay2Line(line, 0) * processDay2Line(line, 1) * processDay2Line(line, 2)
		fmt.Println(`Adding `, multi)
		sum += multi

	}
	fmt.Println(`result `, sum)
	fmt.Println(`hellowlrd`)
}
