package main

import (
	"fmt"

	"droidkfx.com/sudoku/pkg/repository"
)

func main() {
	r, sd := repository.NewSudokuBoardRepo("./data")
	defer sd()

	fmt.Println(r.GetNumber(0))
	fmt.Println(r.GetNumber(900))
	fmt.Println(r.GetNumber(15000))
	fmt.Println(r.GetNumber(50000))
}
