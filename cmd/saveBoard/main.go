package main

import (
	"droidkfx.com/sudoku/pkg/board"
	"droidkfx.com/sudoku/pkg/repository"
)

func main() {
	r, sd := repository.NewSudokuBoardRepo("./data/nyt/med")
	defer sd()

	b := board.FromNumbers([9][9]int{
		{0, 0, 0, 4, 0, 1, 6, 0, 0},
		{0, 0, 2, 5, 0, 0, 0, 0, 3},
		{5, 3, 0, 9, 0, 0, 0, 0, 0},
		{0, 0, 6, 0, 0, 0, 3, 0, 0},
		{0, 0, 0, 0, 8, 0, 0, 4, 1},
		{4, 0, 5, 0, 0, 0, 0, 0, 7},
		{2, 5, 8, 0, 0, 9, 0, 0, 0},
		{0, 0, 9, 0, 4, 2, 0, 0, 0},
		{0, 0, 0, 7, 0, 0, 0, 0, 0},
	})

	r.SaveNew(b)
}
