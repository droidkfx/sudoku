package repository

import (
	"reflect"
	"testing"

	"droidkfx.com/sudoku/pkg/board"
	"github.com/spf13/afero"
)

func Test_sudokuBoardFileRepo_GetRandomBoard(t *testing.T) {
	type args struct {
		initialFileContent []byte
	}
	type want struct {
		sudokuBoard *board.SudokuBoard
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Zero board",
			args: args{
				initialFileContent: []byte{
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
				},
			},
		},
		{
			name: "Empty File",
			args: args{
				initialFileContent: []byte{},
			},
		},
		{
			name: "Beyond File",
			args: args{
				initialFileContent: []byte{
					1, 2, 3, 4, 5, 6, 7, 8, 9,
					10, 11, 12, 13, 14, 15, 16, 17, 18,
					19, 20, 21, 22, 23, 24, 25, 26, 27,
					28, 29, 30, 31, 32, 33, 34, 35, 36,
					37, 38, 39, 40, 41, 42, 43, 44, 45,
					46, 47, 48, 49, 50, 51, 52, 53, 54,
					55, 56, 57, 58, 59, 60, 61, 62, 63,
					64, 65, 66, 67, 68, 69, 70, 71, 72,
					73, 74, 75, 76, 77, 78, 79, 80, 81,
				},
			},
		},
		{
			name: "Bad File",
			args: args{
				initialFileContent: []byte{
					1, 2, 3, 4, 5, 6, 7, 8, 9,
					10, 11, 12, 13, 14, 15, 16, 17, 18,
					19, 20, 21, 22, 23, 24, 25, 26, 27,
					28, 29, 30, 31, 32, 33, 34, 35, 36,
					37, 38, 39, 40, 41, 42, 43, 44, 45,
					46, 47, 48, 49, 50, 51, 52, 53, 54,
					55, 56, 57, 58, 59, 60, 61, 62, 63,
				},
			},
		},
		{
			name: "Full board",
			args: args{
				initialFileContent: []byte{
					1, 2, 3, 4, 5, 6, 7, 8, 9,
					10, 11, 12, 13, 14, 15, 16, 17, 18,
					19, 20, 21, 22, 23, 24, 25, 26, 27,
					28, 29, 30, 31, 32, 33, 34, 35, 36,
					37, 38, 39, 40, 41, 42, 43, 44, 45,
					46, 47, 48, 49, 50, 51, 52, 53, 54,
					55, 56, 57, 58, 59, 60, 61, 62, 63,
					64, 65, 66, 67, 68, 69, 70, 71, 72,
					73, 74, 75, 76, 77, 78, 79, 80, 81,
				},
			},
		},
		{
			name: "Skip board",
			args: args{
				initialFileContent: []byte{
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					1, 2, 3, 4, 5, 6, 7, 8, 9,
					10, 11, 12, 13, 14, 15, 16, 17, 18,
					19, 20, 21, 22, 23, 24, 25, 26, 27,
					28, 29, 30, 31, 32, 33, 34, 35, 36,
					37, 38, 39, 40, 41, 42, 43, 44, 45,
					46, 47, 48, 49, 50, 51, 52, 53, 54,
					55, 56, 57, 58, 59, 60, 61, 62, 63,
					64, 65, 66, 67, 68, 69, 70, 71, 72,
					73, 74, 75, 76, 77, 78, 79, 80, 81,
				},
			},
		},
		{
			name: "Middle board",
			args: args{
				initialFileContent: []byte{
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					1, 2, 3, 4, 5, 6, 7, 8, 9,
					10, 11, 12, 13, 14, 15, 16, 17, 18,
					19, 20, 21, 22, 23, 24, 25, 26, 27,
					28, 29, 30, 31, 32, 33, 34, 35, 36,
					37, 38, 39, 40, 41, 42, 43, 44, 45,
					46, 47, 48, 49, 50, 51, 52, 53, 54,
					55, 56, 57, 58, 59, 60, 61, 62, 63,
					64, 65, 66, 67, 68, 69, 70, 71, 72,
					73, 74, 75, 76, 77, 78, 79, 80, 81,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("VerifyRegion() panicked: %s", r)
				}
			}()

			fs := afero.NewMemMapFs()
			file, _ := fs.Create(dbFileName)
			_, _ = file.Write(tt.args.initialFileContent)
			_ = file.Sync()
			_ = file.Close()

			s, sd := NewSudokuBoardRepoUsingFs(fs)
			_ = s.GetRandomBoard()
			sd()
		})
	}
}

