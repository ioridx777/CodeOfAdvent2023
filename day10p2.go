package main

import (
	"fmt"
	"os"
	"strings"
)

// 1 expand grid to 9 times
// 2 array to store should flood and isflood

//try merge area solution
//Assign each block with unique value, connected pos to one value
//all value doesnt connected to outer value is trapped

var expandSymbol = [...]string{"|", "-", "L", "J", "7", "F", ".", "S"}
var expandSymbolBool = [...][9]bool{
	[...]bool{false, true, false, false, true, false, false, true, false},
	[...]bool{false, false, false, true, true, true, false, false, false},
	[...]bool{false, true, false, false, true, true, false, false, false},
	[...]bool{false, true, false, true, true, false, false, false, false},
	[...]bool{false, false, false, true, true, false, false, true, false},
	[...]bool{false, false, false, false, true, true, false, true, false},
	[...]bool{false, false, false, false, false, false, false, false, false},
	[...]bool{false, false, false, true, true, false, false, true, false}, //depends on your sample
	// [...]bool{false, true, false, true, true, false, false, false, false}, //depends on your input
}

const (
	Up    int = 1
	Down      = 2
	Left      = 3
	Right     = 4
)

var width, height int
var expandedMap []bool
var isFloodedDir, shouldFloodedDir [][4]bool
var floodColor []int

func getExpFromSymbol(symbol string) [9]bool {
	for index, val := range expandSymbol {
		if symbol == val {
			return expandSymbolBool[index]
		}
	}
	panic(`unknown symbol`)
}

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
	height = len(linesOfText) - 1
	width = len(linesOfText[0])
	for _, line := range linesOfText {
		fmt.Println(line)
		if len(line) == 0 {
			continue
		}
		fmt.Println(`is line same `, len(line))
		for i := 0; i < len(line); i++ {
			strMap = append(strMap, line[i:i+1])
		}
	}
	expandedMap = make([]bool, len(strMap)*9)

	for loc, val := range strMap {
		expandedLoc := loc * 9
		symbolMap := getExpFromSymbol(val)
		for i := 0; i < 9; i++ {
			expandedMap[expandedLoc+i] = symbolMap[i]
		}

	}
	expandedWidth := width * 3
	expandedHeight := height * 3
	fmt.Println(`width `, width)
	fmt.Println(`height `, height)
	fmt.Println(`expandedWidth `, expandedWidth)
	fmt.Println(`expandedHeight `, expandedHeight)
	fmt.Println(`extend multi `, (expandedHeight-1)*(expandedWidth-1))
	fmt.Println(`multi `, (height-1)*(width-1))
	fmt.Println(`len `, len(strMap))
	fmt.Println(`extend len `, len(expandedMap))
	for y := 0; y < expandedHeight; y++ {
		composedLine := ""
		for x := 0; x < expandedWidth; x++ {
			// fmt.Println(`x`, x, `y`, y)
			if expandedMap[x+y*expandedWidth] {
				composedLine += " X"
			} else {
				composedLine += " O"
			}
		}
		// fmt.Println(composedLine)
	}

	floodColor = make([]int, len(strMap))
	for i := range strMap {
		fmt.Println(i)
		floodColor[i] = i
	}
	fmt.Println(floodColor)
	fmt.Println(`hellowlrd`)

}
