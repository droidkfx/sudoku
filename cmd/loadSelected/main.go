package main

import (
	"fmt"

	"droidkfx.com/sudoku/pkg/repository"
	"droidkfx.com/sudoku/pkg/solver"
)

func main() {
	r, sd := repository.NewSudokuBoardRepo("./data/nyt/med")
	defer sd()

	b := r.GetNumber(0)
	fmt.Println(b)
	solution := solver.SolveByStrategies(b)
	for _, step := range solution {
		fmt.Println(step)
	}
}