func Test_sudokuBoardFileRepo_GetBoardNumber(t *testing.T) {
	type args struct {
		initialFileContent []byte
		boardNumber        int
	}
	type want struct {
		sudokuBoard *board.SudokuBoard
	}
	tests := []struct {
		name string
		want want
		args args
	}{
		{
			name: "Zero board",
			args: args{
				boardNumber: 0,
				initialFileContent: []byte{
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
				},
			},
			want: want{
				sudokuBoard: board.BoardFromNumbers([9][9]int{}),
			},
		},
		{
			name: "Empty File",
			args: args{
				boardNumber:        0,
				initialFileContent: []byte{},
			},
			want: want{
				sudokuBoard: board.BoardFromNumbers([9][9]int{}),
			},
		},
		{
			name: "Beyond File",
			args: args{
				boardNumber: 1,
				initialFileContent: []byte{
					1, 2, 3, 4, 5, 6, 7, 8, 9,
					10, 11, 12, 13, 14, 15, 16, 17, 18,
					19, 20, 21, 22, 23, 24, 25, 26, 27,
					28, 29, 30, 31, 32, 33, 34, 35, 36,
					37, 38, 39, 40, 41, 42, 43, 44, 45,
					46, 47, 48, 49, 50, 51, 52, 53, 54,
					55, 56, 57, 58, 59, 60, 61, 62, 63,
					64, 65, 66, 67, 68, 69, 70, 71, 72,
					73, 74, 75, 76, 77, 78, 79, 80, 81,
				},
			},
			want: want{
				sudokuBoard: board.BoardFromNumbers([9][9]int{}),
			},
		},
		{
			name: "Bad File",
			args: args{
				boardNumber: 1,
				initialFileContent: []byte{
					1, 2, 3, 4, 5, 6, 7, 8, 9,
					10, 11, 12, 13, 14, 15, 16, 17, 18,
					19, 20, 21, 22, 23, 24, 25, 26, 27,
					28, 29, 30, 31, 32, 33, 34, 35, 36,
					37, 38, 39, 40, 41, 42, 43, 44, 45,
					46, 47, 48, 49, 50, 51, 52, 53, 54,
					55, 56, 57, 58, 59, 60, 61, 62, 63,
				},
			},
			want: want{
				sudokuBoard: board.BoardFromNumbers([9][9]int{}),
			},
		},
		{
			name: "Full board",
			args: args{
				boardNumber: 0,
				initialFileContent: []byte{
					1, 2, 3, 4, 5, 6, 7, 8, 9,
					10, 11, 12, 13, 14, 15, 16, 17, 18,
					19, 20, 21, 22, 23, 24, 25, 26, 27,
					28, 29, 30, 31, 32, 33, 34, 35, 36,
					37, 38, 39, 40, 41, 42, 43, 44, 45,
					46, 47, 48, 49, 50, 51, 52, 53, 54,
					55, 56, 57, 58, 59, 60, 61, 62, 63,
					64, 65, 66, 67, 68, 69, 70, 71, 72,
					73, 74, 75, 76, 77, 78, 79, 80, 81,
				},
			},
			want: want{
				sudokuBoard: board.BoardFromNumbers([9][9]int{
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{10, 11, 12, 13, 14, 15, 16, 17, 18},
					{19, 20, 21, 22, 23, 24, 25, 26, 27},
					{28, 29, 30, 31, 32, 33, 34, 35, 36},
					{37, 38, 39, 40, 41, 42, 43, 44, 45},
					{46, 47, 48, 49, 50, 51, 52, 53, 54},
					{55, 56, 57, 58, 59, 60, 61, 62, 63},
					{64, 65, 66, 67, 68, 69, 70, 71, 72},
					{73, 74, 75, 76, 77, 78, 79, 80, 81},
				}),
			},
		},
		{
			name: "Skip board",
			args: args{
				boardNumber: 1,
				initialFileContent: []byte{
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					1, 2, 3, 4, 5, 6, 7, 8, 9,
					10, 11, 12, 13, 14, 15, 16, 17, 18,
					19, 20, 21, 22, 23, 24, 25, 26, 27,
					28, 29, 30, 31, 32, 33, 34, 35, 36,
					37, 38, 39, 40, 41, 42, 43, 44, 45,
					46, 47, 48, 49, 50, 51, 52, 53, 54,
					55, 56, 57, 58, 59, 60, 61, 62, 63,
					64, 65, 66, 67, 68, 69, 70, 71, 72,
					73, 74, 75, 76, 77, 78, 79, 80, 81,
				},
			},
			want: want{
				sudokuBoard: board.BoardFromNumbers([9][9]int{
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{10, 11, 12, 13, 14, 15, 16, 17, 18},
					{19, 20, 21, 22, 23, 24, 25, 26, 27},
					{28, 29, 30, 31, 32, 33, 34, 35, 36},
					{37, 38, 39, 40, 41, 42, 43, 44, 45},
					{46, 47, 48, 49, 50, 51, 52, 53, 54},
					{55, 56, 57, 58, 59, 60, 61, 62, 63},
					{64, 65, 66, 67, 68, 69, 70, 71, 72},
					{73, 74, 75, 76, 77, 78, 79, 80, 81},
				}),
			},
		},
		{
			name: "Middle board",
			args: args{
				boardNumber: 1,
				initialFileContent: []byte{
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					1, 2, 3, 4, 5, 6, 7, 8, 9,
					10, 11, 12, 13, 14, 15, 16, 17, 18,
					19, 20, 21, 22, 23, 24, 25, 26, 27,
					28, 29, 30, 31, 32, 33, 34, 35, 36,
					37, 38, 39, 40, 41, 42, 43, 44, 45,
					46, 47, 48, 49, 50, 51, 52, 53, 54,
					55, 56, 57, 58, 59, 60, 61, 62, 63,
					64, 65, 66, 67, 68, 69, 70, 71, 72,
					73, 74, 75, 76, 77, 78, 79, 80, 81,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
				},
			},
			want: want{
				sudokuBoard: board.BoardFromNumbers([9][9]int{
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{10, 11, 12, 13, 14, 15, 16, 17, 18},
					{19, 20, 21, 22, 23, 24, 25, 26, 27},
					{28, 29, 30, 31, 32, 33, 34, 35, 36},
					{37, 38, 39, 40, 41, 42, 43, 44, 45},
					{46, 47, 48, 49, 50, 51, 52, 53, 54},
					{55, 56, 57, 58, 59, 60, 61, 62, 63},
					{64, 65, 66, 67, 68, 69, 70, 71, 72},
					{73, 74, 75, 76, 77, 78, 79, 80, 81},
				}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("VerifyRegion() panicked: %s", r)
				}
			}()

			fs := afero.NewMemMapFs()
			file, _ := fs.Create(dbFileName)
			_, _ = file.Write(tt.args.initialFileContent)
			_ = file.Sync()
			_ = file.Close()

			s, sd := NewSudokuBoardRepoUsingFs(fs)
			b := s.GetBoardNumber(tt.args.boardNumber)
			sd()

			if !reflect.DeepEqual(b, tt.want.sudokuBoard) {
				t.Errorf("sudokuBoardFileRepo.SaveNewBoard() = \n%v, want \n%v", b, tt.want.sudokuBoard)
			}
		})
	}
}

