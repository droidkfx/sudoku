package board

/*
VerifyBoard checks if a SudokuBoard is valid. It checks if the board is valid by checking if each column, row, and
region is valid. It returns true if the board is valid and false otherwise. It uses the following helper functions:
  - VerifyColumn
  - VerifyRow
  - VerifyRegion
*/
func VerifyBoard(board *SudokuBoard) bool {
	for i := 0; i < 9; i++ {
		if !VerifyColumn(board, i) {
			return false
		}
		if !VerifyRow(board, i) {
			return false
		}
		if !VerifyRegion(board, i) {
			return false
		}
	}
	return true
}

/*
VerifyColumn checks if a row is valid. A column is considered valid if the row contains no duplicates of the
numbers 1...9.

It returns false if there are any duplicates, it will panic if the board contains any numbers outside the
range [1,9], and true otherwise.
*/
func VerifyColumn(board *SudokuBoard, col int) bool {
	seen := [9]bool{}
	for y := 0; y < 9; y++ {
		val := board.GetAt(col, y)
		if val == 0 {
			continue
		} else if seen[val-1] { // - 1 because we will see 1 thru 9 but the array is zero indexed
			return false
		} else {
			seen[val-1] = true // - 1 because we will see 1 thru 9 but the array is zero indexed
		}
	}
	return true
}

/*
VerifyRow checks if a row is valid. A row is considered valid if the row contains no duplicates of the
numbers 1...9.

It returns false if there are any duplicates, it will panic if the board contains any numbers outside the
range [1,9], and true otherwise.
*/
func VerifyRow(board *SudokuBoard, row int) bool {
	seen := [9]bool{}
	for x := 0; x < 9; x++ {
		val := board.GetAt(x, row)
		if val == 0 {
			continue
		} else if seen[val-1] { // - 1 because we will see 1 thru 9 but the array is zero indexed
			return false
		} else {
			seen[val-1] = true // - 1 because we will see 1 thru 9 but the array is zero indexed
		}
	}
	return true
}

/*
VerifyRegion checks if a region is valid. A region is considered valid if the region contains no duplicates of the
numbers 1...9.

A region is a section of the game board that is 3x3. The regions are numbered as follows:

	0 | 1 | 2
	---------
	3 | 4 | 5
	---------
	6 | 7 | 8

So Region 4 for example contains all elements from (3,3) to (5,5).

It returns false if there are any duplicates, it will panic if the board contains any numbers outside the
range [1,9], and true otherwise.
*/
func VerifyRegion(board *SudokuBoard, region int) bool {
	seen := [9]bool{}
	regionCoord := sudokuRegionMap[region]
	for x := regionCoord.xstart; x <= regionCoord.xend; x++ {
		for y := regionCoord.ystart; y <= regionCoord.yend; y++ {
			val := board.GetAt(x, y)
			if val == 0 {
				continue
			} else if seen[val-1] { // - 1 because we will see 1 thru 9 but the array is zero indexed
				return false
			} else {
				seen[val-1] = true // - 1 because we will see 1 thru 9 but the array is zero indexed
			}
		}
	}
	return true
}

type sudokuRegionCoord struct {
	xstart, ystart, xend, yend int
}

var sudokuRegionMap = map[int]sudokuRegionCoord{
	0: {0, 0, 2, 2},
	1: {3, 0, 5, 2},
	2: {6, 0, 8, 2},
	3: {0, 3, 2, 5},
	4: {3, 3, 5, 5},
	5: {6, 3, 8, 5},
	6: {0, 6, 2, 8},
	7: {3, 6, 5, 8},
	8: {6, 6, 8, 8},
}
