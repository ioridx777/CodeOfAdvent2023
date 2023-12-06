package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"

	"strings"
)


func sumOfPowers(input int) int {
	// sum := 0.0
	// for i := 0; i < input; i++ {
	// 	sum += math.Pow(float64(2), float64(i))
	// }
	if input == 0 {
		return 0
	}
	result := math.Pow(2.0, float64(input-1))
	return int(result)
}
func processLine(line string) int {
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
	return sumOfPowers(hitCount)
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
		pResult := processLine(line)
		sum += pResult
		fmt.Println(index, " process ", pResult)
	}
	fmt.Println(sum)
	fmt.Println(`hellowlrd`)
}
