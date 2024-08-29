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
					0x21, 0x43, 0x65, 0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21,
					0x43, 0x65, 0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21, 0x43,
					0x65, 0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21, 0x43, 0x65,
					0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21, 0x43, 0x65, 0x87,
					0x09,
				},
			},
		},
		{
			name: "Bad File",
			args: args{
				initialFileContent: []byte{
					0x21, 0x43, 0x65, 0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21,
					0x43, 0x65, 0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21, 0x43,
					0x65, 0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21, 0x43, 0x65,
					0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21, 0x43, 0x65, 0x87,
					0x09,
				},
			},
		},
		{
			name: "Full board",
			args: args{
				initialFileContent: []byte{
					0x21, 0x43, 0x65, 0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21,
					0x43, 0x65, 0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21, 0x43,
					0x65, 0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21, 0x43, 0x65,
					0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21, 0x43, 0x65, 0x87,
					0x09,
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
					0x21, 0x43, 0x65, 0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21,
					0x43, 0x65, 0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21, 0x43,
					0x65, 0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21, 0x43, 0x65,
					0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21, 0x43, 0x65, 0x87,
					0x09,
				},
			},
		},
		{
			name: "Middle board",
			args: args{
				initialFileContent: []byte{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00,
					0x21, 0x43, 0x65, 0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21,
					0x43, 0x65, 0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21, 0x43,
					0x65, 0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21, 0x43, 0x65,
					0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21, 0x43, 0x65, 0x87,
					0x09,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00,
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
			_ = s.GetRandom()
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
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00,
				},
			},
			want: want{
				sudokuBoard: board.FromNumbers([9][9]int{}),
			},
		},
		{
			name: "Empty File",
			args: args{
				boardNumber:        0,
				initialFileContent: []byte{},
			},
			want: want{
				sudokuBoard: board.FromNumbers([9][9]int{}),
			},
		},
		{
			name: "Beyond File",
			args: args{
				boardNumber: 1,
				initialFileContent: []byte{
					0x21, 0x43, 0x65, 0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21,
					0x43, 0x65, 0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21, 0x43,
					0x65, 0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21, 0x43, 0x65,
					0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21, 0x43, 0x65, 0x87,
					0x09,
				},
			},
			want: want{
				sudokuBoard: board.FromNumbers([9][9]int{}),
			},
		},
		{
			name: "Bad File",
			args: args{
				boardNumber: 1,
				initialFileContent: []byte{
					0x21, 0x43, 0x65, 0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21,
					0x43, 0x65, 0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21, 0x43,
					0x65, 0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21, 0x43, 0x65,
				},
			},
			want: want{
				sudokuBoard: board.FromNumbers([9][9]int{}),
			},
		},
		{
			name: "Full board",
			args: args{
				boardNumber: 0,
				initialFileContent: []byte{
					0x21, 0x43, 0x65, 0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21,
					0x43, 0x65, 0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21, 0x43,
					0x65, 0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21, 0x43, 0x65,
					0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21, 0x43, 0x65, 0x87,
					0x09,
				},
			},
			want: want{
				sudokuBoard: board.FromNumbers([9][9]int{
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
				}),
			},
		},
		{
			name: "Skip board",
			args: args{
				boardNumber: 1,
				initialFileContent: []byte{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00,
					0x21, 0x43, 0x65, 0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21,
					0x43, 0x65, 0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21, 0x43,
					0x65, 0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21, 0x43, 0x65,
					0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21, 0x43, 0x65, 0x87,
					0x09,
				},
			},
			want: want{
				sudokuBoard: board.FromNumbers([9][9]int{
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
				}),
			},
		},
		{
			name: "Middle board",
			args: args{
				boardNumber: 1,
				initialFileContent: []byte{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00,
					0x21, 0x43, 0x65, 0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21,
					0x43, 0x65, 0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21, 0x43,
					0x65, 0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21, 0x43, 0x65,
					0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21, 0x43, 0x65, 0x87,
					0x09,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00,
				},
			},
			want: want{
				sudokuBoard: board.FromNumbers([9][9]int{
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
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
			b := s.GetNumber(tt.args.boardNumber)
			sd()

			if !reflect.DeepEqual(b, tt.want.sudokuBoard) {
				t.Errorf("sudokuBoardFileRepo.SaveNew() = \n%v, want \n%v", b, tt.want.sudokuBoard)
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
				sudokuBoard: board.FromNumbers([9][9]int{}),
			},
			want: want{
				expectedContent: []byte{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00,
				},
			},
		},
		{
			name: "Always write to end of file",
			args: args{
				sudokuBoard:        board.FromNumbers([9][9]int{}),
				initialFileContent: []byte{255},
			},
			want: want{
				expectedContent: []byte{
					0xFF,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00,
				},
			},
		},
		{
			name: "full list",
			args: args{
				sudokuBoard: board.FromNumbers([9][9]int{
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
				}),
			},
			want: want{
				expectedContent: []byte{
					0x21, 0x43, 0x65, 0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21,
					0x43, 0x65, 0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21, 0x43,
					0x65, 0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21, 0x43, 0x65,
					0x87, 0x19, 0x32, 0x54, 0x76, 0x98, 0x21, 0x43, 0x65, 0x87,
					0x09,
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
			s.SaveNew(tt.args.sudokuBoard)
			sd()

			file, _ = fs.Open(dbFileName)
			content, _ := afero.ReadAll(file)
			if !reflect.DeepEqual(content, tt.want.expectedContent) {
				t.Errorf("sudokuBoardFileRepo.SaveNew() = \n%v, want \n%v", content, tt.want.expectedContent)
			}
		})
	}
}
