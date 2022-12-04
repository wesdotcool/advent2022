package sample

import (
	"fmt"
	"io"
	"strconv"
)

// for testing purposes
// go run advent.go sample
func Run(input io.Reader) (any, error) {
	x := "1000"
	xint, err := strconv.Atoi(x)
	if err != nil {
		return nil, err
	}
	fmt.Println(xint + 10)
	return nil, nil
}
