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

	mux.HandleFunc("GET /board/random", c.GetRandomBoard)
	mux.HandleFunc("GET /board/{id}", c.GetBoardById)
}

type GetBoardByIdResponse struct {
	Id         int       `json:"id"`
	Difficulty string    `json:"difficulty"`
	Board      [9][9]int `json:"board"`
}

func (b *boardController) GetBoardById(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.PathValue("id"))
	if err != nil || id < 0 {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	idGot, nBoard := b.r.GetByNumber(id)
	_ = json.NewEncoder(writer).Encode(GetBoardByIdResponse{
		Id:         idGot,
		Difficulty: "TBD",
		Board:      b.SudokuBoardToResponseBoard(nBoard),
	})
}

func (b *boardController) GetRandomBoard(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	idGot, rBoard := b.r.GetRandom()
	_ = json.NewEncoder(writer).Encode(GetBoardByIdResponse{
		Id:         idGot,
		Difficulty: "TBD",
		Board:      b.SudokuBoardToResponseBoard(rBoard),
	})
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
