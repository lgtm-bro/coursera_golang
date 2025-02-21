package main

import (
	"fmt"
)

func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	return func(t float64) float64 {
		return (a/2)*(t*t) + vo*t + so
	}
}

func getValue(prompt string) float64 {
	var input float64

	fmt.Print(prompt)
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("error reading acceleration value", err)
	}

	return input
}

func main() {
	a := getValue("Please enter an acceleration value: ")
	vo := getValue("Please enter an initial velocity value: ")
	so := getValue("Please enter an initial displacement value: ")

	fn := GenDisplaceFn(a, vo, so)

	t := getValue("Please enter a time value: ")

	displacement := fn(t)

	fmt.Print(displacement)
}
