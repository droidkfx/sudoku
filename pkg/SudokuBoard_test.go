package pkg

import "testing"

func BenchmarkSudokuBoard_GetAt(b *testing.B) {
	s := &SudokuBoard{
		board: [9][9]int{},
	}
	for i := 0; i < 81; i++ {
		s.board[i%9][i/9] = i
	}

	// Ignore setup information
	b.ResetTimer()
	var value int
	for i := 0; i < b.N; i++ {
		// prevent compiler optimizations
		value = s.GetAt(i%9, i%9)

	}
	s.board[0][4] = value
}

func initDefaultBoard() SudokuBoard {
	s := SudokuBoard{
		board: [9][9]int{},
	}
	for i := 0; i < 81; i++ {
		s.board[i/9][i%9] = i
	}
	return s
}

func TestSudokuBoard_GetAt(t *testing.T) {
	type args struct {
		x, y int
	}
	type want struct {
		value       int
		shouldPanic bool
	}
	tests := []struct {
		name  string
		board SudokuBoard
		args  args
		want  want
	}{
		{
			name:  "(0,0)",
			board: initDefaultBoard(),
			args:  args{x: 0, y: 0},
			want:  want{value: 0},
		},
		{
			name:  "(1,0)",
			board: initDefaultBoard(),
			args:  args{x: 1, y: 0},
			want:  want{value: 1},
		},
		{
			name:  "(0,1)",
			board: initDefaultBoard(),
			args:  args{x: 0, y: 1},
			want:  want{value: 9},
		},
		{
			name:  "(5,5)",
			board: initDefaultBoard(),
			args:  args{x: 5, y: 5},
			want:  want{value: 50},
		},
		{
			name:  "(8,8)",
			board: initDefaultBoard(),
			args:  args{x: 8, y: 8},
			want:  want{value: 80},
		},
		{
			name:  "(-1,0)!!",
			board: initDefaultBoard(),
			args:  args{x: -1, y: 0},
			want:  want{shouldPanic: true},
		},
		{
			name:  "(0,-1)!!",
			board: initDefaultBoard(),
			args:  args{x: 0, y: -1},
			want:  want{shouldPanic: true},
		},
		{
			name:  "(9,0)!!",
			board: initDefaultBoard(),
			args:  args{x: 9, y: 0},
			want:  want{shouldPanic: true},
		},
		{
			name:  "(0,9)!!",
			board: initDefaultBoard(),
			args:  args{x: 0, y: 9},
			want:  want{shouldPanic: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					if tt.want.shouldPanic {
						t.Errorf("GetAt() did not panic")
					}
				} else if !tt.want.shouldPanic {
					t.Errorf("GetAt() panicked: %s", r)
				}
			}()

			if got := tt.board.GetAt(tt.args.x, tt.args.y); got != tt.want.value {
				t.Errorf("GetAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSudokuBoard_SetAt(t *testing.T) {
	type args struct {
		x, y int
	}
	type want struct {
		shouldPanic bool
	}
	tests := []struct {
		name  string
		board SudokuBoard
		args  args
		want  want
	}{
		{
			name:  "(0,0)",
			board: initDefaultBoard(),
			args:  args{x: 0, y: 0},
		},
		{
			name:  "(1,0)",
			board: initDefaultBoard(),
			args:  args{x: 1, y: 0},
		},
		{
			name:  "(0,1)",
			board: initDefaultBoard(),
			args:  args{x: 0, y: 1},
		},
		{
			name:  "(5,5)",
			board: initDefaultBoard(),
			args:  args{x: 5, y: 5},
		},
		{
			name:  "(8,8)",
			board: initDefaultBoard(),
			args:  args{x: 8, y: 8},
		},
		{
			name:  "(-1,0)!!",
			board: initDefaultBoard(),
			args:  args{x: -1, y: 0},
			want:  want{shouldPanic: true},
		},
		{
			name:  "(0,-1)!!",
			board: initDefaultBoard(),
			args:  args{x: 0, y: -1},
			want:  want{shouldPanic: true},
		},
		{
			name:  "(9,0)!!",
			board: initDefaultBoard(),
			args:  args{x: 9, y: 0},
			want:  want{shouldPanic: true},
		},
		{
			name:  "(0,9)!!",
			board: initDefaultBoard(),
			args:  args{x: 0, y: 9},
			want:  want{shouldPanic: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					if tt.want.shouldPanic {
						t.Errorf("SetAt() did not panic")
					}
				} else if !tt.want.shouldPanic {
					t.Errorf("SetAt() panicked: %s", r)
				}
			}()
			tt.board.SetAt(tt.args.x, tt.args.y, 99)
			if got := tt.board.GetAt(tt.args.x, tt.args.y); got != 99 {
				t.Errorf("SetAt() = %v, want %v", got, tt.want)
			}
		})
	}
}
