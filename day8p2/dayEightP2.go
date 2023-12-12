package main

import (
	"fmt"
	"os"

	// "regexp"
	// "strconv"
	"strings"
)

var locMap = map[string][2]string{}
var instruct = []int{}

type Path struct {
	hitStep       []int
	startLoopStep int
	endLoopStep   int
}

func getHitAt(path Path, counter int) int {
	if counter < len(path.hitStep) {
		return path.hitStep[counter]
	} else {
		loopCount := counter / len(path.hitStep)
		return (path.hitStep[counter%len(path.hitStep)]) + (path.endLoopStep-path.startLoopStep)*loopCount
	}
}

func main() {
	inputBytes, error := os.ReadFile("inputDay8.txt")

	if error != nil {
		fmt.Println("fail")
		return
	}
	inputText := string(inputBytes)

	linesOfText := strings.Split(inputText, "\n")
	for i := 0; i < len(linesOfText[0]); i++ {
		fmt.Println(`destructing `, linesOfText[0][i:i+1])
		nextInstruct := (linesOfText[0][i : i+1])
		nextInstructNum := 0
		if nextInstruct == "L" {
			nextInstructNum = 0
		} else if nextInstruct == "R" {
			nextInstructNum = 1
		} else {
			continue
		}
		instruct = append(instruct, nextInstructNum)
	}
	fmt.Println(`final instruct `, instruct)
	for lineIndex, line := range linesOfText {
		if lineIndex < 2 {
			continue
		}
		if len(line) == 0 {
			continue
		}
		// fmt.Println(line)
		if len(line) == 0 {
			continue
		}
		locMap[line[0:3]] = [2]string{line[7:10], line[12:15]}
	}
	startingPoss := []string{}
	for pos := range locMap {
		if pos[len(pos)-1:] == "A" {
			startingPoss = append(startingPoss, pos)
		}
	}
	// startingPoss = startingPoss[0:1]
	// startingPoss = []string{"AAA"}
	fmt.Println("start with ", startingPoss)
	pathList := []Path{}
	for _, startingPos := range startingPoss {
		prevPos := startingPos
		endedPosMap := map[string]int{}
		hitStep := []int{}
		step := 1
		loopStep := 0
		loopEndStep := 0
		for {
			for _, ins := range instruct {
				newPos := locMap[prevPos][ins]
				prevPos = newPos
				// if step == 270 {
				//   fmt.Println(270, ": ", prevPos)
				// }
				// if step == 269 {
				//   fmt.Println(269, ": ", prevPos)
				// }
				// if step == 12913 {
				//   fmt.Println(12913, ": ", prevPos)
				//   return
				// }
				// if step == 12912 {
				//   fmt.Println(12912, ": ", prevPos)
				// }
				if prevPos[len(prevPos)-1:] == "Z" {
					// fmt.Println(`appending `, prevPos)
					hitStep = append(hitStep, step)
				}
				step++
			}
			// if prevPos == "KQK" {
			//   fmt.Println("T", step)
			// }
			// fmt.Println(`end walk in`, prevPos)
			// fmt.Println(step - 1)
			// break
			// if step == 12463+1 {
			//   fmt.Println(`goal `, prevPos)
			//   break
			// }
			if endedPosMap[prevPos] > 0 {
				loopStep = endedPosMap[prevPos]
				loopEndStep = step - 1
				fmt.Println(`step `, step-1)
				fmt.Println(`Loop to `, prevPos)
				fmt.Println(`backto `, endedPosMap[prevPos])
				break
			}
			endedPosMap[prevPos] = step - 1
		}
		fmt.Println(`start with `, startingPos)
		fmt.Println(loopEndStep - loopStep)
		fmt.Println(loopStep)
		fmt.Println(loopEndStep)
		fmt.Println(hitStep)
		pathList = append(pathList, Path{
			hitStep:       hitStep,
			startLoopStep: loopStep,
			endLoopStep:   loopEndStep,
		})
	}
	fmt.Println(pathList)
	// return
	counters := make([]int, len(pathList))
	// for s := 0; s < 10; s++ {

	// for _, p := range pathList {
	// 	loopStep := p.endLoopStep - p.startLoopStep
	// 	fmt.Println(loopStep)
	// }

	return
	for {
		compareVals := make([]int, len(pathList))
		for i := 0; i < len(pathList); i++ {
			compareVals[i] = getHitAt(pathList[i], counters[i])
		}
		startVal := compareVals[0]
		minVal := compareVals[0]
		minIndex := 0
		fail := false
		for i := 1; i < len(pathList); i++ {
			// fmt.Println(compareVals[0], ` vs `, minVal)
			if compareVals[i] < minVal {
				minVal = compareVals[i]
				minIndex = i
			}
			if compareVals[i] != startVal {
				fail = true
			}
		}
		fmt.Println(`counters `, counters)
		fmt.Println(`compare value `, compareVals)
		if !fail {
			fmt.Println(`YOu got it `, startVal)
			break
		}
		// fmt.Println(`pumping `, minIndex, " : ", minVal)
		counters[minIndex] = counters[minIndex] + 1
	}
}
