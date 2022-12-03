package sample

import (
	"fmt"
	"strconv"
)

// for testing purposes
// go run advent.go sample
func Run(test bool) error {
	x := "1000"
	xint, err := strconv.Atoi(x)
	if err != nil {
		return err
	}
	fmt.Println(xint + 10)
	return nil
}
