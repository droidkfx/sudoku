package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello, World!"))
	})
	mux.Handle("/", http.FileServer(http.Dir("./web")))

	fmt.Println("Starting server, access at http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		fmt.Println(err)
	}
}
