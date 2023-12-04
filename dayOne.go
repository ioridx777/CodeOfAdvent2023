package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var hello string

var digitEng [9]string = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func processDigit(digitStr string) int {
	if len(digitStr) == 1 {
		digit, _ := strconv.Atoi(digitStr)
		return digit
	}
	for digit, digitWord := range digitEng {
		if digitWord == digitStr {
			return digit + 1
		}
	}
	return 0
}

func processLine(line string) int {
	//r, error := regexp.MatchString(`[0-9]+`, line)
	regexInput := "[0-9]"
	for _, d := range digitEng {
		regexInput += "|(" + d + ")"
	}
	// fmt.Println(`composed `, line)
	r, _ := regexp.Compile(regexInput)
	matchDigitIndex := r.FindStringIndex(line)
	firstDigit := processDigit(line[matchDigitIndex[0]:matchDigitIndex[1]])
	if len(matchDigitIndex) == 0 {
		return 0
	}
	previousMatch := matchDigitIndex
	newMatch := r.FindStringIndex(line[previousMatch[0]+1:])
	for len(newMatch) != 0 {
		previousMatch = []int{newMatch[0] + previousMatch[0] + 1, newMatch[1] + previousMatch[0] + 1}
		newMatch = r.FindStringIndex(line[previousMatch[0]+1:])
	}
	lastDigit := processDigit(line[previousMatch[0]:previousMatch[1]])
	return firstDigit*10 + lastDigit
}

func main() {
	inputBytes, error := os.ReadFile("input1.txt")
	if error != nil {
		fmt.Println("fail")
		return
	}
	inputText := string(inputBytes)

	linesOfText := strings.Split(inputText, "\n")
	sum := 0
	for index, line := range linesOfText {
		// fmt.Println("A")
		if len(line) == 0 {
			break
		}
		// fmt.Println("line ", index)
		// fmt.Println("value ", line)
		// fmt.Println("current sum ", sum)
		sum += processLine(line)
		fmt.Println(index, ".", line, ": ", processLine(line), " = ", sum)
	}
	fmt.Println("Sum ", sum)
}
