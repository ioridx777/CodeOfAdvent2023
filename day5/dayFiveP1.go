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
	seedLine := linesOfText[0]
	numReg, _ := regexp.Compile("[0-9]+")
	seedStrs := numReg.FindAllString(seedLine, -1)
	seedNums := []int{}
	for _, val := range seedStrs {
		numVal, _ := strconv.Atoi(val)
		seedNums = append(seedNums, numVal)
	}
	fmt.Println(`the seed`, seedNums)
	headerReg, _ := regexp.Compile(".*[:]")
	index := 0
	for i, line := range linesOfText {
		// fmt.Println(line)
		if i == 0 {
			continue
		}
		if len(line) == 0 {
			continue
		}
		if len(headerReg.FindString(line)) > 0 {
			fmt.Println("header line ", index, "\n", line)
			index++
		}
	}
	// fmt.Println(`hellowlrd`, headerIndex)
	fmt.Println(`hellowlrd`)
}
