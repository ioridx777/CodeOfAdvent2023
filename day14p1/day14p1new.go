package main

import (
	"fmt"
	"os"
	// "regexp"
	// "strconv"
	"strings"
)

var rockMap [][]int

func getWeight(stackingAt int, stackingRock int) int {
	startWeight := len(rockMap) - stackingAt
	fmt.Println(`start weight `, startWeight)
	return ((startWeight*2 - stackingRock + 1) * stackingRock) / 2
}

func main() {
	inputBytes, error := os.ReadFile("inputDay14.txt")
	// inputBytes, error := os.ReadFile("sampleDay14.txt")

	if error != nil {
		fmt.Println("fail")
		return
	}
	inputText := string(inputBytes)

	linesOfText := strings.Split(inputText, "\n")
	height := len(linesOfText) - 1 // last line is empty
	width := len(linesOfText[0])
	for _, line := range linesOfText {
		// fmt.Println(line)
		if len(line) == 0 {
			continue
		}
		lineOfValue := []int{}
		for i := 0; i < len(line); i++ {
			resultVal := 0
			switchVal := line[i : i+1]
			switch switchVal {
			case "O":
				resultVal = 2
			case "#":
				resultVal = 1
			case ".":
				resultVal = 0
			}
			lineOfValue = append(lineOfValue, resultVal)
		}
		fmt.Println(lineOfValue)
		rockMap = append(rockMap, lineOfValue)
	}
	// sum := 0
	sum := 0
	for i := 0; i < width; i++ {
		// for i := 2; i < 3 && i < width; i++ {
		fmt.Println(`Rock Col `, i)
		for j := 0; j < height; j++ {
			fmt.Print(rockMap[j][i])
		}
		fmt.Println()
		for j := -1; j < height; {
			fmt.Println(`starting j `, j)
			fmt.Println(`current height `, height)
			fmt.Print(i, `th col rolling start at `, j, ` | `)
			//this is a counter
			stackingAt := j
			stackingRock := 0
			j++
			for j < height && rockMap[j][i] != 1 {
				if rockMap[j][i] == 2 {
					fmt.Print(" [", j, "]")
					stackingRock++
				} else {
					fmt.Print(" ", j)
				}
				j++
			}
			fmt.Println()
			pileWeight := getWeight(stackingAt+1, stackingRock)
			fmt.Println(`pile at `, stackingAt, ` : `, stackingRock)
			fmt.Println(`with weight `, pileWeight)
			sum += pileWeight

		}
	}

	fmt.Println(`Result `, sum)

}
