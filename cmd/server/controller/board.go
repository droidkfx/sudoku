package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"droidkfx.com/sudoku/pkg/board"
	"droidkfx.com/sudoku/pkg/repository"
)

func RegisterBoardHandlers(mux *http.ServeMux, r repository.SudokuBoardRepo) {
	c := &boardController{
		r: r,
	}

	mux.HandleFunc("GET /board/{id}", c.GetBoardById)
}

type GetBoardByIdResponse struct {
	Id    int       `json:"id"`
	Board [9][9]int `json:"board"`
}

func (b *boardController) GetBoardById(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.PathValue("id"))
	if err != nil || id < 0 {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(writer).Encode(GetBoardByIdResponse{Id: id, Board: b.SudokuBoardToResponseBoard(b.r.GetNumber(id))})
}

func (b *boardController) SudokuBoardToResponseBoard(brd *board.SudokuBoard) [9][9]int {
	data := [9][9]int{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			data[j][i] = brd.GetAt(i, j)
		}
	}
	return data
}

type boardController struct {
	r repository.SudokuBoardRepo
}
