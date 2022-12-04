package sample

import (
	"fmt"
	"io"
	"log"
)

// for testing purposes
// go run advent.go sample
func Run(input io.Reader) (any, error) {
	output := make([]byte, 0)
	n, err := input.Read(output)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)

	fmt.Println(output)
	input.Read(output)
	fmt.Println(output)
	input.Read(output)
	fmt.Println(output)
	return nil, nil
}
