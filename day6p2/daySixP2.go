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

var matchTimes = []int{}
var matchDistance = []int{}

func getDistance(charge int, total int) int {
	return charge * (total - charge)
}

func main() {
	inputBytes, error := os.ReadFile("inputDay6.txt")

	if error != nil {
		fmt.Println("fail")
		return
	}
	inputText := string(inputBytes)

	linesOfText := strings.Split(inputText, "\n")
	noSpaceTime := strings.ReplaceAll(linesOfText[0], " ", "")
	realTimeStr := strings.ReplaceAll(noSpaceTime, "\t", "")
	noSpaceDist := strings.ReplaceAll(linesOfText[1], " ", "")
	realDistStr := strings.ReplaceAll(noSpaceDist, "\t", "")
	numberReg, _ := regexp.Compile("[0-9]+")
	numberTimeStr := numberReg.FindAllString(realTimeStr, -1)
	numberDistStr := numberReg.FindAllString(realDistStr, -1)

	fmt.Println(`A `, numberTimeStr)
	fmt.Println(`B `, numberDistStr)

	sum := 1
	for i, timeStr := range numberTimeStr {
		time, _ := strconv.Atoi(timeStr)
		dist, _ := strconv.Atoi(numberDistStr[i])
		hitCount := 0
		for j := 1; j < time; j++ {
			newDist := getDistance(j, time)
			// fmt.Println(j, " seconds, ", time-j, " m/s, ", newDist, " / ", dist, " : ", (newDist > dist))
			if newDist > dist {
				hitCount++
			}
		}
		fmt.Println("race ", i, ": ", hitCount)
		sum *= hitCount
	}

	fmt.Println(`hellowlrd `, sum)
}
