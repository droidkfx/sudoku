package solver

import "droidkfx.com/sudoku/pkg/board"

func propagateNumberSetToOptions(opts *[9][9][9]bool, x, y, value int) {
	for i := 0; i < 9; i++ {
		opts[x][i][value-1] = false
		opts[i][y][value-1] = false
		opts[x][y][i] = false
	}

	regionX := x / 3
	regionY := y / 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			opts[regionX*3+i][regionY*3+j][value-1] = false
		}
	}
}

func GetPossibleValues(board *board.SudokuBoard) [9][9][9]bool {
	possibleValues := [9][9][9]bool{}
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if board.GetAt(x, y) != 0 {
				continue
			} else {
				valuesSeen := getIntersectingValues(board, x, y)

				for i := 0; i < 9; i++ {
					possibleValues[x][y][i] = !valuesSeen[i]
				}
			}
		}
	}
	return possibleValues
}

func getIntersectingValues(board *board.SudokuBoard, x int, y int) [9]bool {
	valuesSeen := [9]bool{}
	for i := 0; i < 9; i++ {
		i2 := board.GetAt(x, i) - 1
		if i2 != -1 {
			valuesSeen[i2] = true
		}

		i3 := board.GetAt(i, y) - 1
		if i3 != -1 {
			valuesSeen[i3] = true
		}
	}

	regionX := x / 3
	regionY := y / 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			i2 := board.GetAt(regionX*3+i, regionY*3+j) - 1
			if i2 != -1 {
				valuesSeen[i2] = true
			}
		}
	}
	return valuesSeen
}
