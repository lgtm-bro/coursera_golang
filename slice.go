package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	sli := make([]int, 3)
	var number string
	for {
		fmt.Scanln(&number)
		if number == "x" {
			break
		}
		num, err := strconv.Atoi(number)
		if err != nil {
			fmt.Println("Wrong input")
			continue
		}
		sli = append(sli, num)
		sort.Ints(sli[:])
		fmt.Println(sli)
	}
}
