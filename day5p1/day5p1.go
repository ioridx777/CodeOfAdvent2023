package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	// "regexp"
	// "strconv"
	"strings"
)

func main() {
	inputBytes, error := os.ReadFile("sampleDay5.txt")

	if error != nil {
		fmt.Println("fail")
		return
	}
	inputText := string(inputBytes)

	linesOfText := strings.Split(inputText, "\n")
	numReg, _ := regexp.Compile("[0-9]+")

	vals := numReg.FindAllString(linesOfText[0], -1)
	seedNumbers := []int{}
	for _, val := range vals {
		fmt.Println(val)
		valNum, _ := strconv.Atoi(val)
		seedNumbers = append(seedNumbers, valNum)
	}
	counter := 0
	mapLineReg, _ := regexp.Compile(".*[:]")
	mapVals
	for _, line := range linesOfText {
		fmt.Println(line)
		if len(line) == 0 {
			continue
		}
		if mapLineReg.MatchString(line) {
			counter++
		}
	}

	fmt.Println(seedNumbers)
	fmt.Println(`hellowlrd`)
}
