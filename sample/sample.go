package sample

import (
	"fmt"
	"io"

	mapset "github.com/deckarep/golang-set/v2"
)

// for testing purposes
// go run advent.go sample
func Run(input io.Reader) (any, error) {
	itms := []items{items("HELLO"), items("hello")}
	s1 := mapset.NewSet(itms...)
	s2 := mapset.NewSet(items("hello"), items("hi"))

	fmt.Println(s1.Intersect(s2))
	return nil, nil
}

type items string

func (this items) split() (items, items) {
	return this[0 : len(this)/2], this[len(this)/2:]
}

func priority(b byte) int {
	var result int = int(b) - 96
	if result > 0 {
		return result
	}
	return result + 58
}
