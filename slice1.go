package main

import (
	"fmt"
	"sort"
	"strconv"
	s "strings"
)

func main() {
	var sortedSlice = make([]int, 3)
	var originalSlice = make([]int, 3)
	var input string
	var count int = 0

	fmt.Print("Please enter some numbers. Type 'x' when you are done.")

	for {
		fmt.Print("Enter a number: ")
		fmt.Scan(&input)
		if s.ToLower(input) != "x" {
			n, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Invalid input, please enter a valid number.", err)
				continue
			}
			count += 1
			if count < 4 {
				sortedSlice = append([]int{}, originalSlice...)
				sortedSlice[count-1] = n
				originalSlice[count-1] = n
			} else {
				sortedSlice = append(sortedSlice, n)
			}
			sort.Ints(sortedSlice)
			fmt.Println("Here is the current list of sorted numbers:")
			fmt.Println(sortedSlice)
		} else {
			fmt.Println("Here is the final list of sorted numbers:")
			fmt.Println(sortedSlice)
			break
		}

	}
}
