package main

import (
	"fmt"
	"time"

	"droidkfx.com/sudoku/pkg/board"
	"droidkfx.com/sudoku/pkg/repository"
	"droidkfx.com/sudoku/pkg/solver"
)

func main() {
	r, sd := repository.NewSudokuBoardRepo("./data")
	defer sd()

	// Generate all permutations of the numbers 1-9
	perms := permutations([]int{0, 1, 2, 3, 4, 5, 6, 7, 8})

	startTime := time.Now()
	var totalTime time.Duration
	lastReset := startTime.Add(-1 * time.Hour)
	boards := make([]*board.SudokuBoard, len(perms))
	doneChan := make(chan bool)

	go func() {
		for j := 0; j < len(perms); j++ {
			boards[j] = &board.SudokuBoard{}
			go func(b *board.SudokuBoard, p []int) {
				_ = solver.SolveByGuessing(solver.GuessConfig(solver.NewStaticOrderGuesser(perms[j])), boards[j])
				doneChan <- true
			}(boards[j], perms[j])
		}
	}()

	doneCount := 0
	for range doneChan {
		doneCount++
		if time.Since(lastReset) > time.Millisecond*250 {
			totalTime = time.Since(startTime)
			fmt.Printf("Boards generated: %d (%d/s)\r", doneCount, int(float64(doneCount)/totalTime.Seconds()))
			lastReset = time.Now()
		}
		if doneCount == len(perms) {
			close(doneChan)
		}
	}

	fmt.Printf("Boards generated: %d (%d/s)\r\n", len(perms), int(float64(len(perms))/totalTime.Seconds()))
	fmt.Print("Saving Boards...")
	r.SaveAll(boards)
	fmt.Println("Done!")
	totalTime = time.Since(startTime)
	fmt.Printf("Total time: %v\n", totalTime)
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					arr[i], arr[n-1] = arr[n-1], arr[i]
				} else {
					arr[0], arr[n-1] = arr[n-1], arr[0]
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
