package main

import (
	"errors"
)

func cells_generator(data []byte) ([cellNumber][cellNumber]bool, error) {
	cells := [cellNumber][cellNumber]bool{}
	if len(data) == 2550 {
		i := 0
		for y := 0; y < 50; y++ {
			for x := 0; x <= 50; x++ {
				if x == 50 {
					if data[i] != 10 {
						return cells, errors.New("lines must end with a \\n")
					}
				} else if data[i] == 10 {
					return cells, errors.New("a cell cannot be represented by a \\n")
				} else if data[i] != 46 {
					cells[y][x] = true
				}
				i++
			}
		}
		return cells, nil
	}
	return cells, errors.New("the pattern does not contain the right number of cells")
}
