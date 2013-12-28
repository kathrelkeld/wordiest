package main

import (
	"bytes"
	"fmt"
)

func Parse(line []byte) ([]Tile, error) {
	tiles := []Tile{}
	for _, field := range bytes.Fields(line) {
		if len(field) == 0 {
			continue // Don't think this can happen
		}
		letter, err := normalize(field[0])
		if err != nil {
			return nil, err
		}
		tile := Tile{
			Letter:  letter,
			Value:   defaultValues[letter],
			WordMul: 1,
		}
		badTileErr := fmt.Errorf("Bad tile: %s", string(field))
		switch len(field) {
		case 1:
			tiles = append(tiles, tile)
			continue
		case 3: // below
		default:
			return nil, badTileErr
		}
		mulCh := field[1]
		if mulCh < '2' || mulCh > '9' {
			return nil, badTileErr
		}
		mul := int(mulCh - '0')
		switch field[2] {
		case 'l', 'L':
			tile.Value *= mul
		case 'w', 'W':
			tile.WordMul = mul
		default:
			return nil, badTileErr
		}
		tiles = append(tiles, tile)
	}
	if len(tiles) == 0 {
		return nil, fmt.Errorf("Empty input line")
	}
	return tiles, nil
}

// normalize changes to lowercase, and returns an error if the byte isn't an English alphabet letter.
func normalize(letter byte) (byte, error) {
	switch {
	case letter >= 'a' && letter <= 'z':
		return letter, nil
	case letter >= 'A' && letter <= 'Z':
		return (letter + 'a' - 'A'), nil
	}
	return 0, fmt.Errorf("Non-letter input: %c", letter)
}
