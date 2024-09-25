package main

import (
	"fmt"
	"net/http"

	"droidkfx.com/sudoku/cmd/server/controller"
)

func main() {
	mux := http.NewServeMux()
	controller.RegisterHealthHandlers(mux)
	controller.RegisterBoardHandlers(mux)
	mux.Handle("/", http.FileServer(http.Dir("./web")))

	fmt.Println("Starting server, access at http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		fmt.Println(err)
	}
}
