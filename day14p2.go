package main

import (
	"fmt"
	"os"

	// "regexp"
	// "strconv"
	"strings"
)

var rockMap [][]int

type Direction int64

const (
	Up    Direction = 0
	Down  Direction = 1
	Left  Direction = 2
	Right Direction = 4
)

func composeStoneSeq(rollingStoneCount int, totalLen int) []int {
	result := []int{}
	for j := 0; j < totalLen; j++ {
		if rollingStoneCount > j {
			result = append(result, 2)
		} else {
			result = append(result, 0)
		}
	}
	return result
}

func stackLineToStart(line []int) []int {
	result := []int{}
	pivot := -1
	rollingStoneCount := 0
	for i := 0; i < len(line); i++ {
		// fmt.Println(`i `, i)
		if line[i] == 2 {
			rollingStoneCount++
		}
		if line[i] == 1 {
			//Update pivot stone
			// fmt.Println(`pivot `, pivot)
			if pivot != -1 {
				result = append(result, 1)
			}
			result = append(result, composeStoneSeq(rollingStoneCount, i-pivot-1)...)
			pivot = i
			rollingStoneCount = 0
		}
	}

	if pivot != -1 {
		result = append(result, 1)
	}
	result = append(result, composeStoneSeq(rollingStoneCount, len(line)-pivot-1)...)
	// return []int{1, 2, 3, 4}
	return result
}

func moveUp(inputGrid [][]int) [][]int {
	result := [][]int{}
	for j := 0; j < height; j++ {
		result = append(result, make([]int, width))
	}
	// Horizontal
	for i := 0; i < width; i++ {
		// Vertical
		col := []int{}
		for j := 0; j < height; j++ {
			col = append(col, inputGrid[j][i])
		}
		// fmt.Println(`line 76 is for `, col)
		// fmt.Println(col)
		resultCol := stackLineToStart(col)
		// fmt.Println(`thelen `, len(resultCol))
		// fmt.Println(`start col `, col)

		// fmt.Println(`end col`, resultCol)
		for j := 0; j < len(resultCol); j++ {
			result[j][i] = resultCol[j]
		}
	}
	return result
}

func moveDown(inputGrid [][]int) [][]int {
	result := [][]int{}
	for j := 0; j < height; j++ {
		result = append(result, make([]int, width))
	}
	// Horizontal
	for i := 0; i < width; i++ {
		// Vertical
		col := []int{}
		for j := height - 1; j >= 0; j-- {
			col = append(col, inputGrid[j][i])
		}
		// fmt.Println(`line 76 is for `, col)
		// fmt.Println(col)
		resultCol := stackLineToStart(col)
		// fmt.Println(`thelen `, len(resultCol))
		// fmt.Println(`start col `, col)

		// fmt.Println(`end col`, resultCol)
		// for j := height - 1; j >= 0; j-- {
		for j := 0; j < height; j++ {
			result[j][i] = resultCol[height-1-j]
		}
	}
	return result
}

func moveLeft(inputGrid [][]int) [][]int {
	result := [][]int{}
	for j := 0; j < height; j++ {
		result = append(result, make([]int, width))
	}
	for j := 0; j < height; j++ {
		row := inputGrid[j]
		resultRow := stackLineToStart(row)
		// fmt.Println(`result row `, row)

		for i := 0; i < width; i++ {
			result[j][i] = resultRow[i]
		}
	}
	return result
}

func moveRight(inputGrid [][]int) [][]int {
	result := [][]int{}
	for j := 0; j < height; j++ {
		result = append(result, make([]int, width))
	}
	for j := 0; j < height; j++ {
		row := []int{}
		for i := 0; i < width; i++ {
			row = append(row, inputGrid[j][width-i-1])
		}
		resultRow := stackLineToStart(row)
		// fmt.Println(`result row `, row)

		for i := 0; i < width; i++ {
			result[j][i] = resultRow[width-i-1]
		}
	}
	return result
}

func getWeight(rockMap [][]int) int {
	weightSum := 0
	for j := 0; j < len(rockMap); j++ {
		for i := 0; i < len(rockMap[0]); i++ {
			if rockMap[j][i] == 2 {
				weightSum += len(rockMap) - j
			}
		}
	}
	return weightSum
}

var width int = 0
var height int = 0

var stackGrid [][][]int

func main() {

	// testline := []int{2, 2, 0, 2, 0, 1, 0, 2, 0, 1, 2, 2, 2, 2}
	// // testline := []int{1, 0, 1, 0, 2, 0, 1, 1, 0, 1, 2}
	// fmt.Println(testline)
	// test := stackLineToStart(testline)
	// fmt.Println(test)
	// return
	inputBytes, error := os.ReadFile("inputDay14.txt")
	// inputBytes, error := os.ReadFile("sampleDay14.txt")

	if error != nil {
		fmt.Println("fail")
		return
	}
	inputText := string(inputBytes)

	linesOfText := strings.Split(inputText, "\n")
	height = len(linesOfText) - 1 // last line is empty
	width = len(linesOfText[0])
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
	fmt.Println(`start working `)
	// upedMap := moveUp(rockMap)
	// fmt.Println(upedMap)
	// downedMap := moveRight(rockMap)
	stackGrid = [][][]int{}
	hitId := -1
	for i := 0; i < 1000000000; i++ {
		// for i := 0; i < 1; i++ {
		stackGrid = append(stackGrid, rockMap)
		rockMap = moveUp(rockMap)
		rockMap = moveLeft(rockMap)
		rockMap = moveDown(rockMap)
		rockMap = moveRight(rockMap)
		for stackId := 0; stackId < len(stackGrid) && hitId == -1; stackId++ {
			sameMap := true
			for y := 0; y < height && hitId == -1; y++ {
				for x := 0; x < width; x++ {
					if rockMap[y][x] != stackGrid[stackId][y][x] {
						sameMap = false
						break
					}
				}
			}
			if sameMap {
				hitId = stackId
				fmt.Println(`hit at `, stackId)
				break
			}
		}
		if hitId != -1 {
			fmt.Println(`same map `, hitId, ` and `, i)
			break
		}
	}
	// for stackId := 0; stackId < len(stackGrid); stackId++ {
	// 	fmt.Println(stackId, " stack \n", stackGrid[stackId])
	// }
	fmt.Println("repeating \n", rockMap)
	// at 1000000000
	fmt.Println(`hit at `, hitId)
	fmt.Println(`repeat at `, len(stackGrid))
	loopLength := len(stackGrid) - hitId
	fmt.Println(`loop length `, loopLength)
	targetPos := (1000000000-hitId)%loopLength + hitId
	fmt.Println(`position `, targetPos)
	fmt.Println(stackGrid[targetPos])
	// fmt.Println(getWeight(stackGrid[targetPos-1]))
	fmt.Println(getWeight(stackGrid[targetPos]))
	// fmt.Println(getWeight(stackGrid[targetPos+1]))
}
