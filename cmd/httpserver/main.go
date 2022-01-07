package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)


func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	err := http.ListenAndServe(":8080", r)
	fmt.Println(err)
}