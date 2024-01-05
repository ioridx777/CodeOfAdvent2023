package main

import (
	"fmt"
	// "os"
	// "regexp"
	// "strconv"
	// "strings"
)

func doHash(strInput string) int {
	// for i := 0; i < len(strInput); i++ {
	// 	currentChar := strInput[i : i+1]
	//    strconv.Atoi()
	// }
	sum := 0
	for _, val := range strInput {
		addVal := (sum + val) * 17
		remainderVal := addVal % 256
		sum += int(remainderVal)
		fmt.Println(sum)
	}
	return 1
}
func main() {
	// inputBytes, error := os.ReadFile("inputDay15.txt")

	// if error != nil {
	// 	fmt.Println("fail")
	// 	return
	// }
	// inputText := string(inputBytes)

	// linesOfText := strings.Split(inputText, "\n")
	// for _, line := range linesOfText {
	// 	fmt.Println(line)
	// 	if len(line) == 0 {
	// 		continue
	// 	}

	// }
	inputStr := "HASH"
	hashVal := doHash(inputStr)
	fmt.Println("helloworld ", hashVal)
}
