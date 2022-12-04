package main

// https://adventofcode.com/2022/

import (
	"advent2022wesdotcool/dec1"
	"advent2022wesdotcool/sample"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var m = map[string]func(io.Reader) (any, error){
	"sample": sample.Run,
	"dec1":   dec1.Run,
}

// Expects command line arguments to select day to run
// go run advent.go dec1
func main() {
	if len(os.Args) == 1 {
		fmt.Println("Must include problem to run. Such as dec1")
	}
	arg := os.Args[1]
	filename := arg + "/input"
	if len(os.Args) > 2 {
		filename += "_" + os.Args[2]
	}
	filename += ".txt"
	input, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer input.Close()

	function, ok := m[arg]
	if !ok {
		fmt.Println("I don't understand argument", arg)
		return
	}
	fmt.Println("======================================================")
	fmt.Println("======================================================")
	fmt.Printf("Advent of Code 2022! Running code for %v with file %v\n", arg, filename)
	fmt.Println("======================================================")
	fmt.Println("======================================================")

	startTime := time.Now()
	answer, err := function(input)
	endTime := time.Now()
	fmt.Println("Runtime:", endTime.Sub(startTime))
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(answer)
	}
}
