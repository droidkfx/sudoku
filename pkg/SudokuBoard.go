package pkg

type SudokuBoard struct {
	board [9][9]int
}

func (s *SudokuBoard) GetAt(x, y int) int {
	return s.board[y][x]
}

func (s *SudokuBoard) SetAt(x, y, value int) {
	s.board[y][x] = value
}
