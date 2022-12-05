package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(res http.ResponseWriter, req *http.Request){
		if req.Method == http.MethodGet {
			res.Write([]byte("hello"))
		}
	})
	http.ListenAndServe(":8080", nil)
}