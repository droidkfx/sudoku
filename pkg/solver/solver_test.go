package solver

import (
	"testing"

	"droidkfx.com/sudoku/pkg/board"
)

func BenchmarkSolveByGuessing(b *testing.B) {
	var metrics SolveMetrics
	for i := 0; i < b.N; i++ {
		metrics = SolveByGuessing(DefaultGuessConfig(), &board.SudokuBoard{})
	}
	b.Log(metrics)
}
