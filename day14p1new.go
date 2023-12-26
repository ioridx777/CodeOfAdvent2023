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
	return 0
}

func main() {
	// inputBytes, error := os.ReadFile("inputDay14.txt")
	inputBytes, error := os.ReadFile("sampleDay14.txt")

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
	// for i := 0; i < width; i++ {
	sum := 0
	for i := 0; i < 2 && i < width; i++ {
		fmt.Println(`Rock Col `, i)
		for j := 0; j < height; j++ {
			fmt.Print(rockMap[j][i])
		}
		fmt.Println()
		for j := 0; j < height; j++ {
			if rockMap[j][i] == 1 || j == 0 {
				fmt.Println(`hitting at `, j)
				if j == 0 || rockMap[j][i] == 1 {
					//this is a counter
					stackingAt := j
					stackingRock := 0
					j++
					for j < height && rockMap[j][i] != 2 {
						fmt.Println(`checking at `, j)
						if rockMap[j][i] == 1 {
							stackingRock++
							break
						}
						j++
					}
					pileWeight := getWeight(stackingAt, stackingRock)
					fmt.Println(`stacking at `, stackingAt, ` : `, stackingRock)
					sum += pileWeight
				}
			}
		}
	}

	fmt.Println(`Result `, sum)

}
