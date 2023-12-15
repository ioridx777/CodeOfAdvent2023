package main

import (
	"fmt"
	"os"
	// "regexp"
	// "strconv"
	"strings"
)

func tempMain() {
	inputBytes, error := os.ReadFile("inputDay2.txt")

	if error != nil {
		fmt.Println("fail")
		return
	}
	inputText := string(inputBytes)

	linesOfText := strings.Split(inputText, "\n")
	for _, line := range linesOfText {
		fmt.Println(line)
		if len(line) == 0 {
			continue
		}

	}
	fmt.Println(`hellowlrd`)
}
