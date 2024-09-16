package solver

import (
	"reflect"
	"testing"

	"droidkfx.com/sudoku/pkg/board"
)

// func TestSolverTest(T *testing.T) {
// 	v := SolveByStrategies(&board.SudokuBoard{})
// 	for _, step := range v {
// 		fmt.Printf("%v\n", step)
// 	}
// }

func TestStrategies(t *testing.T) {
	type args struct {
		b     *board.SudokuBoard
		strat StrategyMethod
	}
	tests := []struct {
		name string
		args args
		want []*StrategyStep
	}{
		{
			name: "Psychic empty",
			args: args{
				b:     &board.SudokuBoard{},
				strat: PsychicStrategy,
			},
			want: []*StrategyStep{
				{
					actions: []StrategyAction{{set: true, opts: false, x: 0, y: 0, value: 1}},
					name:    StrategyNamePsychicStrategy,
				},
				{
					actions: []StrategyAction{{set: true, opts: false, x: 1, y: 0, value: 2}},
					name:    StrategyNamePsychicStrategy,
				},
				{
					actions: []StrategyAction{{set: true, opts: false, x: 2, y: 0, value: 3}},
					name:    StrategyNamePsychicStrategy,
				},
			},
		},
		{
			name: "Psychic finds next element",
			args: args{
				b: board.FromNumbers([9][9]int{
					{2, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
				}),
				strat: PsychicStrategy,
			},
			want: []*StrategyStep{
				{
					actions: []StrategyAction{{set: true, opts: false, x: 1, y: 0, value: 1}},
					name:    StrategyNamePsychicStrategy,
				},
				{
					actions: []StrategyAction{{set: true, opts: false, x: 2, y: 0, value: 3}},
					name:    StrategyNamePsychicStrategy,
				},
				{
					actions: []StrategyAction{{set: true, opts: false, x: 3, y: 0, value: 4}},
					name:    StrategyNamePsychicStrategy,
				},
			},
		},
		{
			name: "Psychic Finishes Board",
			args: args{
				b: board.FromNumbers([9][9]int{
					{0, 1, 5, 4, 2, 6, 7, 8, 9},
					{4, 2, 6, 7, 8, 9, 3, 1, 5},
					{7, 8, 9, 0, 1, 5, 4, 2, 6},
					{1, 3, 4, 5, 6, 2, 8, 9, 7},
					{5, 6, 2, 8, 9, 7, 1, 0, 4},
					{8, 9, 7, 1, 3, 4, 5, 6, 2},
					{2, 5, 3, 6, 4, 1, 9, 7, 8},
					{6, 4, 1, 9, 7, 0, 2, 5, 3},
					{9, 7, 8, 2, 5, 3, 6, 4, 0},
				}),
				strat: PsychicStrategy,
			},
			want: []*StrategyStep{
				{
					actions: []StrategyAction{{set: true, opts: false, x: 0, y: 0, value: 3}},
					name:    StrategyNamePsychicStrategy,
				},
				{
					actions: []StrategyAction{{set: true, opts: false, x: 3, y: 2, value: 3}},
					name:    StrategyNamePsychicStrategy,
				},
				{
					actions: []StrategyAction{{set: true, opts: false, x: 7, y: 4, value: 3}},
					name:    StrategyNamePsychicStrategy,
				},
				{
					actions: []StrategyAction{{set: true, opts: false, x: 5, y: 7, value: 8}},
					name:    StrategyNamePsychicStrategy,
				},
				{
					actions: []StrategyAction{{set: true, opts: false, x: 8, y: 8, value: 1}},
					name:    StrategyNamePsychicStrategy,
				},
				nil,
			},
		},
		{
			name: "LastCandidate Finishes Board",
			args: args{
				b: board.FromNumbers([9][9]int{
					{0, 1, 5, 4, 2, 6, 7, 8, 9},
					{4, 2, 6, 7, 8, 9, 3, 1, 5},
					{7, 8, 9, 0, 1, 5, 4, 2, 6},
					{1, 3, 4, 5, 6, 2, 8, 9, 7},
					{5, 6, 2, 8, 9, 7, 1, 0, 4},
					{8, 9, 7, 1, 3, 4, 5, 6, 2},
					{2, 5, 3, 6, 4, 1, 9, 7, 8},
					{6, 4, 1, 9, 7, 0, 2, 5, 3},
					{9, 7, 8, 2, 5, 3, 6, 4, 0},
				}),
				strat: LastCandidateStrategy,
			},
			want: []*StrategyStep{
				{
					actions: []StrategyAction{{set: true, opts: false, x: 0, y: 0, value: 3}},
					name:    StrategyNameLastCandidateStrategy,
				},
				{
					actions: []StrategyAction{{set: true, opts: false, x: 3, y: 2, value: 3}},
					name:    StrategyNameLastCandidateStrategy,
				},
				{
					actions: []StrategyAction{{set: true, opts: false, x: 7, y: 4, value: 3}},
					name:    StrategyNameLastCandidateStrategy,
				},
				{
					actions: []StrategyAction{{set: true, opts: false, x: 5, y: 7, value: 8}},
					name:    StrategyNameLastCandidateStrategy,
				},
				{
					actions: []StrategyAction{{set: true, opts: false, x: 8, y: 8, value: 1}},
					name:    StrategyNameLastCandidateStrategy,
				},
				nil,
			},
		},
		{
			name: "LastCandidate Stops",
			args: args{
				b: board.FromNumbers([9][9]int{
					{0, 0, 0, 4, 2, 6, 0, 0, 0},
					{0, 0, 0, 7, 8, 9, 0, 0, 0},
					{0, 0, 0, 0, 1, 5, 0, 0, 0},
					{0, 0, 0, 5, 6, 2, 0, 0, 0},
					{0, 0, 0, 8, 9, 7, 0, 0, 0},
					{0, 0, 0, 1, 3, 4, 0, 0, 0},
					{0, 0, 0, 6, 4, 1, 0, 0, 0},
					{0, 0, 0, 9, 7, 0, 0, 0, 0},
					{0, 0, 0, 2, 5, 3, 0, 0, 0},
				}),
				strat: LastCandidateStrategy,
			},
			want: []*StrategyStep{
				{
					actions: []StrategyAction{{set: true, opts: false, x: 3, y: 2, value: 3}},
					name:    StrategyNameLastCandidateStrategy,
				},
				{
					actions: []StrategyAction{{set: true, opts: false, x: 5, y: 7, value: 8}},
					name:    StrategyNameLastCandidateStrategy,
				},
				nil,
			},
		},
		{
			name: "LastInColumnStrategy Finishes Board",
			args: args{
				b: board.FromNumbers([9][9]int{
					{0, 1, 5, 4, 2, 6, 7, 8, 9},
					{4, 2, 6, 7, 8, 9, 3, 1, 5},
					{7, 8, 9, 0, 1, 5, 4, 2, 6},
					{1, 3, 4, 5, 6, 2, 8, 9, 7},
					{5, 6, 2, 8, 9, 7, 1, 0, 4},
					{8, 9, 7, 1, 3, 4, 5, 6, 2},
					{2, 5, 3, 6, 4, 1, 9, 7, 8},
					{6, 4, 1, 9, 7, 0, 2, 5, 3},
					{9, 7, 8, 2, 5, 3, 6, 4, 0},
				}),
				strat: LastInColumnStrategy,
			},
			want: []*StrategyStep{
				{
					actions: []StrategyAction{{set: true, opts: false, x: 0, y: 0, value: 3}},
					name:    StrategyNameLastInColumnStrategy,
				},
				{
					actions: []StrategyAction{{set: true, opts: false, x: 3, y: 2, value: 3}},
					name:    StrategyNameLastInColumnStrategy,
				},
				{
					actions: []StrategyAction{{set: true, opts: false, x: 5, y: 7, value: 8}},
					name:    StrategyNameLastInColumnStrategy,
				},
				{
					actions: []StrategyAction{{set: true, opts: false, x: 7, y: 4, value: 3}},
					name:    StrategyNameLastInColumnStrategy,
				},

				{
					actions: []StrategyAction{{set: true, opts: false, x: 8, y: 8, value: 1}},
					name:    StrategyNameLastInColumnStrategy,
				},
				nil,
			},
		},
		{
			name: "LastInColumnStrategy Stops",
			args: args{
				b: board.FromNumbers([9][9]int{
					{0, 0, 0, 4, 2, 6, 0, 0, 0},
					{0, 0, 0, 7, 8, 9, 0, 0, 0},
					{0, 0, 0, 0, 1, 5, 0, 0, 0},
					{0, 0, 0, 5, 6, 2, 0, 0, 0},
					{0, 0, 0, 8, 9, 7, 0, 0, 0},
					{0, 0, 0, 1, 3, 4, 0, 0, 0},
					{0, 0, 0, 6, 4, 1, 0, 0, 0},
					{0, 0, 0, 9, 7, 0, 0, 0, 0},
					{0, 0, 0, 2, 5, 3, 0, 0, 0},
				}),
				strat: LastInColumnStrategy,
			},
			want: []*StrategyStep{
				{
					actions: []StrategyAction{{set: true, opts: false, x: 3, y: 2, value: 3}},
					name:    StrategyNameLastInColumnStrategy,
				},
				{
					actions: []StrategyAction{{set: true, opts: false, x: 5, y: 7, value: 8}},
					name:    StrategyNameLastInColumnStrategy,
				},
				nil,
			},
		},
		{
			name: "LastInRowStrategy Finishes Board",
			args: args{
				b: board.FromNumbers([9][9]int{
					{0, 1, 5, 4, 2, 6, 7, 8, 9},
					{4, 2, 6, 7, 8, 9, 3, 1, 5},
					{7, 8, 9, 0, 1, 5, 4, 2, 6},
					{1, 3, 4, 5, 6, 2, 8, 9, 7},
					{5, 6, 2, 8, 9, 7, 1, 0, 4},
					{8, 9, 7, 1, 3, 4, 5, 6, 2},
					{2, 5, 3, 6, 4, 1, 9, 7, 8},
					{6, 4, 1, 9, 7, 0, 2, 5, 3},
					{9, 7, 8, 2, 5, 3, 6, 4, 0},
				}),
				strat: LastInRowStrategy,
			},
			want: []*StrategyStep{
				{
					actions: []StrategyAction{{set: true, opts: false, x: 0, y: 0, value: 3}},
					name:    StrategyNameLastInRowStrategy,
				},
				{
					actions: []StrategyAction{{set: true, opts: false, x: 3, y: 2, value: 3}},
					name:    StrategyNameLastInRowStrategy,
				},
				{
					actions: []StrategyAction{{set: true, opts: false, x: 7, y: 4, value: 3}},
					name:    StrategyNameLastInRowStrategy,
				},
				{
					actions: []StrategyAction{{set: true, opts: false, x: 5, y: 7, value: 8}},
					name:    StrategyNameLastInRowStrategy,
				},
				{
					actions: []StrategyAction{{set: true, opts: false, x: 8, y: 8, value: 1}},
					name:    StrategyNameLastInRowStrategy,
				},
				nil,
			},
		},
		{
			name: "LastInRowStrategy Stops",
			args: args{
				b: board.FromNumbers([9][9]int{
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{7, 8, 9, 0, 1, 5, 4, 2, 6},
					{1, 3, 4, 5, 6, 2, 8, 9, 7},
					{5, 6, 2, 8, 9, 7, 1, 0, 4},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
				}),
				strat: LastInRowStrategy,
			},
			want: []*StrategyStep{
				{
					actions: []StrategyAction{{set: true, opts: false, x: 3, y: 2, value: 3}},
					name:    StrategyNameLastInRowStrategy,
				},
				{
					actions: []StrategyAction{{set: true, opts: false, x: 7, y: 4, value: 3}},
					name:    StrategyNameLastInRowStrategy,
				},
				nil,
			},
		},
		{
			name: "LastInRegionStrategy Finishes Board",
			args: args{
				b: board.FromNumbers([9][9]int{
					{0, 1, 5, 4, 2, 6, 7, 8, 9},
					{4, 2, 6, 7, 8, 9, 3, 1, 5},
					{7, 8, 9, 0, 1, 5, 4, 2, 6},
					{1, 3, 4, 5, 6, 2, 8, 9, 7},
					{5, 6, 2, 8, 9, 7, 1, 0, 4},
					{8, 9, 7, 1, 3, 4, 5, 6, 2},
					{2, 5, 3, 6, 4, 1, 9, 7, 8},
					{6, 4, 1, 9, 7, 0, 2, 5, 3},
					{9, 7, 8, 2, 5, 3, 6, 4, 0},
				}),
				strat: LastInRegionStrategy,
			},
			want: []*StrategyStep{
				{
					actions: []StrategyAction{{set: true, opts: false, x: 0, y: 0, value: 3}},
					name:    StrategyNameLastInRegionStrategy,
				},
				{
					actions: []StrategyAction{{set: true, opts: false, x: 3, y: 2, value: 3}},
					name:    StrategyNameLastInRegionStrategy,
				},
				{
					actions: []StrategyAction{{set: true, opts: false, x: 7, y: 4, value: 3}},
					name:    StrategyNameLastInRegionStrategy,
				},
				{
					actions: []StrategyAction{{set: true, opts: false, x: 5, y: 7, value: 8}},
					name:    StrategyNameLastInRegionStrategy,
				},
				{
					actions: []StrategyAction{{set: true, opts: false, x: 8, y: 8, value: 1}},
					name:    StrategyNameLastInRegionStrategy,
				},
				nil,
			},
		},
		{
			name: "LastInRegionStrategy Stops",
			args: args{
				b: board.FromNumbers([9][9]int{
					{0, 0, 0, 4, 2, 6, 0, 0, 0},
					{0, 0, 0, 7, 8, 9, 0, 0, 0},
					{0, 0, 0, 0, 1, 5, 0, 0, 0},
					{0, 0, 0, 5, 6, 2, 0, 0, 0},
					{0, 0, 0, 8, 9, 7, 0, 0, 0},
					{0, 0, 0, 1, 3, 4, 0, 0, 0},
					{0, 0, 0, 6, 4, 1, 0, 0, 0},
					{0, 0, 0, 9, 7, 0, 0, 0, 0},
					{0, 0, 0, 2, 5, 3, 0, 0, 0},
				}),
				strat: LastInRegionStrategy,
			},
			want: []*StrategyStep{
				{
					actions: []StrategyAction{{set: true, opts: false, x: 3, y: 2, value: 3}},
					name:    StrategyNameLastInRegionStrategy,
				},
				{
					actions: []StrategyAction{{set: true, opts: false, x: 5, y: 7, value: 8}},
					name:    StrategyNameLastInRegionStrategy,
				},
				nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := GetPossibleValues(tt.args.b)
			for _, step := range tt.want {
				if got := tt.args.strat(tt.args.b, &opts); !reflect.DeepEqual(got, step) {
					t.Errorf("got %v, want %v", got, step)
				}
				if step != nil {
					ApplyStep(tt.args.b, *step, &opts)
				}
			}
		})
	}
}
