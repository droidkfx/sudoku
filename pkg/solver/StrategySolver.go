package solver

import "droidkfx.com/sudoku/pkg/board"

type StrategyAction struct {
	set   bool
	opts  bool
	x     int
	y     int
	value int
}

type StrategyName string

const (
	StrategyNamePsychicStrategy       StrategyName = "Psychic"
	StrategyNameLastCandidateStrategy StrategyName = "LastCandidate"
	StrategyNameLastInRowStrategy     StrategyName = "LastInRow"
	StrategyNameLastInColumnStrategy  StrategyName = "LastInColumn"
	StrategyNameLastInRegionStrategy  StrategyName = "LastInRegion"
)

type StrategyDifficulty uint8

const (
	StrategyDifficultyEasy StrategyDifficulty = iota
	StrategyDifficultyMedium
	StrategyDifficultyHard
	StrategyDifficultyVeryHard
	StrategyDifficultyImpossible
)

type StrategyStep struct {
	actions []StrategyAction
	name    StrategyName
}

type StrategyMethod func(b *board.SudokuBoard, opts *[9][9][9]bool) *StrategyStep

var strategyDifficultyMap = map[StrategyName]StrategyDifficulty{
	StrategyNameLastInColumnStrategy:  StrategyDifficultyEasy,
	StrategyNameLastInRegionStrategy:  StrategyDifficultyEasy,
	StrategyNameLastInRowStrategy:     StrategyDifficultyEasy,
	StrategyNameLastCandidateStrategy: StrategyDifficultyMedium,
	StrategyNamePsychicStrategy:       StrategyDifficultyImpossible,
}

// List of strategies, this list should be sorted by difficulty since they will be tried in order
var strategies = []StrategyMethod{
	LastInRowStrategy,
	LastInColumnStrategy,
	LastInRegionStrategy,
	LastCandidateStrategy,
	PsychicStrategy,
}

func SolveByStrategies(b *board.SudokuBoard) []StrategyStep {
	opts := GetPossibleValues(b)
	var steps []StrategyStep
	for !board.IsSolved(b) {
		step := *SolveNextStep(b, &opts)
		steps = append(steps, step)
		ApplyStep(b, step, &opts)
	}

	return steps
}

func ApplyStep(b *board.SudokuBoard, step StrategyStep, opts *[9][9][9]bool) {
	for _, action := range step.actions {
		if action.set {
			if action.opts {
				opts[action.x][action.y][action.value-1] = false
			} else {
				b.SetAt(action.x, action.y, action.value)
				propagateNumberSetToOptions(opts, action.x, action.y, action.value)
			}
		} else {
			if action.opts {
				opts[action.x][action.y][action.value-1] = true
			} else {
				b.SetAt(action.x, action.y, 0)
			}
		}
	}
}

func SolveNextStep(b *board.SudokuBoard, opts *[9][9][9]bool) *StrategyStep {
	for _, strategy := range strategies {
		step := strategy(b, opts)
		if step != nil {
			return step
		}
	}
	return nil
}

func findNextEmpty(x, y int, b *board.SudokuBoard) (int, int, bool) {
	if x >= 9 {
		x = 0
		y++
	}

	for y < 9 {
		if b.GetAt(x, y) == 0 {
			return x, y, true
		}

		x++
		if x >= 9 {
			x = 0
			y++
		}
	}
	return 0, 0, false
}

var tmpPsychicBoard *board.SudokuBoard = nil

func PsychicStrategy(b *board.SudokuBoard, _ *[9][9][9]bool) *StrategyStep {
	if tmpPsychicBoard == nil || tmpPsychicBoard != b {
		tmpPsychicBoard = b.Copy()
		SolveByGuessing(DefaultGuessConfig(), tmpPsychicBoard)
	}
	x, y, hasNext := findNextEmpty(0, 0, b)

	for hasNext {
		solvedValue := tmpPsychicBoard.GetAt(x, y)
		if b.GetAt(x, y) == 0 && solvedValue != 0 {
			return &StrategyStep{
				name: StrategyNamePsychicStrategy,
				actions: []StrategyAction{
					{set: true, opts: false, x: x, y: y, value: tmpPsychicBoard.GetAt(x, y)},
				},
			}
		}
		x, y, hasNext = findNextEmpty(x+1, y, b)
	}

	return nil
}

func LastCandidateStrategy(b *board.SudokuBoard, opts *[9][9][9]bool) *StrategyStep {
	x, y, hasNext := findNextEmpty(0, 0, b)

	for hasNext {
		candidates := 0
		lastCandidate := 0
		for i := 0; i < 9 && candidates <= 1; i++ {
			if opts[x][y][i] {
				candidates++
				lastCandidate = i + 1
			}
		}
		if candidates == 1 {
			return &StrategyStep{
				name: StrategyNameLastCandidateStrategy,
				actions: []StrategyAction{
					{set: true, opts: false, x: x, y: y, value: lastCandidate},
				},
			}
		}
		x, y, hasNext = findNextEmpty(x+1, y, b)
	}

	return nil
}

func LastInRowStrategy(_ *board.SudokuBoard, opts *[9][9][9]bool) *StrategyStep {
	for y := 0; y < 9; y++ {
		seenCount := [9]int{}
		lastSeenX := [9]int{}
		for x := 0; x < 9; x++ {
			for opt := 0; opt < 9; opt++ {
				if opts[x][y][opt] {
					seenCount[opt]++
					lastSeenX[opt] = x
				}
			}
		}

		for i := 0; i < 9; i++ {
			if seenCount[i] == 1 {
				return &StrategyStep{
					name: StrategyNameLastInRowStrategy,
					actions: []StrategyAction{
						{set: true, opts: false, x: lastSeenX[i], y: y, value: i + 1},
					},
				}
			}
		}
	}
	return nil
}

func LastInColumnStrategy(_ *board.SudokuBoard, opts *[9][9][9]bool) *StrategyStep {
	for x := 0; x < 9; x++ {
		seenCount := [9]int{}
		lastSeenY := [9]int{}
		for y := 0; y < 9; y++ {
			for opt := 0; opt < 9; opt++ {
				if opts[x][y][opt] {
					seenCount[opt]++
					lastSeenY[opt] = y
				}
			}
		}

		for i := 0; i < 9; i++ {
			if seenCount[i] == 1 {
				return &StrategyStep{
					name: StrategyNameLastInColumnStrategy,
					actions: []StrategyAction{
						{set: true, opts: false, x: x, y: lastSeenY[i], value: i + 1},
					},
				}
			}
		}
	}
	return nil
}

func LastInRegionStrategy(_ *board.SudokuBoard, opts *[9][9][9]bool) *StrategyStep {
	for region := 0; region < 9; region++ {
		regionX := (region % 3) * 3
		regionY := (region / 3) * 3

		seenCount := [9]int{}
		lastSeenX := [9]int{}
		lastSeenY := [9]int{}
		for x := regionX; x < regionX+3; x++ {
			for y := regionY; y < regionY+3; y++ {
				for opt := 0; opt < 9; opt++ {
					if opts[x][y][opt] {
						seenCount[opt]++
						lastSeenX[opt] = x
						lastSeenY[opt] = y
					}
				}
			}
		}

		for i := 0; i < 9; i++ {
			if seenCount[i] == 1 {
				return &StrategyStep{
					name: StrategyNameLastInRegionStrategy,
					actions: []StrategyAction{
						{set: true, opts: false, x: lastSeenX[i], y: lastSeenY[i], value: i + 1},
					},
				}
			}
		}
	}
	return nil
}
