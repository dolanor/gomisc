package main

import (
	"encoding/json"
	"io"
	"net/http"
	"log"
)

func handle[Req, Res any](method string, f func(Req) (Res, error)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, "bad method", http.StatusMethodNotAllowed)
			return
		}

		var req Req

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res, err := f(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err =json.NewEncoder(w).Encode(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func main() {
	http.HandleFunc("/uppercase", handle(http.MethodPost, UpperCase))
	http.HandleFunc("/snakecase", handle(http.MethodPost, SnakeCase))

	log.Println("listening on :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
