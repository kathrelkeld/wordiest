package main

import (
	"fmt"
	"os"

	"github.com/cespare/argf"
)

type Tile struct {
	Letter  byte // a-z
	Value   int  // 1+
	WordMul int  // 1+
}

type WordAndValue struct {
	Word  string
	Value int
}

func (w WordAndValue) String() string {
	return fmt.Sprintf("%s (%d)", w.Word, w.Value)
}

type Solution struct {
	Word1 WordAndValue
	Word2 WordAndValue
	Total int
}

func (s Solution) String() string {
	return fmt.Sprintf("%d: %s %s", s.Total, s.Word1, s.Word2)
}

func main() {
	for argf.Scan() {
		line := argf.Bytes()
		tiles, err := Parse(line)
		if err != nil {
			fatal(err)
		}
		bestNaive := SolveNaive(tiles)
		fmt.Println("naive:", bestNaive)
	}
	if err := argf.Error(); err != nil {
		fatal(err)
	}
}

func fatal(args ...interface{}) {
	fmt.Println(args...)
	os.Exit(-1)
}
