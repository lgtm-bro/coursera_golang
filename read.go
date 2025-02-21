package main

import (
	"bufio"
	"fmt"
	"os"
	s "strings"
)

func main() {
	var filename string
	type Name struct {
		fname string
		lname string
	}
	var names []Name

	fmt.Print("Please enter the file path: ")
	_, filenameErr := fmt.Scan(&filename)
	if filenameErr != nil {
		fmt.Println("error reading filename", filenameErr)
	}

	file, fileErr := os.Open(filename)
	if fileErr != nil {
		fmt.Println("error reading the file", fileErr)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		words := s.Fields(line)
		name := Name{fname: words[0], lname: words[1]}
		names = append(names, name)
		fmt.Print()
	}

	for _, name := range names {
		fmt.Println(name.fname, name.lname)
	}
}
