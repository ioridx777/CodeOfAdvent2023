package main

import (
	"fmt"
	"os"
	// "regexp"
	// "strconv"
	"strings"
)

var rockMap [][]int

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
	sum := 0
	// for i := 0; i < width; i++ {
	for i := 0; i < 1 && i < width; i++ {

		rollingRockCount := 0
		// prevRollingTo:=0
		rollingTo := 0
		newRollingLine := make([]int, height)
		for j := 0; j < height; j++ {
			fmt.Println(j, "J ")
			fmt.Println(rockMap[j][i])

			switch rockMap[j][i] {
			case 2:
				rollingRockCount++
			case 1:
				fmt.Println(`the rolling cout`, rollingRockCount)
				for k := 0; k < j-rollingTo; k++ {
					if rollingRockCount > k {
						newRollingLine[k+rollingTo] = 2
					} else {
						newRollingLine[k+rollingTo] = 0
					}
				}
				rollingRockCount = 0
				newRollingLine[j] = 1
				rollingTo = j + 1
			case 0:
			}

			if j == height-1 && rockMap[j][i] != 1 {
				fmt.Println(`here`)
				for k := 0; k < j-rollingTo; k++ {
					if rollingRockCount > k {
						newRollingLine[k+rollingTo] = 2
					} else {
						newRollingLine[k+rollingTo] = 0
					}
				}
			}
		}
		fmt.Println(`start sum `, sum)
		lineAdding := 0
		for i, val := range newRollingLine {
			if val == 2 {

				fmt.Println(i, ` adding `, height-i)
				lineAdding += height - i
			}
		}
		fmt.Println(i, ` line adding `, lineAdding)
		sum += lineAdding
		fmt.Println(`end sum `, sum)
		fmt.Println(`newRollingLine `, newRollingLine)
	}
	fmt.Println(`hellowlrd`)

}
