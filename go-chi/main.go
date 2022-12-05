package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
    chi := chi.NewRouter()
    chi.Use(middleware.Logger)
    chi.Get("/hello", func(res http.ResponseWriter, req *http.Request) {
        res.Write([]byte("hello"))
    })
    http.ListenAndServe(":8080", chi)
}
