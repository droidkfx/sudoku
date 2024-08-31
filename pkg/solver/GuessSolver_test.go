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

func TestSolveByGuessing(t *testing.T) {
	type args struct {
		cfg   GuessSolverConfig
		board *board.SudokuBoard
	}
	tests := []struct {
		name string
		args args
		want SolveMetrics
	}{
		{
			name: "zero board",
			args: args{
				cfg:   DefaultGuessConfig(),
				board: &board.SudokuBoard{},
			},
			want: SolveMetrics{tryCount: 81, resetCount: 1},
		},
		{
			name: "solved board",
			args: args{
				cfg: DefaultGuessConfig(),
				board: board.FromNumbers([9][9]int{
					{3, 9, 8, 4, 6, 2, 5, 7, 1},
					{4, 6, 2, 5, 7, 1, 3, 9, 8},
					{5, 7, 1, 3, 9, 8, 4, 6, 2},
					{9, 3, 4, 8, 2, 6, 7, 1, 5},
					{8, 2, 6, 7, 1, 5, 9, 3, 4},
					{7, 1, 5, 9, 3, 4, 8, 2, 6},
					{6, 8, 3, 2, 4, 9, 1, 5, 7},
					{2, 4, 9, 1, 5, 7, 6, 8, 3},
					{1, 5, 7, 6, 8, 3, 2, 4, 9},
				}),
			},
			want: SolveMetrics{tryCount: 0, resetCount: 0},
		},
		{
			name: "row missing board",
			args: args{
				cfg: DefaultGuessConfig(),
				board: board.FromNumbers([9][9]int{
					{3, 9, 8, 4, 6, 2, 5, 7, 1},
					{4, 6, 2, 5, 7, 1, 3, 9, 8},
					{5, 7, 1, 3, 9, 8, 4, 6, 2},
					{9, 3, 4, 8, 2, 6, 7, 1, 5},
					{8, 2, 6, 7, 1, 5, 9, 3, 4},
					{7, 1, 5, 9, 3, 4, 8, 2, 6},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{6, 8, 3, 2, 4, 9, 1, 5, 7},
					{2, 4, 9, 1, 5, 7, 6, 8, 3},
				}),
			},
			want: SolveMetrics{tryCount: 9, resetCount: 0},
		},
		{
			name: "double row missing board",
			args: args{
				cfg: DefaultGuessConfig(),
				board: board.FromNumbers([9][9]int{
					{3, 9, 8, 4, 6, 2, 5, 7, 1},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{5, 7, 1, 3, 9, 8, 4, 6, 2},
					{9, 3, 4, 8, 2, 6, 7, 1, 5},
					{8, 2, 6, 7, 1, 5, 9, 3, 4},
					{7, 1, 5, 9, 3, 4, 8, 2, 6},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{6, 8, 3, 2, 4, 9, 1, 5, 7},
					{2, 4, 9, 1, 5, 7, 6, 8, 3},
				}),
			},
			want: SolveMetrics{tryCount: 18, resetCount: 0},
		},
		{
			name: "double row & col missing board",
			args: args{
				cfg: DefaultGuessConfig(),
				board: board.FromNumbers([9][9]int{
					{3, 9, 0, 4, 6, 2, 0, 7, 1},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{5, 7, 0, 3, 9, 8, 0, 6, 2},
					{9, 3, 0, 8, 2, 6, 0, 1, 5},
					{8, 2, 0, 7, 1, 5, 0, 3, 4},
					{7, 1, 0, 9, 3, 4, 0, 2, 6},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{6, 8, 0, 2, 4, 9, 0, 5, 7},
					{2, 4, 0, 1, 5, 7, 0, 8, 3},
				}),
			},
			want: SolveMetrics{tryCount: 32, resetCount: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SolveByGuessing(tt.args.cfg, tt.args.board)
			if got.tryCount < tt.want.tryCount {
				t.Errorf("SolveByGuessing().tryCount = %v, want at least %v try", got.tryCount, tt.want.tryCount)
			}
			if got.resetCount < tt.want.resetCount {
				t.Errorf("SolveByGuessing().resetCount = %v, want at least %v reset", got.resetCount, tt.want.resetCount)
			}
			if !board.IsSolved(tt.args.board) {
				t.Errorf("SolveByGuessing() did not solve the board")
			}
		})
	}
}
