package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Please enter your name:")
	scanner.Scan()
	name := scanner.Text()

	fmt.Print("Please enter your street address: ")
	scanner.Scan()
	address := scanner.Text()

	users := map[string]string{
		"name":    name,
		"address": address,
	}

	data, jsonErr := json.Marshal(users)
	if jsonErr != nil {
		fmt.Println("error formatting json: ", jsonErr)
	}

	fmt.Println(string(data))
}
