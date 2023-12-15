package main

import (
	"bufio"
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

var walkedMap [][]int

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
	inputBytes, error := os.ReadFile("sampleDay10p2.txt")

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
	for i := 0; i < height; i++ {
		walkedMap = append(walkedMap, make([]int, width))
	}
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
	nextDir := Down
	moveCounter := 0
	for nextDir > 0 {
		walkedMap[nextY][nextX] = 3
		prevX, prevY := nextX, nextY
		prevDir := nextDir
		nextX, nextY = getNextCoor(nextX, nextY, nextDir)
		nextDir = checkDir(nextDir, strMap[nextX+nextY*width])
		moveCounter++
		fmt.Println(moveCounter, ": ", prevX, ",", prevY, " = ", prevDir, " => ", nextX, ",", nextY)
		fmt.Println("next dir ", nextDir)
		fmt.Println(`counter `, moveCounter)
	}
	for _, v := range walkedMap {
		fmt.Println(v)
	}
	distSum := 0

	// VH=3,V=2,H=1, None=0

	// for j := 0; j < 2; j++ {
	for j := 0; j < height; j++ {

		for i := 0; i < width; i++ {

			if walkedMap[j][i] > 0 {
				hDist := 1
				vDist := 1

				//go right
				if walkedMap[j][i]%2 == 1 {
					for hDist = 1; hDist+i < width; hDist++ {
						if walkedMap[j][i+hDist]%2 == 1 {
							walkedMap[j][i+hDist] -= 1
							break
						}
					}
				}
				//go down
				if walkedMap[j][i] >= 2 {
					for vDist = 1; vDist+j < height; vDist++ {
						if walkedMap[j+vDist][i] >= 2 {
							walkedMap[j+vDist][i] -= 2
							break
						}
					}
				}
				if hDist+i == width {
					hDist = 1
				}
				if vDist+j == height {
					vDist = 1
				}
				// if hDist > 1 {

				walkedMap[j][i] = 0

				fmt.Println(`hitting `, i, " , ", j)
				if hDist > 1 {
					fmt.Println(`h hit `, i+hDist, ",", j)
				}
				if vDist > 1 {
					fmt.Println(`v hit `, i, ",", j+vDist)
				}
				// }
				fmt.Println(`starting sum`, distSum)
				fmt.Println(`x `, hDist)
				fmt.Println(`y`, vDist)
				distSum += hDist - 1
				distSum += vDist - 1
				fmt.Println(`ending sum`, distSum)
				if i == 1 && j == 2 {
					fmt.Println(`new map `, distSum)
					for _, v := range walkedMap {
						fmt.Println(v)
					}
				}
			}
		}
		fmt.Println(`new map `, distSum)
		for _, v := range walkedMap {
			fmt.Println(v)
		}
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		reader.ReadString('\n')
	}
	for _, v := range walkedMap {
		fmt.Println(v)
	}
	fmt.Println(`area `, distSum)
	fmt.Println(`hellowlrd`)

}
