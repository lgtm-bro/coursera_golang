package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter a word: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error reading input: ", err)
	} else {
		lowerWord := strings.TrimSpace(strings.Replace(strings.ToLower(input), " ", "", -1))
		fmt.Println(lowerWord)

		if strings.HasPrefix(lowerWord, "i") &&
			strings.HasSuffix(lowerWord, "n") &&
			strings.Contains(lowerWord, "a") {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}
	}
}
