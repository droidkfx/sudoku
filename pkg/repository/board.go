package repository

import (
	"encoding/binary"
	"math/rand"
	osConst "os"

	"droidkfx.com/sudoku/pkg/board"
	"github.com/spf13/afero"
)

const dbFileName = "sudoku.db"

type SudokuBoardRepo interface {
	GetRandomBoard() *board.SudokuBoard
	GetBoardNumber(n int) *board.SudokuBoard
	SaveNewBoard(sudokuBoard *board.SudokuBoard)
}

func NewSudokuBoardRepo(dbLocation string) (SudokuBoardRepo, func()) {
	return NewSudokuBoardRepoUsingFs(afero.NewBasePathFs(afero.NewOsFs(), dbLocation))
}

func NewSudokuBoardRepoUsingFs(fileSystem afero.Fs) (SudokuBoardRepo, func()) {
	dbFile, err := fileSystem.OpenFile(dbFileName, osConst.O_CREATE|osConst.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	s := &sudokuBoardFileRepo{
		home: dbFile,
	}

	return s, s.shutdown
}

type sudokuBoardFileRepo struct {
	home afero.File
}

func (s *sudokuBoardFileRepo) shutdown() {
	_ = s.home.Sync()
	_ = s.home.Close()
}

func (s *sudokuBoardFileRepo) GetRandomBoard() *board.SudokuBoard {
	dataLength := int64(binary.Size(byte(0)) * 81)
	fStat, err := s.home.Stat()
	if err != nil {
		panic(err)
	}

	fSize := fStat.Size()
	if fSize < dataLength {
		return board.BoardFromNumbers([9][9]int{})
	}

	return s.dataToBoard(s.loadLine(rand.Intn(int(fSize / dataLength))))
}

func (s *sudokuBoardFileRepo) GetBoardNumber(n int) *board.SudokuBoard {
	return s.dataToBoard(s.loadLine(n))
}

func (s *sudokuBoardFileRepo) SaveNewBoard(sudokuBoard *board.SudokuBoard) {
	s.saveLineAtEnd(s.boardToData(sudokuBoard))
}

func (s *sudokuBoardFileRepo) loadLine(n int) []byte {
	data := make([]byte, 81)
	dataSize := int64(binary.Size(byte(0)) * 81)
	offset := dataSize * int64(n)

	fStat, err := s.home.Stat()
	if err != nil {
		panic(err)
	}

	if fStat.Size() <= offset {
		return data // there is no such index
	}
	_, err = s.home.ReadAt(data, offset)
	if err != nil {
		panic(err)
	}

	return data
}

func (s *sudokuBoardFileRepo) saveLineAtEnd(data []byte) {
	fStat, err := s.home.Stat()
	if err != nil {
		panic(err)
	}

	_, err = s.home.WriteAt(data, fStat.Size())
	if err != nil {
		panic(err)
	}
}

func (s *sudokuBoardFileRepo) boardToData(sudokuBoard *board.SudokuBoard) []byte {
	result := make([]byte, 81)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			result[i+(j*9)] = byte(sudokuBoard.GetAt(i, j))
		}
	}
	return result
}

func (s *sudokuBoardFileRepo) dataToBoard(data []byte) *board.SudokuBoard {
	b := board.SudokuBoard{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			b.SetAt(i, j, int(data[i+(j*9)]))
		}
	}
	return &b
}
