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
	NumberOrder GuessOrderProvider
}

type GuessOrderProvider func(i, j, v int) int

func NewStaticOrderGuesser(order []int) GuessOrderProvider {
	return func(i, j, v int) int {
		return order[v]
	}
}

func DefaultGuessConfig() GuessSolverConfig {
	return GuessSolverConfig{
		NumberOrder: NewStaticOrderGuesser([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}),
	}
}

func GuessConfig(provider GuessOrderProvider) GuessSolverConfig {
	return GuessSolverConfig{
		NumberOrder: provider,
	}
}

func (s SolveMetrics) String() string {
	return fmt.Sprintf("Try Count: %d, Reset Count: %d", s.tryCount, s.resetCount)
}

func SolveByGuessing(cfg GuessSolverConfig, board *board.SudokuBoard) SolveMetrics {
	possibleValues := GetPossibleValues(board)
	metrics := SolveMetrics{}
	solveByGuessing(&cfg, board, possibleValues, 0, 0, &metrics)
	return metrics
}

func solveByGuessing(cfg *GuessSolverConfig, board *board.SudokuBoard, values [9][9][9]bool, x, y int,
	metrics *SolveMetrics) bool {
	if board.GetAt(x, y) == 0 {
		anyPossibleValues := false
		for v := 0; v < 9; v++ {
			number := cfg.NumberOrder(x, y, v)
			if values[x][y][number] {
				anyPossibleValues = true
				metrics.tryCount++
				if tryValue(cfg, board, values, x, y, number+1, metrics) {
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
	} else {
		nextX, nextY, hasNext := getNextCoords(x, y)
		if !hasNext {
			return true
		}
		return solveByGuessing(cfg, board, values, nextX, nextY, metrics)
	}

	return board.GetAt(x, y) != 0
}

func tryValue(cfg *GuessSolverConfig, board *board.SudokuBoard, values [9][9][9]bool, x, y, value int,
	metrics *SolveMetrics) bool {
	board.SetAt(x, y, value)
	nextX, nextY, hasNext := getNextCoords(x, y)
	if !hasNext {
		return true
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

func getNextCoords(x, y int) (int, int, bool) {
	nextX := x + 1
	nextY := y
	if nextX == 9 {
		nextX = 0
		nextY++
		if nextY == 9 {
			return 0, 0, false
		}
	}
	return nextX, nextY, true
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
