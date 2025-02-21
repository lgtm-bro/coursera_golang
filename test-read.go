package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	stdin_reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the file path: ")
	file_path, _ := stdin_reader.ReadString('\n')
	file_path = strings.TrimRight(file_path, "\n")

	names_file, _ := os.Open(file_path)
	defer names_file.Close()

	the_scanner := bufio.NewScanner(names_file)

	var names_structs []name

	for the_scanner.Scan() {
		names := strings.Fields(the_scanner.Text())
		the_name := name{fname: names[0], lname: names[1]}
		names_structs = append(names_structs, the_name)
	}

	for _, name_struct := range names_structs {
		fmt.Printf("first name: %s, last name: %s \n", name_struct.fname, name_struct.lname)
	}
}
