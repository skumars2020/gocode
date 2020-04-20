package main

import (
	"encoding/json"
	"net/http"
	"D:/Go/gocode/src/Project/structs"
)

func main() {

	mux := http.NewServeMux("/ping/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := structs.Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			json.NewEncoder(w).Encode(data)
		}
	})

	http.ListenAndServe(":3000", mux)

}
