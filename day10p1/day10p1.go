package main

import (
	"fmt"
	"os"

	// "regexp"
	// "strconv"
	"strings"
)

const (
	Up    int = 1
	Down      = 2
	Left      = 3
	Right     = 4
)

var width, height int

func reverseDir(dir int) int {
	if dir%2 == 0 {
		return dir - 1
	} else {
		return dir + 1
	}
}

func pipeDir(from int, dir1 int, dir2 int) int {
	if from == reverseDir(dir1) {
		return dir2
	} else if from == reverseDir(dir2) {
		return dir1
	}
	return -1
}

func getNextCoor(x int, y int, dir int) (int, int) {
	if dir < 1 || dir > 5 {
		//wrong direction
		return -1, -1
	}
	switch dir {
	case Up:
		if y-1 >= 0 {
			return x, y - 1
		}
	case Down:
		if y+1 < height {
			return x, y + 1
		}
	case Left:
		if x-1 >= 0 {
			return x - 1, y
		}
	case Right:
		if x+1 < width {
			return x + 1, y
		}
	}
	return -1, -1
}

func checkDir(from int, val string) int {
	fmt.Println(`checkdir `, from, ` val `, val)
	switch val {
	case ".":
		return -1
	case "|":
		return pipeDir(from, Up, Down)
	case "-":
		return pipeDir(from, Right, Left)
	case "L":
		return pipeDir(from, Up, Right)
	case "J":
		return pipeDir(from, Up, Left)
	case "7":
		return pipeDir(from, Left, Down)
	case "F":
		return pipeDir(from, Right, Down)
	case "S":
		println(`Success`)
		return -1
	}
	fmt.Println(`fall off`)
	return -1
}

func main() {
	inputBytes, error := os.ReadFile("inputDay10.txt")

	if error != nil {
		fmt.Println("fail")
		return
	}
	inputText := string(inputBytes)

	linesOfText := strings.Split(inputText, "\n")
	strMap := []string{}
	startingPoint := [2]int{}
	height = len(linesOfText)
	width = len(linesOfText[0])
	for y, line := range linesOfText {
		fmt.Println(line)
		if len(line) == 0 {
			continue
		}
		for i := 0; i < len(line); i++ {
			strMap = append(strMap, line[i:i+1])
			if line[i:i+1] == "S" {
				startingPoint[0] = i
				startingPoint[1] = y
			}
		}
	}
	fmt.Println(`starting at `, startingPoint)
	nextX, nextY := startingPoint[0], startingPoint[1]
	nextDir := Up
	moveCounter := 0
	for nextDir > 0 {
		prevX, prevY := nextX, nextY
		prevDir := nextDir
		nextX, nextY = getNextCoor(nextX, nextY, nextDir)
		nextDir = checkDir(nextDir, strMap[nextX+nextY*width])
		moveCounter++
		fmt.Println(moveCounter, ": ", prevX, ",", prevY, " = ", prevDir, " => ", nextX, ",", nextY)
		fmt.Println("next dir ", nextDir)
		fmt.Println(`counter `, moveCounter)
	}
	fmt.Println(`final `, moveCounter/2)
	fmt.Println(`hellowlrd`)
}
