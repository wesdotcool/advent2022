package main

// https://adventofcode.com/2022/

import (
	"advent2022wesdotcool/dec1"
	"advent2022wesdotcool/dec10"
	"advent2022wesdotcool/dec2"
	"advent2022wesdotcool/dec3"
	"advent2022wesdotcool/dec4"
	"advent2022wesdotcool/dec5"
	"advent2022wesdotcool/dec6"
	"advent2022wesdotcool/dec7"
	"advent2022wesdotcool/dec8"
	"advent2022wesdotcool/dec9"
	"advent2022wesdotcool/sample"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type info struct {
	function  func(io.Reader) (any, error)
	directory string
}

var m = map[string]info{
	"sample":  {sample.Run, "sample/"},
	"dec1":    {dec1.Run, "dec1/"},
	"dec1.2":  {dec1.Run2, "dec1/"},
	"dec2":    {dec2.Run, "dec2/"},
	"dec2.2":  {dec2.Run2, "dec2/"},
	"dec3":    {dec3.Run, "dec3/"},
	"dec3.2":  {dec3.Run2, "dec3/"},
	"dec4":    {dec4.Run, "dec4/"},
	"dec4.2":  {dec4.Run2, "dec4/"},
	"dec5":    {dec5.Run, "dec5/"},
	"dec5.2":  {dec5.Run2, "dec5/"},
	"dec6":    {dec6.Run, "dec6/"},
	"dec6.2":  {dec6.Run2, "dec6/"},
	"dec7":    {dec7.Run, "dec7/"},
	"dec7.2":  {dec7.Run2, "dec7/"},
	"dec8":    {dec8.Run, "dec8/"},
	"dec8.2":  {dec8.Run2, "dec8/"},
	"dec9":    {dec9.Run, "dec9/"},
	"dec9.2":  {dec9.Run2, "dec9/"},
	"dec10":   {dec10.Run, "dec10/"},
	"dec10.2": {dec10.Run2, "dec10/"},
}

// Expects command line arguments to select day to run
// go run advent.go dec1
func main() {
	if len(os.Args) == 1 {
		fmt.Println("Must include problem to run. Such as dec1")
	}
	arg := os.Args[1]
	filename := m[arg].directory + "input"
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

	fmt.Println("======================================================")
	fmt.Println("======================================================")
	fmt.Printf("Advent of Code 2022! Running code for %v with file %v\n", arg, filename)
	fmt.Println("======================================================")
	fmt.Println("======================================================")

	startTime := time.Now()
	answer, err := m[arg].function(input)
	endTime := time.Now()
	fmt.Println("Runtime:", endTime.Sub(startTime))
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(answer)
	}
}
