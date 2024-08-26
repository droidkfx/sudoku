package board

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
