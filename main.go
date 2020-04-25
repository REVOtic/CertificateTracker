package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./")))
	err := http.ListenAndServe(":9080", nil)
	if err != nil {
		panic(err)
	}
}
