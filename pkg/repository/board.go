package repository

import (
	"encoding/binary"
	"math/rand"
	osConst "os"

	"droidkfx.com/sudoku/pkg/board"
	"github.com/spf13/afero"
)

const dbBoardFile = "boards.bin"
const boardDataBytes = 41
const lowMask uint8 = 0b0000_1111
const highMask uint8 = 0b1111_0000

type SudokuBoardRepo interface {
	GetRandom() (int, *board.SudokuBoard)
	GetByNumber(n int) (int, *board.SudokuBoard)
	SaveNew(sudokuBoard *board.SudokuBoard)
	SaveAll(sudokuBoards []*board.SudokuBoard)
}

func NewSudokuBoardRepo(dbLocation string) (SudokuBoardRepo, func()) {
	return NewSudokuBoardRepoUsingFs(afero.NewBasePathFs(afero.NewOsFs(), dbLocation))
}

func NewSudokuBoardRepoUsingFs(fileSystem afero.Fs) (SudokuBoardRepo, func()) {
	dbFile, err := fileSystem.OpenFile(dbBoardFile, osConst.O_CREATE|osConst.O_RDWR, 0666)
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

func (s *sudokuBoardFileRepo) GetRandom() (int, *board.SudokuBoard) {
	dataLength := int64(binary.Size(byte(0)) * 81)
	fStat, err := s.home.Stat()
	if err != nil {
		panic(err)
	}

	fSize := fStat.Size()
	if fSize < dataLength {
		return 0, board.FromNumbers([9][9]int{})
	}

	loadedBoard, id := s.loadBoard(rand.Intn(int(fSize / dataLength)))
	return id, s.dataToBoard(loadedBoard)
}

func (s *sudokuBoardFileRepo) GetByNumber(n int) (int, *board.SudokuBoard) {
	loadedBoard, id := s.loadBoard(n)
	return id, s.dataToBoard(loadedBoard)
}

func (s *sudokuBoardFileRepo) SaveNew(sudokuBoard *board.SudokuBoard) {
	s.SaveAll([]*board.SudokuBoard{sudokuBoard})
}

func (s *sudokuBoardFileRepo) SaveAll(sudokuBoards []*board.SudokuBoard) {
	data := make([]byte, 0, boardDataBytes*len(sudokuBoards))
	for _, b := range sudokuBoards {
		data = append(data, s.boardToData(b)...)
	}
	s.saveToEnd(data)
}

func (s *sudokuBoardFileRepo) loadBoard(n int) ([]byte, int) {
	data := make([]byte, 41)
	dataSize := int64(41)
	offset := dataSize * int64(n)

	fStat, err := s.home.Stat()
	if err != nil {
		panic(err)
	}

	if fStat.Size() <= offset {
		return data, 0 // there is no such index
	}
	_, err = s.home.ReadAt(data, offset)
	if err != nil {
		panic(err)
	}

	return data, n
}

func (s *sudokuBoardFileRepo) saveToEnd(data []byte) {
	fStat, err := s.home.Stat()
	if err != nil {
		panic(err)
	}

	_, err = s.home.WriteAt(data, fStat.Size())
	if err != nil {
		panic(err)
	}
}

func (s *sudokuBoardFileRepo) boardToData(b *board.SudokuBoard) []byte {
	// 9x9 is 81 bytes, to make it divisible by 2 we add 1 -> 41
	result := make([]byte, 0, boardDataBytes)

	i, j := 0, 0
	var currentByte byte
	var high bool
	for j < 9 {
		if !high {
			currentByte = lowMask
			currentByte &= byte(b.GetAt(i, j))
			high = true
		} else {
			currentByte |= highMask & (byte(b.GetAt(i, j)) << 4)
			result = append(result, currentByte)
			high = false
		}

		i++
		if i >= 9 {
			j++
			i = 0
		}
	}
	result = append(result, currentByte)

	return result
}

func (s *sudokuBoardFileRepo) dataToBoard(data []byte) *board.SudokuBoard {
	b := board.SudokuBoard{}
	for j := 0; j < 9; j++ {
		for i := 0; i < 9; i++ {
			index := i + (j * 9)
			// even index -> high byte
			if index%2 == 0 {
				b.SetAt(i, j, int(data[index/2]&lowMask))
			} else {
				b.SetAt(i, j, int(data[index/2]&highMask>>4))
			}
		}
	}
	return &b
}