func Test_sudokuBoardFileRepo_SaveNewBoard(t *testing.T) {
	type args struct {
		sudokuBoard        *board.SudokuBoard
		initialFileContent []byte
	}
	type want struct {
		expectedContent []byte
	}
	tests := []struct {
		name string
		want want
		args args
	}{
		{
			name: "Zero board",
			args: args{
				sudokuBoard: board.BoardFromNumbers([9][9]int{}),
			},
			want: want{
				expectedContent: []byte{
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
				},
			},
		},
		{
			name: "Always write to end of file",
			args: args{
				sudokuBoard:        board.BoardFromNumbers([9][9]int{}),
				initialFileContent: []byte{255},
			},
			want: want{
				expectedContent: []byte{
					255,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
				},
			},
		},
		{
			name: "full list",
			args: args{
				sudokuBoard: board.BoardFromNumbers([9][9]int{
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{10, 11, 12, 13, 14, 15, 16, 17, 18},
					{19, 20, 21, 22, 23, 24, 25, 26, 27},
					{28, 29, 30, 31, 32, 33, 34, 35, 36},
					{37, 38, 39, 40, 41, 42, 43, 44, 45},
					{46, 47, 48, 49, 50, 51, 52, 53, 54},
					{55, 56, 57, 58, 59, 60, 61, 62, 63},
					{64, 65, 66, 67, 68, 69, 70, 71, 72},
					{73, 74, 75, 76, 77, 78, 79, 80, 81},
				}),
			},
			want: want{
				expectedContent: []byte{
					1, 2, 3, 4, 5, 6, 7, 8, 9,
					10, 11, 12, 13, 14, 15, 16, 17, 18,
					19, 20, 21, 22, 23, 24, 25, 26, 27,
					28, 29, 30, 31, 32, 33, 34, 35, 36,
					37, 38, 39, 40, 41, 42, 43, 44, 45,
					46, 47, 48, 49, 50, 51, 52, 53, 54,
					55, 56, 57, 58, 59, 60, 61, 62, 63,
					64, 65, 66, 67, 68, 69, 70, 71, 72,
					73, 74, 75, 76, 77, 78, 79, 80, 81,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("VerifyRegion() panicked: %s", r)
				}
			}()

			fs := afero.NewMemMapFs()
			file, _ := fs.Create(dbFileName)
			_, _ = file.Write(tt.args.initialFileContent)
			_ = file.Sync()
			_ = file.Close()

			s, sd := NewSudokuBoardRepoUsingFs(fs)
			s.SaveNewBoard(tt.args.sudokuBoard)
			sd()

			file, _ = fs.Open(dbFileName)
			content, _ := afero.ReadAll(file)
			if !reflect.DeepEqual(content, tt.want.expectedContent) {
				t.Errorf("sudokuBoardFileRepo.SaveNewBoard() = \n%v, want \n%v", content, tt.want.expectedContent)
			}
		})
	}
}
