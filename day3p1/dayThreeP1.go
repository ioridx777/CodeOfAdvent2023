package main

import (
	"fmt"
	"os"
	"strconv"

	// "regexp"
	// "strconv"
	"strings"
)

var hitmap = []bool{}

func isNumberChar(val string) bool {
	return (val >= "0" && val <= "9")
}
func getCoord(x int, y int, width int) int {
	return x + y*width
}
func updateHit(x int, y int, width int, height int, hitmapInput []bool) {
	isLeft := x-1 >= 0
	isRight := x+1 < width
	isUp := y-1 >= 0
	isDown := y+1 < height
	// fmt.Println(`isLeft `, isLeft)
	// fmt.Println(`isRight `, isRight)
	// fmt.Println(`isUp `, isUp)
	// fmt.Println(`isDown `, isDown)
	coord := getCoord(x, y, width)
	// fmt.Println(`starting `, coord)
	// fmt.Println(`width `, width)
	xAxis := []bool{isLeft, true, isRight}
	yAxis := []bool{isUp, true, isDown}
	for i, xCoor := range xAxis {

		// fmt.Println(i, ": ", xCoor)
		if !xCoor {
			continue
		}
		for j, yCoor := range yAxis {
			// fmt.Println(j, ": ", yCoor)
			if i == 1 && j == 1 {
				continue
			}
			if !yCoor {
				continue
			}
			markingCoor := coord + i - 1 + (j-1)*width
			// fmt.Println(`working with `, markingCoor)
			hitmapInput[markingCoor] = true
		}
	}
}

func mainOld() {
	inputBytes, error := os.ReadFile("inputDay3.txt")

	if error != nil {
		fmt.Println("fail")
		return
	}
	inputText := string(inputBytes)
	linesOfText := strings.Split(inputText, "\n")
	width := len(linesOfText[0])
	height := len(linesOfText)
	fmt.Println(`width `, width)
	fmt.Println(`height `, height)
	hitmap = make([]bool, width*height)

	for j, line := range linesOfText {
		fmt.Println(line)
		if len(line) == 0 {
			continue
		}
		for i := 0; i < len(line); i++ {
			if line[i:i+1] == "." || isNumberChar(line[i:i+1]) {
				continue
			}
			updateHit(i, j, width, height, hitmap)
		}
	}
	for i, val := range hitmap {
		if val {
			fmt.Print("Y")
		} else {
			fmt.Print(".")
		}
		if i%width == width-1 {
			fmt.Println()
		}
	}
	fmt.Println()
	// for i, val := range hitmap {
	// 	j = i / width
	//
	// }
	sum := 0
	for i := 0; i < len(hitmap); i++ {
		x := i % width
		y := i / width
		if len(linesOfText[y]) == 0 {
			break
		}
		targetChar := linesOfText[y][x : x+1]
		if isNumberChar(targetChar) {
			bufferNumberString := targetChar
			i++
			x := i % width
			targetChar = linesOfText[y][x : x+1]
			for isNumberChar(targetChar) && x < width {
				bufferNumberString += targetChar
				i++
				x := i % width
				targetChar = linesOfText[y][x : x+1]
			}
			numberLen := len(bufferNumberString)
			isHit := false

			fmt.Print(`checking `)
			for c := 0; c < numberLen; c++ {
				fmt.Print(hitmap[i-1-c], " ")
				if hitmap[i-c-1] {
					isHit = true
					break
				}
			}
			fmt.Println()
			if isHit {
				fmt.Println(`buffer number`, bufferNumberString)
				fmt.Println(`val of i `, x)
				bufferNumber, _ := strconv.Atoi(bufferNumberString)
				sum += bufferNumber
			} else {
				fmt.Println(`miss number`, bufferNumberString)
			}
		}
	}
	fmt.Println(`hellowlrd `, sum)
}
