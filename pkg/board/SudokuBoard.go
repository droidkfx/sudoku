package board

import "fmt"

type SudokuBoard struct {
	board [9][9]int
}

func BoardFromNumbers(data [9][9]int) *SudokuBoard {
	board := [9][9]int{}
	for i, list := range data {
		for j, num := range list {
			board[i][j] = num
		}
	}
	return &SudokuBoard{board: board}
}

func (s *SudokuBoard) GetAt(x, y int) int {
	return s.board[y][x]
}

func (s *SudokuBoard) SetAt(x, y, value int) {
	s.board[y][x] = value
}

func (s *SudokuBoard) String() string {
	result := ""
	for _, list := range s.board {
		for _, num := range list {
			result += fmt.Sprintf("%d ", num)
		}
		result += "\n"
	}
	return result
}
