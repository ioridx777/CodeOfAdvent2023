package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"strings"
)

var cardPoint = map[int]int{}
var cardConnect = map[int][]int{}

func getCardPoint(cardId int) int {
	if cardPoint[cardId] != 0 {
		return cardPoint[cardId]
	} else if len(cardConnect[cardId]) == 0 {
		cardPoint[cardId] = 1
		return 1
	} else {
		sum := 1
		for _, childId := range cardConnect[cardId] {
			sum += getCardPoint(childId)
		}
		cardPoint[cardId] = sum
		return sum
	}

}

func sumOfPowers(input int, startCard int) {
	cardConnect[startCard] = []int{}
	for i := 1; i <= input; i++ {
		cardConnect[startCard] = append(cardConnect[startCard], startCard+i)
	}
}
func processLine(line string, index int) {
	hitCount := 0
	headerReg, _ := regexp.Compile(`Card[ |\t]+[0-9]+:`)
	removeHeaderLine := headerReg.ReplaceAllString(line, "")
	splitedLine := strings.Split(removeHeaderLine, "|")
	fmt.Println(`A `, splitedLine[0])
	fmt.Println(`B `, splitedLine[1])
	numberReg, _ := regexp.Compile(`[0-9]+`)
	aNumberStrings := numberReg.FindAllString(splitedLine[0], -1)
	aNumArray := []int{}
	for _, val := range aNumberStrings {
		aNum, _ := strconv.Atoi(val)
		aNumArray = append(aNumArray, aNum)
	}
	bNumberStrings := numberReg.FindAllString(splitedLine[1], -1)
	// bNumArray := []int{}
	for _, valB := range bNumberStrings {
		bNum, _ := strconv.Atoi(valB)
		for _, aNum := range aNumArray {
			if aNum == bNum {
				fmt.Println(`hitting on `, aNum)
				hitCount++
				break
			}
		}
		// bNumArray = append(bNumArray, bNum)
	}
	sumOfPowers(hitCount, index)
}
func main() {
	inputBytes, error := os.ReadFile("inputDay4.txt")

	if error != nil {
		fmt.Println("fail")
		return
	}
	inputText := string(inputBytes)

	linesOfText := strings.Split(inputText, "\n")
	sum := 0
	for index, line := range linesOfText {
		fmt.Println(line)
		if len(line) == 0 {
			continue
		}
		processLine(line, index)
	}
	// sum += len(linesOfText)
	for cardId := 0; cardId < len(linesOfText)-1; cardId++ {
		fmt.Println(`card `, cardId)
		sum += getCardPoint(cardId)
	}
	fmt.Println(`card connect `, cardConnect)
	fmt.Println(`card point `, cardPoint)
	fmt.Println(sum)
	fmt.Println(`hellowlrd`)
}
