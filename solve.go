package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var dictionary []string

func init() {
	const dictFilename = "sowpods.txt"
	f, err := os.Open(dictFilename)
	if err != nil {
		fatal(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		word := scanner.Text()
		if len(word) == 0 {
			continue
		}
		dictionary = append(dictionary, strings.ToLower(word))
	}
	if err := scanner.Err(); err != nil {
		fatal(err)
	}
	fmt.Printf("Dictionary loaded (%d words).\n", len(dictionary))
}

type Word []Tile

// TODO: double-check this logic is how the game actually works.
func (w Word) Value() int {
	v := 0
	mul := 1
	for _, tile := range w {
		v += tile.Value
		mul *= tile.WordMul
	}
	return v * mul
}

func SolveNaive(tiles []Tile) *Solution {
	// Use a very simplistic heuristic instead of finding the absolute best solution:
	// * Find (one of) the highest-value words we can make using given letters
	// * Remove those letters from the set, and repeat
	best, remaining := findHighest(tiles)
	secondBest, _ := findHighest(remaining)
	return &Solution{
		Word1: best,
		Word2: secondBest,
		Total: best.Value + secondBest.Value,
	}
}

// findHighest goes through the dictionary and finds the highest-valued word that is made with these tiles. It
// returns the highest result and the remaining tiles.
// Note that it may not actually return the best result because it could use one tile (say, 'A') where another
// A tile with a better letter/word multiplier would've been worth more points.
func findHighest(tiles []Tile) (highest WordAndValue, remaining []Tile) {
	highest = WordAndValue{"", 0}
	var highestRemaining []Tile
wordLoop:
	for _, word := range dictionary {
		if len(word) > len(tiles) {
			continue
		}
		remaining = make([]Tile, len(tiles))
		copy(remaining, tiles)
		wordTiles := Word{}
	letterLoop:
		for _, letter := range word {
			for i, tile := range remaining {
				if tile.Letter == byte(letter) {
					wordTiles = append(wordTiles, tile)
					// Remove tile from remaining
					remaining = append(remaining[:i], remaining[i+1:]...)
					continue letterLoop
				}
			}
			// Can't form this word
			continue wordLoop
		}
		// We are able to make the word
		value := wordTiles.Value()
		if value > highest.Value {
			highest = WordAndValue{word, value}
			highestRemaining = remaining
		}
	}
	return highest, highestRemaining
}
