package main

import (
	"net/http"
)

func main() {

	mux := http.NewServeMux("/ping/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := structs.Response{
				Code: http.StatusOK,
				Body: "pong",
			}
		}
	})

	http.ListenAndServe(":3000", mux)

}
