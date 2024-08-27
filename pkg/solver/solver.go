package solver

import (
	"fmt"

	"droidkfx.com/sudoku/pkg/board"
)

type SolveMetrics struct {
	tryCount   int
	resetCount int
}

type GuessSolverConfig struct {
	NumberOrder []int
}

func DefaultGuessConfig() GuessSolverConfig {
	return GuessSolverConfig{
		NumberOrder: []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
	}
}

func (s SolveMetrics) String() string {
	return fmt.Sprintf("Try Count: %d, Reset Count: %d", s.tryCount, s.resetCount)
}

func SolveByGuessing(cfg GuessSolverConfig, board *board.SudokuBoard) SolveMetrics {
	possibleValues := GetPossibleValues(board)
	metrics := SolveMetrics{}
	solveByGuessing(cfg, board, possibleValues, 0, 0, &metrics)
	return metrics
}

func solveByGuessing(cfg GuessSolverConfig, board *board.SudokuBoard, values [9][9][9]bool, x, y int,
	metrics *SolveMetrics) bool {
	if board.GetAt(x, y) == 0 {
		anyPossibleValues := false
		for _, i := range cfg.NumberOrder {
			if values[x][y][i] {
				anyPossibleValues = true
				if tryValue(cfg, board, values, x, y, i+1, metrics) {
					return true
				} else {
					board.SetAt(x, y, 0)
					metrics.resetCount++
				}
			}
		}
		if !anyPossibleValues {
			return false
		}
	}

	return board.GetAt(x, y) != 0
}

func tryValue(cfg GuessSolverConfig, board *board.SudokuBoard, values [9][9][9]bool, x, y, value int,
	metrics *SolveMetrics) bool {
	board.SetAt(x, y, value)
	metrics.tryCount++
	nextX := x + 1
	nextY := y
	if nextX == 9 {
		nextX = 0
		nextY++
		if nextY == 9 {
			return true
		}
	}

	for i := 0; i < 9; i++ {
		values[x][i][value-1] = false
		values[i][y][value-1] = false
	}

	regionX := x / 3
	regionY := y / 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			values[regionX*3+i][regionY*3+j][value-1] = false
		}
	}

	return solveByGuessing(cfg, board, values, nextX, nextY, metrics)
}

func untryValue(board *board.SudokuBoard, values [9][9][9]bool, x, y int) {
	value := board.GetAt(x, y)

	for i := 0; i < 9; i++ {
		valuesSeen := getIntersectingValues(board, x, i)
		if !valuesSeen[value-1] {
			values[x][i][value-1] = true
		}
	}

	for i := 0; i < 9; i++ {
		valuesSeen := getIntersectingValues(board, i, y)
		if !valuesSeen[value-1] {
			values[i][y][value-1] = true
		}
	}

	regionX := x / 3
	regionY := y / 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			valuesSeen := getIntersectingValues(board, regionX*3+i, regionY*3+j)
			if !valuesSeen[value-1] {
				values[regionX*3+i][regionY*3+j][value-1] = true
			}
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
