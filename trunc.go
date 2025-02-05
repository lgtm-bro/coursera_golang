package main

import (
	"fmt"
	"math"
)

func main() {
	var userInp float64
	fmt.Print("Enter a number with a decimal: ")
	if _, err := fmt.Scanf("%f", &userInp); err != nil {
		fmt.Println("error truncating number:", err)
	} else {
		truncInt := math.Trunc(userInp)
		fmt.Println(truncInt)
	}
}
