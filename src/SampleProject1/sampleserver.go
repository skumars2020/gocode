package main

import (
	"fmt"
	"net/http"
)

func main() {
	// multiplixcer or router
	mux := http.NewServeMux()
	mux.HandleFunc("/go", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
		fmt.Println("request receved")
	})
	http.ListenAndServe("localhost:3000", mux)

}
