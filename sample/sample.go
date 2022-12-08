package sample

import (
	"fmt"
	"io"
	"sort"

	"golang.org/x/exp/slices"
)

// for testing purposes
// go run advent.go sample
func Run(input io.Reader) (any, error) {
	runes := []rune("hello")
	fmt.Println(runes)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	fmt.Println(runes)
	fmt.Println(slices.Compact(runes))
	return nil, nil
}
