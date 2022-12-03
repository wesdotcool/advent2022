package main

// https://adventofcode.com/2022/

import (
	"advent2022wesdotcool/dec1"
	"advent2022wesdotcool/sample"
	"fmt"
	"log"
	"os"
)

var m = map[string]func(bool) error{
	"sample": sample.Run,
	"dec1":   dec1.Run,
}

// Expects command line arguments to select day to run
// go run advent.go dec1
func main() {
	arg := os.Args[1]
	test := false
	if len(os.Args) > 2 {
		if os.Args[2] == "true" || os.Args[2] == "test" {
			test = true
		}
	}

	function, ok := m[arg]
	if !ok {
		fmt.Println("I don't understand argument", arg)
		return
	}
	fmt.Println("======================================================")
	fmt.Println("======================================================")
	fmt.Printf("Advent of Code 2022! Running code for %v\n", arg)
	fmt.Println("======================================================")
	fmt.Println("======================================================")

	err := function(test)
	if err != nil {
		log.Fatal(err)
	}
}
